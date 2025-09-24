# AWS IAM Connector

## Overview
The **AWS IAM Connector** is a Python-based integration project that fetches AWS IAM users and displays their attached and inline policies. It demonstrates how external systems can connect to AWS IAM via APIs and read user permissions.

**Key Features:**
- Fetch all IAM users in an AWS account.
- Fetch attached and inline policies for each user.
- Display permissions in a readable tabular format.
- Optional CLI arguments to fetch permissions for a specific user.
- Optional Flask-based web dashboard to visualize user permissions.

---

## Prerequisites
Before running this project, ensure you have:

1. **AWS Account** (Free Tier works).
2. **IAM User** with programmatic access:
   - Create an IAM user via AWS Console (e.g., `iam-connector-user`).
   - Attach `IAMReadOnlyAccess` policy.
   - Generate Access Key & Secret Key and download the `.csv`.
3. **AWS CLI** installed and configured:
```bash
aws configure
# Enter Access Key, Secret Key, default region (e.g., ap-south-1), and output format (json)
Python Environment (Python 3.x recommended):

bash
Copy code
python3 -m venv venv
# Activate the environment
# Linux / Mac
source venv/bin/activate
# Windows
venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt
Project Structure
bash
Copy code
aws-iam-connector/
├── iam_connector.py     # Main Python script
├── requirements.txt     # Python dependencies
└── README.md
Installation & Usage
1. Run Basic Script
bash
Copy code
python iam_connector.py
Expected Output:

pgsql
Copy code
User        Permissions
---------   ----------------------------
alice       AmazonS3ReadOnlyAccess
bob         CustomInlinePolicy
test-user   No Policies
2. Run for Specific User (Optional)
bash
Copy code
python iam_connector.py --user alice
Fetch permissions for a specific user only.

3. Optional: Flask Dashboard
Create a simple Flask app to display users and permissions in a web interface.

Run:

bash
Copy code
export FLASK_APP=app.py   # Linux / Mac
set FLASK_APP=app.py      # Windows
flask run
Implementation Details
Connect to AWS

python
Copy code
import boto3
iam = boto3.client('iam')
Fetch all users

python
Copy code
users = iam.list_users()['Users']
Fetch user permissions

python
Copy code
def get_user_permissions(user_name):
    policies = []

    # Attached policies
    attached = iam.list_attached_user_policies(UserName=user_name)
    for p in attached['AttachedPolicies']:
        policies.append(p['PolicyName'])
    
    # Inline policies
    inline = iam.list_user_policies(UserName=user_name)
    for p in inline['PolicyNames']:
        policies.append(p)
    
    return policies
Display in table

python
Copy code
from tabulate import tabulate

table = []
for user in users:
    uname = user['UserName']
    perms = get_user_permissions(uname)
    table.append([uname, ", ".join(perms) if perms else "No Policies"])

print(tabulate(table, headers=["User", "Permissions"]))