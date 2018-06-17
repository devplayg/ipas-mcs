var mapIpas = document.getElementById( "map-ipas" ),
    ipasReportMap = document.getElementById( "map-rpt-ipas" ),
    ipasReportOrgId = 0, // 기관 ID
    ipasReportEquipId = "", // 장비 ID
    ipasReportLog = null, // 로그 객체
    ipasReportDate = null, // 검색 기준 날짜
    ipasReportPastDays = 0; // 검색 시작 날짜
// var content = '<div class="customoverlay">' +
//     '  <a href="http://map.daum.net/link/map/11394059" target="_blank">' +
//     '    <span class="title">구의야구공원</span>' +
//     '  </a>' +
//     '</div>';



$( "#modal-ipas-report" )
    .on( "shown.bs.modal", function (e) { // 로그 화면에서 진입할 경우

        var encoded = $( e.relatedTarget ).data( "encoded" ),
            log = JSON.parse( decodeURI( encoded ) );

        ipasReportOrgId = log.org_id;
        ipasReportEquipId = log.equip_id;
        ipasReportLog = log;
        if ( ipasReportLog.date !== undefined ) {
            ipasReportDate = ipasReportLog.date;
            $( ".btn-rpt-theday" )
                .removeClass( "hide" )
                .text( ipasReportDate.substr(0, 10) )
                .data( "date", ipasReportDate );
            $( ".btn-rpt-theday" ).removeClass( "default" ).addClass( "green" );
        } else {
            $( ".btn-rpt-theday" ).removeClass( "green" ).addClass( "default" );
        }
        showReport( ipasReportOrgId, ipasReportEquipId );
    })
    .on( "hidden.bs.modal", function (e) {
        clearIpasReport();
    });


function showReport( orgId, equipId ) {
    
    // 로딩 이미지 보이기
    $( "#modal-ipas-report .modal-body" ).waitMe({
        effect: "win8",
        text: "Loading",
    });

    var param = {};
    if ( ipasReportLog !== null ) {
        param.date = ipasReportDate;
        param.past_days = ipasReportPastDays;
    }

    var url = "/report/evt/org/"+orgId+"/eqid/" + equipId + "?" + $.param(param);
    $.ajax({
        type:  "GET",
        async: true,
        url:  url
    }).done( function( rpt ) {
        console.log(rpt);
        $( ".btn-rpt-today" ).data( "date", rpt.date.today );
        $( ".rpt-startDate" ).text( moment( rpt.date.from ).format("lll") );
        $( ".rpt-endDate" ).text( moment( rpt.date.to ).format("lll") );

        var a = moment( rpt.date.from ).format("lll");
        console.log(a);

        // console.log(rpt);
        var equipTypeImg = "",
            equipType = "";

        // 태그 타입
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

        // IPAS 정보
        $( ".rpt-ipas-orgName" ).text( rpt.ipas.org_name );
        $( ".rpt-ipas-tag" ).html( rpt.ipas.org_name + '<i class="fa fa-chevron-right mlr10 font-grey-salsa"></i>' + rpt.ipas.equip_id );
        $( ".rpt-ipas-equipId" ).text( equipId );
        $( ".rpt-ipas-created" ).text( rpt.ipas.created );
        $( ".rpt-ipas-updated" ).text( rpt.ipas.updated );
        $( ".rpt-ipas-location" ).text( rpt.ipas.latitude + " / " + rpt.ipas.longitude );

        // 이벤트 발생 수 정보
        $( ".rpt-counts-startup" ).text( rpt.counts.startup );
        $( ".rpt-counts-shock" ).text( rpt.counts.shock );
        $( ".rpt-counts-speeding" ).text( rpt.counts.speeding );
        $( ".rpt-counts-proximity" ).text( rpt.counts.proximity );

        // 조회기간 서정
        $( ".btn-rpt-period" ).each(function( i, b ) {
            var period = $( b ).data( "period" );
            if ( ipasReportPastDays === period ) {
                $( b ).removeClass( "default" ).addClass( "blue" );
                return;
            }
        });

        // 자취 & 이벤트
        if ( rpt.events !== null && rpt.events.length > 0 ) {
            updateTracks( rpt.events, rpt.status );
            // 지도 중심을 이동 시킵니다
            // map.setCenter( new daum.maps.LatLng(33.452613, 126.570888) );

            $( "#table-rpt-events" ).bootstrapTable( "load", rpt.events );
            $( "#table-rpt-events" ).removeClass( "hide" );
;       }

        if ( ipasReportLog !== null ) { // 로그 정보
            var eventType = "";
            if ( ipasReportLog.event_type === StartupEvent ) {
                eventType = felang.startup;

            } else if ( ipasReportLog.event_type === ShockEvent ) {
                eventType = felang.shock;

            } else if ( ipasReportLog.event_type === SpeedingEvent ) {
                eventType = felang.speeding;

            } else if ( ipasReportLog.event_type === ProximityEvent ) {
                eventType = felang.proximity;

            } else {
                eventType = "N/A (" + ipasReportLog.event_type + ")"  ;
            }
            $( ".rpt-log-eventType" ).text( eventType );
            $( ".rpt-log-location" ).text( ipasReportLog.latitude + " / " + ipasReportLog.longitude );
            $( ".rpt-log-snr" ).html( snrFormatter( ipasReportLog.snr, null, null ) );
            $( ".rpt-log-snr-value" ).html( ipasReportLog.snr );
            $( ".rpt-log-speed" ).text( ipasReportLog.speed );

            $( ".rpt-log" ).removeClass( "hide" );
        }

    }).always(function() {
        // 로딩 이미지 삭제
        $( "#modal-ipas-report .modal-body" ).waitMe( "hide" );

        // 지도 맞춤

    });
}


function updateTracks( eventTracks, statusTracks ) {
    $( "#map-rpt-ipas" ).removeClass( "hide" );

    // var imageSrc = 'http://t1.daumcdn.net/localimg/localimages/07/mapapidoc/marker_red.png', // 마커이미지의 주소입니다
    //     imageSize = new daum.maps.Size(64, 69), // 마커이미지의 크기입니다
    //     imageOption = {offset: new daum.maps.Point(27, 69)}; // 마커이미지의 옵션입니다. 마커의 좌표와 일치시킬 이미지 안에서의 좌표를 설정합니다.
    // var markerImage = new daum.maps.MarkerImage(imageSrc, imageSize, imageOption);
    // 마커 이미지의 이미지 크기 입니다
    var imageSrc = "http://t1.daumcdn.net/localimg/localimages/07/mapapidoc/markerStar.png",
        imageSize = new daum.maps.Size(24, 35);
        markerImage = new daum.maps.MarkerImage(imageSrc, imageSize);

    var mapOption = {
        center: new daum.maps.LatLng( eventTracks[eventTracks.length-1].latitude, eventTracks[eventTracks.length-1].longitude ), // 지도의 중심좌표
        level: 3 // 지도의 확대 레벨
    };
    var map = new daum.maps.Map( ipasReportMap, mapOption );

    // 상태 자취 표시
    for ( var i=0; i<statusTracks.length; i++ ) {
        var marker = new daum.maps.Marker({
            map: map, // 마커를 표시할 지도
            position: new daum.maps.LatLng( statusTracks[i].latitude+0.003, statusTracks[i].longitude+0.003 ),
            title : statusTracks[i].date,
            // image: markerImage
        });
    }

    // 이벤트 자취 표시
    for ( var i=0; i<eventTracks.length; i++ ) {
        // var marker = new daum.maps.Marker({
        //     map: map, // 마커를 표시할 지도
        //     position: new daum.maps.LatLng( eventTracks[i].latitude, eventTracks[i].longitude ),
        //     title : eventTracks[i].date,
        //     // image: markerImage
        // });
        var content = '';
        if ( eventTracks[i].event_type == ShockEvent ) {
            content = '<div class="markerimg blue"></div><button class="btn blue btn-xs"><i class="fa fa-bolt"></i> ' + felang.shock + '</button>';
        } else if ( eventTracks[i].event_type == SpeedingEvent ) {
            content = '<div class="markerimg yellow"></div><button class="btn btn-warning btn-xs" style="background-color: #f4902f"><i class="icon-speedometer"></i> ' + felang.speeding + '</button>';
        } else if ( eventTracks[i].event_type == ProximityEvent ) {
            content = '<div class="markerimg red"></div><button class="btn red btn-xs"><i class="icon-size-actual"></i> ' + felang.proximity + '</button>';
        }
        // content += '<div class="markerimg red"></div>';

        // content = '<span class="fa-stack fa-lg font-red-haze"><i class="fa fa-square fa-stack-2x"></i><i class="fa fa-map-marker fa-stack-1x fa-inverse"></i></span> ';
        // <span class="fa-stack fa-lg"><i class="fa fa-square fa-stack-2x"></i><i class="fa fa-map-marker fa-stack-1x fa-inverse"></i></span>

        var customOverlay = new daum.maps.CustomOverlay({
            map: map, // 마커를 표시할 지도
            position: new daum.maps.LatLng( eventTracks[i].latitude, eventTracks[i].longitude ),
            title : eventTracks[i].date,
            // content: '<div class="customoverlay"><a href="#"><span class="title">' + eventTracks[i].event_type + '</span></a></div>',
            content: content,
            // xAnchor: 0.3,
            yAnchor: 1.00,
            image: markerImage
        });
    }




    // 확대 & 축소 컨트럴 추가
    var zoomControl = new daum.maps.ZoomControl();
    map.addControl(zoomControl, daum.maps.ControlPosition.RIGHT);

    // 뷰 변환 컨트럴 추가
    var mapTypeControl = new daum.maps.MapTypeControl();
    map.addControl(mapTypeControl, daum.maps.ControlPosition.TOPRIGHT);

    // 지도 이동
    map.setCenter( new daum.maps.LatLng(ipasReportLog.latitude, ipasReportLog.longitude) );
}


$( ".btn-rpt-period" ).click(function(e) {
    e.preventDefault();

    // 기간 설정
    ipasReportPastDays = $( this ).data( "period" );
    clearIpasReport();
    showReport( ipasReportOrgId, ipasReportEquipId );
});


$( ".btn-rpt-date" ).click(function(e) {
    e.preventDefault();

    // 기간 설정
    var date = $( this ).data( "date" );
        ipasReportDate = date;
    $( ".btn-rpt-date" ).each(function( i, b ) {
        var $btn = $( this );
        if ( $btn.data("date").substr( 0, 10 ) === date.substr( 0, 10 ) ) {
            $btn.removeClass( "default" ).addClass( "green" );
        } else {
            $btn.removeClass( "green" ).addClass( "default" );
        }
    });

    showReport( ipasReportOrgId, ipasReportEquipId );
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
    $( "#map-rpt-ipas" ).empty();

    // 로그 감추기
    $( "#table-rpt-events" ).bootstrapTable( "removeAll" );
    // $( "#table-rpt-events" ).addClass( "hide" );
}