var mapIpas = document.getElementById( "map-ipas" );

$( "#modal-ipas-report" )
    .on( "shown.bs.modal", function (e) {
        var encoded = $( e.relatedTarget ).data( "encoded" ),
            log = JSON.parse( decodeURI( encoded ) );

        showReport( log.org_id, log.equip_id, log    );
        // console.log(333);
        // $( ".ipasreport-equip-id" ).text( log.equip_id );
        // console.log( log );
            // encData =link.data( "encoded" )
        // log = JSON.parse( decodeURI( encData ) );
        // console.log(encoded);
    })
    .on( "hidden.bs.modal", function (e) {
        clearIpasReport();
    });

http://127.0.0.1

function showReport( orgId, equipId, log ) {
    $.ajax({
        type:  "GET",
        async: true,
        url:  "/evtreport/" + equipId + "/org/" + orgId + "/since/20"
    }).done( function( rpt ) {
        console.log(rpt);
    });
}


$( "#modal-ipas-map" ).on( "shown.bs.modal", function (e) {
    var latitude = $( e.relatedTarget ).data( "latitude" ),
        longitude = $( e.relatedTarget ).data( "longitude" ),
        $modal = $( this );

    $( ".latitude", $modal ).text( latitude );
    $( ".longitude", $modal ).text( longitude );
    console.log(latitude + "/" + longitude);
    var options = {
        center: new daum.maps.LatLng( latitude, longitude ),
        level: 3
    };
    var map = new daum.maps.Map( mapIpas, options );
    // 일반 지도와 스카이뷰로 지도 타입을 전환할 수 있는 지도타입 컨트롤을 생성합니다
    var mapTypeControl = new daum.maps.MapTypeControl();
    // 지도 타입 컨트롤을 지도에 표시합니다
    map.addControl(mapTypeControl, daum.maps.ControlPosition.TOPRIGHT);
    var markerPosition  = new daum.maps.LatLng(latitude, longitude);
    var marker = new daum.maps.Marker({
        position: markerPosition
    });
    marker.setMap(map);
});

function clearIpasReport() {
    $( ".ipasreport-data" ).empty();
    $( ".ipasreport-log" ).addClass( "hide" );
}