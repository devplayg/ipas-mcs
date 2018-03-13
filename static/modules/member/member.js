$(function() {

    function fillTestData() {
        $( "#form-member-add input[name=username]" ).val("wsahn");
        $( "#form-member-add input[name=password]" ).val("sniper123!@#");
        $( "#form-member-add input[name=password_confirm]" ).val("sniper123!@#");
        $( "#form-member-add input[name=name]" ).val("안원석");
        $( "#form-member-add input[name=email]" ).val("wsahn@wins21.co.kr");
    }

    /**
     * 1. 초기화
     *
     */

    //     // 로그 테이블
    var $table      = $( "#table-log" ),
        tableKey    = getTableKey( $table, reqVars.ctrl, reqVars.act ); // 테이블 고유키

    // 테이블 컬럼속성 복원
    restoreTableColumns( $table, tableKey );


    // 필터 유효성 체크
    $( "#form-member-add" ).validate({
        submitHandler: function( form, e ) {
            e.preventDefault();

            $.ajax({
                type: "POST",
                async: true,
                url: "/members",
                data: $( form ).serialize()
            }).done( function ( result ) {
                console.log(result);
                if ( result.state ) {
                    $( ".alert", $( form ) ).addClass( "hidden" );
                    // $( "#modal-member-add" ).modal( "hide" );
                } else {
                    $( ".alert .message", $( form ) ).text( result.message );
                    $( ".alert", $( form ) ).removeClass( "hidden" );
                }
            }).always( function() {
            });

        },
        ignore: "input[type='hidden']",
        errorClass: "help-block",
        rules: {
            },
        },
        messages: {
            SrcPortStart: "0 ~ 65535",
            SrcPortEnd  : "0 ~ 65535",
        },
        highlight: function( element ) {
            $( element ).closest( ".form-group" ).addClass( "has-error" );
        },
        unhighlight: function( element ) {
            $( element ).closest( ".form-group" ).removeClass( "has-error" );
        }
    });


    /**
     * 2. 이벤트
     *
     */
    $table.on( "column-switch.bs.table", function( e, field, checked ) { // 테이블 컬럼 보기/숨기기 속성이 변경되는 경우
        captureTableColumns( $( this ), tableKey );
    });

    $( "#modal-member-add" ).on( "shown.bs.modal", function () {
        $( "#form-member-add input[name=username]" ).focus();
        fillTestData();
    });

    $( ".modal-member" )
        // .on( "shown.bs.modal", function () {
        //     var $form = $( this ).find( "form" );
        //     if ( $form.attr( "id" ) == "form-member-add" ) {
        //         $( "input[name=Username]", $form ).focus().select();
        //     } else if ( $form.attr( "id" ) == "form-members-edit" ) {
        //         $( "input[name=Name]", $form ).focus().select();
        //     } else if ( $form.attr( "id" ) == "form-members-password" ) {
        //         $( "input[name=NewPassword]", $form ).focus();
        //     }
        // })
        .on( "hidden.bs.modal", function () {
            // Reset form
            var $form = $( this ).find( "form" );
            $form.validate().resetForm();
            $form.get( 0 ).reset();
            $( ".alert", $form ).addClass( "hidden" );
            $( ".alert .message", $form ).empty(    );

            // Refresh the member table
            $( "#table-member" ).bootstrapTable( "refresh" );
        });




    /**
     * 3. 함수
     *
     */


});

