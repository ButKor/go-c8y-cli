﻿. $PSScriptRoot/imports.ps1

Describe -Name "Get-FirmwarePatchCollection" {
    BeforeEach {

    }

    It -Skip "Get a list of firmware patches related to a firmware package" {
        $Response = PSc8y\Get-FirmwarePatchCollection -Firmware 12345
        $LASTEXITCODE | Should -Be 0
        $Response | Should -Not -BeNullOrEmpty
    }

    It -Skip "Get a list of firmware patches where the dependency version starts with "1."" {
        $Response = PSc8y\Get-FirmwarePatchCollection -Firmware 12345 -Dependency '1.*'
        $LASTEXITCODE | Should -Be 0
        $Response | Should -Not -BeNullOrEmpty
    }


    AfterEach {

    }
}

