$(function() {

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    var assets = { },
        interval = 3000;
    initializeAssets();
    updateStats();




    setInterval(function() {
        // updateStats();
    }, interval);




    /**
     * 2. 이벤트
     *
     */

    $( "#select-assets" ).on( "change", function() {
        var selected = $( "#select-assets :selected" ).val();
        // arr = val.split('/');
        // console.log(selected);
        // asset_idx = asset.indexOf(val);
        // client_last_update = '';
        // changeAsset(asset[asset_idx]);
    });


    /**
     * 3. 함수
     *
     */

    // 자산 초기화
    function initializeAssets() {
        $.ajax({
            type  : "GET",
            async : true,
            url   : "/userassetclass/1/children"
        }).done( function( result ) {
            // 기관
            $.each( result, function( idx, org ) {
                $( "#select-assets" ).append(
                    $( "<option>", {
                        value: org.asset_id + "/-1",
                        text: org.name
                    })
                );

                // 그룹
                $.each( org.children, function( i, group ) {
                    $( "#select-assets" ).append(
                        $( "<option>", {
                            value: org.asset_id + "/" + group.asset_id,
                            text: "- " + group.name,
                            class: "ml20"
                        })
                    );
                });
            });


        }).always( function() {
            $( "#select-assets" ).selectpicker( "refresh" );
        });
    }


    function updateStats() {
        updateRankings();
    }


    function updateRankings() {
        $( ".table-ranking" ).each(function( idx, obj ) {
            var param = {
                assetKey: $( "#select-assets :selected" ).val()
            };
            var url = $( this ).data( "query" ) + "?" + $.param( param );
            $( this ).bootstrapTable( "refresh", { url: url, silent: true } );
        });
    }

});

