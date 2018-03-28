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
            console.log(selected);
            $( "button.btn-asset-manage" ).each(function( idx ) {
                var $el = $( this ),
                    perm = $el.data( "perm" );

                if ( perm & (1 << selected.original.type) ) {
                    $el.removeClass( "hide" );
                    // console.log("perm:" + perm);
                    // console.log("selected.original.type:" + (1 << selected.original.type));
                } else {
                    $el.addClass( "hide" );
                }
            });
            // console.log(selected.original.type);
            // if ( selected.original.type == Root || selected.original.type == Org ) {
            //     $( ".btn-asset-add" ).removeClass( "hide" );
            // } else {
            //     $( ".btn-asset-add" ).addClass( "hide" );
            // }
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
        var root = $tree.jstree(true).get_node("j1_1").state;
        if ( root.opened ) {
            $tree.jstree( "close_all" );
        } else {
            $tree.jstree( "open_all" );
        }
    });
    $( ".btn-tree-collapse" ).click( function( e ) {
        $tree.jstree( "close_all" );
    });

    $( "#modal-asset-add" )
        .on( "shown.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $( "input[name=name]", $form ).focus();
        }).on( "hidden.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $form.get( 0 ).reset();
        });

    $( "#modal-asset-edit" )
        .on( "shown.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $( "input[name=name]", $form ).focus().select();

        }).on( "hidden.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $form.get( 0 ).reset();
        });


    $( ".btn-asset-add" ).click(function() {
        var target ;
        if ( selected.original.type === Root ) {
            target = felang[ "org" ];
        } else if ( selected.original.type === Org ) {
            target = felang[ "group" ];
        } else {
            return;
        }

        $( "#form-asset-add input[name=parent_id]" ).val( selected.original.asset_id );
        $( "#form-asset-add input[name=type1]" ).val( selected.original.type + 1 );
        $( "#form-asset-add .target" ).text( target );
        $( "#form-asset-add .name" ).text( selected.original.name );
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
                swal(result.message, "", "error");
            }
        });
    });


    $( ".btn-asset-edit" ).click(function() {
        $.ajax({
            type: "GET",
            async: true,
            url: "/assets/" + selected.original.asset_id,
        }).done( function( result ) {
            if ( result.state ) {
                var target ;
                if ( selected.original.type === Org ) {
                    target = felang[ "org" ];
                } else if ( selected.original.type === Group ) {
                    target = felang[ "group" ];
                } else {
                    return;
                }
                $( "#form-asset-edit .target" ).text( target );
                $( "#form-asset-edit input[name=name]" ).val( result.data.name );
                $( "#form-asset-edit input[name=asset_id]" ).val( selected.original.asset_id );
                $( "#modal-asset-edit" ).modal( "show" );

            } else {
                swal(result.message, "", "error");
            }
        });
    });


    $( ".btn-asset-remove" ).click(function() {
        swal({
            title: felang[ "msg.confirm_delete" ],
            type: "warning",
            showCancelButton: true,
            confirmButtonClass: "btn red",
            confirmButtonText: felang[ "yes" ],
            cancelButtonText: felang[ "no" ]
        }).then((result) => {
            if (result.value) {
                $.ajax({
                    type: "DELETE",
                    async: true,
                    url: "/assets/" + selected.original.asset_id,
                }).done( function( result ) {
                    if ( result.state ) {
                        $tree.jstree( true ).refresh( -1 );
                    } else {
                        swal(result.message, "", "error");
                    }
                });
            }
        });
    });


    $( "#form-asset-edit" ).submit(function( e ) {
        e.preventDefault();

        var $form = $( this );

        $.ajax({
            type: "patch",
            async: true,
            url: "/assets/" + $( "input[name=asset_id]", $form ).val(),
            data: $form.serialize()
        }).done( function( result ) {
            console.log(result);
            if ( result.state ) {
                $( "#modal-asset-edit" ).modal( "hide" );
                $tree.jstree( true ).refresh( -1 );
            } else {
                swal(result.message, "", "error");
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

