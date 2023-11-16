import boto3
import json
import sys

def invoke_lambda():
    args = sys.argv
    if len(args) < 2:
        print('error: missing email to send data')
        return

    client = boto3.client('lambda')
    function = 'analyse_transactions-prod'

    payload_data = {
        "send_to": args[1]
    }

    try:
        payload = json.dumps(payload_data)
        response = client.invoke(
            FunctionName=function,
            Payload=payload
        )

        result = response['Payload'].read().decode('utf-8')
        print(f"Invoke result: {result}")

    except Exception as e:
        print(f"An error accurs on lambda calling: {e}")

if __name__ == "__main__":
    invoke_lambda()
