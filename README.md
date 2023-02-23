### What is SIMU-cleaner &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ![alt text](SIMU-cleaner.gif "SIMU")
SIMU stands for S3 Incomplete Multipart Uplods Cleaner. It is an intreactive CLI tool to clean AWS accounts that contain incomplete chunks of S3 data. 

### Problem
When large files are uploaded to S3 , they are split into smaller parts by the S3 client. Each individual part is upload and when all parts are uploaded, they are concatenated to form the final S3 object in the bucket. Sometimes uploads are aborted due to user actions or network outages. Such data accumulates and incurs storage cost for customers. This can be a significant amount for heavy S3 users. 

### What can EC2C do
1. Shows you a list of EC2 instances from the AWS account and chosen Region as per AWS credentials
2. Creates Amaxzon Machine Images for chosen EC2 instances.
3. Waits untill the images are in "available" status. Time taken depends on time taken for a snapshot to be created. Depends on the size of disk. For timeliness timeout occurs after 40 minutes.
4. Asks for target AWS account number. Thsi is the AWS account where you would want the EC2 instances to be migrated.
5. Changes permissions on newly created images. Adds the target account as a shared account for the image.

### Prerequisites
1. AWS credentials from source account with appropriate persmissions to create an AMI 
2. Target AWS account number

### How to get it
Chose the right artifact for your CPU architecture and OS type from https://github.com/bit-cloner/ec2c/releases
```
wget https://github.com/bit-cloner/ec2c/releases/download/0.9.1/ec2c-0.9.1-linux-amd64.tar.gz
```
```
tar -xvzf ec2c-0.9.1-linux-amd64.tar.gz
```
### How to use it
```
chmod +x ./ec2c
```
```
./ec2c
```
### Future improvements
1. Have an import mode so that the same CLI can be used to launch EC2s from shared AMI images
2. Handle KMS keys and permissions propogation to target account
3. Another Other features/ bugs , please file an issue
4. Code optimization to take advantage of go routines and functions

<h3 align="left">Support:</h3>
<p><a href="https://www.buymeacoffee.com/welldone"> <img align="left" src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="50" width="210" alt="welldone" /></a></p><br><br>


Disclaimer: Some of the code in this repo has been sugeested by ChatGPT3 and Github's Copilot