package monredt

var tableHTML = `
<div class="w3-container">
    <div>
        <p>Block: {{.BlockNumber}} ({{.Elapsed}} sec) {{.Timestamp}}</br>
        Gas: {{.GasUsed}} of {{.GasLimit}}</br>
        Txs: {{.BlockNumTxs}} Total: {{.AllNumTxs}} </p>
    </div>
    <div class="w3-responsive w3-card-4">
        <table class="w3-table w3-striped w3-bordered">
            <thead>
                <tr class="w3-theme">
                    <th>Author</th>
                    <th>Signer</th>
                    <th>Name</th>
                </tr>
            </thead>
            <tbody>
                {{range .Signers }}
				{{if .NextProposer}}
                <tr class="w3-blue">
				{{else}}
                <tr>
				{{end}}
					{{if .AsProposer}}
						{{if .CurrentProposer}}
		                    <td><span class='w3-badge'>{{.AsProposer}}</span></td>
						{{else}}
							<td>{{.AsProposer}}</td>
						{{end}}
					{{else}}
	                    <td><span class='w3-badge w3-red'>{{.AsProposer}}</span></td>
					{{end}}

					{{if .AsSigner}}
						{{if .CurrentSigner}}
		                    <td><span class='w3-badge'>{{.AsSigner}}</span></td>
						{{else}}
							<td>{{.AsSigner}}</td>
						{{end}}
					{{else}}
	                    <td><span class='w3-badge w3-red'>{{.AsSigner}}</span></td>
					{{end}}
                    <td>{{.Name}}</td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
    <div>
        <p>Next: {{.NextProposerName}}</p>
    </div>
</div>
`
