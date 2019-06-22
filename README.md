# GitPuller
this tool to pull js project and automatic build. identical to:
```
cd to\js\path\project
git pull 
#if new push there
npm run build

```

you can setting path of js source and delay time.

## How to Use

  1. install
     ```go get github.com/shouva/gitpuller```
  2. create config
      ```
      vim setting.json
      ```
      add this text:
      ```
      {
        "path":"/home/mosleim/js/apitest",
        "delay": 60
      }
   
  3. run it!
  ```
  gitpuller
  ```
