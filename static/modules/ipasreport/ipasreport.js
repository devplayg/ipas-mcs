var mapIpas = document.getElementById( "map-ipas" ),
    ipasReportMap = document.getElementById( "map-rpt-ipas" ),
    ipasReportPeriod = 0,
    ipasReportOrgId = 0,
    ipasReportEquipId = "",
    ipasReportLog = null,
    ipasReportMapMarker = null;

$( "#modal-ipas-report" )
    .on( "show.bs.modal", function (e) {

    })
    .on( "shown.bs.modal", function (e) {

        var encoded = $( e.relatedTarget ).data( "encoded" ),
            log = JSON.parse( decodeURI( encoded ) );


        ipasReportOrgId = log.org_id;
        ipasReportEquipId = log.equip_id;
        ipasReportLog = log;
        showReport( ipasReportOrgId, ipasReportEquipId, ipasReportLog );
        // console.log(333);
        // console.log( log );
            // encData =link.data( "encoded" )
        // log = JSON.parse( decodeURI( encData ) );
        // console.log(encoded);
    })
    .on( "hidden.bs.modal", function (e) {
        clearIpasReport();
    });


function showReport( orgId, equipId, log ) {
    
    // 로딩 이미지 보이기
    $( "#modal-ipas-report .modal-body" ).waitMe({
        effect: "win8",
        text: "Loading",
    });

    $.ajax({
        type:  "GET",
        async: true,
        url:  "/evtreport/" + equipId + "/org/" + orgId + "/since/" + ipasReportPeriod
    }).done( function( rpt ) {
        $( ".rpt-startDate" ).text( rpt.date.from );
        $( ".rpt-endDate" ).text( rpt.date.to );

        console.log(rpt);
        var equipTypeImg = "",
            equipType = "";
        if ( rpt.ipas.equip_type == PedestrianTag ) {
            equipTypeImg = "pt";
            equipType = "Pedestrian Tag";
        } else if ( rpt.ipas.equip_type == ZoneTag ) {
            equipTypeImg = "zt";
            equipType = "Zone Tag";
        } else if ( rpt.ipas.equip_type == VehicleTag ) {
            equipTypeImg = "vt";
            equipType = "Vehicle Tag";
        }
        $( ".rpt-ipas-equipType" ).text( equipType );
        $( ".rpt-ipas-usim" ).text( rpt.ipas.usim );
        $( "#rpt-img-equipType" ).attr( "src", "/static/assets/img/obj/" + equipTypeImg + ".png" );

        // IPAS
        $( ".rpt-ipas-orgName" ).text( rpt.ipas.org_name );
        $( ".rpt-ipas-tag" ).html( rpt.ipas.org_name + '<i class="fa fa-chevron-right mlr10 font-grey-salsa"></i>' + rpt.ipas.equip_id );
        $( ".rpt-ipas-equipId" ).text( equipId );
        $( ".rpt-ipas-created" ).text( rpt.ipas.created );
        $( ".rpt-ipas-updated" ).text( rpt.ipas.updated );
        $( ".rpt-ipas-location" ).text( rpt.ipas.latitude + " / " + rpt.ipas.longitude );

        // 건수 정보
        $( ".rpt-counts-startup" ).text( rpt.counts.startup );
        $( ".rpt-counts-shock" ).text( rpt.counts.shock );
        $( ".rpt-counts-speeding" ).text( rpt.counts.speeding );
        $( ".rpt-counts-proximity" ).text( rpt.counts.proximity );

        // 조회기간 서정
        $( ".btn-rpt-period" ).each(function( i, b ) {
            var period = $( b ).data( "period" );
            if ( ipasReportPeriod === period ) {
                $( b ).removeClass( "default" ).addClass( "blue" );
                return;
            }
        });

        // 자취 & 이벤트
        if ( rpt.events.length > 0 ) {
            $( "#map-rpt-ipas" ).removeClass( "hide" );
            updateTracks( rpt.events );
            $( "#table-rpt-events" ).bootstrapTable( "load", rpt.events );
            $( "#table-rpt-events" ).removeClass( "hide" );
;       }

    }).always(function() {
        // 로딩 이미지 삭제
        $( "#modal-ipas-report .modal-body" ).waitMe( "hide" );
    });
}


function updateTracks( tracks ) {
    var imageSrc = 'http://t1.daumcdn.net/localimg/localimages/07/mapapidoc/marker_red.png', // 마커이미지의 주소입니다
        imageSize = new daum.maps.Size(64, 69), // 마커이미지의 크기입니다
        imageOption = {offset: new daum.maps.Point(27, 69)}; // 마커이미지의 옵션입니다. 마커의 좌표와 일치시킬 이미지 안에서의 좌표를 설정합니다.
    var markerImage = new daum.maps.MarkerImage(imageSrc, imageSize, imageOption);
    var mapOption = {
            center: new daum.maps.LatLng( tracks[tracks.length-1].latitude, tracks[tracks.length-1].longitude ), // 지도의 중심좌표
            level: 3 // 지도의 확대 레벨
        };

    var map = new daum.maps.Map( ipasReportMap, mapOption );

    // 자취 그리기
    var positions = [];
    for ( var i=0; i<tracks.length; i++ ) {
        positions.push({
            title: tracks[i].date,
            latlng: new daum.maps.LatLng( tracks[i].latitude, tracks[i].longitude )
        });
    }
    for (var i = 0; i < positions.length; i ++) {
        var marker = new daum.maps.Marker({
            map: map, // 마커를 표시할 지도
            position: positions[i].latlng,
            title : positions[i].title,
            image: markerImage
        });
    }

    // 확대 & 축소 컨트럴 추가
    var zoomControl = new daum.maps.ZoomControl();
    map.addControl(zoomControl, daum.maps.ControlPosition.RIGHT);

    // 뷰 변환 컨트럴 추가
    var mapTypeControl = new daum.maps.MapTypeControl();
    map.addControl(mapTypeControl, daum.maps.ControlPosition.TOPRIGHT);
}


$( ".btn-rpt-period" ).click(function(e) {
    e.preventDefault();

    ipasReportPeriod = $( this ).data( "period" );
    clearIpasReport();
    showReport( ipasReportOrgId, ipasReportEquipId, ipasReportLog );
});


$( "#modal-ipas-map" ).on( "shown.bs.modal", function (e) {
    var latitude = $( e.relatedTarget ).data( "latitude" ),
        longitude = $( e.relatedTarget ).data( "longitude" ),
        $modal = $( this );

    $( ".latitude", $modal ).text( latitude );
    $( ".longitude", $modal ).text( longitude );
    // console.log(latitude + "/" + longitude);

    var options = {
        center: new daum.maps.LatLng( latitude, longitude ),
        level: 3
    };
    var map = new daum.maps.Map( mapIpas, options );
    var markerPosition  = new daum.maps.LatLng(latitude, longitude);
    var marker = new daum.maps.Marker({
        position: markerPosition
    });

    // 확대 & 축소 컨트럴 추가
    var zoomControl = new daum.maps.ZoomControl();
    map.addControl(zoomControl, daum.maps.ControlPosition.RIGHT);

    // 뷰 변환 컨트럴 추가
    var mapTypeControl = new daum.maps.MapTypeControl();
    map.addControl(mapTypeControl, daum.maps.ControlPosition.TOPRIGHT);

    marker.setMap(map);
});

function clearIpasReport() {
    $( ".rpt-data" ).empty();
    $( ".rpt-log" ).addClass( "hide" );
    $( "#rpt-img-equipType" ).attr( "src", "" );
    $( ".btn-rpt-period" ).removeClass( "blue" ).addClass( "default" );

    // 지도 감추기
    $( "#map-rpt-ipas" ).addClass( "hide" );

    // 로그 감추기
    $( "#table-rpt-events" ).bootstrapTable( "removeAll" );
    $( "#table-rpt-events" ).addClass( "hide" );
}