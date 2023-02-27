### What is SIMU-cleaner &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ![alt text](SIMU-cleaner.gif "SIMU")
SIMU stands for S3 Incomplete Multipart Uplods Cleaner. It is an intreactive CLI tool to clean AWS accounts that contain incomplete chunks of S3 data. 

### Problem
When large files are uploaded to S3 , they are split into smaller parts by the S3 client. Each individual part is upload and when all parts are uploaded, they are concatenated to form the final S3 object in the bucket. Sometimes uploads are aborted due to user action or network outages. Such data accumulates and incurs storage cost for customers. This can be a significant amount for heavy S3 users. 

Disclaimer: Some of the code in this repo has been sugeested by ChatGPT3 and Github's Copilot

### What can SIMU do
1. Scans buckets in AWS account in selected region for incomplete multipart uploads
2. If the size of multipart uploads is more than 0 MB and the uploads are older than 24 hours
3. Shows the size od orphaned data and the approximate cost it is incuring
4. Asks user to either skip or clean the data
5. If user selects clean, deletes incomplete chunks. Reduced AWS bill 

### Prerequisites
1. AWS credentials with appropriate permissions to perform S3 actions
### How to get it
Chose the right artifact for your CPU architecture and OS type from https://github.com/bit-cloner/simu-cleaner/releases
```
wget https://github.com/bit-cloner/simu-cleaner/releases/download/0.6/simu-0.6-linux-amd64.tar.gz
```
tar -xvzf simu-0.6-linux-amd64.tar.gz
```
### How to use it
```
chmod +x ./simu
```
```
./simu
```
<h3 align="left">Support:</h3>
<p><a href="https://www.buymeacoffee.com/welldone"> <img align="left" src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="50" width="210" alt="welldone" /></a></p><br><br>