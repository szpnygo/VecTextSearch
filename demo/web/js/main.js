$(document).ready(function () {
    // 搜索按钮点击事件
    $("#search-btn").on("click", function () {
        const searchText = $("#search-text").val().trim();

        if (searchText.length === 0) {
            M.toast({html: "请输入查询内容"});
            return;
        }

        // 显示加载动画
        $("#loading").show();

        // 发送API请求
        $.ajax({
            url: "http://127.0.0.1:8000/search-similar-texts",
            method: "POST",
            contentType: "application/json",
            data: JSON.stringify({
                content: searchText
            }),
            success: function (response) {
                // 隐藏加载动画
                $("#loading").hide();

                if (!response || !Array.isArray(response)) {
                    M.toast({html: "搜索结果有误，请稍后再试"});
                    return;
                }

                // 清空搜索结果表格
                $("#results-container").empty();

                // 显示搜索结果
                response.forEach(function (item) {
                    const card = `
                        <div class="result-card">
                            <div class="content">${item.content}</div>
                            <div class="progress">
                                <div class="determinate" style="width: ${(item.certainty * 100).toFixed(2)}%;"></div>
                            </div>
                            <div class="similarity">${(item.certainty * 100).toFixed(2)}%</div>
                        </div>
                    `;
                    $("#results-container").append(card);
                });

                // 显示搜索结果表格
                $("#results").show();
            },
            error: function (error) {
                // 隐藏加载动画
                $("#loading").hide();

                // 显示错误信息
                M.toast({html: "搜索失败，请重试"});
            }
        });
    });
});

$(".preset-btn").on("click", function () {
    const presetQuery = $(this).text();
    $("#search-text").val(presetQuery).focus();;
});
