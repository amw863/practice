/*
 * @Author : wb
 * @Date : 2022/2/3 11:45 上午
 */

package code

// 建造者模式解决参数过多和特殊要求的问题，golang function-options 模式也支持
type Builder struct {
	A string
	B string
}

func (b *Builder) Build() {

}
