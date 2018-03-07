$(function() {
/**
 * 1. Initialize
 *
 */
$( "#form-signin input[name=Username]" ).focus();

// Ajax
var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
        }
        options.headers = headers;
        return ajax(url, options);
    }
});

/**
 * 2. Events
 *
 */
// Validation
$( "#form-login" ).validate({
    submitHandler: function( form, e ) {
        e.preventDefault();

        var $form = $( form ),
            username = $( "input[name=username]", $form ).val().toLowerCase(),
            password = $( "input[name=password]", $form ).val();

        $.ajax({
            type: "get",
            async: true,
            url: "/signin/" + username + "/salt"
        }).done( function( result ) {
            if ( result.state ) {
                $.ajax({
                    type: "POST",
                    async: true,
                    url: "/signin",
                    data: {
                        username: username,
                        encPassword: getHash( getHash( username + password ) + result.data )
                    }
                }).done( function( result2 ) {
                    if ( result2.state ) {
                        window.location.href = result2.data.redirectUrl;
                    } else {
                    //     $( ".note", $form ).text( result2.Message + " / Code: " + result2.Code ).removeClass( "hidden" );
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
        username: {
            required: true
        },
        password: {
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
