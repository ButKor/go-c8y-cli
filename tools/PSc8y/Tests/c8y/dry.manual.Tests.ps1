. $PSScriptRoot/../imports.ps1

Describe -Name "c8y dry" {

    Context "multi-form data" {
        It "Displays multi form data" {

            $output = c8y binaries create --file=./Makefile --dry
            $LASTEXITCODE | Should -Be 0
            $output | Should -Not -BeNullOrEmpty
        }
    }

    Context "POST" {
        It "Shows the body contents" {

            $output = c8y devices create --name test01 --dry
            $LASTEXITCODE | Should -Be 0
            $output | Should -Not -BeNullOrEmpty
        }

        It "Shows the body contents with a custom body" {

            $output = c8y devices create --name test01 --data "test=1" --dry
            $LASTEXITCODE | Should -Be 0
            $output | Should -Not -BeNullOrEmpty
        }

        It "Hides sensitive information" {
            $env:C8Y_LOGGER_HIDE_SENSITIVE = "true"
            $output = c8y devices create --name test01 --data "test=1" --dry
            $LASTEXITCODE | Should -Be 0
            $output | Should -Not -BeNullOrEmpty
            $output | Should -Not -Match 

            $env:C8Y_LOGGER_HIDE_SENSITIVE = $null
        }

        
    }

    AfterEach {
    }
}
