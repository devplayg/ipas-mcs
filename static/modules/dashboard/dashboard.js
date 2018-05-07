$(function() {

    /**
     * 1. 초기화
     *
     */
    // 자산 (기관 / 그룹)
    var assets = { },
        interval = 2000,
        timer = null;
    initializeAssets();
    updateStats();
    startTimer();







    /**
     * 2. 이벤트
     *
     */

    $( "#select-assets" ).on( "change", function() {
        stopTimer();
        updateStats();
        startTimer();
    });

    $( ".btn-start" ).click(function(e) {
        if ( timer === null ) {
            startTimer();
        } else {
            stopTimer();
        }
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
                // console.log(org.asset_id + "/-1");

                // 그룹
                $.each( org.children, function( i, group ) {
                    $( "#select-assets" ).append(
                        $( "<option>", {
                            value: org.asset_id + "/" + group.asset_id,
                            text: "- " + group.name,
                            class: "ml20"
                        })
                    );
                    // console.log(org.asset_id + "/-1");
                });
            });


        }).always( function() {
            $( "#select-assets" ).selectpicker( "refresh" );
        });
    }

    function startTimer() {  // use a one-off timer
        $( ".btn-start" ).removeClass( "default" ).addClass( "blue" );
        $( ".btn-start .text" ).html( "<i class='fa fa-circle-o-notch fa-spin'></i>" );
        timer = setTimeout(updateStats, interval);
    }

    function stopTimer() {
        $( ".btn-start" ).removeClass( "blue" ).addClass( "default" );
        $( ".btn-start .text" ).html( "<i class='fa fa-play'></i>" );
        clearTimeout( timer );
        timer = null;
    }


    function updateStats() {
        updateRankings();
    }


    function updateRankings() {
        $( ".table-ranking" ).each(function( idx, obj ) {
            var asset = $( "#select-assets :selected" ).val().split( "/", 2 ),
                url = $( this ).data( "query" ) + "/org/" + asset[0] + "/group/" + asset[1];
            // console.log(url);
            $( this ).bootstrapTable( "refresh", { url: url, silent: true } );
        });
    }

});

