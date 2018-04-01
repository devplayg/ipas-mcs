function shockCountFormatter(val, row, idx) {
    if ( row.shock_count >= 10 ) {
        return '<button class="btn red btn-xs">' +  val + '</button>';

    } else if ( row.shock_count >= 8 ) {
        return '<button class="btn btn-warning btn-xs">' +  val + '</button>';
    }
}


function rowStyle(row, idx) {
    if ( row.shock_count >= 10 ) {
        return {
            classes: "row-danger"
        };
    } else if ( row.shock_count >= 8 ) {
        // return {
        //     classes: "row-warning"
        // };
    }

    return {};
}