// function shockCountFormatter( val, row, idx ) {
//     if ( row.shock_count >= 10 ) {
//         return '<button class="btn red btn-xs">' +  val + '</button>';
//
//     } else if ( row.shock_count >= 8 ) {
//         return '<button class="btn btn-warning btn-xs">' +  val + '</button>';
//     }
// }
//
//
// function ipaslogLocationFormatter(val, row, idx) {
//     var loc = '';
//     loc += '<a href="#" class="tooltips"  data-container="body" data-placement="top" data-original-title="Tooltip in top" data-toggle="modal" data-target="#modal-ipas-map" data-latitude="' + row.latitude + '" data-longitude="' + row.longitude + '"><i class="fa fa-map-marker s18"></i></a>';
//     // loc += '<small class="ml5">' + row.latitude + ", " + row.longitude + '</small>';
//
//     return loc;
// }
//
//
// function ipaslogEventTypeFormatter(val, row, idx) {
//     if ( val === StartupEvent ) {
//         return felang[ "startup" ] + ' <span class="pull-right"><i class="icon-power"></i></span>';
//
//     } else if ( val === ShockEvent ) {
//         return felang[ "shock" ] + ' <span class="pull-right"><i class="fa fa-bolt"></i></span></span>';
//
//     } else if ( val === SpeedingEvent ) {
//         return felang[ "speeding" ] + ' <span class="pull-right"><i class="icon-speedometer"></i></span>';
//         // return felang[ "speeding" ] + ' <span class="pull-right"><i class="fa fa-long-arrow-up"></i></span>';
//
//     } else if ( val === ProximityEvent) {
//         // return felang[ "proximity" ] + ' <span class="pull-right"><i class="fa fa-warning font-red"></i></span>';
//         return felang[ "proximity" ] + ' <span class="pull-right"><i class="icon-size-actual"></i></span>';
//     }
// }
//
// function equipIdFormatter( val, row, idx ) {
//     return getIpasTag( val );
// }
//
// function ipaslogTargetsFormatter( val, row, idx ) {
//     if ( row.event_type === ProximityEvent ) {
//         var list = val.split(","),
//             tags = '';
//
//         for (var i = 0; i < list.length; i++) {
//             tags += equipIdFormatter(list[i]);
//         }
//
//         return tags;
//     }
// }
//
// function ipaslogDistanceFormatter( val, row, idx ) {
//     if ( row.event_type === ProximityEvent ) {
//         return val;
//     }
// }
//
// function ipaslogSpeedingFormatter( val, row, idx ) {
//     var threshold = 12;
//     if ( row.event_type === SpeedingEvent && val > threshold) {
//         return val + '<span class="pull-right font-red bold s12">+' + ( val - threshold ) + '</span>';
//     }
// }
//
