# example 1
dotnet new console -o May

# example 2
dotnet publish `
-p:Configuration=Release `
-p:PublishSingleFile=true `
-p:RuntimeIdentifier=win-x64 `
-p:SelfContained=false

# example 3
dotnet run
