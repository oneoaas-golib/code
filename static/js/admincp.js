function control_nav(id){
    $(".bs-docs-sidenav").children("li").eq(id).toggleClass("active");
}
/**
 * 显示添加标签的对话框
 * @param  {object} ele 当前点击的对象
 * @param  {int}    id  文章的id
 * @return {[type]}     [description]
 */
function createArticleTag(ele, id)
{
    $(ele).css('color', '#D60000');
    var str = '<h5 class="dialog-item-title">输入标签(2至8个字符)</h5><br/>';
        str += '<input type="text" id="tag" name="name" class="input-item-text" autofocus />';
    var d = dialog({
        title: '添加标签 ： ',
        content: str,
        okValue: '确定',
        ok: function () {
            name = $('#tag').val();
            insertArticleTag(ele, id, name);
        },
        cancelValue: '取消',
        cancel: function(){
            $(ele).css('color', '#666');
        }
    });
    d.showModal();
}

/**
 * 将标签添加到数据库
 * @param  {object} ele  当前点击的对象
 * @param  {int}    id   文章id
 * @param  {string} name 标签名称
 * @return {[type]}      [description]
 */
function insertArticleTag(ele, id, name)
{
    $('#tips').html('正在添加...').css('color', '#666');
    $.ajax({
        type: 'POST',
        dataType: 'json',
        data: { id : id, name : name},
        url: '/manager/tag/createAction',
        success: function(response){
            var d = dialog({
                title: '提示 ： ',
                content: response.info
            });
            d.showModal();
            setTimeout(function(){
                d.close().remove();
            }, 2000);
            if ( response.code == 'success' ) {
                var str  = '\r\n<em class="input-item-submit">' + response.data.name + '<a href="javascript:;" title="从该文章中移去该标签"';
                str += 'onClick="deleteArticleTag(this, ' + id + ',' + response.data.id + ')">×</a></em>\r\n';
                $(ele).before(str).css('color', '#666');
            }
        }
    });
}

/**
 * 删除标签的请求
 * @param  {object} ele 当前点击的对象
 * @param  {int}    id  标签的id
 * @return {[type]}     [description]
 */
function deleteTagAction(ele, aid, id)
{
    $.ajax({
        type: 'POST',
        dataType: 'json',
        data: { aid : aid, id : id },
        url: '/manager/tag/deleteAction',
        success: function(response){
            var d = dialog({
                title: '提示 ： ',
                content: response.info
            });
            d.showModal();
            setTimeout(function(){
                d.close().remove();
            },2000);
            if ( response.code == 'success' ) {
                $(ele).closest('em').remove();
            }
        }
    });
}

/**
 * 显示删除标签的对话框
 * @param  {object} ele 当前点击的对象
 * @param  {int}    id  标签的id
 * @return {[type]}     [description]
 */
function deleteArticleTag(ele, aid, id) {
    var d = dialog({
        title: '提示 ： ',
        content: '真的要将该标签从该文章删除吗？',
        okValue: '确定',
        ok: function(){
            deleteTagAction(ele, aid, id);
        },
        cancelValue: '取消',
        cancel: function(){}
    });
    d.showModal();
}

/**
 * 上传图片
 * @param  {object} ele 当前点击的对象
 * @return {[type]}     [description]
 */
function uploadAttachAction(ele) {
    $(ele).closest('form').ajaxSubmit({
        url: '/upload/image',
        dataType: 'json',
        type: 'POST',
        beforeSubmit: function() {
            $('.upload-attach-wait').show();
            $('#upload-attach-butn').addClass('button-disabled');
            $('#upload-attach-tips').html('正在上传，请稍候...').css('color', '#666');
        },
        success: function(response) {
            var d = dialog({
                title: '提示 ： ',
                content: response.info
            });
            d.showModal();
            setTimeout(function () {
                d.close().remove();
            }, 2000);
            if( response.code === 'success' ) {
                var num = $('#upload-attach-data').find('li').length;
                var str = '<li class="upload-item-' + (num + 1) + ' clearfix">';
                    str += '<img src="' + response.data.file_path + '" class="imageview fl" title="' + response.data.file_name + '" />';
                    str += '<div class="imageinfo fl">';
                    str += '<h5>' + response.data.file_name + '</h5>';
                    str += '<p><a href="javascript:;" data-nums="0" onClick="insertAttachImage(this, \'' + response.data.file_name + '\')">插入到文章</a>';
                    str += '&nbsp;&nbsp;&nbsp;<a href="javascript:;" onClick="deleteAttachImage(this, \'' + response.data.file_name + '\')">删除</a></p>';
                    str += '</div>';
                    str += '</li>';
                $('#upload-attach-data').append(str);
            }
            $('.upload-attach-wait').hide();
            $('#upload-attach-butn').removeClass('button-disabled');
            $('#upload-attach-tips').html(response.info).css('color', (response.code === 'success') ? 'green' : '#d60000');
        }
    });
}

/**
 * 将上传的文章图片插入到文章里面
 * @param  {object} ele   当前点击的对象
 * @param  {string} image 图片的名字
 * @return {[type]}       [description]
 */
function insertAttachImage(ele, image) {
    var nums = $(ele).attr('data-nums');
    if( nums > 0 ) {
        var d = dialog({
            title: '提示：',
            content: '该图片已经存在文章中，还要再插入一次吗？',
            okValue: '确定',
            ok: function () {
                moveInsertAttachImage(image);
            },
            cancelValue: '取消',
            cancel: function () {}
        });
        d.showModal();
    } else {
        moveInsertAttachImage(image);
        $(ele).closest('li').addClass('inserted');
    }
    $(ele).attr('data-nums', parseInt(nums) + 1);
}

/**
 * 移动上传的文章图片到文章的文件夹
 * @param  {string} image 图片的名字
 * @return {[type]}       [description]
 */
function moveInsertAttachImage(image)
{
    $.ajax({
        type: 'POST',
        url: '/manager/article/moveImageAction',
        data: { filename: image },
        dataType: 'json',
        success: function(response)
        {
            var d = dialog({
                title: '提示 ： ',
                content: response.info
            });
            d.showModal();
            setTimeout(function () {
                d.close().remove();
            }, 2000);
            if ( response.code == 'success' ) {
                editor.execCommand('insertImage', { src : response.data.file_path });
            }
        }
    });
}

/**
 * 删除上传的文章图片
 * @param  {object} ele   当前点击的对象
 * @param  {string} image 图片的名称
 * @return {[type]}       [description]
 */
function deleteAttachImage(ele, image) {
    var nums = $(ele).prev('a').attr('data-nums');
    if( nums > 0 ) {
        var d = dialog({
            title: '提示：',
            content: '该图片已经插入到文章中，如果删除之后在文章中使用将无法看到图片，确定要删除吗？',
            okValue: '确定',
            ok: function () {
                deleteImage(ele, image);
            },
            cancelValue: '取消',
            cancel: function () {}
        });
        d.showModal();
    } else {
        deleteImage(ele, image);
    }
}

/**
 * 处理删除的请求，真正的删除
 * @param  {object} ele   当前点击的对象
 * @param  {string} image 图片的名称
 * @return {[type]}       [description]
 */
function deleteImage(ele, image) {
    $(ele).html('删除中...');
    $.getJSON( '/manager/article/deleteImageAction', { filename: image }, function(response) {
        var d = dialog({
            title: '提示 ： ',
            content: response.info
        }).showModal();
        setTimeout(function () {
            d.close().remove();
        }, 2000);
        if( response.code == 'success' ) {
            $(ele).closest('li').remove();
            $('#upload-attach-tips').html('请重新上传！');
        }
        $(ele).html('删除');
    });
}