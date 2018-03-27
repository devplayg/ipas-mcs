$(function() {

//$("#modal-sensor-add").modal("show");
//$("#form-sensor-add select[name=Type1]").val(1).selectpicker( "refresh" );
//$("#form-sensor-add input[name=Name]").val("No name");
//$("#form-sensor-add input[name=Hostname]").val("10.10.10.10");
//$("#form-sensor-add input[name=Port]").val(4000);

    /**
     * 1. 초기화
     *
     */

    // Tree
    var $tree = $( "#tree-assets" ),
        selected = null,
        Root = 0,
        Org = 1,
        Group = 2;

    $tree.jstree({
        "plugins" : [
            "types", "state", "sort"
        ],

        "core" : {
            "data" : {
                "url" : "/assetclass/1/root/0",
            },
        },
        "types" : {
            1: {
                icon: "fa fa-building-o"
            },
            2: {
                icon: "fa fa-folder font-blue-sharp"
            }
        },

    }).on( "changed.jstree", function( e, obj ) {
        if ( obj.action == "select_node" ) {
            selected = obj.node;
        }
    });


    /**
     * 2. 이벤트
     *
     */
    $( ".btn-tree-refresh" ).click( function( e ) {
        $tree.jstree( true ).refresh( -1 );
    });
    $( ".btn-tree-expand" ).click( function( e ) {
        $tree.jstree( "open_all" );
    });
    $( ".btn-tree-collapse" ).click( function( e ) {
        $tree.jstree( "close_all" );
    });

    $( "#modal-asset-add" )
        .on( "shown.bs.modal", function () {
            $( "#form-asset-add input[name=name]" ).focus();
        }).on( "hidden.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $form.get( 0 ).reset();
        });

    $( ".btn-asset-add" ).click(function() {
        var name ;
        if ( selected.original.type === Root ) {
            name = felang[ "org" ];
        } else if ( selected.original.type === Org ) {
            name = felang[ "group" ];
        } else if ( selected.original.type === Group ) {
            return;
        } else {
            return;
        }

        $( "#form-asset-add input[name=parent_id]" ).val( selected.original.asset_id );
        $( "#form-asset-add input[name=type1]" ).val( selected.original.type + 1 );
        $( "#form-asset-add .name" ).text( name );
        $( "#modal-asset-add" ).modal( "show" );
    });

    $( "#form-asset-add" ).submit(function( e ) {
        e.preventDefault();
        var $form = $( this );

        $.ajax({
            type: "POST",
            async: true,
            url: "/assets",
            data: $form.serialize()
        }).done( function( result ) {
            if ( result.state ) {
                $( "#modal-asset-add" ).modal( "hide" );
                $tree.jstree( true ).refresh( -1 );
            } else {
                swal("fail", result.message, "error");
            }
        });

    });

//$( "#tree-assets').on( "changed.jstree", function (e, data) {
//    console.log(3);
//    $( ".btn-page').text(1);
//    $( "#form-filter input[name=offset]').val(0);

//    //if (defined(data.node.original)) {
//    if (data.node != undefined) {
//        // Nav. text
//        var nav = data.node.original.text_origin;
//        if (data.node.parents.length == 2) {
//            nav = data.instance.get_node(data.node.parent).original.text_origin + " <span class='mlr10'> &gt; </span> " + nav;
//        }
//        if (data.node.parents.length == 3) {
//            nav = data.instance.get_node(data.node.parent).original.text_origin + " <span class='mlr10'> &gt; </span> " + nav;
//            nav = data.instance.get_node(data.instance.get_node(data.node.parent).parent).original.text_origin + " <span class='mlr10'> &gt; </span> " + nav;
//        }
//        $( ".btn-org-tag').html(nav);
//        var param = {
//            depth: data.node.parents.length,
//            id: data.node.id,
//            keyword: $.trim($( "#form-filter input[name=keyword]').val())
//        };

//        var url = '?mod=agent&act=procGetAgent&with_version=1&' + filter_query + '&' + $.param(param);
//        $( "#table_agent').bootstrapTable( "refresh', {url: url} );
//    } else {
//        started = 1;
//    }
//});






    /**
     * 3. 함수
     *
     */

//     var delete_node = function( node ) {
//         swal({
//             title: "Are you sure?",
//             text: "You will not be able to recover this imaginary file!",
//             type: "warning",
//             showLoaderOnConfirm: true,
//             showCancelButton: true,
//             confirmButtonText: "Yes, delete it!",
//             cancelButtonText: "No, keep it",
//             confirmButtonColor: "#E7505A",
//         }).then(function() {
//             $.ajax({
//                 type: "DELETE",
//                 url: "/assets/" + node.original.AssetId,
//                 async: true,
//                 beforeSend: function(xhr) {
//                     xhr.setRequestHeader("X-CSRFToken", getXsrsToken());
//                 },
//             }).done( function( result ) {
//                 if ( result.state ) {
// //                $tree.jstree( true ).refresh( -1 );
//                 } else {
//                     swal( "Fail", result.Message, "error" );
//                 }
//             }).fail( function() {
//             }).always( function() {
//                 $tree.jstree( true ).refresh( -1 );
//             });
//         })
//     }


});

