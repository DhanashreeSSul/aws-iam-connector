from flask import Flask, render_template_string
import boto3

app = Flask(__name__)
iam = boto3.client('iam')

def get_user_permissions(user_name):
    permissions = []
    attached = iam.list_attached_user_policies(UserName=user_name)
    for policy in attached['AttachedPolicies']:
        permissions.append(policy['PolicyName'])
    inline = iam.list_user_policies(UserName=user_name)
    for policy in inline['PolicyNames']:
        permissions.append(policy)
    return permissions

@app.route("/")
def index():
    users = iam.list_users()['Users']
    rows = []
    for user in users:
        perms = get_user_permissions(user['UserName'])
        rows.append((user['UserName'], ", ".join(perms) if perms else "No Policies"))

    # Simple inline HTML
    html = """
    <h2>AWS IAM Connector</h2>
    <table border="1" cellpadding="5">
      <tr><th>User</th><th>Permissions</th></tr>
      {% for u, p in rows %}
      <tr><td>{{ u }}</td><td>{{ p }}</td></tr>
      {% endfor %}
    </table>
    """
    return render_template_string(html, rows=rows)

if __name__ == "__main__":
    app.run(debug=True)
