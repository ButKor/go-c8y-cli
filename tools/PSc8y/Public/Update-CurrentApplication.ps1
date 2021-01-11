﻿# Code generated from specification version 1.0.0: DO NOT EDIT
Function Update-CurrentApplication {
<#
.SYNOPSIS
Update current application

.DESCRIPTION
Required authentication with bootstrap user

.EXAMPLE
PS> Update-CurrentApplication -Data @{ mycustomProp = @{ value1 = 1}}

Update custom properties of the current application (requires using application credentials)


#>
    [cmdletbinding(SupportsShouldProcess = $true,
                   PositionalBinding=$true,
                   HelpUri='',
                   ConfirmImpact = 'High')]
    [Alias()]
    [OutputType([object])]
    Param(
        # data
        [Parameter()]
        [object]
        $Data,

        # Name of application
        [Parameter()]
        [string]
        $Name,

        # Shared secret of application
        [Parameter()]
        [string]
        $Key,

        # Application will be applied to this type of documents, possible values [ALARM, AUDIT, EVENT, MEASUREMENT, OPERATION, *].
        [Parameter()]
        [ValidateSet('MARKET','PRIVATE')]
        [string]
        $Availability,

        # contextPath of the hosted application
        [Parameter()]
        [string]
        $ContextPath,

        # URL to application base directory hosted on an external server
        [Parameter()]
        [string]
        $ResourcesUrl,

        # authorization username to access resourcesUrl
        [Parameter()]
        [string]
        $ResourcesUsername,

        # authorization password to access resourcesUrl
        [Parameter()]
        [string]
        $ResourcesPassword,

        # URL to the external application
        [Parameter()]
        [string]
        $ExternalUrl,

        # Cumulocity processing mode
        [Parameter()]
        [AllowNull()]
        [AllowEmptyString()]
        [ValidateSet("PERSISTENT", "QUIESCENT", "TRANSIENT", "CEP", "")]
        [string]
        $ProcessingMode,

        # Template (jsonnet) file to use to create the request body.
        [Parameter()]
        [string]
        $Template,

        # Variables to be used when evaluating the Template. Accepts a file path, json or json shorthand, i.e. "name=peter"
        [Parameter()]
        [string]
        $TemplateVars,

        # Show the full (raw) response from Cumulocity including pagination information
        [Parameter()]
        [switch]
        $Raw,

        # Write the response to file
        [Parameter()]
        [string]
        $OutputFile,

        # Ignore any proxy settings when running the cmdlet
        [Parameter()]
        [switch]
        $NoProxy,

        # Specifiy alternative Cumulocity session to use when running the cmdlet
        [Parameter()]
        [string]
        $Session,

        # TimeoutSec timeout in seconds before a request will be aborted
        [Parameter()]
        [double]
        $TimeoutSec,

        # Don't prompt for confirmation
        [Parameter()]
        [switch]
        $Force
    )

    Begin {
        $Parameters = @{}
        if ($PSBoundParameters.ContainsKey("Data")) {
            $Parameters["data"] = ConvertTo-JsonArgument $Data
        }
        if ($PSBoundParameters.ContainsKey("Name")) {
            $Parameters["name"] = $Name
        }
        if ($PSBoundParameters.ContainsKey("Key")) {
            $Parameters["key"] = $Key
        }
        if ($PSBoundParameters.ContainsKey("Availability")) {
            $Parameters["availability"] = $Availability
        }
        if ($PSBoundParameters.ContainsKey("ContextPath")) {
            $Parameters["contextPath"] = $ContextPath
        }
        if ($PSBoundParameters.ContainsKey("ResourcesUrl")) {
            $Parameters["resourcesUrl"] = $ResourcesUrl
        }
        if ($PSBoundParameters.ContainsKey("ResourcesUsername")) {
            $Parameters["resourcesUsername"] = $ResourcesUsername
        }
        if ($PSBoundParameters.ContainsKey("ResourcesPassword")) {
            $Parameters["resourcesPassword"] = $ResourcesPassword
        }
        if ($PSBoundParameters.ContainsKey("ExternalUrl")) {
            $Parameters["externalUrl"] = $ExternalUrl
        }
        if ($PSBoundParameters.ContainsKey("ProcessingMode")) {
            $Parameters["processingMode"] = $ProcessingMode
        }
        if ($PSBoundParameters.ContainsKey("Template") -and $Template) {
            $Parameters["template"] = $Template
        }
        if ($PSBoundParameters.ContainsKey("TemplateVars") -and $TemplateVars) {
            $Parameters["templateVars"] = $TemplateVars
        }
        if ($PSBoundParameters.ContainsKey("OutputFile")) {
            $Parameters["outputFile"] = $OutputFile
        }
        if ($PSBoundParameters.ContainsKey("NoProxy")) {
            $Parameters["noProxy"] = $NoProxy
        }
        if ($PSBoundParameters.ContainsKey("Session")) {
            $Parameters["session"] = $Session
        }
        if ($PSBoundParameters.ContainsKey("TimeoutSec")) {
            $Parameters["timeout"] = $TimeoutSec * 1000
        }

        if ($env:C8Y_DISABLE_INHERITANCE -ne $true) {
            # Inherit preference automatic variables
            if (!$WhatIfPreference.IsPresent) {
                $WhatIfPreference = $PSCmdlet.SessionState.PSVariable.get("WhatIfPreference").Value
            }
        
            # Inherit custom parameters
            $Stack = Get-PSCallStack | Select-Object -Skip 1 -First 1
            $InheritVariables = @(@{Source="Force"; Target="Force"}, @{Source="WhatIf"; Target="WhatIfPreference"})
            foreach ($iVariable in $InheritVariables) {
                if (-Not $PSBoundParameters.ContainsKey($iVariable.Source)) {
                    if ($null -ne $Stack -and $Stack.InvocationInfo.BoundParameters.ContainsKey($iVariable.Source)) {
                        Set-Variable -Name $iVariable.Target -Value $Stack.InvocationInfo.BoundParameters[$iVariable.Source] -WhatIf:$false
                    }
                }
            }
        }
    }

    Process {
        foreach ($item in @("")) {

            if (!$Force -and
                !$WhatIfPreference -and
                !$PSCmdlet.ShouldProcess(
                    (PSc8y\Get-C8ySessionProperty -Name "tenant"),
                    (Format-ConfirmationMessage -Name $PSCmdlet.MyInvocation.InvocationName -InputObject $item)
                )) {
                continue
            }

            Invoke-ClientCommand `
                -Noun "currentApplication" `
                -Verb "update" `
                -Parameters $Parameters `
                -Type "application/vnd.com.nsn.cumulocity.application+json" `
                -ItemType "" `
                -ResultProperty "" `
                -Raw:$Raw
        }
    }

    End {}
}
