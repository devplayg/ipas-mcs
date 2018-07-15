function rankFormatter( val, row, idx ) {
    var btn_class = 'btn-default';
    if (val == 1) {
        btn_class = 'blue';
    }

    return '<button class="btn ' + btn_class + ' btn-xs">' + val + '</button>';
}


function dashboardOrgGroupNameOfStartupEventFormatter( val, row, idx ) {
    return dashboardOrgGroupNameFormatter( row, StartupEvent );
}


function dashboardOrgGroupNameOfShockEventFormatter( val, row, idx ) {
    return dashboardOrgGroupNameFormatter( row, ShockEvent );
}


function dashboardOrgGroupNameOfSpeedingEventFormatter( val, row, idx ) {
    return dashboardOrgGroupNameFormatter( row, SpeedingEvent );
}


function dashboardOrgGroupNameOfProximityEventFormatter( val, row, idx ) {
    return dashboardOrgGroupNameFormatter( row, ProximityEvent );
}


function dashboardOrgGroupNameFormatter( row, eventType ) {
    var asset = row.item.split( "/", 2 ),
        groupName = row.group_name;
    if ( row.item.endsWith( "/0" ) ) {
        groupName = '<span class="font-grey-silver">Default</span>';
    }
    var attr =  getIpasLogLinkAttr({
        stats_mode:  true,
        org_id: asset[0],
        group_id: asset[1],
        event_type: eventType,
    });
    return '<a href="#" style="color: #555555" ' + attr + '>' +  '<span class="s12">' + row.org_name + '</span>' + '<i class="fa fa-angle-right mlr5"></i>' + groupName + '</a>';
}


function dashboardIpasEquipIdOfStartupEventFormatter( val, row, idx ) {
    var attr =  getIpasLogLinkAttr({
        stats_mode:  true,
        equip_id: val,
        event_type: StartupEvent
    });
    return '<a href="#"' + attr + '>' +  getIpasTag( val ) + '</a>';
}


function dashboardIpasEquipIdOfShockEventFormatter( val, row, idx ) {
    var attr = getIpasLogLinkAttr({
        stats_mode:  true,
        equip_id: val,
        event_type: ShockEvent
    });
    var a = '<a href="#"' + attr + '>' +  getIpasTag( val ) + '</a>';
    return a;
}


function dashboardIpasEquipIdOfSpeedingEventFormatter( val, row, idx ) {
    var attr = getIpasLogLinkAttr({
        stats_mode:  true,
        equip_id: val,
        event_type: SpeedingEvent
    });
    return '<a href="#"' + attr + '>' +  getIpasTag( val ) + '</a>';
}


function dashboardIpasEquipIdOfProximityEventFormatter( val, row, idx ) {
    var attr = getIpasLogLinkAttr({
        stats_mode:  true,
        equip_id: val,
        event_type: ProximityEvent
    });
    return '<a href="#"' + attr + '>' +  getIpasTag( val ) + '</a>';
}


function getIpasLogLinkAttr( req ) {
    req.fast_paging = "on";
    var attr = ' class="btn-show-ipaslog-on-modal" data-query="'+ $.param( req ) + '" ';
    return attr
}


function dashboardEventDescriptionFormatter( val, row, idx ) {
    var eventName,
        // eventIcon,
        lineIcon,

        prefix = '<a href="#" data-toggle="modal" data-target="#modal-ipas-report" data-org-id="' + row.org_id + '" data-equip-id="' + row.equip_id + '" data-encoded="' + encodeURI( JSON.stringify( row ) ) + '" ><span>';
        suffix = "</a>";
        // lineIcon = '<div class="label label-sm label-success"><i class="fa fa-bell-o"></i></div>';

    if ( row.event_type === StartupEvent ) {
        // lineIcon = '<div class="label label-sm label-success"><i class="fa fa-bell-o"></i></div>';
        eventName = felang.startup;
        lineIcon = '<div class="label label-sm label-success"><i class="icon-power label-icon-w"></i></div>';

    } else if ( row.event_type === ShockEvent ) {
        // lineIcon = '<i class="fa fa-warning"></i>';
        eventName = felang.shock;
        lineIcon = '<div class="label label-sm label-primary" ><i class="fa fa-bolt label-icon-w"></i></div>';

    } else if ( row.event_type === SpeedingEvent ) {
        eventName = felang.speeding;
        lineIcon = '<div class="label label-sm label-warning bg-orange"><i class="icon-speedometer label-icon-w"></i></div>';

    } else if ( row.event_type === ProximityEvent ) {
        eventName = felang.proximity;
        lineIcon = '<div class="label label-sm label-danger bg-red"><i class="icon-size-actual label-icon-w"></i></div>';
        prefix += '<span class="font-red">';

    } else {
        eventName = "Unknown";
        lineIcon = '<i class="icon-info"></i>';
    }

    if ( lang === "ko-kr" ){
        return prefix + lineIcon + " <i>" + row.org_name + "</i> 의 <i>" + row.group_name + "</i> 에서 " + " " + eventName + " 이벤트가 발생하였습니다" + suffix;
    } else {
        return prefix + lineIcon + " " + eventName + " event in <i>" + row.group_name + ", " + row.org_name + "</i>" + suffix;
    }
}


function dashboardDateAgoFormatter( val, row, idx ) {
    return '<span class="s12"><i>' + val + '</i></span>';
}

function optimeFormatter( val, row, idx ) {
    return sec2humanReadable( val );
}

function sec2humanReadable(duration) {
    var hour = 0;
    var min = 0;
    var sec = 0;

    if (duration) {
        if (duration >= 60) {
            min = Math.floor(duration / 60);
            sec = duration % 60;
        }
        else {
            sec = duration;
        }

        if (min >= 60) {
            hour = Math.floor(min / 60);
            min = min - hour * 60;
        }

        if (hour < 10) {
            hour = '0' + hour;
        }
        if (min < 10) {
            min = '0' + min;
        }
        if (sec < 10) {
            sec = '0' + sec;
        }
    }

    return hour +":"+ min +":"+ sec;
}



// function dashboardIpasEquipIdFormatter( equipId, row, eventType ) {
//     var param = {
//         stats_mode:  true,
//         fast_paging: "on",
//         equip_id:    equipId,
//         start_date:  null,
//         end_date:    null,
//         event_type:  eventType
//     };
//
//     if ( $("#start_date").val() !== undefined && ! $("#start_date").prop( "disabled" ) ) {
//         param.start_date = $("#start_date").val();
//
//         // 검색 시작날짜가 설정된 경우
//         if ( $("#end_date").val() !== undefined && ! $("#end_date").prop( "disabled" ) ) { // 시작/종료 날짜가 설정된 경우(기간 검색인 경우)
//             param.end_date = $("#end_date").val();
//         } else { // 시작 날짜만 설졍된 경우 (특정 날짜를 검색하는 경우)
//
//         }
//     }
//
//     var prefix = '<a href="#" class="btn-show-ipaslog-on-modal" style="color: inherit; " data-query="'+ $.param( param ) + '">',
//         suffix = '</a>';
//     return prefix + getIpasTag( equipId ) + suffix;
// }