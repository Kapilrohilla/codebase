<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Form + PHP + MySql</title>
</head>

<body>
    <h1>Form + PHP + MySql</h1>
    <button onclick="displayForm()">Show Form</button> <button onclick="displayData()">Show Previously Submitted
        Data</button>

    <div class="form" style="display: none">
        <h2>Form</h2>
        <form action="connect.php" method="POST" onsubmit="(e)=>e.preventDefault();">
            <label for="name">name: </label>
            <input type="text" name="name" id="name"> <br /><br />
            <label for="email">email: </label>
            <input type="email" name="email" id="email"> <br /><br />
            <label for="phone">mobile: </label>
            <input type="text" name="phone" id="phone"> <br /><br />
            <label for="pswd">password: </label>
            <input type="password" name="pswd" id="pswd"> <br /><br />

            <input type="submit" value="submit" />
        </form>
    </div>

    <div class="data" style="display: none">
        <?php
        $name = $email = $phone = $pswd = "";
        if ($_SERVER['REQUEST_METHOD'] === "POST" and $_POST['name'] !== "") {
            $name = $_POST['name'];
            $email = $_POST['email'];
            $phone = $_POST['phone'];
            $pswd = $_POST['pswd'];
        }
        ?>
        <h2>Data: </h2>
        <?php
        if ($name !== "") {
            echo "<p>Name: $name</p>
            <p>Email:$email </p>
            <p>Mobile: $phone</p>
            <p>Password: $pswd</p>";
        } else {
            echo "Firsly, submit some data <br/>";
        }
        ?>
    </div>
    <script>
        function displayForm() {
            const form = document.getElementsByClassName('form')[0];
            const btn = document.getElementsByTagName('button')[0];

            if (btn.innerHTML === "Show Form") {
                form.setAttribute('style', 'display: unset');
                btn.innerHTML = "Hide Form";
            } else {
                form.setAttribute('style', 'display: none ');
                btn.innerHTML = "Show Form";
            }
        }
        function displayData() {
            const form = document.getElementsByClassName('data')[0];
            const btn = document.getElementsByTagName('button')[1];

            if (btn.innerHTML === "Show Previously Submitted Data") {
                form.setAttribute('style', 'display: unset');
                btn.innerHTML = "Hide Previously Submitted Data";
            } else {
                form.setAttribute('style', 'display: none ');
                btn.innerHTML = "Show Previously Submitted Data";
            }
        }
    </script>
</body>

</html>