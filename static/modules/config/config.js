$(function() {

    /**
     * 1. Initialize
     *
     */


    /**
     * 2. Events
     *
     */

    // Validation
    $( "#form-config" ).validate({
        submitHandler: function( form ) {
            $.ajax({
                type: "patch",
                async: true,
                url: "/config",
                data: $( form ).serialize()
            }).done( function( result ) {
                if ( result.state ) {
                    swal( result.message, "", "success");
                } else {
                    swal( result.message, "", "warning");
                }
            }).always(function() {
            });
        },
        ignore: "input[type='hidden']",
        errorClass: "help-block",
        rules: {
            data_retention_days: {
                required: true,
                number: true,
                min: 60,
                max: 9999

            },
            max_failed_login_attempts: {
                required: true,
                number: true,
                min: 0,
                max: 999
            }
        },
        highlight: function(element) {
            $( element ).closest( ".form-group ").addClass( "has-error" );
        },
        unhighlight: function(element) {
            $( element ).closest( ".form-group" ).removeClass( "has-error" );
        }
    });




    /**
     * 3. Functions
     *
     */





});
