<!-- IP Card -->
<div id="modal-ipcard" class="modal fade" tabindex="-1" role="dialog">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <form role="form" id="form-ipcard-edit" class="form-ipcard">
                <input type="hidden" class="form-control" name="Ip" readonly>

                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4 class="modal-title">
                        <i class="fa fa-vcard-o"></i> IP Card
                    </h4>
                </div>

                <div class="modal-body">
                    <div class="row">
                        <div class="col-lg-4 col-md-4 col-sm-4 form-group">
                            <label class="control-label">IP</label>
                            <input type="text" class="form-control" name="IpStr" readonly>
                        </div>
                        <div class="col-lg-3 col-md-3 col-sm-3 form-group">
                            <label class="control-label">Name</label>
                            <input type="text" class="form-control" name="Name">
                        </div>
                        <div class="col-lg-5 col-md-5 col-sm-5 form-group">
                            <label class="control-label">Position</label>
                            <input type="text" class="form-control" name="Position">
                        </div>

                    </div>

                    <div class="row">
                        <div class="col-lg-4 col-md-4 col-sm-4 form-group">
                            <label class="control-label">Dept1</label>
                            <input type="text" class="form-control" name="Dept1">
                        </div>
                        <div class="col-lg-8 col-md-8 col-sm-8 form-group">
                            <label class="control-label"><i class="fa fa-envelope-o"></i> Email</label>
                            <input type="text" class="form-control" name="Email">
                        </div>
                    </div>

                    <div class="row">
                        <div class="col-lg-4 col-md-4 col-sm-4 form-group">
                            <label class="control-label">Dept2</label>
                            <input type="text" class="form-control" name="Dept2">
                        </div>
                        <div class="col-lg-4 col-md-4 col-sm-4 form-group">
                            <label class="control-label"><i class="fa fa-phone"></i> Phone1</label>
                            <input type="text" class="form-control" name="Phone1">
                        </div>
                        <div class="col-lg-4 col-md-4 col-sm-4 form-group">
                            <label class="control-label"><i class="fa fa-building-o"></i> Office</label>
                            <input type="text" class="form-control" name="Phone2">
                        </div>
                    </div>

                    <div class="note note-danger hidden"></div>
                </div> <!-- .modal-body -->
                <div class="modal-footer">
                    <button type="submit" class="btn btn-primary">Update</button>
                    <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                </div>
            </form>
        </div> <!-- .modal-contents -->
    </div> <!-- .modal-dialog -->
</div> <!-- .modal -->
<div class="mt-element-ribbon bg-grey-steel ipcard" style="postion:absolute; width:300px;">
    <div class="ribbon ribbon-color-primary uppercase">Ribbon</div>
    <p class="ribbon-content">Duis mollis, est non commodo luctus, nisi erat porttitor ligula</p>
</div>


<script>
//<!--$( ".btn-ipcard-show" ).click(function(e) {-->
//$('body').on('click', '.btn-ipcard-show', function() {


$( "#modal-ipcard" )
    .on( "shown.bs.modal", function ( e ) {
        var $modal = $( this );
        $( ".modal-body", $modal ).waitMe({
            effect: "roundBounce",
            text: "Loading",
        });

        var $form = $( this ).find( "form" ),
            $btn = $( e.relatedTarget ),
            ip_str = $btn.data( "ip" );

        $.ajax({
            type: "GET",
            async: true,
            url: "/ipcard/" + inet_aton( ip_str ),
        }).done( function( ipcard ) {
            $( "input[name=Ip]", $form ).val( inet_aton( ip_str ) );
            $( "input[name=IpStr]", $form ).val( ip_str );
            $( "input[name=Name]", $form ).val( ipcard.Name );
            $( "input[name=Position]", $form ).val( ipcard.Position );
            $( "input[name=Dept1]", $form ).val( ipcard.Dept1 );
            $( "input[name=Dept2]", $form ).val( ipcard.Dept2 );
            $( "input[name=Phone1]", $form ).val( ipcard.Phone1 );
            $( "input[name=Phone2]", $form ).val( ipcard.Phone2 );
            $( "input[name=Email]", $form ).val( ipcard.Email );
        }).always( function() {
            $( ".modal-body", $modal ).waitMe( "hide" );
            $( "input[name=Name]", $form ).focus().select();
        });
    })
    .on( "hidden.bs.modal", function () {
        var $form = $( this ).find( "form" );
        $form.get(0).reset();
        $form.find( ".has-error" ).removeClass( "has-error" );
    });

$( "#form-ipcard-edit" ).validate({
    submitHandler: function() {
        var $form = $( this.currentForm ),
            $modal = $form.closest( ".modal" ),
            $btn = $(this.submitButton),
            ip = $( "input[name=Ip]", $form ).val();

        $.ajax({
            type: "PATCH",
            async: true,
            url: "/ipcard/" + ip,
            data: $form.serialize()
        }).done( function ( result ) {
            if ( result.State ) {
                $modal.modal( "hide" );
            } else {
                $( ".note", $form ).removeClass( "hidden" ).text( result.Message );
            }
        }).always( function() {
            $btn.button( "reset" );
        });


    },
    ignore: "input[type='hidden']",
    errorClass: "help-block",
    rules: {
        Name: {
            required: true,
            maxlength: 64,
        },
        Position: {
            maxlength: 64,
        },
        Dept1: {
            maxlength: 64,
        },
        Dept2: {
            maxlength: 64,
        },
        Phone1: {
            maxlength: 32,
        },
        Phone2: {
            maxlength: 32,
        },
        Email: {
            email: true,
        }
    },
    messages: {
    },
    highlight: function( elem ) {
        $( elem ).closest( ".form-group ").addClass( "has-error" );
    },
    unhighlight: function( elem ) {
        $( elem ).closest( ".form-group" ).removeClass( "has-error" );
    }
});

</script>
