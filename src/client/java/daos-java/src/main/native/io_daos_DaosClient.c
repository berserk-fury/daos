/*
 * (C) Copyright 2018-2021 Intel Corporation.
 *
 * SPDX-License-Identifier: BSD-2-Clause-Patent
 */

#include "io_daos_DaosClient.h"
#include <sys/stat.h>
#include <gurt/common.h>
#include <libgen.h>
#include <stdio.h>
#include <daos.h>
#include <daos_jni_common.h>
#include <fcntl.h>

/**
 * JNI method to open pool with given \a poolId.
 *
 * \param[in]	env		JNI environment
 * \param[in]	clientClass	class of DaosFsClient
 * \param[in]	poolId		pool UUID
 * \param[in]	serverGroup	server group name
 * \param[in]	ranks		ranks separated by ':'
 *
 * \return	copied pool handle in long
 */
JNIEXPORT jlong JNICALL
Java_io_daos_DaosClient_daosOpenPool(JNIEnv *env,
				     jclass clientClass, jstring poolId,
				     jstring serverGroup,
				     jint flags)
{
	const char *pool_str = (*env)->GetStringUTFChars(env, poolId, 0);
	const char *server_group = (*env)->GetStringUTFChars(env, serverGroup,
								0);
	uuid_t pool_uuid;
	jlong ret;
	daos_handle_t poh;
	int rc;

	uuid_parse(pool_str, pool_uuid);
	rc = daos_pool_connect(pool_uuid,
			       server_group,
			       flags,
			       &poh /* returned pool handle */,
			       NULL /* returned pool info */,
			       NULL /* event */);
	if (rc) {
		char *msg = NULL;

		asprintf(&msg, "Failed to connect to pool (%s)",
			 pool_str);
		throw_base(env, msg, rc, 1, 0);
		ret = -1;
	} else {
		memcpy(&ret, &poh, sizeof(poh));
	}
	(*env)->ReleaseStringUTFChars(env, poolId, pool_str);
	if (serverGroup != NULL) {
		(*env)->ReleaseStringUTFChars(env, serverGroup, server_group);
	}
	return ret;
}

/**
 * JNI method to close pool denoted by \a poolHandle.
 *
 * \param[in]	env		JNI environment
 * \param[in]	clientClass	class of DaosFsClient
 * \param[in]	poolHandle	pool handle
 */
JNIEXPORT void JNICALL
Java_io_daos_DaosClient_daosClosePool(JNIEnv *env,
				      jclass clientClass, jlong poolHandle)
{
	daos_handle_t poh;

	memcpy(&poh, &poolHandle, sizeof(poh));
	int rc = daos_pool_disconnect(poh, NULL);

	if (rc) {
		printf("Failed to close pool rc: %d\n", rc);
		printf("error msg: %s\n", d_errstr(rc));
	}
}

/**
 * JNI method to open container with given \a contUuid.
 *
 * \param[in]	env		JNI environment
 * \param[in]	clientClass	class of DaosFsClient
 * \param[in]	poolHandle	pool handle
 * \param[in]	contUuid	container UUID
 * \param[in]	mode		container mode
 *
 * \return	copied container handle in long
 */
JNIEXPORT jlong JNICALL
Java_io_daos_DaosClient_daosOpenCont(JNIEnv *env,
				     jclass clientClass, jlong poolHandle,
				     jstring contUuid, jint mode)
{
	daos_handle_t poh;
	daos_cont_info_t co_info;
	const char *cont_str = (*env)->GetStringUTFChars(env, contUuid, NULL);
	uuid_t cont_uuid;
	daos_handle_t coh;
	jlong ret = -1;

	uuid_parse(cont_str, cont_uuid);
	memcpy(&poh, &poolHandle, sizeof(poh));
	int rc = daos_cont_open(poh, cont_uuid, mode, &coh, &co_info, NULL);

	if (rc) {
		char *msg = NULL;

		asprintf(&msg, "Failed to open container (id: %s)",
			 cont_str);
		throw_base(env, msg, rc, 1, 0);
		ret = -1;
	} else {
		memcpy(&ret, &coh, sizeof(coh));
	}
	(*env)->ReleaseStringUTFChars(env, contUuid, cont_str);
	return ret;
}

/**
 * JNI method to close container denoted by \a contHandle.
 *
 * \param[in]	env		JNI environment
 * \param[in]	clientClass	class of DaosFsClient
 * \param[in]	contHandle	container handle
 */
JNIEXPORT void JNICALL
Java_io_daos_DaosClient_daosCloseContainer(JNIEnv *env,
					   jclass clientClass,
					   jlong contHandle)
{
	daos_handle_t coh;

	memcpy(&coh, &contHandle, sizeof(coh));
	int rc = daos_cont_close(coh, NULL);

	if (rc) {
		printf("Failed to close container rc: %d\n", rc);
		printf("error msg: %.256s\n", d_errstr(rc));
	}
}

JNIEXPORT jlong JNICALL
Java_io_daos_DaosClient_createEventQueue(JNIEnv *env, jclass clientClass,
					 jint nbrOfEvents)
{
	daos_handle_t eqhdl;
	int rc = daos_eq_create(&eqhdl);
	int i;
	int count;

	if (rc) {
		throw_const(env, "Failed to create EQ", rc);
		return -1;
	}

	event_queue_wrapper_t *eq = (event_queue_wrapper_t *)calloc(1,
			sizeof(event_queue_wrapper_t));
	if (eq == NULL) {
		goto fail;
	}
	eq->events = (daos_event_t **)malloc(
		nbrOfEvents * sizeof(daos_event_t *));
	eq->polled_events = (daos_event_t **)malloc(
                nbrOfEvents * sizeof(daos_event_t *));
	if (eq->events == NULL || eq->polled_events == NULL) {
		char *msg = NULL;

		asprintf(&msg,
			 "Failed to allocate event array with length, %d",
			 nbrOfEvents);
		rc = 1;
		throw_base(env, msg, rc, 1, 0);
		goto fail;
	}
	eq->nbrOfEvents = nbrOfEvents;
	eq->eqhdl = eqhdl;
	for (i = 0; i < nbrOfEvents; i++) {
		eq->events[i] = (daos_event_t *)malloc(sizeof(daos_event_t));
		if (eq->events[i] == NULL) {
			char *msg = NULL;

			asprintf(&msg, "Failed to allocate %d th event.",
				 i);
			rc = 1;
			throw_base(env, msg, rc, 1, 0);
			goto fail;
		}
		rc = daos_event_init(eq->events[i], eqhdl, NULL);
		if (rc) {
			char *msg = NULL;

			asprintf(&msg, "Failed to init event %d",
				 i);
			throw_base(env, msg, rc, 1, 0);
			goto fail;
		}
		eq->events[i]->ev_debug = i;
	}

fail:
	if (rc) {
		count = i;
		while (i >= 0) {
			if (eq->events[i] && i < count) {
				daos_event_fini(eq->events[i]);
			}
			i--;
		}
		daos_eq_destroy(eqhdl, 1);
		for (i = 0; i <= count; i++) {
			if (eq->events[i]) {
				free(eq->events[i]);
			}
		}
		free(eq->polled_events);
		free(eq->events);
		free(eq);
	}
	return *(jlong *)&eq;
}

JNIEXPORT void JNICALL
Java_io_daos_DaosClient_pollCompleted(JNIEnv *env, jclass clientClass,
		jlong eqWrapperHdl, jlong memAddress,
		jint nbrOfEvents, jlong timeoutMs)
{
	event_queue_wrapper_t *eq = *(event_queue_wrapper_t **)&eqWrapperHdl;
	char *buffer = (char *)memAddress;
	uint16_t idx;
	int i;
	int rc = daos_eq_poll(eq->eqhdl, 1, timeoutMs * 1000, nbrOfEvents,
				eq->polled_events);

	if (rc < 0) {
		char *msg = NULL;

		asprintf(&msg,
			 "Failed to poll completed events, max events: %d",
			 nbrOfEvents);
		throw_base(env, msg, rc, 1, 0);
		return;
	}
	if (rc > nbrOfEvents) {
		char *msg = NULL;

		asprintf(&msg,
			 "More (%d) than expected (%d) events returned.",
			 rc, nbrOfEvents);
		throw_base(env, msg, rc, 1, 0);
		return;
	}
	idx = rc;
	memcpy(buffer, &idx, 2);
	buffer += 2;
	for (i = 0; i < rc; i++) {
		idx = eq->polled_events[i]->ev_debug;
		memcpy(buffer, &idx, 2);
		buffer += 2;
	}
}

JNIEXPORT void JNICALL
Java_io_daos_DaosClient_abortEvent(JNIEnv *env, jclass clientClass,
                jlong eqWrapperHdl, jshort eid)
{
        event_queue_wrapper_t *eq = *(event_queue_wrapper_t **)&eqWrapperHdl;
        daos_event_t *event = eq->events[eid];
        int rc;

        if (event->ev_error != EVENT_IN_USE) {
                return;
        }
        rc = daos_event_abort(event);
        event->ev_error = 0;
        if (rc) {
                char *msg = NULL;

                asprintf(&msg, "Failed to abort event (%d)",
                         event->ev_debug);
                throw_base(env, msg, rc, 1, 0);
        }
}

JNIEXPORT void JNICALL
Java_io_daos_DaosClient_destroyEventQueue(JNIEnv *env, jclass clientClass,
		jlong eqWrapperHdl)
{
	event_queue_wrapper_t *eq = *(event_queue_wrapper_t **)&eqWrapperHdl;
	int i;
	int rc;
	int count = 0;
	daos_event_t *ev;

	while (daos_eq_poll(eq->eqhdl, 1, 1000, eq->nbrOfEvents, eq->polled_events)) {
		count++;
		if (count > 4) {
			break;
		}
	}
	if (eq->events) {
		for (i = 0; i < eq->nbrOfEvents; i++) {
			ev = eq->events[i];
			if (!ev) {
				continue;
			}
			rc = daos_event_fini(ev);
			if (rc) {
				char *msg = NULL;

				asprintf(&msg,
					 "Failed to finalize %d th event.",
				         i);
				throw_base(env, msg, rc, 1, 0);
				goto fin;
			}
		}
	}
	if (eq->eqhdl.cookie) {
		rc = daos_eq_destroy(eq->eqhdl, 0);
		if (rc) {
			throw_const_obj(env,
					"Failed to destroy EQ.",
					rc);
			goto fin;
		}
	}
fin:
	if (eq->events) {
		for (i = 0; i < eq->nbrOfEvents; i++) {
		ev = eq->events[i];
			if (ev) {
				free(ev);
			}
		}
		free(eq->events);
	}
	if (eq->polled_events) {
	    free(eq->polled_events);
	}
	if (eq) {
		free(eq);
	}
}

/**
 * JNI method to finalize DAOS.
 *
 * \param[in]	env		JNI environment
 * \param[in]	clientClass	class of DaosFsClient
 */
JNIEXPORT void JNICALL
Java_io_daos_DaosClient_daosFinalize(JNIEnv *env,
					jclass clientClass)
{
	int rc = daos_eq_lib_fini();

	if (rc) {
		printf("Failed to finalize EQ lib rc: %d\n", rc);
		printf("error msg: %.256s\n", d_errstr(rc));
	}

	rc = daos_fini();
	if (rc) {
		printf("Failed to finalize daos rc: %d\n", rc);
		printf("error msg: %.256s\n", d_errstr(rc));
	}
}
