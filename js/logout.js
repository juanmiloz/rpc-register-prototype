$(document).ready(function() {

    $("#logout_btn").click(function() {
        $.ajax({
            url: "/logout",
            method: "POST", 
            beforeSend:function() {
                $("#logout_btn").attr('disabled', true);
                $("#logout_btn").val('Conectando...');
            },
            success:function(response) {
                $('#logout_btn').removeAttr('disabled');
                $("#logout_btn").val('Logout');
                if(response == "1") {
                    $("#logoutModal").modal("toggle");
                    redirectLogin()
                } else {
                    window.location.replace("/");
                }
            }
        });
    });

    async function redirectLogin() {
        await sleep(5000);
        window.location.replace("/");
    }

    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

});