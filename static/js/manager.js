$(function(){
    var isEmail = function(email) {
        var re = /^(\w-*\.*)+@(\w-?)+(\.\w{2,})+$/;
        if (re.test(email)) {
            return true;
        }
        return false;
    }
$('#usercreateform').ajaxForm({
    dataType: 'json',
    beforeSubmit: function() {
        if ($.trim($('#email').val()) == '') {
            $('#formtip').css("color", "red").html("请输入电子邮件!");
            return false;
        }
        if (!isEmail($.trim($('#email').val()))) {
            $('#formtip').css("color", "red").html("电子邮件格式不正确，请重新输入!");
            return false;
        }
        if ($.trim($('#name').val()) == '') {
            $('#formtip').css("color", "red").html("请输入用户姓名!");
            return false;
        }
        if ($.trim($('#password').val()) == '') {
            $('#formtip').css("color", "red").html("请输入用户密码!");
            return false;
        }
        if ($.trim($('#password').val()).length < 6) {
            $('#formtip').css("color", "red").html("请输入 6-16 位用户密码!");
            return false;
        }
        if ($.trim($('#repassword').val()) == '') {
            $('#formtip').css("color", "red").html("请再次输入用户密码!");
            return false;
        }
        if ($.trim($('#password').val()) != $.trim($('#repassword').val())) {
            $('#formtip').css("color", "red").html("两次输入的密码不一致!")
            return false;
        };
        if ($.trim($('#description').val()) == '') {
            $('#formtip').css("color", "red").html("请输入用户的简单介绍!");
            return false;
        }
        $('#formtip
                ').removeClass('
                color ').html('
                ');
                $("#save").attr({
                    "disabled": "disabled"
                }); $("#back").attr({
                    "disabled": "disabled"
                });
            },
            success: function(response) {
                $("#save").removeAttr("disabled");
                $("#back").removeAttr("disabled");
                $('#formtip').css("color", "red").html(response.info).addClass(response.code);
                var d = dialog({
                    title: '提示 ： ',
                    content: response.info
                });
                d.showModal();
                setTimeout(function() {
                    d.close().remove();
                }, 2000);
                if (response.code === 'success') {
                    $("#save").attr({
                        "disabled": "disabled"
                    });
                    $("#back").attr({
                        "disabled": "disabled"
                    });
                    setTimeout(function() {
                        window.location = "{{ site_url('manager/user') }}";
                    }, 2200)
                }
            }
    });
$('#articleform').ajaxForm({
    dataType: "json",
    beforeSubmit: function() {
        if ($.trim($('#title').val()) == "") {
            $("#formtip").css("color", "red").html("请输入文章标题！");
            return false;
        }
        if ($.trim($("#content").val()) == "") {
            $("#formtip").css("color", "red").html("请输入文章内容！");
            return false
        }
        $("#formtip").removeClass('color').html('');
        $("#save").attr({
            "disabled": "disabled"
        });
        $("#back").attr({
            "disabled": "disabled"
        });
    },
    success: function(response) {
        $("#formtip").css("color", "red").html(response.info);
        if (response.code == "success") {
            setTimeout(function() {
                window.location = "/manager/article";
            }, 200);
        } else {
            $("#save").removeAttr('disabled');
            $("#back").removeAttr('disabled');
        }
    }
});
$('#articleeditform').ajaxForm({
    dataType: "json",
    beforeSubmit: function() {
        if ($.trim($('#title').val()) == "") {
            $("#formtip").css("color", "red").html("请输入文章标题！");
            return false;
        }
        if ($.trim($("#content").val()) == "") {
            $("#formtip").css("color", "red").html("请输入文章内容！");
            return false
        }
        $("#formtip").removeClass('color').html('');
        $("#save").attr({
            "disabled": "disabled"
        });
        $("#back").attr({
            "disabled": "disabled"
        });
    },
    success: function(response) {
        $("#formtip").css("color", "red").html(response.info);
        if (response.code == "success") {
            setTimeout(function() {
                window.location = "/manager/article";
            }, 200);
        } else {
            $("#save").removeAttr('disabled');
            $("#back").removeAttr('disabled');
        }
    }
});
//删除文章
var deleteFunction = function(id) {
    $.ajax({
        type: 'POST',
        dataType: 'json',
        url: '/manager/article/del',
        data: {
            id: id
        },
        success: function(response) {
            var d = dialog({
                title: '提示 ： ',
                content: response.info
            });
            d.showModal();
            setTimeout(function() {
                d.close().remove();
            }, 2000);
            if (response.code == 'success') {
                setTimeout(function() {
                    window.location.reload();
                }, 2200);
            }
        }
    });
}
$('.delete').on('click', function() {
    var id = $(this).attr('data-id');
    var d = dialog({
        title: '提示：',
        content: '真的要删除这篇文章吗！',
        okValue: '确定',
        ok: function() {
            deleteFunction(id);
        },
        cancelValue: '取消',
        cancel: function() {}
    });
    d.showModal();
});
//从回收站恢复
var returnTrash = function(id) {
        $.ajax({
            url: '/manager/article/return',
            type: 'POST',
            dataType: 'json',
            data: {
                id: id
            },
            success: function(response) {
                var d = dialog({
                    title: '提示 ： ',
                    content: response.info
                });
                d.showModal();
                setTimeout(function() {
                    d.close().remove();
                }, 2000);
                if (response.code == 'success') {
                    setTimeout(function() {
                        window.location = "/manager/article/trash";
                    }, 2200);
                }
            }
        });
    }
    //
$('.return').on('click', function() {
    var id = $(this).attr('data-id');
    var d = dialog({
        title: '提示：',
        content: '恢复前请检查文章信息是否完整，恢复后文章将直接发布到互联网，确定要操作吗？',
        okValue: '确定',
        ok: function() {
            returnTrash(id);
        },
        cancelValue: '取消',
        cancel: function() {}
    });
    d.showModal();
});

var removeTrash = function(id) {
    $.ajax({
        url: '/manager/article/remove',
        type: 'POST',
        dataType: 'json',
        data: {
            id: id
        },
        success: function(response) {
            var d = dialog({
                title: '提示 ： ',
                content: response.info
            });
            d.showModal();
            setTimeout(function() {
                d.close().remove();
            }, 2000);
            if (response.code == 'success') {
                setTimeout(function() {
                    window.location = "/manager/article";
                }, 2200);
            }
        }
    });
}
$('.trash').on('click', function() {
    var id = $(this).attr('data-id');
    var d = dialog({
        title: '提示：',
        content: '在回收站的文章不会在网站显示！确定将这篇文章放到回收站吗？',
        okValue: '确定',
        ok: function() {
            removeTrash(id);
        },
        cancelValue: '取消',
        cancel: function() {}
    });
    d.showModal();
});
$('#categorycreateform').ajaxForm({
    dataType: 'json',
    beforeSubmit: function() {
        if ($.trim($('#name').val()) === '') {
            $('#formtip').css("color", "red").html("请输入分类名称!");
            return false;
        }
        if ($.trim($('#description').val()) === '') {
            $('#formtip').css("color", "red").html("请输入分类描述!");
            return false;
        }
        $('#formtip').removeClass('color').html('');
        $("#save").attr({
            "disabled": "disabled"
        });
        $("#back").attr({
            "disabled": "disabled"
        });
    },
    success: function(response) {
        $('#formtip').css("color", "red").html(response.info).addClass(response.code);
        if (response.code === 'success') {
            setTimeout(function() {
                window.location = "/manager/category";
            }, 2000)
        } else {
            $("#save").removeAttr("disabled");
            $("#back").removeAttr("disabled");
        }
    }
});
$('#categoryeditform').ajaxForm({
    dataType: 'json',
    beforeSubmit: function() {
        if ($.trim($('#name').val()) === '') {
            $('#formtip').css("color", "red").html("请输入分类名称!");
            return false;
        }
        if ($.trim($('#description').val()) === '') {
            $('#formtip').css("color", "red").html("请输入分类描述!");
            return false;
        }
        $('#formtip').removeClass('color').html('');
        $("#save").attr({
            "disabled": "disabled"
        });
        $("#back").attr({
            "disabled": "disabled"
        });
    },
    success: function(response) {
        $('#formtip').css("color", "red").html(response.info).addClass(response.code);
        if (response.code === 'success') {
            $("#save").attr({
                "disabled": "disabled"
            });
            $("#back").attr({
                "disabled": "disabled"
            });
            setTimeout(function() {
                window.location = "/manager/category";
            }, 2000)
        } else {
            $("#save").removeAttr("disabled");
            $("#back").removeAttr("disabled");
        }
    }
});
var deleteCreate = function(id) {
    $.ajax({
        type: 'POST',
        dataType: 'json',
        url: '/manager/category/del',
        data: {
            id: id
        },
        success: function(response) {
            var d = dialog({
                title: '提示 ： ',
                content: response.info
            });
            d.showModal();
            setTimeout(function() {
                d.close().remove();
            }, 2000);
            if (response.code == 'success') {
                setTimeout(function() {
                    window.location.reload();
                }, 2200);
            }
        }
    });
}
$('.delete').on('click', function() {
    var id = $(this).attr('data-id');
    var d = dialog({
        title: '提示：',
        content: '真的要删除这个分类吗！',
        okValue: '确定',
        ok: function() {
            deleteCreate(id);
        },
        cancelValue: '取消',
        cancel: function() {}
    });
    d.showModal();
});
$('#usereditform').ajaxForm({
    dataType: 'json',
    beforeSubmit: function() {
        if ($.trim($('#email').val()) == '') {
            $('#formtip').css("color", "red").html("请输入电子邮件!");
            return false;
        }
        if (!isEmail($.trim($('#email').val()))) {
            $('#formtip').css("color", "red").html("电子邮件格式不正确，请重新输入!");
            return false;
        }
        if ($.trim($('#name').val()) == '') {
            $('#formtip').css("color", "red").html("请输入用户姓名!");
            return false;
        }
        if ($.trim($('#password').val()) != '') {
            if ($.trim($('#password').val()).length < 6) {
                $('#formtip').css("color", "red").html("请输入 6-16 位用户密码!");
                return false;
            }
            if ($.trim($('#repassword').val()) == '') {
                $('#formtip').css("color", "red").html("请再次输入用户密码!");
                return false;
            }
            if ($.trim($('#password').val()) != $.trim($('#repassword').val())) {
                $('#formtip').css("color", "red").html("两次输入的密码不一致!")
                return false;
            }
        }
        if ($.trim($('#description').val()) == '') {
            $('#formtip').css("color", "red").html("请输入用户的简单介绍!");
            return false;
        }
        $('#formtip').removeClass('color').html('');
        $("#save").attr({
            "disabled": "disabled"
        });
        $("#back").attr({
            "disabled": "disabled"
        });
    },
    success: function(response) {
        $("#save").removeAttr("disabled");
        $("#back").removeAttr("disabled");
        $('#formtip').css("color", "red").html(response.info).addClass(response.code);
        var d = dialog({
            title: '提示 ： ',
            content: response.info
        });
        d.showModal();
        setTimeout(function() {
            d.close().remove();
        }, 2000);
        if (response.code === 'success') {
            $("#save").attr({
                "disabled": "disabled"
            });
            $("#back").attr({
                "disabled": "disabled"
            });
            setTimeout(function() {
                window.location = "{{ site_url('manager/user') }}";
            }, 2200)
        }
    }
});
});