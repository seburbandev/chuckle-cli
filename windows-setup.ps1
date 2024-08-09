# Get the current directory
$currentDir = (Get-Location).Path
$exeName = "chuckle.exe"

# Build the Go application
go build -o $exeName main.go

# Add the current directory to the user PATH
$envPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
if (-Not ($envPath -contains $currentDir)) {
    [System.Environment]::SetEnvironmentVariable("Path", "$envPath;$currentDir", "User")
    Write-Host "Added $currentDir to user PATH" -ForegroundColor Green
} else {
    Write-Host "$currentDir is already in user PATH" -ForegroundColor Green
}

Write-Host "Setup completed. You can now run 'chuckle' from the command line." -ForegroundColor Green