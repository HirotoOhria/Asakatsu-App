function_name = FetchActivitiesFromSlackBatch
trigger_topic_name = fetch-activities-from-slack-batch-executer
cmd_directory_name = fetch_activities_from_slack
runtime_go_version = go116

export_dotenv:
	source ../../src/.env
run_local:
	cd ../../src && \
 	go run ./app/fetch_activities_from_slack/main.go && \
 	cd ../cmd/$(cmd_directory_name)
create_function:
	make export_dotenv && \
	cd ../../src && \
	gcloud functions deploy $(function_name) \
		--region asia-northeast1 \
		--runtime $(runtime_go_version) \
		--trigger-topic $(trigger_topic_name) \
		--set-env-vars GCP_PROJECT_ID=$(GCP_PROJECT_ID) \
		--set-env-vars SLACK_BOT_OAUTH_TOKEN=$(SLACK_BOT_OAUTH_TOKEN) \
		--set-env-vars SLACK_ASAKATSU_CHANNEL_ID=$(SLACK_ASAKATSU_CHANNEL_ID) && \
	 cd ../cmd/$(cmd_directory_name)
delete_function:
	gcloud functions delete $(function_name) \
		--region asia-northeast1 \
		--quiet
deploy:
	cd ../../src && \
	gcloud functions deploy $(function_name) \
		--region asia-northeast1 \
		--trigger-topic $(trigger_topic_name) \
		--runtime $(runtime_go_version) && \
	cd ../cmd/$(cmd_directory_name)
delete_deploy:
	make delete_function;
	make deploy;
build_zip:
	cd ../../src && \
	zip -r fetch_activities_from_slach_batch.zip . && \
	cd ../cmd/$(cmd_directory_name)