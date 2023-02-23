### What is SIMU-cleaner &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ![alt text](SIMU-cleaner.gif "SIMU")
SIMU stands for S3 Incomplete Multipart Uplods Cleaner. It is an intreactive CLI tool to clean AWS accounts that contain incomplete chunks of S3 data. 

### Problem
When large files are uploaded to S3 , they are split into smaller parts by the S3 client. Each individual part is upload and when all parts are uploaded, they are concatenated to form the final S3 object in the bucket. Sometimes uploads are aborted due to user action or network outages. Such data accumulates and incurs storage cost for customers. This can be a significant amount for heavy S3 users. 

### What can SIMU do
1. 
<h3 align="left">Support:</h3>
<p><a href="https://www.buymeacoffee.com/welldone"> <img align="left" src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="50" width="210" alt="welldone" /></a></p><br><br>


Disclaimer: Some of the code in this repo has been sugeested by ChatGPT3 and Github's Copilot