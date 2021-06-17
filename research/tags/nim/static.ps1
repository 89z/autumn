nim c `
'--dynlibOverride:crypto-' `
'--dynlibOverride:ssl-' `
'--passl:-lssl' `
'--passl:-lcrypto' `
'--passl:-lws2_32' `
'--passl:-static' `
'-d:noOpenSSLHacks' `
'-d:ssl' `
'-d:sslVersion:(' `
a.nim
