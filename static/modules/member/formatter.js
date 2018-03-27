function memberPositionFormatter( val, row, idx ) {
    var marks = '';
    if ( val >= positions["Administrator"] ) {
        var star = '';
        if ( val & positions["Superman"] ) {
            star += '<i class="fa fa-star"></i> ';
        }
        marks += '<button type="button" class="btn red btn-xs">Administrator ' + star + '</button>';
    }

    if ( val & positions["Observer"] ) {
        marks += '<button type="button" class="btn blue btn-xs clear">Objserver <i class="fa fa-search"></i></button>';
    }
    if ( val & positions["User"] ) {
        marks += '<button type="button" class="btn green btn-xs">User <i class="fa fa-user"></i></button>';
    }
    return marks;
}

function memberCommandFormatter(val, row, idx) {
    var str = '<a class="edit" href="javascript:void(0)" title="Edit">'
            + '<i class="icon-pencil"></i>'
            + '</a>';

    if (row.position < positions['Administrator']) {
        str += '<a class="acl" href="javascript:void(0)" title="Access control list">'
            + '<i class="icon-layers ml5"></i>'
            + '</a>';
    }

    if ( (row.position & positions['Superman']) == 0) {
        str +=    '<a class="remove ml5 " href="javascript:void(0)" title="Remove">'
            + '<i class="fa fa-trash-o s18 font-red"></i>'
            + '</a>';
    }
    return str;

}

function memberAllowedIpFormatter(val, row, idx) {
    if (val !== null) {
        var list = val.replace(/\/32/g, '').split(",");
        var str = '';
        for (var i=0; i<list.length; i++) {
            if (list[i].indexOf("/") > -1) {
                str += list[i];
            } else {
                // var s = namecardTagOnlyFormatter(list[i], row, index) + " " + list[i];
                str += list[i];
            }
            str += "<br>";
        }

        return str;
    }
}


function memberAssetSummaryFormatter(val, row, idx) {
    if ( val === null || val == "" ) return;

    var list = val.split( "__//__" );
    var tags = '';
    for (var i=0; i<list.length; i++) {
        var asset = list[i].split( "__--__" );
        tags += asset[0] + '<i class="fa fa-chevron-right mlr10 font-grey-cascade"></i>' + asset[1] + "<br>";
    }
    return tags;
}