﻿# Code generated from specification version 1.0.0: DO NOT EDIT
Function Get-Software {
<#
.SYNOPSIS
Get software

.DESCRIPTION
Get an existing software package (managedObject)

.LINK
https://reubenmiller.github.io/go-c8y-cli/docs/cli/c8y/software_get

.EXAMPLE
PS> Get-Software -Id $mo.id

Get a software package

.EXAMPLE
PS> Get-ManagedObject -Id $mo.id | Get-Software

Get a software package (using pipeline)


#>
    [cmdletbinding(PositionalBinding=$true,
                   HelpUri='')]
    [Alias()]
    [OutputType([object])]
    Param(
        # Software package (managedObject) id (required)
        [Parameter(Mandatory = $true,
                   ValueFromPipeline=$true,
                   ValueFromPipelineByPropertyName=$true)]
        [object[]]
        $Id
    )
    DynamicParam {
        Get-ClientCommonParameters -Type "Get"
    }

    Begin {

        if ($env:C8Y_DISABLE_INHERITANCE -ne $true) {
            # Inherit preference variables
            Use-CallerPreference -Cmdlet $PSCmdlet -SessionState $ExecutionContext.SessionState
        }

        $c8yargs = New-ClientArgument -Parameters $PSBoundParameters -Command "software get"
        $ClientOptions = Get-ClientOutputOption $PSBoundParameters
        $TypeOptions = @{
            Type = "application/vnd.com.nsn.cumulocity.inventory+json"
            ItemType = ""
            BoundParameters = $PSBoundParameters
        }
    }

    Process {

        if ($ClientOptions.ConvertToPS) {
            $Id `
            | Group-ClientRequests `
            | c8y software get $c8yargs `
            | ConvertFrom-ClientOutput @TypeOptions
        }
        else {
            $Id `
            | Group-ClientRequests `
            | c8y software get $c8yargs
        }
        
    }

    End {}
}
