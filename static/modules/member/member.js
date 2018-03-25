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

    // 로그 테이블
    var $table      = $( "#table-member" ),
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
                // console.log(result);
                if ( result.state ) {
                    $( ".alert", $( form ) ).addClass( "hidden" );
                    $( "#modal-member-add" ).modal( "hide" );
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
            username : {
                required: true,
                username: true
            },
            password : {
                required: true,
                minlength: 8,
                maxlength: 16,
                password: true,
            },
            password_confirm : {
                equalTo: "#form-member-add input[name=password]"
            },
            email : {
                required: true,
                email: true,
            },

        },
        highlight: function( element ) {
            $( element ).closest( ".form-group" ).addClass( "has-error" );
        },
        unhighlight: function( element ) {
            $( element ).closest( ".form-group" ).removeClass( "has-error" );
        }
    });


    // 필터 유효성 체크
    $( "#form-member-edit" ).validate({
        submitHandler: function( form, e ) {
            e.preventDefault();

            $.ajax({
                type: "PATCH",
                async: true,
                url: "/members/" + $( "input[name=member_id]", $( form ) ).val(),
                data: $( form ).serialize()
            }).done( function ( result ) {
                // console.log(result);
                if ( result.state ) {
                    $( ".alert", $( form ) ).addClass( "hidden" );
                    $( "#modal-member-edit" ).modal( "hide" );
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
            username : {
                required: true,
                username: true
            },
            password : {
                minlength: 8,
                maxlength: 16,
                password: true,
            },
            password_confirm : {
                equalTo: "#form-member-edit input[name=password]"
            },
            email : {
                required: true,
                email: true,
            },

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
        console.log(tableKey);
        captureTableColumns( $( this ), tableKey );
    });

    $( "#modal-member-add" ).on( "shown.bs.modal", function () {
        $( "#form-member-add input[name=username]" ).focus();
        fillTestData();
    });

    $( "#modal-member-edit" ).on( "shown.bs.modal", function () {
        $( "#form-member-edit input[name=name]" ).focus().select();
        fillTestData();
    });

    $( ".modal-member" )
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

    window.memberActionEvents = {
        "click .edit": function( e, val, row, idx ) {
            showForm( row, "edit" );
        },
        "click .acl": function( e, val, row, idx ) {
            showForm( row, "acl" );
        },
        "click .remove": function( e, val, row, idx ) {
            swal({
                title: felang[ "msg.confirm_delete" ],
                type: "warning",
                showCancelButton: true,
                confirmButtonClass: "btn red",
                confirmButtonText: felang[ "yes" ],
                cancelButtonText: felang[ "no" ]
            }).then((result) => {
                if (result.value) {
                    $.ajax({
                        type: "DELETE",
                        async: true,
                        url: "/members/" + row.member_id,
                    }).done( function( result ) {
                        if ( result.state ) {
                            $( "#table-member" ).bootstrapTable( "refresh" );
                        } else {
                            swal("fail", result.message, "error");
                        }
                    });
                }
            })

        }
    };


    /**
     * 3. 함수
     *
     */
    function showForm(row, mode) {
        if (mode == "edit") { // 사용자 정보 수정
            $.ajax({
                type: "GET",
                async: true,
                url: "/members/" + row.member_id,
            }).done(function(result) {
                if ( ! result.state ) {
                    return;
                }

                var m = result.data;
                var $form = $( "#form-member-edit" );

                $( "input[name=member_id]", $form ).val( m.member_id );
                $( "input[name=username]", $form ).val( result.data.username );
                $( "input[name=name]", $form ).val( result.data.name );
                $( "input[name=email]", $form ).val( result.data.email );
                if (result.data.allowed_ip != null) {
                    $( "textarea[name=allowed_ip]", $form ).val(m.allowed_ip.replace(/\/32/g, "" ).replace(/,/g, "\n" ));
                }

                // 권한 설정
                $form.find( "input[name=user_groups]" ).each(function(i) {
                    if ( m.position & ( 1 << $( this ).val() ) ) {
                        $( this ).prop( "checked", true );
                    }
                });

                // 타임존
                $( "select[name=timezone]", $form ).val( m.timezone ).selectpicker( "refresh" );
                $( "#modal-member-edit" ).modal( "show" );
            });

        } else if (mode == "acl") { // 사용자 권한 설정
            $.ajax({
                type: "GET",
                async: true,
                url: "/members/" + row.member_id + "/acl",
            }).done(function(result) {
                $( "#modal-member-acl" ).modal( "show" );
            });
        }
    }


});

