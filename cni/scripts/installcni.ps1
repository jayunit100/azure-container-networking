Write-Host $env:CONTAINER_SANDBOX_MOUNT_POINT

$sourceCNI = $env:CONTAINER_SANDBOX_MOUNT_POINT + "\azure-container-networking\cni\network\plugin\azure-vnet.exe"
$sourceConflist = $env:CONTAINER_SANDBOX_MOUNT_POINT + "\azure-container-networking\cni\azure-windows-swift.conflist"

Rename-Item -Path "C:\k\azurecni\bin\azure-vnet.exe" -NewName "azure-vnet-old.exe"
Copy-Item $sourceCNI -Destination "C:\k\azurecni\bin"
Copy-Item $sourceConflist -Destination "C:\k\azurecni\netconf\10-azure.conflist"
Rename-Item -Path $sourceConflist -NewName "10-azure.conflist"
