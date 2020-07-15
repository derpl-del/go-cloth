$(document).ready(function () {
    $.get("/val/GetUserSession", function (data) {
        var response = jQuery.parseJSON(data);
        val = response.ErrorMessage;
        role = response.ErrorRole;
        user = response.User;
        if (val == true) {
            $("#login-logout").text("logout");

        } else {
            $("#login-logout").text("login");

        }
        if (role == true) {
            $("#role-btn").text("Shopkeeper");
        } else {
            window.location.replace("/");
        }
    });

    $("#login-logout").click(function () {
        if (val == true) {
            $.get("/val/RemoveUserSession", function (data) {
                document.location.reload();
            });
        } else {
            window.location.replace("/login");
        }
    });
});