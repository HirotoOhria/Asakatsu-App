function_name = GetActivitiesFromSlackUidFunction
cmd_directory_name = get_activities_from_slack_uid
runtime_go_version = go116
slack_uid = U01DG785DT4

export_dotenv:
	source ../../src/.env
run_local:
	cd ../../src && \
 	go run ./app/$(cmd_directory_name)/main.go $(slack_uid) && \
 	cd ../cmd/$(cmd_directory_name)
# CloudFunction をデプロイする場合、go.mod が存在するディレクトリに移動する必要がある
create_function:
	make export_dotenv && \
	cd ../../src && \
	gcloud functions deploy $(function_name) \
		--region asia-northeast1 \
		--runtime $(runtime_go_version) \
		--trigger-http \
		--allow-unauthenticated \
		--set-env-vars GCP_PROJECT_ID=$(GCP_PROJECT_ID) \
		--set-env-vars SLACK_BOT_OAUTH_TOKEN=$(SLACK_BOT_OAUTH_TOKEN) \
		--set-env-vars SLACK_ASAKATSU_CHANNEL_ID=$(SLACK_ASAKATSU_CHANNEL_ID) && \
	 cd ../cmd/$(cmd_directory_name)
delete_function:
	gcloud functions delete $(function_name) \
		--region asia-northeast1 \
		--quiet
# CloudFunction をデプロイする場合、go.mod が存在するディレクトリに移動する必要がある
deploy:
	cd ../../src && \
	gcloud functions deploy $(function_name) \
		--region asia-northeast1 \
		--runtime $(runtime_go_version) \
		--trigger-http \
        --allow-unauthenticated && \
	cd ../cmd/$(cmd_directory_name)
delete_deploy:
	make delete_function;
	make deploy;
build_zip:
	cd ../../src && \
	zip -r fetch_activities_from_slach_batch.zip . && \
	cd ../cmd/$(cmd_directory_name)