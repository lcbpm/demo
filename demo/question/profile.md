# 修改环境变量的步骤

要在 Mac 上修改环境变量，可以按照以下步骤操作：

1. **打开终端**
    - 您可以在 Spotlight 中搜索“终端”或者在“应用程序” > “实用工具”中找到它。

2. **编辑或创建环境变量**
    - 使用 `nano` 或 `vi` 等文本编辑器打开或创建 `.bash_profile` 或 `.zshrc` 文件。
    - 例如，使用 `nano ~/.bash_profile`。
    - 在文件中添加环境变量。例如，添加 `export MY_VAR="my_value"`。

3. **保存并应用更改**
    - 如果使用 `nano`，可以使用 `Ctrl + O` 保存文件，然后按 `Enter` 关闭编辑器。
    - 使用 `source ~/.bash_profile` 或 `source ~/.zshrc` 使更改生效。

这样，您的环境变量就会被修改并应用到当前的终端会话中。
