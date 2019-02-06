// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/type/color.proto

package google_type

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a color in the RGBA color space. This representation is designed
// for simplicity of conversion to/from color representations in various
// languages over compactness; for example, the fields of this representation
// can be trivially provided to the constructor of "java.awt.Color" in Java; it
// can also be trivially provided to UIColor's "+colorWithRed:green:blue:alpha"
// method in iOS; and, with just a little work, it can be easily formatted into
// a CSS "rgba()" string in JavaScript, as well. Here are some examples:
//
// Example (Java):
//
//      import com.google.type.Color;
//
//      // ...
//      public static java.awt.Color fromProto(Color protocolor) {
//        float alpha = protocolor.hasAlpha()
//            ? protocolor.getAlpha().getValue()
//            : 1.0;
//
//        return new java.awt.Color(
//            protocolor.getRed(),
//            protocolor.getGreen(),
//            protocolor.getBlue(),
//            alpha);
//      }
//
//      public static Color toProto(java.awt.Color color) {
//        float red = (float) color.getRed();
//        float green = (float) color.getGreen();
//        float blue = (float) color.getBlue();
//        float denominator = 255.0;
//        Color.Builder resultBuilder =
//            Color
//                .newBuilder()
//                .setRed(red / denominator)
//                .setGreen(green / denominator)
//                .setBlue(blue / denominator);
//        int alpha = color.getAlpha();
//        if (alpha != 255) {
//          result.setAlpha(
//              FloatValue
//                  .newBuilder()
//                  .setValue(((float) alpha) / denominator)
//                  .build());
//        }
//        return resultBuilder.build();
//      }
//      // ...
//
// Example (iOS / Obj-C):
//
//      // ...
//      static UIColor* fromProto(Color* protocolor) {
//         float red = [protocolor red];
//         float green = [protocolor green];
//         float blue = [protocolor blue];
//         FloatValue* alpha_wrapper = [protocolor alpha];
//         float alpha = 1.0;
//         if (alpha_wrapper != nil) {
//           alpha = [alpha_wrapper value];
//         }
//         return [UIColor colorWithRed:red green:green blue:blue alpha:alpha];
//      }
//
//      static Color* toProto(UIColor* color) {
//          CGFloat red, green, blue, alpha;
//          if (![color getRed:&red green:&green blue:&blue alpha:&alpha]) {
//            return nil;
//          }
//          Color* result = [Color alloc] init];
//          [result setRed:red];
//          [result setGreen:green];
//          [result setBlue:blue];
//          if (alpha <= 0.9999) {
//            [result setAlpha:floatWrapperWithValue(alpha)];
//          }
//          [result autorelease];
//          return result;
//     }
//     // ...
//
//  Example (JavaScript):
//
//     // ...
//
//     var protoToCssColor = function(rgb_color) {
//        var redFrac = rgb_color.red || 0.0;
//        var greenFrac = rgb_color.green || 0.0;
//        var blueFrac = rgb_color.blue || 0.0;
//        var red = Math.floor(redFrac * 255);
//        var green = Math.floor(greenFrac * 255);
//        var blue = Math.floor(blueFrac * 255);
//
//        if (!('alpha' in rgb_color)) {
//           return rgbToCssColor_(red, green, blue);
//        }
//
//        var alphaFrac = rgb_color.alpha.value || 0.0;
//        var rgbParams = [red, green, blue].join(',');
//        return ['rgba(', rgbParams, ',', alphaFrac, ')'].join('');
//     };
//
//     var rgbToCssColor_ = function(red, green, blue) {
//       var rgbNumber = new Number((red << 16) | (green << 8) | blue);
//       var hexString = rgbNumber.toString(16);
//       var missingZeros = 6 - hexString.length;
//       var resultBuilder = ['#'];
//       for (var i = 0; i < missingZeros; i++) {
//          resultBuilder.push('0');
//       }
//       resultBuilder.push(hexString);
//       return resultBuilder.join('');
//     };
//
//     // ...
type Color struct {
	// The amount of red in the color as a value in the interval [0, 1].
	Red float32 `protobuf:"fixed32,1,opt,name=red" json:"red,omitempty"`
	// The amount of green in the color as a value in the interval [0, 1].
	Green float32 `protobuf:"fixed32,2,opt,name=green" json:"green,omitempty"`
	// The amount of blue in the color as a value in the interval [0, 1].
	Blue float32 `protobuf:"fixed32,3,opt,name=blue" json:"blue,omitempty"`
	// The fraction of this color that should be applied to the pixel. That is,
	// the final pixel color is defined by the equation:
	//
	//   pixel color = alpha * (this color) + (1.0 - alpha) * (background color)
	//
	// This means that a value of 1.0 corresponds to a solid color, whereas
	// a value of 0.0 corresponds to a completely transparent color. This
	// uses a wrapper message rather than a simple float scalar so that it is
	// possible to distinguish between a default value and the value being unset.
	// If omitted, this color object is to be rendered as a solid color
	// (as if the alpha value had been explicitly given with a value of 1.0).
	Alpha *google_protobuf.FloatValue `protobuf:"bytes,4,opt,name=alpha" json:"alpha,omitempty"`
}

func (m *Color) Reset()                    { *m = Color{} }
func (m *Color) String() string            { return proto.CompactTextString(m) }
func (*Color) ProtoMessage()               {}
func (*Color) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Color) GetRed() float32 {
	if m != nil {
		return m.Red
	}
	return 0
}

func (m *Color) GetGreen() float32 {
	if m != nil {
		return m.Green
	}
	return 0
}

func (m *Color) GetBlue() float32 {
	if m != nil {
		return m.Blue
	}
	return 0
}

func (m *Color) GetAlpha() *google_protobuf.FloatValue {
	if m != nil {
		return m.Alpha
	}
	return nil
}

func init() {
	proto.RegisterType((*Color)(nil), "google.type.Color")
}

func init() { proto.RegisterFile("google/type/color.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0xce, 0xcf, 0xc9, 0x2f, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x48, 0xe8, 0x81, 0x24, 0xa4, 0xe4, 0xa0, 0xaa, 0xc0, 0x52,
	0x49, 0xa5, 0x69, 0xfa, 0xe5, 0x45, 0x89, 0x05, 0x05, 0xa9, 0x45, 0xc5, 0x10, 0xc5, 0x4a, 0x65,
	0x5c, 0xac, 0xce, 0x20, 0xbd, 0x42, 0x02, 0x5c, 0xcc, 0x45, 0xa9, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x4c, 0x41, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x7a, 0x51, 0x6a, 0x6a, 0x9e, 0x04, 0x13,
	0x58, 0x0c, 0xc2, 0x11, 0x12, 0xe2, 0x62, 0x49, 0xca, 0x29, 0x4d, 0x95, 0x60, 0x06, 0x0b, 0x82,
	0xd9, 0x42, 0x86, 0x5c, 0xac, 0x89, 0x39, 0x05, 0x19, 0x89, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0xd2, 0x7a, 0x50, 0x17, 0xc0, 0x2c, 0xd5, 0x73, 0xcb, 0xc9, 0x4f, 0x2c, 0x09, 0x4b, 0xcc,
	0x29, 0x4d, 0x0d, 0x82, 0xa8, 0x74, 0x52, 0xe5, 0xe2, 0x4f, 0xce, 0xcf, 0xd5, 0x43, 0x72, 0xaa,
	0x13, 0x17, 0xd8, 0x21, 0x01, 0x20, 0x3d, 0x01, 0x8c, 0x8b, 0x98, 0x98, 0xdd, 0x43, 0x02, 0x92,
	0xd8, 0xc0, 0x46, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x62, 0x47, 0x62, 0x2a, 0xed, 0x00,
	0x00, 0x00,
}
