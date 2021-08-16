function_name = FetchActivitiesFromSlackBatch
trigger_topic_name = fetch-activities-from-slack-batch-executer

export_dotenv:
	set -a;
	source .env;
	set +a;
run_local:
	go run ./cmd/local/main.go
create_function:
	gcloud functions deploy $(function_name) \
		--region asia-northeast1 \
		--runtime go113 \
		--trigger-topic $(trigger_topic_name) \
		--set-env-vars GCP_PROJECT_ID=$(GCP_PROJECT_ID) \
		--set-env-vars SLACK_BOT_OAUTH_TOKEN=$(SLACK_BOT_OAUTH_TOKEN) \
		--set-env-vars SLACK_ASAKATSU_CHANNEL_ID=$(SLACK_ASAKATSU_CHANNEL_ID) 
delete_function:
	gcloud functions delete $(function_name) \
		--region asia-northeast1 \
		--quiet
delete_deploy:
	make delete_function;
	make deploy;
deploy:
	gcloud functions deploy $(function_name) \
		--region asia-northeast1 \
		--trigger-topic $(trigger_topic_name)
build_zip:
	zip -r fetch_activities_from_slach_batch.zip .