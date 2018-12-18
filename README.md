# request ID from Apex

	/aws/lambda/requestid_r START RequestId: 69b10604-027c-11e9-a485-1d345e710f4f Version: 1
	/aws/lambda/requestid_r {"fields":{"RequestID":"69b10604-027c-11e9-a485-1d345e710f4f"},"level":"info","timestamp":"2018-12-18T04:21:38.703721059Z","message":"Got the handle"}
	/aws/lambda/requestid_r {"fields":{"RequestID":"69b10604-027c-11e9-a485-1d345e710f4f"},"level":"info","timestamp":"2018-12-18T04:21:38.703844058Z","message":"Hello from Apex"}
	/aws/lambda/requestid_r END RequestId: 69b10604-027c-11e9-a485-1d345e710f4f
	/aws/lambda/requestid_r REPORT RequestId: 69b10604-027c-11e9-a485-1d345e710f4f  Duration: 47.83 ms      Billed Duration: 100 ms         Memory Size: 128 MB   Max Memory Used: 28 MB


# request ID from Up

There doesn't appear to be a context to tap into, so one each http request
handler one needs to do something like:

	ctx := log.WithFields(log.Fields{
		"request-id": r.Header.Get("X-Request-Id"),
	})
