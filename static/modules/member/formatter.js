function memberPositionFormatter( val, row, idx ) {
    var marks = '';
    if ( val >= positions["Administrator"] ) {
        var star = '';
        if ( val & positions["Superman"] ) {
            star += '<i class="fa fa-star"></i> ';
        }
        marks += '<button type="button" class="btn btn-danger btn-xs">Administrator ' + star + '</button>';
    }

    if ( val & positions["Observer"] ) {
        marks += '<button type="button" class="btn btn-primary btn-xs clear">Objserver <i class="fa fa-search"></i></button>';
    }
    if ( val & positions["User"] ) {
        marks += '<button type="button" class="btn btn-success btn-xs">User <i class="fa fa-user"></i></button>';
    }
    return marks;
}

function memberCommandFormatter(val, row, idx) {
    var str = '<a class="edit" href="javascript:void(0)" title="Edit">'
            + '<i class="fa fa-edit s18"></i>'
            + '</a>'

            + '<a class="reset_pwd ml5" href="javascript:void(0)" title="Reset password">'
            + '<i class="fa fa-key s18"></i>'
            + '</a>';

    if ( (row.position & positions['Superman']) == 0) {
        str +=    '<a class="remove ml5 " href="javascript:void(0)" title="Remove">'
            + '<i class="fa fa-remove s18 font-red"></i>'
            + '</a>';
    //
    //         + '<a class="ippool ml5" href="javascript:void(0)" title="IP Pool">'
    //         + '<i class="fa fa-th s18"></i>'
    //         + '</a>'
    }
    return str;

}