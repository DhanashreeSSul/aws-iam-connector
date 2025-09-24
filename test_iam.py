import boto3

# Create IAM client (this uses your AWS CLI credentials automatically)
iam = boto3.client('iam')

# List IAM users
response = iam.list_users()

print("IAM Users in your account:")
for user in response['Users']:
    print(f"- {user['UserName']}")



from tabulate import tabulate

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

# Fetch all users and permissions
users = iam.list_users()['Users']
table = []
for user in users:
    uname = user['UserName']
    perms = get_user_permissions(uname)
    table.append([uname, ", ".join(perms) if perms else "No Policies"])

print(tabulate(table, headers=["User", "Permissions"]))
