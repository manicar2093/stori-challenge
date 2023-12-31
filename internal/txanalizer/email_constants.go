package txanalizer

const accountStatusHtmlTemplate = `<!DOCTYPE html>
<html lang="en" xmlns:v="urn:schemas-microsoft-com:vml">

<head>
  <meta charset="utf-8">
  <meta name="x-apple-disable-message-reformatting">
  <meta http-equiv="x-ua-compatible" content="ie=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="format-detection" content="telephone=no, date=no, address=no, email=no">
  <meta http-equiv="Content-Type" content="text/html charset=UTF-8">
  <meta name="color-scheme" content="light dark">
  <meta name="supported-color-schemes" content="light dark">
  <!--[if mso]>
      <noscript>
        <xml>
          <o:OfficeDocumentSettings
            xmlns:o="urn:schemas-microsoft-com:office:office"
          >
            <o:PixelsPerInch>96</o:PixelsPerInch>
          </o:OfficeDocumentSettings>
        </xml>
      </noscript>
      <style>
        td,
        th,
        div,
        p,
        a,
        h1,
        h2,
        h3,
        h4,
        h5,
        h6 {
          font-family: "Segoe UI", sans-serif;
          mso-line-height-rule: exactly;
        }
      </style>
    <![endif]-->
  <title>Stori - Estado de Cuenta</title>
  <style>
    :root {
      color-scheme: light dark;
      supported-color-schemes: light dark;
    }
  </style>
  <style>
    .hover-bg-blue-600:hover {
      background-color: #2563eb !important;
    }

    .hover-underline:hover {
      text-decoration: underline !important;
    }

    @media (max-width: 600px) {
      .sm-w-full {
        width: 100% !important;
      }

      .sm-py-32 {
        padding-top: 32px !important;
        padding-bottom: 32px !important;
      }

      .sm-px-24 {
        padding-left: 24px !important;
        padding-right: 24px !important;
      }

      .sm-leading-32 {
        line-height: 32px !important;
      }
    }

    @media (prefers-color-scheme: dark) {
      .dark-mode-bg-gray-999 {
        background-color: #1d1f22 !important;
      }

      .dark-mode-bg-gray-989 {
        background-color: #2d2d2d !important;
      }

      .dark-mode-text-gray-979 {
        color: #a9a9a9 !important;
      }

      .dark-mode-text-white {
        color: #ffffff !important;
      }
    }
  </style>
</head>

<body class="dark-mode-bg-gray-999"
  style="margin: 0; width: 100%; padding: 0; word-break: break-word; -webkit-font-smoothing: antialiased; background-color: #f3f4f6;">
  <div role="article" aria-roledescription="email" aria-label="Your receipt for order 12345" lang="en">
    <table class="sm-w-full" align="center" style="width: 600px;" cellpadding="0" cellspacing="0" role="presentation">
      <tr>
        <td class="sm-py-32 sm-px-24" style="padding: 48px; text-align: center;">
          <a target="_blank" href="https://www.storicard.com/">
            <img src="https://www.storicard.com/_next/static/media/complete-logo.0f6b7ce5.svg" width="200"
              alt="Your Logo" style="max-width: 100%; vertical-align: middle; line-height: 100%; border: 0;">
          </a>
        </td>
      </tr>
    </table>
    <table style="width: 100%; font-family: ui-sans-serif, system-ui, -apple-system, 'Segoe UI', sans-serif;"
      cellpadding="0" cellspacing="0" role="presentation">
      <tr>
        <td align="center" class="dark-mode-bg-gray-999" style="background-color: #f3f4f6;">
          <table class="sm-w-full" style="width: 600px;" cellpadding="0" cellspacing="0" role="presentation">
            <tr>
              <td align="center" class="sm-px-24">
                <table style="margin-bottom: 48px; width: 100%;" cellpadding="0" cellspacing="0" role="presentation">
                  <tr>
                    <td class="dark-mode-bg-gray-989 dark-mode-text-gray-979 sm-px-24"
                      style="background-color: #ffffff; padding: 48px; text-align: left; font-size: 16px; line-height: 24px; color: #1f2937;">
                      <p class="sm-leading-32 dark-mode-text-white"
                        style="margin: 0; margin-bottom: 36px; font-family: ui-serif, Georgia, Cambria, 'Times New Roman', Times, serif; font-size: 24px; font-weight: 600; color: #000000;">
                        Los detalles de tu cuenta
                      </p>
                      <p style="margin: 0; margin-bottom: 24px;">
                        ¡Este correo contiene la información de tu cuenta!
                      </p>
                      <table style="margin-bottom: 32px; width: 100%;" cellpadding="0" cellspacing="0"
                        role="presentation">
                        <thead>
                          <th>Mes</th>
                          <th># de Transacciones</th>
                        </thead>
                        <tbody>
                          {{range $key, $value := .TransactionByMonth}}
                            <tr>
                              <td>{{getMonthName $key}}</td>
                              <td>{{$value}}</td>
                            </tr>
                          {{end}}
                        </tbody>
                      </table>

                      <table style="margin-bottom: 32px; width: 100%;" cellpadding="0" cellspacing="0"
                        role="presentation">
                        <thead>
                          <th>Monto promedio de crédito</th>
                          <th>Monto promedio de débito</th>
                          <th>Balance total</th>
                        </thead>
                        <tbody>
                          <tr>
                            <td>${{formatAmount .AverageCreditAmount}}</td>
                            <td>${{formatAmount .AverageDebitAmount}}</td>
                            <td>${{formatAmount .TotalBalance}}</td>
                          </tr>
                        </tbody>
                      </table>
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
    <table style="width: 100%; font-family: ui-sans-serif, system-ui, -apple-system, 'Segoe UI', sans-serif;"
      cellpadding="0" cellspacing="0" role="presentation">
      <tr>
        <td style="padding-left: 24px; padding-right: 24px; text-align: center; font-size: 12px; color: #4b5563;">
        </td>
      </tr>
    </table>
  </div>
</body>

</html>`
