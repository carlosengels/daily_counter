import json
import base64
import requests
import os

github_token = os.getenv("")
github_api_url = f"https://github.com/carlosengels/daily_counter/blob/main/files/counter.txt"

def lambda_handler(event, context):
    headers = {"Authorization": f"token {github_token}", "Accept": "application/vnd.github.v3+json"}

    response = requests.get(github_api_url, headers=headers)
    if response.status_code != 200:
        return {"statusCode": response.status_code, "body": "Error fetching file from GitHub"}

    print(response.status_code)

    file_data = response.json()
    content = base64.b64decode(file_data["content"]).decode("utf-8").strip()

    try:
        counter_value = int(content)
    except ValueError:
        return {"statusCode": 400, "body": "Invalid file content"}

    counter_value += 1
    updated_content = str(counter_value)
    encoded_content = base64.b64encode(updated_content.encode("utf-8")).decode("utf-8")
    
    # Update the file on GitHub
    update_data = {
        "message": "Update counter",
        "content": encoded_content,
        "sha": file_data["sha"],
    }
    
    update_response = requests.put(github_api_url, headers=headers, json=update_data)
    if update_response.status_code != 200:
        return {"statusCode": update_response.status_code, "body": "Error updating file on GitHub"}
    
    return {"statusCode": 200, "body": "File updated successfully", "new_value": counter_value}

