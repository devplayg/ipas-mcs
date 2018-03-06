$(function() {
/**
 * 1. Initialize
 *
 */
$( "#form-signin input[name=Username]" ).focus();


/**
 * 2. Events
 *
 */
// Validation
$( "#form-signin" ).validate({
    submitHandler: function( form ) {
        var $form = $( form );
        $.ajax({
            type: "get",
            async: true,
            url: "/signin/" + $( "input[name=Username]", $form ).val() + "/salt"
        }).done( function( result ) {

            if ( result.State == true ) {
                var username    = $( "input[name=Username]", $form ).val().toLowerCase(),
                    password    = $( "input[name=Password]", $form ).val();
                $( "input[name=EncPassword]", $form ).val( getHash( getHash( username + password ) + result.Code ) );
                $( "input[name=Salt]", $form ).val(result.Code);
                $.ajax({
                    type: "post",
                    async: true,
                    url: "/signin",
                    data: $form.serialize()
                }).done( function( result2 ) {
                    if ( result2.State ) {
                        window.location.href = result2.Code;
                    } else {
                        $( ".note", $form ).text( result2.Message + " / Code: " + result2.Code ).removeClass( "hidden" );
                    }

                }).always(function() {

                });
            } else {
                $( ".note", $form ).text( result.Message + " / Code: " + result.Code ).removeClass( "hidden" );
            }
        });
    },
    ignore: "input[type='hidden']",
    errorClass: "help-block",
    rules: {
        Username: {
            required: true
        },
        Password: {
            required: true
        }
    },
    messages: {
    },
    highlight: function(element) {
        $(element).closest( ".form-group ").addClass( "has-error" );
    },
    unhighlight: function(element) {
        $(element).closest( ".form-group" ).removeClass( "has-error" );
    }
});

/**
 * 3. Functions
 *
 */
function getHash(str) {
    return CryptoJS.SHA256(str).toString();
}



});
