package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["llgtorque"] = LLGTORQUE }

const LLGTORQUE = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00001b85_00000000-9_llgtorque.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/mx3/gpu/ptx/llgtorque.cu"
	.file	3 "/usr/local/cuda/bin/../include/device_functions.h"

.visible .entry llgtorque(
	.param .u64 llgtorque_param_0,
	.param .u64 llgtorque_param_1,
	.param .u64 llgtorque_param_2,
	.param .u64 llgtorque_param_3,
	.param .u64 llgtorque_param_4,
	.param .u64 llgtorque_param_5,
	.param .u64 llgtorque_param_6,
	.param .u64 llgtorque_param_7,
	.param .u64 llgtorque_param_8,
	.param .f32 llgtorque_param_9,
	.param .u32 llgtorque_param_10
)
{
	.reg .pred 	%p<2>;
	.reg .s32 	%r<18>;
	.reg .f32 	%f<35>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd10, [llgtorque_param_0];
	ld.param.u64 	%rd11, [llgtorque_param_1];
	ld.param.u64 	%rd12, [llgtorque_param_2];
	ld.param.u64 	%rd13, [llgtorque_param_3];
	ld.param.u64 	%rd14, [llgtorque_param_4];
	ld.param.u64 	%rd15, [llgtorque_param_5];
	ld.param.u64 	%rd16, [llgtorque_param_6];
	ld.param.u64 	%rd17, [llgtorque_param_7];
	ld.param.u64 	%rd18, [llgtorque_param_8];
	ld.param.f32 	%f1, [llgtorque_param_9];
	ld.param.u32 	%r2, [llgtorque_param_10];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd18;
	cvta.to.global.u64 	%rd5, %rd17;
	cvta.to.global.u64 	%rd6, %rd16;
	cvta.to.global.u64 	%rd7, %rd15;
	cvta.to.global.u64 	%rd8, %rd14;
	cvta.to.global.u64 	%rd9, %rd13;
	.loc 2 9 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 2 10 1
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	.loc 2 12 1
	mul.wide.s32 	%rd19, %r1, 4;
	add.s64 	%rd20, %rd9, %rd19;
	add.s64 	%rd21, %rd8, %rd19;
	add.s64 	%rd22, %rd7, %rd19;
	.loc 2 13 1
	add.s64 	%rd23, %rd6, %rd19;
	add.s64 	%rd24, %rd5, %rd19;
	add.s64 	%rd25, %rd4, %rd19;
	.loc 2 12 1
	ld.global.f32 	%f2, [%rd21];
	.loc 2 13 1
	ld.global.f32 	%f3, [%rd25];
	ld.global.f32 	%f4, [%rd24];
	.loc 2 12 1
	ld.global.f32 	%f5, [%rd22];
	.loc 2 15 1
	mul.f32 	%f6, %f5, %f4;
	mul.f32 	%f7, %f3, %f2;
	sub.f32 	%f8, %f6, %f7;
	.loc 2 13 1
	ld.global.f32 	%f9, [%rd23];
	.loc 2 12 1
	ld.global.f32 	%f10, [%rd20];
	.loc 2 15 1
	mul.f32 	%f11, %f10, %f3;
	mul.f32 	%f12, %f9, %f5;
	sub.f32 	%f13, %f11, %f12;
	mul.f32 	%f14, %f2, %f9;
	mul.f32 	%f15, %f4, %f10;
	sub.f32 	%f16, %f14, %f15;
	fma.rn.f32 	%f17, %f1, %f1, 0f3F800000;
	mov.f32 	%f18, 0fBF800000;
	.loc 3 2399 3
	div.rn.f32 	%f19, %f18, %f17;
	.loc 2 15 1
	mul.f32 	%f20, %f5, %f13;
	mul.f32 	%f21, %f16, %f2;
	sub.f32 	%f22, %f20, %f21;
	mul.f32 	%f23, %f10, %f16;
	mul.f32 	%f24, %f8, %f5;
	sub.f32 	%f25, %f23, %f24;
	mul.f32 	%f26, %f2, %f8;
	mul.f32 	%f27, %f13, %f10;
	sub.f32 	%f28, %f26, %f27;
	fma.rn.f32 	%f29, %f22, %f1, %f8;
	fma.rn.f32 	%f30, %f25, %f1, %f13;
	fma.rn.f32 	%f31, %f28, %f1, %f16;
	mul.f32 	%f32, %f19, %f29;
	mul.f32 	%f33, %f19, %f30;
	mul.f32 	%f34, %f19, %f31;
	.loc 2 17 1
	add.s64 	%rd26, %rd3, %rd19;
	st.global.f32 	[%rd26], %f32;
	.loc 2 18 1
	add.s64 	%rd27, %rd2, %rd19;
	st.global.f32 	[%rd27], %f33;
	.loc 2 19 1
	add.s64 	%rd28, %rd1, %rd19;
	st.global.f32 	[%rd28], %f34;

BB0_2:
	.loc 2 21 2
	ret;
}


`
