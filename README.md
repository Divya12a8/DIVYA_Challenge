# DIVYA_Challenge
This repository contains go lang code for json transformer based on criteria given in the challenge 

## Repo Link
https://github.com/Divya12a8/DIVYA_Challenge.git

## Execution

Go to - https://replit.com/~
Click on "+ Create Repl"
On the top right corner of pop-up, click Import from Github
Click "from URL" out of the tabs from the pop-up and paste https://github.com/Divya12a8/DIVYA_Challenge.git
Go to shell and navigate to main.go file located in submission folder by using "cd submission" command
Enter "go run ." command to see the output in the console 

OR

navigate to https://replit.com/@divya12a8/DIVYAChallenge using browser and login (the repo in repl is public so anyone should be able to access else raise permission)
Go to shell and navigate to main.go file located in submission folder by using "cd submission" command
Enter "go run ." command to see the output in the console 

## Steps used for statisc website deployment (all files are included in the repo):

Create EC2 Instances:
Launch two EC2 instances. One will be used for Ansible, and the other  will serve as server.
Make sure to configure security groups to allow SSH access to the Ansible instance and to the servers.
2.Install Ansible EC2 Instance:
Connect to the Ansible EC2 instance using SSH.
Run the following commands to install Ansible:

sudo yum update -y
sudo yum install -y epel-release
sudo yum install -y ansible

Verify the installation by running:
ansible --version

3.Configure Ansible Hosts File:
Edit the Ansible hosts file to define the servers to be managed by Ansible.

sudo nano /etc/ansible/hosts

Add the following lines to the hosts file:

[servers]
Webserver ansible_host=<Webserver_Public_IP>

Optionally, set up other variables like Python interpreter path, user, and SSH private key.

4.Create Playbooks and Files:

Create a directory for your Ansible playbooks.

mkdir playbooks
cd playbooks
Create a playbook file (playbook.yml) and add your playbook code to it.
Create necessary files like index.html for your web page and ssl.conf for SSL configuration.

Run Ansible Playbook:

Execute the playbook to configure the servers.

sudo ansible-playbook playbook.yml
When prompted, enter y or yes to confirm.
Adjust Permissions (Optional):

If necessary, adjust permissions for the SSH private key file.
bash
Copy code
sudo chmod 400 /home/ec2-user/server1.pem
This step may be necessary to ensure that Ansible can use the private key for SSH authentication.

5.Verify Configuration:
Check if the configuration was successful by accessing the server's public IP address in a web browser and verifying that the web page is served correctly.
Also, verify SSL configuration to ensure HTTPS redirection is working correctly.

Configure Security Group:
Optionally, configure security groups for your servers to allow only necessary inbound and outbound traffic, such as HTTP and HTTPS.
