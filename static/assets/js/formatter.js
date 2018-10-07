function ipWithTagFormatter( val, row, idx ) {
    var ip = inet_ntoa( val );
    if ( row.SrcIpCard.length > 0 ) {
        var ipcard = $.parseJSON( row.SrcIpCard );
        var head = '<button class="btn default btn-xs pull-right ipcard" data-toggle="modal" data-target="#modal-ipcard" data-ip="' + ip + '"><i class="fa fa-id-card-o"></i> ',
            body = ipcard.Name,
            foot = '</button>';
        return head + body + foot + ip;
    } else {
        var head = '<a href="#"><i class="fa fa-tag pull-right ipcard" data-toggle="modal" data-target="#modal-ipcard" data-ip="' + ip + '"></i> ',
            body = '',
            foot = '</a>';
        return head + body + foot + ip;
    }
}

function ip2intFormatter( val, row, idx ) {
    return inet_aton(val);
}

function int2ipFormatter( val, row, idx ) {
    return inet_ntoa(val);
}

function dateFormatter( val, row, idx ) {
    if ( val === undefined ) return;

    var m = moment( val );
    // return '<span class="">' + m.format( "MMMM D YYYY, hh:mm:ss" ) + '</span>';
    return '<span class="">' + m.format() + '</span>';
}

function shortDateFormatter( val, row, idx ) {
    var m = moment( val ),
        // prefix = '<span class="tooltips" data-container="body" data-placement="top" data-original-title="' + m.format( "MMMM D YYYY, h:mm:ss a" ) + '" title="' + m.format( "MMMM D YYYY, h:mm:ss a" ) + '">',
        prefix = '<span class="tooltips" data-container="body" data-placement="top" data-original-title="' + m.format( "MMMM D YYYY, h:mm:ss a" ) + '" title="' + m.format() + '">',
        suffix = '</span>';
    return prefix + m.format("HH:mm") + suffix;
}

function numberFormatter( val, row, idx ) {
    return val.toLocaleString();
}


function ipasEquipIdFormatter( val, row, idx ) {
    var header = '',
        body = '',
        footer = '';

    header += '<a href="#" data-toggle="modal" data-target="#modal-ipas-report" data-org-id="' + row.org_id + '" data-equip-id="' + row.equip_id + '" data-encoded="' + encodeURI( JSON.stringify( row ) ) + '" >';
    body += getIpasTag( row.equip_id );
    return header + body + footer;
}

function orgNameFormatter( val, row, idx ) {
    // return '<span class="tooltips" data-container="body" data-placement="top" data-original-title="' + row.org_id + '">' + row.org_name + '</span>'
    return '<span class="tooltips" title="Org ID: ' + row.org_id + '">' + row.org_name + '</span>'
}

function groupNameFormatter( val, row, idx ) {
    if ( row.group_id == 0 ) {
        return;
    }
    return row.group_name;
}

function orgGroupNameFormatter( val, row, idx ) {
    var orgName = orgNameFormatter( row.org_id, row, null );
    if ( row.group_id == 0 ) {
        return row.org_name + ' / <span class="font-grey-silver">Default</span>';

    }
    return orgName + " / " + row.group_name;
}



function ipasEventRowStyle( row, idx ) {
    if ( row.event_type == ProximityEvent ) {
        return {
            classes: "row-danger"
        };
    }
    // if ( row.shock_count >= 10 ) {
    //     return {
    //         classes: "row-danger"
    //     };
    // } else if ( row.shock_count >= 8 ) {
    //     return {
    //         classes: "row-warning"
    //     };
    // }

    return {};
}

function snrFormatter(val, row, idx) {
    var c  = '\u2759',
        color = ['red-mint', 'red-mint', 'yellow-crusta', 'yellow-crusta', 'blue', 'blue'],
        level = Math.ceil(val / 2);

    if ( level > 6 ) {
        level = 6;
    }

    var str = "";
    for (var i=0; i<level; i++) {
        str += c;
    }

    return '<span class="font-' + color[level-1] + '">' + str + '</span>';
}


function shockCountFormatter( val, row, idx ) {
    if ( row.shock_count >= 10 ) {
        return '<button class="btn red btn-xs">' +  val + '</button>';

    } else if ( row.shock_count >= 8 ) {
        return '<button class="btn btn-warning btn-xs">' +  val + '</button>';
    }
}


function ipaslogLocationFormatter(val, row, idx) {
    var loc = '';
    loc += '<a href="#" class="tooltips"  data-container="body" data-placement="top" data-original-title="Tooltip in top" data-toggle="modal" data-target="#modal-ipas-map" data-latitude="' + row.latitude + '" data-longitude="' + row.longitude + '"><i class="fa fa-map-marker s18"></i></a>';
    // loc += '<small class="ml5">' + row.latitude + ", " + row.longitude + '</small>';

    return loc;
}


function ipaslogEventTypeFormatter(val, row, idx) {
    if ( val === StartupEvent ) {
        return felang[ "startup" ] + ' <span class="pull-right"><i class="icon-power"></i></span>';

    } else if ( val === ShockEvent ) {
        return felang[ "shock" ] + ' <span class="pull-right"><i class="fa fa-bolt"></i></span></span>';

    } else if ( val === SpeedingEvent ) {
        return felang[ "speeding" ] + ' <span class="pull-right"><i class="icon-speedometer"></i></span>';
        // return felang[ "speeding" ] + ' <span class="pull-right"><i class="fa fa-long-arrow-up"></i></span>';

    } else if ( val === ProximityEvent) {
        // return felang[ "proximity" ] + ' <span class="pull-right"><i class="fa fa-warning font-red"></i></span>';
        return felang[ "proximity" ] + ' <span class="pull-right"><i class="icon-size-actual"></i></span>';
    }
}

function equipIdFormatter( val, row, idx ) {
    return getIpasTag( val );
}

function ipaslogTargetsFormatter( val, row, idx ) {
    if ( row.event_type === ProximityEvent ) {
        var list = val.split(","),
            tags = '';

        for (var i = 0; i < list.length; i++) {
            tags += equipIdFormatter(list[i]);
        }

        return tags;
    }
}

function ipaslogDistanceFormatter( val, row, idx ) {
    if ( row.event_type === ProximityEvent ) {
        return val;
    }
}

function ipaslogSpeedingFormatter( val, row, idx ) {
    var threshold = 12;
    if ( row.event_type === SpeedingEvent && val > threshold) {
        return val + '<span class="pull-right font-red bold s12">+' + ( val - threshold ) + '</span>';
    }
}

function ipasEquipTypeFormatter( val, row, idx ) {
    if ( val === PedestrianTag ) {
        return "PT";

    } else if ( val === ZoneTag ) {
        return "ZT";

    } else if ( val === VehicleTag ) {
        return "VT";

    } else {
        return val;
    }
}