# Shot
Automate the process of capturing screenshots from a list of URLs.

# Dependencies
		// Install chromedp dependencies first using these commands below or you can download manually.
		$ wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
		$ sudo dpkg -i google-chrome-stable_current_amd64.deb; sudo apt-get -fy install


Usage Example
		
		$ go run main.go urls.txt

		2023/07/31 10:54:46 Screenshot saved for URL: http://www.example.com/
		2023/07/31 10:54:46 Screenshot saved for URL: http://example.com/
		2023/07/31 10:54:48 Screenshot saved for URL: https://example.com/
		.....

		$ tree .

		├── aHR0cDovL21lQGV4YW1wbGUuY29tLw.jpg
		├── aHR0cDovL2V4YW1wbGUuY29t.jpg
		├── aHR0cDovL2V4YW1wbGUuY29tLw.jpg
		├── aHR0cDovL3d3dy5leGFtcGxlLmNvbQ.jpg
		├── aHR0cDovL3d3dy5leGFtcGxlLmNvbS8.jpg
		.....
		
		Or Compile package and access using "shot" 
		$ go install
		$ shot urls.txt


	### To view the result images you can simply use sublime text or any  other image viewer applications.
		$ subl . 

	### To get the url you can simply decode the filename to base64
		$ echo "aHR0cDovL21lQGV4YW1wbGUuY29tLw==" |base64 --decode
		http://me@example.com/

		Note: Remember to append "==" sign while decoding. 


	### To remove or move all the generated images as once you can use bash command
		$ rm *.jpg
		$ mv *.jpg




# Feel Free To Contribute
