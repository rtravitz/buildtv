# Build TV
This project display different images on the build monitor based on the build's current status. Currently you can update the status with three curl requests to receive a random, relevant image.


To use, clone the project and run the binary. After cloning, boot the server and navigate to the root. Curl requests currently change the status. Posting success or fail will cycle through a series of random images. Neutral is always Chillest Monkey.

### Curl requests
* **success**: curl -X POST -F  "status=fail" http://localhost:1234/status
* **fail**: curl -X POST -F  "status=fail" http://localhost:1234/status
* **neutral**: curl -X POST -F  "status=neutral" http://localhost:1234/status