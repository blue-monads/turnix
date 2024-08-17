

export const samplePreview = `

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <script src="//unpkg.com/alpinejs" defer></script>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
    <script>
        let pendingMessages = {}
        let messageId = 0;
        window.addEventListener("message", (ev) => {
            console.log("message", ev);
            if (ev.data !== "transfer_port") {
                console.log("wrong message", ev);
                return;
            }
            const port = ev.ports[0];
            port.onmessage = handleMessage;
            window["__parent_port__"] = port;
        }, false);
        const handleMessage = (ev) => {
            const data = ev.data;
            const messageId = data.msgId;
            const resolve = pendingMessages[messageId];
            delete pendingMessages[messageId];
            resolve(data.data);
            console.log("data", data);
        }
        const sendMessage = async (data) => {
            console.log("sendMessage", data);
            const msgId = messageId + 1;
            messageId = msgId;
            const p = new Promise((resolve) => {
                pendingMessages[msgId] = resolve;
            });
            data["msgId"] = msgId;
            const port = window["__parent_port__"]
            port?.postMessage(data);
            return p;
        }
    </script>
</head>
<body>
    <div class="container">
        <h1>Reports Editor Example</h1> 
        <p>This is a sample preview</p>
        
    </div>
    <div class="container-xl mt-5" x-data="accountTable()" id="ac_table">
        
    <div class="container">
        <button id="start" class="btn btn-primary" type="button" x-on:click="fetchAccounts">
            <span class="icon icon-thumbs-up"></span>
            Start
        </button>    
    </div>


        <div x-show="!loading" class="table-responsive">
            <table class="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Type</th>
                        <th>Info</th>
                        <th>Created At</th>
                        <th>Updated At</th>
                    </tr>
                </thead>
                <tbody>
                    <template x-for="account in accounts" :key="account.id">
                        <tr>
                            <td x-text="account.id"></td>
                            <td x-text="account.name"></td>
                            <td x-text="account.acc_type" class="text-capitalize"></td>
                            <td x-text="account.info"></td>
                            <td x-text="formatDate(account.created_at)"></td>
                            <td x-text="formatDate(account.updated_at)"></td>
                        </tr>
                    </template>
                </tbody>
            </table>
        </div>
    </div>

    <script>
    function accountTable() {
      return {
          accounts: [],
          loading: true,
          async fetchAccounts() {
              this.loading = true;
              try {
                  const data = await sendMessage({
                      type: "sql_query",
                      name: "list_accounts",
                      data: [],
                  });
                  console.log("accounts", data);

                  this.accounts = data.data;
              } catch (error) {
                  console.error("Error fetching accounts:", error);
              } finally {
                  this.loading = false;
              }
          },
          formatDate(dateString) {
              return new Date(dateString).toLocaleDateString();
          }
      }
  }


    </script>
</body>
</html>







`;
