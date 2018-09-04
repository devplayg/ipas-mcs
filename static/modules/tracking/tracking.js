$(function() {
    // $('.ui.search')
    //     //     .search({
    //     //         type: 'category',
    //     //         source: categoryContent
    //     //     })
    //     // ;
    //

    // var options = {
    //
    //     url: "/static/plugins/EasyAutocomplete/ex.json",
    //
    //     getValue: "name"
    // };
    //
    // $("#countries").easyAutocomplete(options);
    /**
     * 1. 초기화
     *
     */

    updateLiveSearch();


    /**
     * 2. 이벤트
     *
     */

    // 날짜
    $( ".datetime" ).datetimepicker({
        format:         "yyyy-mm-dd hh:ii",
        pickerPosition: "bottom-left",
        todayHighlight: 1,
        minView:        2,
        maxView:        4,
        autoclose:      true
    });

    // 필터 유효성 체크
    $( "#form-filter" ).validate({
        submitHandler: function( form, e ) {
            // console.log(3);
            // console.log($(form).serializeArray());
            e.preventDefault();
            //
            // if ( ! $( "input[name=fast_paging]", form ).is( ":checked" ) ) {
            //     $( form ).addHidden( "fast_paging", "off" );
            // }
            // $( form ).addHidden( "sort", $table.bootstrapTable("getOptions").sortName );
            // $( form ).addHidden( "order", $table.bootstrapTable("getOptions").sortOrder );
            //
            form.submit();
        },
        ignore: "input[type='hidden']",
        errorClass: "help-block",
        rules: {
            guid: {
                minlength: 2,
                maxlength: 5
            },
        },
        messages: {
            SrcPortStart: "0 ~ 65535",
            SrcPortEnd:   "0 ~ 65535",
        },
        highlight: function( element ) {
            $( element ).closest( ".form-group" ).addClass( "has-error" );
        },
        unhighlight: function( element ) {
            $( element ).closest( ".form-group" ).removeClass( "has-error" );
        }
    });

    // 장비가 선택됐을 경우
    $( "#select-equipId" ).change(function() {
        var a = $( "#select-equipId option:selected" ).val();
        console.log( a );
    });


    /**
     * 3. 함수
     *
     */
    // 장비목록 출력
    function updateLiveSearch() {

        $( "#select-equipId" ).empty().append( $( "<option>", {
            value: 0,
            html: "Tag"
        }));;

        $.ajax({
            type:  "GET",
            async: true,
            url:   "/ipaslist"
        }).done( function( result ) {
            $.each( result, function( i, r ) {
                $( "#select-equipId" ).append( $( "<option>", {
                    value: r.org_id + "/" + r.equip_id,
                    html: r.org_name + " / " + r.equip_id
                }));
            });
        }).always( function() {
            $( "#select-equipId" ).selectpicker( "val", 0 ).selectpicker( "refresh" );
        });

        // $( "#search-equipId" ).append($("<option>", { value: 1, html: "abc" }));
    }

});