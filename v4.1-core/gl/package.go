// Copyright (c) 2010 Khronos Group.
// This material may be distributed subject to the terms and conditions
// set forth in the Open Publication License, v 1.0, 8 June 1999.
// http://opencontent.org/openpub/.
//
// Copyright (c) 1991-2006 Silicon Graphics, Inc.
// This document is licensed under the SGI Free Software B License.
// For details, see http://oss.sgi.com/projects/FreeB.

// Package gl implements Go bindings to OpenGL.
//
// This package was automatically generated using Glow:
//  http://github.com/go-gl/glow
//
// Generated based on the OpenGL XML specification:
//  SVN revision 27695
package gl

// #cgo darwin  LDFLAGS: -framework OpenGL
// #cgo linux   LDFLAGS: -lGL
// #cgo windows LDFLAGS: -lopengl32
// #if defined(_WIN32) && !defined(APIENTRY) && !defined(__CYGWIN__) && !defined(__SCITECH_SNAP__)
// #ifndef WIN32_LEAN_AND_MEAN
// #define WIN32_LEAN_AND_MEAN 1
// #endif
// #include <windows.h>
// #endif
// #ifndef APIENTRY
// #define APIENTRY
// #endif
// #ifndef APIENTRYP
// #define APIENTRYP APIENTRY *
// #endif
// #ifndef GLAPI
// #define GLAPI extern
// #endif
// #include <stddef.h>
// #ifndef GLEXT_64_TYPES_DEFINED
// /* This code block is duplicated in glxext.h, so must be protected */
// #define GLEXT_64_TYPES_DEFINED
// /* Define int32_t, int64_t, and uint64_t types for UST/MSC */
// /* (as used in the GL_EXT_timer_query extension). */
// #if defined(__STDC_VERSION__) && __STDC_VERSION__ >= 199901L
// #include <inttypes.h>
// #elif defined(__sun__) || defined(__digital__)
// #include <inttypes.h>
// #if defined(__STDC__)
// #if defined(__arch64__) || defined(_LP64)
// typedef long int int64_t;
// typedef unsigned long int uint64_t;
// #else
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #endif /* __arch64__ */
// #endif /* __STDC__ */
// #elif defined( __VMS ) || defined(__sgi)
// #include <inttypes.h>
// #elif defined(__SCO__) || defined(__USLC__)
// #include <stdint.h>
// #elif defined(__UNIXOS2__) || defined(__SOL64__)
// typedef long int int32_t;
// typedef long long int int64_t;
// typedef unsigned long long int uint64_t;
// #elif defined(_WIN32) && defined(__GNUC__)
// #include <stdint.h>
// #elif defined(_WIN32)
// typedef __int32 int32_t;
// typedef __int64 int64_t;
// typedef unsigned __int64 uint64_t;
// #else
// /* Fallback if nothing above works */
// #include <inttypes.h>
// #endif
// #endif
// typedef unsigned int GLenum;
// typedef unsigned char GLboolean;
// typedef unsigned int GLbitfield;
// typedef void GLvoid;
// typedef signed char GLbyte;
// typedef short GLshort;
// typedef int GLint;
// typedef int GLclampx;
// typedef unsigned char GLubyte;
// typedef unsigned short GLushort;
// typedef unsigned int GLuint;
// typedef int GLsizei;
// typedef float GLfloat;
// typedef float GLclampf;
// typedef double GLdouble;
// typedef char GLchar;
// typedef GLint GLfixed;
// typedef ptrdiff_t GLintptr;
// typedef ptrdiff_t GLsizeiptr;
// typedef int64_t GLint64;
// typedef uint64_t GLuint64;
// typedef uint64_t GLuint64EXT;
// typedef struct __GLsync *GLsync;
// struct _cl_context;
// struct _cl_event;
// typedef void (APIENTRY *GLDEBUGPROC)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCARB)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// typedef void (APIENTRY *GLDEBUGPROCKHR)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
// extern void glowDebugCallback_glcore41(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam);
// static void APIENTRY glowCDebugCallback(GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar* message, const void* userParam) {
//   glowDebugCallback_glcore41(source, type, id, severity, length, message, userParam);
// }
// typedef void  (APIENTRYP GPACCUMXOES)(GLenum  op, GLfixed  value);
// typedef void  (APIENTRYP GPACTIVEPROGRAMEXT)(GLuint  program);
// typedef void  (APIENTRYP GPACTIVESHADERPROGRAM)(GLuint  pipeline, GLuint  program);
// typedef void  (APIENTRYP GPACTIVESHADERPROGRAMEXT)(GLuint  pipeline, GLuint  program);
// typedef void  (APIENTRYP GPACTIVETEXTURE)(GLenum  texture);
// typedef void  (APIENTRYP GPALPHAFUNCXOES)(GLenum  func, GLfixed  ref);
// typedef void  (APIENTRYP GPATTACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPBEGINCONDITIONALRENDER)(GLuint  id, GLenum  mode);
// typedef void  (APIENTRYP GPBEGINPERFMONITORAMD)(GLuint  monitor);
// typedef void  (APIENTRYP GPBEGINPERFQUERYINTEL)(GLuint  queryHandle);
// typedef void  (APIENTRYP GPBEGINQUERY)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBEGINQUERYINDEXED)(GLenum  target, GLuint  index, GLuint  id);
// typedef void  (APIENTRYP GPBEGINTRANSFORMFEEDBACK)(GLenum  primitiveMode);
// typedef void  (APIENTRYP GPBINDATTRIBLOCATION)(GLuint  program, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDBUFFER)(GLenum  target, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERBASE)(GLenum  target, GLuint  index, GLuint  buffer);
// typedef void  (APIENTRYP GPBINDBUFFERRANGE)(GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPBINDBUFFERSBASE)(GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers);
// typedef void  (APIENTRYP GPBINDBUFFERSRANGE)(GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizeiptr * sizes);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATION)(GLuint  program, GLuint  color, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAGDATALOCATIONINDEXED)(GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name);
// typedef void  (APIENTRYP GPBINDFRAMEBUFFER)(GLenum  target, GLuint  framebuffer);
// typedef void  (APIENTRYP GPBINDIMAGETEXTURE)(GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format);
// typedef void  (APIENTRYP GPBINDIMAGETEXTURES)(GLuint  first, GLsizei  count, const GLuint * textures);
// typedef void  (APIENTRYP GPBINDPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPBINDPROGRAMPIPELINEEXT)(GLuint  pipeline);
// typedef void  (APIENTRYP GPBINDRENDERBUFFER)(GLenum  target, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPBINDSAMPLER)(GLuint  unit, GLuint  sampler);
// typedef void  (APIENTRYP GPBINDSAMPLERS)(GLuint  first, GLsizei  count, const GLuint * samplers);
// typedef void  (APIENTRYP GPBINDTEXTURE)(GLenum  target, GLuint  texture);
// typedef void  (APIENTRYP GPBINDTEXTUREUNIT)(GLuint  unit, GLuint  texture);
// typedef void  (APIENTRYP GPBINDTEXTURES)(GLuint  first, GLsizei  count, const GLuint * textures);
// typedef void  (APIENTRYP GPBINDTRANSFORMFEEDBACK)(GLenum  target, GLuint  id);
// typedef void  (APIENTRYP GPBINDVERTEXARRAY)(GLuint  array);
// typedef void  (APIENTRYP GPBINDVERTEXBUFFER)(GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride);
// typedef void  (APIENTRYP GPBINDVERTEXBUFFERS)(GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides);
// typedef void  (APIENTRYP GPBITMAPXOES)(GLsizei  width, GLsizei  height, GLfixed  xorig, GLfixed  yorig, GLfixed  xmove, GLfixed  ymove, const GLubyte * bitmap);
// typedef void  (APIENTRYP GPBLENDBARRIERKHR)();
// typedef void  (APIENTRYP GPBLENDBARRIERNV)();
// typedef void  (APIENTRYP GPBLENDCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPBLENDCOLORXOES)(GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha);
// typedef void  (APIENTRYP GPBLENDEQUATION)(GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONEXT)(GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATE)(GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEI)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONSEPARATEIARB)(GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha);
// typedef void  (APIENTRYP GPBLENDEQUATIONI)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDEQUATIONIARB)(GLuint  buf, GLenum  mode);
// typedef void  (APIENTRYP GPBLENDFUNC)(GLenum  sfactor, GLenum  dfactor);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATE)(GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEI)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCSEPARATEIARB)(GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha);
// typedef void  (APIENTRYP GPBLENDFUNCI)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLENDFUNCIARB)(GLuint  buf, GLenum  src, GLenum  dst);
// typedef void  (APIENTRYP GPBLENDPARAMETERINV)(GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPBLITFRAMEBUFFER)(GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBLITNAMEDFRAMEBUFFER)(GLuint  readFramebuffer, GLuint  drawFramebuffer, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter);
// typedef void  (APIENTRYP GPBUFFERDATA)(GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPBUFFERPAGECOMMITMENTARB)(GLenum  target, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPBUFFERSTORAGE)(GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags);
// typedef void  (APIENTRYP GPBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data);
// typedef GLenum  (APIENTRYP GPCHECKFRAMEBUFFERSTATUS)(GLenum  target);
// typedef GLenum  (APIENTRYP GPCHECKNAMEDFRAMEBUFFERSTATUS)(GLuint  framebuffer, GLenum  target);
// typedef void  (APIENTRYP GPCLAMPCOLOR)(GLenum  target, GLenum  clamp);
// typedef void  (APIENTRYP GPCLEAR)(GLbitfield  mask);
// typedef void  (APIENTRYP GPCLEARACCUMXOES)(GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha);
// typedef void  (APIENTRYP GPCLEARBUFFERDATA)(GLenum  target, GLenum  internalformat, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARBUFFERSUBDATA)(GLenum  target, GLenum  internalformat, GLintptr  offset, GLsizeiptr  size, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARBUFFERFI)(GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil);
// typedef void  (APIENTRYP GPCLEARBUFFERFV)(GLenum  buffer, GLint  drawbuffer, const GLfloat * value);
// typedef void  (APIENTRYP GPCLEARBUFFERIV)(GLenum  buffer, GLint  drawbuffer, const GLint * value);
// typedef void  (APIENTRYP GPCLEARBUFFERUIV)(GLenum  buffer, GLint  drawbuffer, const GLuint * value);
// typedef void  (APIENTRYP GPCLEARCOLOR)(GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha);
// typedef void  (APIENTRYP GPCLEARCOLORXOES)(GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha);
// typedef void  (APIENTRYP GPCLEARDEPTH)(GLdouble  depth);
// typedef void  (APIENTRYP GPCLEARDEPTHF)(GLfloat  d);
// typedef void  (APIENTRYP GPCLEARDEPTHFOES)(GLclampf  depth);
// typedef void  (APIENTRYP GPCLEARDEPTHXOES)(GLfixed  depth);
// typedef void  (APIENTRYP GPCLEARNAMEDBUFFERDATA)(GLuint  buffer, GLenum  internalformat, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARNAMEDBUFFERSUBDATA)(GLuint  buffer, GLenum  internalformat, GLintptr  offset, GLsizei  size, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERFI)(GLuint  framebuffer, GLenum  buffer, const GLfloat  depth, GLint  stencil);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERFV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLfloat * value);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERIV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLint * value);
// typedef void  (APIENTRYP GPCLEARNAMEDFRAMEBUFFERUIV)(GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLuint * value);
// typedef void  (APIENTRYP GPCLEARSTENCIL)(GLint  s);
// typedef void  (APIENTRYP GPCLEARTEXIMAGE)(GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data);
// typedef void  (APIENTRYP GPCLEARTEXSUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data);
// typedef GLenum  (APIENTRYP GPCLIENTWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// typedef void  (APIENTRYP GPCLIPCONTROL)(GLenum  origin, GLenum  depth);
// typedef void  (APIENTRYP GPCLIPPLANEFOES)(GLenum  plane, const GLfloat * equation);
// typedef void  (APIENTRYP GPCLIPPLANEXOES)(GLenum  plane, const GLfixed * equation);
// typedef void  (APIENTRYP GPCOLOR3XOES)(GLfixed  red, GLfixed  green, GLfixed  blue);
// typedef void  (APIENTRYP GPCOLOR3XVOES)(const GLfixed * components);
// typedef void  (APIENTRYP GPCOLOR4XOES)(GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha);
// typedef void  (APIENTRYP GPCOLOR4XVOES)(const GLfixed * components);
// typedef void  (APIENTRYP GPCOLORMASK)(GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha);
// typedef void  (APIENTRYP GPCOLORMASKI)(GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a);
// typedef void  (APIENTRYP GPCOMPILESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPCOMPILESHADERINCLUDEARB)(GLuint  shader, GLsizei  count, const GLchar *const* path, const GLint * length);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE1D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXIMAGE3D)(GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCOMPRESSEDTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data);
// typedef void  (APIENTRYP GPCONVOLUTIONPARAMETERXOES)(GLenum  target, GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPCONVOLUTIONPARAMETERXVOES)(GLenum  target, GLenum  pname, const GLfixed * params);
// typedef void  (APIENTRYP GPCOPYBUFFERSUBDATA)(GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPCOPYIMAGESUBDATA)(GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth);
// typedef void  (APIENTRYP GPCOPYNAMEDBUFFERSUBDATA)(GLuint  readBuffer, GLuint  writeBuffer, GLintptr  readOffset, GLintptr  writeOffset, GLsizei  size);
// typedef void  (APIENTRYP GPCOPYTEXIMAGE1D)(GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLint  border);
// typedef void  (APIENTRYP GPCOPYTEXIMAGE2D)(GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCOPYTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPCREATEBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPCREATEFRAMEBUFFERS)(GLsizei  n, GLuint * framebuffers);
// typedef void  (APIENTRYP GPCREATEPERFQUERYINTEL)(GLuint  queryId, GLuint * queryHandle);
// typedef GLuint  (APIENTRYP GPCREATEPROGRAM)();
// typedef void  (APIENTRYP GPCREATEPROGRAMPIPELINES)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPCREATEQUERIES)(GLenum  target, GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPCREATERENDERBUFFERS)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPCREATESAMPLERS)(GLsizei  n, GLuint * samplers);
// typedef GLuint  (APIENTRYP GPCREATESHADER)(GLenum  type);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMEXT)(GLenum  type, const GLchar * string);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMV)(GLenum  type, GLsizei  count, const GLchar *const* strings);
// typedef GLuint  (APIENTRYP GPCREATESHADERPROGRAMVEXT)(GLenum  type, GLsizei  count, const GLchar ** strings);
// typedef GLsync  (APIENTRYP GPCREATESYNCFROMCLEVENTARB)(struct _cl_context * context, struct _cl_event * event, GLbitfield  flags);
// typedef void  (APIENTRYP GPCREATETEXTURES)(GLenum  target, GLsizei  n, GLuint * textures);
// typedef void  (APIENTRYP GPCREATETRANSFORMFEEDBACKS)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPCREATEVERTEXARRAYS)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPCULLFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACK)(GLDEBUGPROC  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACKARB)(GLDEBUGPROCARB  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECALLBACKKHR)(GLDEBUGPROCKHR  callback, const void * userParam);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROL)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROLARB)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGECONTROLKHR)(GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERT)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERTARB)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDEBUGMESSAGEINSERTKHR)(GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf);
// typedef void  (APIENTRYP GPDELETEBUFFERS)(GLsizei  n, const GLuint * buffers);
// typedef void  (APIENTRYP GPDELETEFENCESNV)(GLsizei  n, const GLuint * fences);
// typedef void  (APIENTRYP GPDELETEFRAMEBUFFERS)(GLsizei  n, const GLuint * framebuffers);
// typedef void  (APIENTRYP GPDELETENAMEDSTRINGARB)(GLint  namelen, const GLchar * name);
// typedef void  (APIENTRYP GPDELETEPERFMONITORSAMD)(GLsizei  n, GLuint * monitors);
// typedef void  (APIENTRYP GPDELETEPERFQUERYINTEL)(GLuint  queryHandle);
// typedef void  (APIENTRYP GPDELETEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPDELETEPROGRAMPIPELINES)(GLsizei  n, const GLuint * pipelines);
// typedef void  (APIENTRYP GPDELETEPROGRAMPIPELINESEXT)(GLsizei  n, const GLuint * pipelines);
// typedef void  (APIENTRYP GPDELETEQUERIES)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETERENDERBUFFERS)(GLsizei  n, const GLuint * renderbuffers);
// typedef void  (APIENTRYP GPDELETESAMPLERS)(GLsizei  count, const GLuint * samplers);
// typedef void  (APIENTRYP GPDELETESHADER)(GLuint  shader);
// typedef void  (APIENTRYP GPDELETESYNC)(GLsync  sync);
// typedef void  (APIENTRYP GPDELETETEXTURES)(GLsizei  n, const GLuint * textures);
// typedef void  (APIENTRYP GPDELETETRANSFORMFEEDBACKS)(GLsizei  n, const GLuint * ids);
// typedef void  (APIENTRYP GPDELETEVERTEXARRAYS)(GLsizei  n, const GLuint * arrays);
// typedef void  (APIENTRYP GPDEPTHFUNC)(GLenum  func);
// typedef void  (APIENTRYP GPDEPTHMASK)(GLboolean  flag);
// typedef void  (APIENTRYP GPDEPTHRANGE)(GLdouble  xnear, GLdouble  xfar);
// typedef void  (APIENTRYP GPDEPTHRANGEARRAYV)(GLuint  first, GLsizei  count, const GLdouble * v);
// typedef void  (APIENTRYP GPDEPTHRANGEINDEXED)(GLuint  index, GLdouble  n, GLdouble  f);
// typedef void  (APIENTRYP GPDEPTHRANGEF)(GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPDEPTHRANGEFOES)(GLclampf  n, GLclampf  f);
// typedef void  (APIENTRYP GPDEPTHRANGEXOES)(GLfixed  n, GLfixed  f);
// typedef void  (APIENTRYP GPDETACHSHADER)(GLuint  program, GLuint  shader);
// typedef void  (APIENTRYP GPDISABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPDISABLEVERTEXARRAYATTRIB)(GLuint  vaobj, GLuint  index);
// typedef void  (APIENTRYP GPDISABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPDISABLEI)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTE)(GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTEGROUPSIZEARB)(GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z, GLuint  group_size_x, GLuint  group_size_y, GLuint  group_size_z);
// typedef void  (APIENTRYP GPDISPATCHCOMPUTEINDIRECT)(GLintptr  indirect);
// typedef void  (APIENTRYP GPDRAWARRAYS)(GLenum  mode, GLint  first, GLsizei  count);
// typedef void  (APIENTRYP GPDRAWARRAYSINDIRECT)(GLenum  mode, const void * indirect);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCED)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDBASEINSTANCE)(GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWARRAYSINSTANCEDEXT)(GLenum  mode, GLint  start, GLsizei  count, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWBUFFER)(GLenum  buf);
// typedef void  (APIENTRYP GPDRAWBUFFERS)(GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPDRAWELEMENTS)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWELEMENTSBASEVERTEX)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINDIRECT)(GLenum  mode, GLenum  type, const void * indirect);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCED)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEINSTANCE)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEX)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance);
// typedef void  (APIENTRYP GPDRAWELEMENTSINSTANCEDEXT)(GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTS)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices);
// typedef void  (APIENTRYP GPDRAWRANGEELEMENTSBASEVERTEX)(GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACK)(GLenum  mode, GLuint  id);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKINSTANCED)(GLenum  mode, GLuint  id, GLsizei  instancecount);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKSTREAM)(GLenum  mode, GLuint  id, GLuint  stream);
// typedef void  (APIENTRYP GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED)(GLenum  mode, GLuint  id, GLuint  stream, GLsizei  instancecount);
// typedef void  (APIENTRYP GPENABLE)(GLenum  cap);
// typedef void  (APIENTRYP GPENABLEVERTEXARRAYATTRIB)(GLuint  vaobj, GLuint  index);
// typedef void  (APIENTRYP GPENABLEVERTEXATTRIBARRAY)(GLuint  index);
// typedef void  (APIENTRYP GPENABLEI)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENDCONDITIONALRENDER)();
// typedef void  (APIENTRYP GPENDPERFMONITORAMD)(GLuint  monitor);
// typedef void  (APIENTRYP GPENDPERFQUERYINTEL)(GLuint  queryHandle);
// typedef void  (APIENTRYP GPENDQUERY)(GLenum  target);
// typedef void  (APIENTRYP GPENDQUERYINDEXED)(GLenum  target, GLuint  index);
// typedef void  (APIENTRYP GPENDTRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPEVALCOORD1XOES)(GLfixed  u);
// typedef void  (APIENTRYP GPEVALCOORD1XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPEVALCOORD2XOES)(GLfixed  u, GLfixed  v);
// typedef void  (APIENTRYP GPEVALCOORD2XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPFEEDBACKBUFFERXOES)(GLsizei  n, GLenum  type, const GLfixed * buffer);
// typedef GLsync  (APIENTRYP GPFENCESYNC)(GLenum  condition, GLbitfield  flags);
// typedef void  (APIENTRYP GPFINISH)();
// typedef void  (APIENTRYP GPFINISHFENCENV)(GLuint  fence);
// typedef void  (APIENTRYP GPFLUSH)();
// typedef void  (APIENTRYP GPFLUSHMAPPEDBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPFLUSHMAPPEDNAMEDBUFFERRANGE)(GLuint  buffer, GLintptr  offset, GLsizei  length);
// typedef void  (APIENTRYP GPFOGXOES)(GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPFOGXVOES)(GLenum  pname, const GLfixed * param);
// typedef void  (APIENTRYP GPFRAMEBUFFERPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPFRAMEBUFFERRENDERBUFFER)(GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE1D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE2D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURE3D)(GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset);
// typedef void  (APIENTRYP GPFRAMEBUFFERTEXTURELAYER)(GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer);
// typedef void  (APIENTRYP GPFRONTFACE)(GLenum  mode);
// typedef void  (APIENTRYP GPFRUSTUMFOES)(GLfloat  l, GLfloat  r, GLfloat  b, GLfloat  t, GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPFRUSTUMXOES)(GLfixed  l, GLfixed  r, GLfixed  b, GLfixed  t, GLfixed  n, GLfixed  f);
// typedef void  (APIENTRYP GPGENBUFFERS)(GLsizei  n, GLuint * buffers);
// typedef void  (APIENTRYP GPGENFENCESNV)(GLsizei  n, GLuint * fences);
// typedef void  (APIENTRYP GPGENFRAMEBUFFERS)(GLsizei  n, GLuint * framebuffers);
// typedef void  (APIENTRYP GPGENPERFMONITORSAMD)(GLsizei  n, GLuint * monitors);
// typedef void  (APIENTRYP GPGENPROGRAMPIPELINES)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPGENPROGRAMPIPELINESEXT)(GLsizei  n, GLuint * pipelines);
// typedef void  (APIENTRYP GPGENQUERIES)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENRENDERBUFFERS)(GLsizei  n, GLuint * renderbuffers);
// typedef void  (APIENTRYP GPGENSAMPLERS)(GLsizei  count, GLuint * samplers);
// typedef void  (APIENTRYP GPGENTEXTURES)(GLsizei  n, GLuint * textures);
// typedef void  (APIENTRYP GPGENTRANSFORMFEEDBACKS)(GLsizei  n, GLuint * ids);
// typedef void  (APIENTRYP GPGENVERTEXARRAYS)(GLsizei  n, GLuint * arrays);
// typedef void  (APIENTRYP GPGENERATEMIPMAP)(GLenum  target);
// typedef void  (APIENTRYP GPGENERATETEXTUREMIPMAP)(GLuint  texture);
// typedef void  (APIENTRYP GPGETACTIVEATOMICCOUNTERBUFFERIV)(GLuint  program, GLuint  bufferIndex, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETACTIVEATTRIB)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINENAME)(GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINEUNIFORMNAME)(GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVESUBROUTINEUNIFORMIV)(GLuint  program, GLenum  shadertype, GLuint  index, GLenum  pname, GLint * values);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORM)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKNAME)(GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMBLOCKIV)(GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMNAME)(GLuint  program, GLuint  uniformIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformName);
// typedef void  (APIENTRYP GPGETACTIVEUNIFORMSIV)(GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETATTACHEDSHADERS)(GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders);
// typedef GLint  (APIENTRYP GPGETATTRIBLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETBOOLEANI_V)(GLenum  target, GLuint  index, GLboolean * data);
// typedef void  (APIENTRYP GPGETBOOLEANV)(GLenum  pname, GLboolean * data);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERI64V)(GLenum  target, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETBUFFERPOINTERV)(GLenum  target, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETBUFFERSUBDATA)(GLenum  target, GLintptr  offset, GLsizeiptr  size, void * data);
// typedef void  (APIENTRYP GPGETCLIPPLANEFOES)(GLenum  plane, GLfloat * equation);
// typedef void  (APIENTRYP GPGETCLIPPLANEXOES)(GLenum  plane, GLfixed * equation);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXIMAGE)(GLenum  target, GLint  level, void * img);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXTUREIMAGE)(GLuint  texture, GLint  level, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETCOMPRESSEDTEXTURESUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETCONVOLUTIONPARAMETERXVOES)(GLenum  target, GLenum  pname, GLfixed * params);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOG)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOGARB)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef GLuint  (APIENTRYP GPGETDEBUGMESSAGELOGKHR)(GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog);
// typedef void  (APIENTRYP GPGETDOUBLEI_V)(GLenum  target, GLuint  index, GLdouble * data);
// typedef void  (APIENTRYP GPGETDOUBLEV)(GLenum  pname, GLdouble * data);
// typedef GLenum  (APIENTRYP GPGETERROR)();
// typedef void  (APIENTRYP GPGETFENCEIVNV)(GLuint  fence, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETFIRSTPERFQUERYIDINTEL)(GLuint * queryId);
// typedef void  (APIENTRYP GPGETFIXEDVOES)(GLenum  pname, GLfixed * params);
// typedef void  (APIENTRYP GPGETFLOATI_V)(GLenum  target, GLuint  index, GLfloat * data);
// typedef void  (APIENTRYP GPGETFLOATV)(GLenum  pname, GLfloat * data);
// typedef GLint  (APIENTRYP GPGETFRAGDATAINDEX)(GLuint  program, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETFRAGDATALOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(GLenum  target, GLenum  attachment, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETFRAMEBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUS)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSARB)();
// typedef GLenum  (APIENTRYP GPGETGRAPHICSRESETSTATUSKHR)();
// typedef void  (APIENTRYP GPGETHISTOGRAMPARAMETERXVOES)(GLenum  target, GLenum  pname, GLfixed * params);
// typedef GLuint64  (APIENTRYP GPGETIMAGEHANDLEARB)(GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format);
// typedef void  (APIENTRYP GPGETINTEGER64I_V)(GLenum  target, GLuint  index, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGER64V)(GLenum  pname, GLint64 * data);
// typedef void  (APIENTRYP GPGETINTEGERI_V)(GLenum  target, GLuint  index, GLint * data);
// typedef void  (APIENTRYP GPGETINTEGERV)(GLenum  pname, GLint * data);
// typedef void  (APIENTRYP GPGETINTERNALFORMATI64V)(GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint64 * params);
// typedef void  (APIENTRYP GPGETINTERNALFORMATIV)(GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETLIGHTXOES)(GLenum  light, GLenum  pname, GLfixed * params);
// typedef void  (APIENTRYP GPGETLIGHTXVOES)(GLenum  light, GLenum  pname, GLfixed * params);
// typedef void  (APIENTRYP GPGETMAPXVOES)(GLenum  target, GLenum  query, GLfixed * v);
// typedef void  (APIENTRYP GPGETMATERIALXOES)(GLenum  face, GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPGETMATERIALXVOES)(GLenum  face, GLenum  pname, GLfixed * params);
// typedef void  (APIENTRYP GPGETMULTISAMPLEFV)(GLenum  pname, GLuint  index, GLfloat * val);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPARAMETERI64V)(GLuint  buffer, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPARAMETERIV)(GLuint  buffer, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERPOINTERV)(GLuint  buffer, GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETNAMEDBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizei  size, void * data);
// typedef void  (APIENTRYP GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV)(GLuint  framebuffer, GLenum  attachment, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDFRAMEBUFFERPARAMETERIV)(GLuint  framebuffer, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETNAMEDRENDERBUFFERPARAMETERIV)(GLuint  renderbuffer, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNAMEDSTRINGARB)(GLint  namelen, const GLchar * name, GLsizei  bufSize, GLint * stringlen, GLchar * string);
// typedef void  (APIENTRYP GPGETNAMEDSTRINGIVARB)(GLint  namelen, const GLchar * name, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNEXTPERFQUERYIDINTEL)(GLuint  queryId, GLuint * nextQueryId);
// typedef void  (APIENTRYP GPGETOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTLABELEXT)(GLenum  type, GLuint  object, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABEL)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETOBJECTPTRLABELKHR)(const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label);
// typedef void  (APIENTRYP GPGETPERFCOUNTERINFOINTEL)(GLuint  queryId, GLuint  counterId, GLuint  counterNameLength, GLchar * counterName, GLuint  counterDescLength, GLchar * counterDesc, GLuint * counterOffset, GLuint * counterDataSize, GLuint * counterTypeEnum, GLuint * counterDataTypeEnum, GLuint64 * rawCounterMaxValue);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERDATAAMD)(GLuint  monitor, GLenum  pname, GLsizei  dataSize, GLuint * data, GLint * bytesWritten);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERINFOAMD)(GLuint  group, GLuint  counter, GLenum  pname, void * data);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERSTRINGAMD)(GLuint  group, GLuint  counter, GLsizei  bufSize, GLsizei * length, GLchar * counterString);
// typedef void  (APIENTRYP GPGETPERFMONITORCOUNTERSAMD)(GLuint  group, GLint * numCounters, GLint * maxActiveCounters, GLsizei  counterSize, GLuint * counters);
// typedef void  (APIENTRYP GPGETPERFMONITORGROUPSTRINGAMD)(GLuint  group, GLsizei  bufSize, GLsizei * length, GLchar * groupString);
// typedef void  (APIENTRYP GPGETPERFMONITORGROUPSAMD)(GLint * numGroups, GLsizei  groupsSize, GLuint * groups);
// typedef void  (APIENTRYP GPGETPERFQUERYDATAINTEL)(GLuint  queryHandle, GLuint  flags, GLsizei  dataSize, GLvoid * data, GLuint * bytesWritten);
// typedef void  (APIENTRYP GPGETPERFQUERYIDBYNAMEINTEL)(GLchar * queryName, GLuint * queryId);
// typedef void  (APIENTRYP GPGETPERFQUERYINFOINTEL)(GLuint  queryId, GLuint  queryNameLength, GLchar * queryName, GLuint * dataSize, GLuint * noCounters, GLuint * noInstances, GLuint * capsMask);
// typedef void  (APIENTRYP GPGETPIXELMAPXV)(GLenum  map, GLint  size, GLfixed * values);
// typedef void  (APIENTRYP GPGETPOINTERV)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPOINTERVKHR)(GLenum  pname, void ** params);
// typedef void  (APIENTRYP GPGETPROGRAMBINARY)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary);
// typedef void  (APIENTRYP GPGETPROGRAMINFOLOG)(GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMINTERFACEIV)(GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEINFOLOG)(GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEINFOLOGEXT)(GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEIV)(GLuint  pipeline, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMPIPELINEIVEXT)(GLuint  pipeline, GLenum  pname, GLint * params);
// typedef GLuint  (APIENTRYP GPGETPROGRAMRESOURCEINDEX)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATION)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETPROGRAMRESOURCELOCATIONINDEX)(GLuint  program, GLenum  programInterface, const GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCENAME)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name);
// typedef void  (APIENTRYP GPGETPROGRAMRESOURCEIV)(GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params);
// typedef void  (APIENTRYP GPGETPROGRAMSTAGEIV)(GLuint  program, GLenum  shadertype, GLenum  pname, GLint * values);
// typedef void  (APIENTRYP GPGETPROGRAMIV)(GLuint  program, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYINDEXEDIV)(GLenum  target, GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTI64V)(GLuint  id, GLenum  pname, GLint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTIV)(GLuint  id, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUI64V)(GLuint  id, GLenum  pname, GLuint64 * params);
// typedef void  (APIENTRYP GPGETQUERYOBJECTUIV)(GLuint  id, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETQUERYIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETRENDERBUFFERPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIIV)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIUIV)(GLuint  sampler, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETSHADERINFOLOG)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog);
// typedef void  (APIENTRYP GPGETSHADERPRECISIONFORMAT)(GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision);
// typedef void  (APIENTRYP GPGETSHADERSOURCE)(GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source);
// typedef void  (APIENTRYP GPGETSHADERIV)(GLuint  shader, GLenum  pname, GLint * params);
// typedef const GLubyte * (APIENTRYP GPGETSTRING)(GLenum  name);
// typedef const GLubyte * (APIENTRYP GPGETSTRINGI)(GLenum  name, GLuint  index);
// typedef GLuint  (APIENTRYP GPGETSUBROUTINEINDEX)(GLuint  program, GLenum  shadertype, const GLchar * name);
// typedef GLint  (APIENTRYP GPGETSUBROUTINEUNIFORMLOCATION)(GLuint  program, GLenum  shadertype, const GLchar * name);
// typedef void  (APIENTRYP GPGETSYNCIV)(GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values);
// typedef void  (APIENTRYP GPGETTEXENVXVOES)(GLenum  target, GLenum  pname, GLfixed * params);
// typedef void  (APIENTRYP GPGETTEXGENXVOES)(GLenum  coord, GLenum  pname, GLfixed * params);
// typedef void  (APIENTRYP GPGETTEXIMAGE)(GLenum  target, GLint  level, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERFV)(GLenum  target, GLint  level, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERIV)(GLenum  target, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXLEVELPARAMETERXVOES)(GLenum  target, GLint  level, GLenum  pname, GLfixed * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIUIV)(GLenum  target, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERFV)(GLenum  target, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERIV)(GLenum  target, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXPARAMETERXVOES)(GLenum  target, GLenum  pname, GLfixed * params);
// typedef GLuint64  (APIENTRYP GPGETTEXTUREHANDLEARB)(GLuint  texture);
// typedef void  (APIENTRYP GPGETTEXTUREIMAGE)(GLuint  texture, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETTEXTURELEVELPARAMETERFV)(GLuint  texture, GLint  level, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXTURELEVELPARAMETERIV)(GLuint  texture, GLint  level, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIIV)(GLuint  texture, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIUIV)(GLuint  texture, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERFV)(GLuint  texture, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETTEXTUREPARAMETERIV)(GLuint  texture, GLenum  pname, GLint * params);
// typedef GLuint64  (APIENTRYP GPGETTEXTURESAMPLERHANDLEARB)(GLuint  texture, GLuint  sampler);
// typedef void  (APIENTRYP GPGETTEXTURESUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKVARYING)(GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKI64_V)(GLuint  xfb, GLenum  pname, GLuint  index, GLint64 * param);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKI_V)(GLuint  xfb, GLenum  pname, GLuint  index, GLint * param);
// typedef void  (APIENTRYP GPGETTRANSFORMFEEDBACKIV)(GLuint  xfb, GLenum  pname, GLint * param);
// typedef GLuint  (APIENTRYP GPGETUNIFORMBLOCKINDEX)(GLuint  program, const GLchar * uniformBlockName);
// typedef void  (APIENTRYP GPGETUNIFORMINDICES)(GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices);
// typedef GLint  (APIENTRYP GPGETUNIFORMLOCATION)(GLuint  program, const GLchar * name);
// typedef void  (APIENTRYP GPGETUNIFORMSUBROUTINEUIV)(GLenum  shadertype, GLint  location, GLuint * params);
// typedef void  (APIENTRYP GPGETUNIFORMDV)(GLuint  program, GLint  location, GLdouble * params);
// typedef void  (APIENTRYP GPGETUNIFORMFV)(GLuint  program, GLint  location, GLfloat * params);
// typedef void  (APIENTRYP GPGETUNIFORMIV)(GLuint  program, GLint  location, GLint * params);
// typedef void  (APIENTRYP GPGETUNIFORMUIV)(GLuint  program, GLint  location, GLuint * params);
// typedef void  (APIENTRYP GPGETVERTEXARRAYINDEXED64IV)(GLuint  vaobj, GLuint  index, GLenum  pname, GLint64 * param);
// typedef void  (APIENTRYP GPGETVERTEXARRAYINDEXEDIV)(GLuint  vaobj, GLuint  index, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETVERTEXARRAYIV)(GLuint  vaobj, GLenum  pname, GLint * param);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIUIV)(GLuint  index, GLenum  pname, GLuint * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBLDV)(GLuint  index, GLenum  pname, GLdouble * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBLUI64VARB)(GLuint  index, GLenum  pname, GLuint64EXT * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBPOINTERV)(GLuint  index, GLenum  pname, void ** pointer);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBDV)(GLuint  index, GLenum  pname, GLdouble * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBFV)(GLuint  index, GLenum  pname, GLfloat * params);
// typedef void  (APIENTRYP GPGETVERTEXATTRIBIV)(GLuint  index, GLenum  pname, GLint * params);
// typedef void  (APIENTRYP GPGETNCOMPRESSEDTEXIMAGEARB)(GLenum  target, GLint  lod, GLsizei  bufSize, void * img);
// typedef void  (APIENTRYP GPGETNTEXIMAGEARB)(GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * img);
// typedef void  (APIENTRYP GPGETNUNIFORMDVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFV)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMFVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIV)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIVARB)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPGETNUNIFORMUIVKHR)(GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params);
// typedef void  (APIENTRYP GPHINT)(GLenum  target, GLenum  mode);
// typedef void  (APIENTRYP GPINDEXXOES)(GLfixed  component);
// typedef void  (APIENTRYP GPINDEXXVOES)(const GLfixed * component);
// typedef void  (APIENTRYP GPINSERTEVENTMARKEREXT)(GLsizei  length, const GLchar * marker);
// typedef void  (APIENTRYP GPINVALIDATEBUFFERDATA)(GLuint  buffer);
// typedef void  (APIENTRYP GPINVALIDATEBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizeiptr  length);
// typedef void  (APIENTRYP GPINVALIDATEFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPINVALIDATENAMEDFRAMEBUFFERDATA)(GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments);
// typedef void  (APIENTRYP GPINVALIDATENAMEDFRAMEBUFFERSUBDATA)(GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPINVALIDATESUBFRAMEBUFFER)(GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPINVALIDATETEXIMAGE)(GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPINVALIDATETEXSUBIMAGE)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef GLboolean  (APIENTRYP GPISBUFFER)(GLuint  buffer);
// typedef GLboolean  (APIENTRYP GPISENABLED)(GLenum  cap);
// typedef GLboolean  (APIENTRYP GPISENABLEDI)(GLenum  target, GLuint  index);
// typedef GLboolean  (APIENTRYP GPISFENCENV)(GLuint  fence);
// typedef GLboolean  (APIENTRYP GPISFRAMEBUFFER)(GLuint  framebuffer);
// typedef GLboolean  (APIENTRYP GPISIMAGEHANDLERESIDENTARB)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISNAMEDSTRINGARB)(GLint  namelen, const GLchar * name);
// typedef GLboolean  (APIENTRYP GPISPROGRAM)(GLuint  program);
// typedef GLboolean  (APIENTRYP GPISPROGRAMPIPELINE)(GLuint  pipeline);
// typedef GLboolean  (APIENTRYP GPISPROGRAMPIPELINEEXT)(GLuint  pipeline);
// typedef GLboolean  (APIENTRYP GPISQUERY)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISRENDERBUFFER)(GLuint  renderbuffer);
// typedef GLboolean  (APIENTRYP GPISSAMPLER)(GLuint  sampler);
// typedef GLboolean  (APIENTRYP GPISSHADER)(GLuint  shader);
// typedef GLboolean  (APIENTRYP GPISSYNC)(GLsync  sync);
// typedef GLboolean  (APIENTRYP GPISTEXTURE)(GLuint  texture);
// typedef GLboolean  (APIENTRYP GPISTEXTUREHANDLERESIDENTARB)(GLuint64  handle);
// typedef GLboolean  (APIENTRYP GPISTRANSFORMFEEDBACK)(GLuint  id);
// typedef GLboolean  (APIENTRYP GPISVERTEXARRAY)(GLuint  array);
// typedef void  (APIENTRYP GPLABELOBJECTEXT)(GLenum  type, GLuint  object, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPLIGHTMODELXOES)(GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPLIGHTMODELXVOES)(GLenum  pname, const GLfixed * param);
// typedef void  (APIENTRYP GPLIGHTXOES)(GLenum  light, GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPLIGHTXVOES)(GLenum  light, GLenum  pname, const GLfixed * params);
// typedef void  (APIENTRYP GPLINEWIDTH)(GLfloat  width);
// typedef void  (APIENTRYP GPLINEWIDTHXOES)(GLfixed  width);
// typedef void  (APIENTRYP GPLINKPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPLOADMATRIXXOES)(const GLfixed * m);
// typedef void  (APIENTRYP GPLOADTRANSPOSEMATRIXXOES)(const GLfixed * m);
// typedef void  (APIENTRYP GPLOGICOP)(GLenum  opcode);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLENONRESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKEIMAGEHANDLERESIDENTARB)(GLuint64  handle, GLenum  access);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLENONRESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAKETEXTUREHANDLERESIDENTARB)(GLuint64  handle);
// typedef void  (APIENTRYP GPMAP1XOES)(GLenum  target, GLfixed  u1, GLfixed  u2, GLint  stride, GLint  order, GLfixed  points);
// typedef void  (APIENTRYP GPMAP2XOES)(GLenum  target, GLfixed  u1, GLfixed  u2, GLint  ustride, GLint  uorder, GLfixed  v1, GLfixed  v2, GLint  vstride, GLint  vorder, GLfixed  points);
// typedef void * (APIENTRYP GPMAPBUFFER)(GLenum  target, GLenum  access);
// typedef void * (APIENTRYP GPMAPBUFFERRANGE)(GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access);
// typedef void  (APIENTRYP GPMAPGRID1XOES)(GLint  n, GLfixed  u1, GLfixed  u2);
// typedef void  (APIENTRYP GPMAPGRID2XOES)(GLint  n, GLfixed  u1, GLfixed  u2, GLfixed  v1, GLfixed  v2);
// typedef void * (APIENTRYP GPMAPNAMEDBUFFER)(GLuint  buffer, GLenum  access);
// typedef void * (APIENTRYP GPMAPNAMEDBUFFERRANGE)(GLuint  buffer, GLintptr  offset, GLsizei  length, GLbitfield  access);
// typedef void  (APIENTRYP GPMATERIALXOES)(GLenum  face, GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPMATERIALXVOES)(GLenum  face, GLenum  pname, const GLfixed * param);
// typedef void  (APIENTRYP GPMEMORYBARRIER)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMEMORYBARRIERBYREGION)(GLbitfield  barriers);
// typedef void  (APIENTRYP GPMINSAMPLESHADING)(GLfloat  value);
// typedef void  (APIENTRYP GPMINSAMPLESHADINGARB)(GLfloat  value);
// typedef void  (APIENTRYP GPMULTMATRIXXOES)(const GLfixed * m);
// typedef void  (APIENTRYP GPMULTTRANSPOSEMATRIXXOES)(const GLfixed * m);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYS)(GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  drawcount);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSEXT)(GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  primcount);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSINDIRECT)(GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWARRAYSINDIRECTCOUNTARB)(GLenum  mode, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTS)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSBASEVERTEX)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount, const GLint * basevertex);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSEXT)(GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  primcount);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSINDIRECT)(GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTIDRAWELEMENTSINDIRECTCOUNTARB)(GLenum  mode, GLenum  type, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride);
// typedef void  (APIENTRYP GPMULTITEXCOORD1BOES)(GLenum  texture, GLbyte  s);
// typedef void  (APIENTRYP GPMULTITEXCOORD1BVOES)(GLenum  texture, const GLbyte * coords);
// typedef void  (APIENTRYP GPMULTITEXCOORD1XOES)(GLenum  texture, GLfixed  s);
// typedef void  (APIENTRYP GPMULTITEXCOORD1XVOES)(GLenum  texture, const GLfixed * coords);
// typedef void  (APIENTRYP GPMULTITEXCOORD2BOES)(GLenum  texture, GLbyte  s, GLbyte  t);
// typedef void  (APIENTRYP GPMULTITEXCOORD2BVOES)(GLenum  texture, const GLbyte * coords);
// typedef void  (APIENTRYP GPMULTITEXCOORD2XOES)(GLenum  texture, GLfixed  s, GLfixed  t);
// typedef void  (APIENTRYP GPMULTITEXCOORD2XVOES)(GLenum  texture, const GLfixed * coords);
// typedef void  (APIENTRYP GPMULTITEXCOORD3BOES)(GLenum  texture, GLbyte  s, GLbyte  t, GLbyte  r);
// typedef void  (APIENTRYP GPMULTITEXCOORD3BVOES)(GLenum  texture, const GLbyte * coords);
// typedef void  (APIENTRYP GPMULTITEXCOORD3XOES)(GLenum  texture, GLfixed  s, GLfixed  t, GLfixed  r);
// typedef void  (APIENTRYP GPMULTITEXCOORD3XVOES)(GLenum  texture, const GLfixed * coords);
// typedef void  (APIENTRYP GPMULTITEXCOORD4BOES)(GLenum  texture, GLbyte  s, GLbyte  t, GLbyte  r, GLbyte  q);
// typedef void  (APIENTRYP GPMULTITEXCOORD4BVOES)(GLenum  texture, const GLbyte * coords);
// typedef void  (APIENTRYP GPMULTITEXCOORD4XOES)(GLenum  texture, GLfixed  s, GLfixed  t, GLfixed  r, GLfixed  q);
// typedef void  (APIENTRYP GPMULTITEXCOORD4XVOES)(GLenum  texture, const GLfixed * coords);
// typedef void  (APIENTRYP GPNAMEDBUFFERDATA)(GLuint  buffer, GLsizei  size, const void * data, GLenum  usage);
// typedef void  (APIENTRYP GPNAMEDBUFFERPAGECOMMITMENTARB)(GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPNAMEDBUFFERPAGECOMMITMENTEXT)(GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit);
// typedef void  (APIENTRYP GPNAMEDBUFFERSTORAGE)(GLuint  buffer, GLsizei  size, const void * data, GLbitfield  flags);
// typedef void  (APIENTRYP GPNAMEDBUFFERSUBDATA)(GLuint  buffer, GLintptr  offset, GLsizei  size, const void * data);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERDRAWBUFFER)(GLuint  framebuffer, GLenum  buf);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERDRAWBUFFERS)(GLuint  framebuffer, GLsizei  n, const GLenum * bufs);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERPARAMETERI)(GLuint  framebuffer, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERREADBUFFER)(GLuint  framebuffer, GLenum  src);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERRENDERBUFFER)(GLuint  framebuffer, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERTEXTURE)(GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level);
// typedef void  (APIENTRYP GPNAMEDFRAMEBUFFERTEXTURELAYER)(GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer);
// typedef void  (APIENTRYP GPNAMEDRENDERBUFFERSTORAGE)(GLuint  renderbuffer, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE)(GLuint  renderbuffer, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPNAMEDSTRINGARB)(GLenum  type, GLint  namelen, const GLchar * name, GLint  stringlen, const GLchar * string);
// typedef void  (APIENTRYP GPNORMAL3XOES)(GLfixed  nx, GLfixed  ny, GLfixed  nz);
// typedef void  (APIENTRYP GPNORMAL3XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPOBJECTLABEL)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTLABELKHR)(GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABEL)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPOBJECTPTRLABELKHR)(const void * ptr, GLsizei  length, const GLchar * label);
// typedef void  (APIENTRYP GPORTHOFOES)(GLfloat  l, GLfloat  r, GLfloat  b, GLfloat  t, GLfloat  n, GLfloat  f);
// typedef void  (APIENTRYP GPORTHOXOES)(GLfixed  l, GLfixed  r, GLfixed  b, GLfixed  t, GLfixed  n, GLfixed  f);
// typedef void  (APIENTRYP GPPASSTHROUGHXOES)(GLfixed  token);
// typedef void  (APIENTRYP GPPATCHPARAMETERFV)(GLenum  pname, const GLfloat * values);
// typedef void  (APIENTRYP GPPATCHPARAMETERI)(GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPAUSETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPPIXELMAPX)(GLenum  map, GLint  size, const GLfixed * values);
// typedef void  (APIENTRYP GPPIXELSTOREF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPPIXELSTOREI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPPIXELSTOREX)(GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPPIXELTRANSFERXOES)(GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPPIXELZOOMXOES)(GLfixed  xfactor, GLfixed  yfactor);
// typedef void  (APIENTRYP GPPOINTPARAMETERF)(GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERFV)(GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPPOINTPARAMETERI)(GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERIV)(GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPPOINTPARAMETERXOES)(GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPPOINTPARAMETERXVOES)(GLenum  pname, const GLfixed * params);
// typedef void  (APIENTRYP GPPOINTSIZE)(GLfloat  size);
// typedef void  (APIENTRYP GPPOINTSIZEXOES)(GLfixed  size);
// typedef void  (APIENTRYP GPPOLYGONMODE)(GLenum  face, GLenum  mode);
// typedef void  (APIENTRYP GPPOLYGONOFFSET)(GLfloat  factor, GLfloat  units);
// typedef void  (APIENTRYP GPPOLYGONOFFSETXOES)(GLfixed  factor, GLfixed  units);
// typedef void  (APIENTRYP GPPOPDEBUGGROUP)();
// typedef void  (APIENTRYP GPPOPDEBUGGROUPKHR)();
// typedef void  (APIENTRYP GPPOPGROUPMARKEREXT)();
// typedef void  (APIENTRYP GPPRIMITIVERESTARTINDEX)(GLuint  index);
// typedef void  (APIENTRYP GPPRIORITIZETEXTURESXOES)(GLsizei  n, const GLuint * textures, const GLfixed * priorities);
// typedef void  (APIENTRYP GPPROGRAMBINARY)(GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPPROGRAMPARAMETERI)(GLuint  program, GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPROGRAMPARAMETERIEXT)(GLuint  program, GLenum  pname, GLint  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1D)(GLuint  program, GLint  location, GLdouble  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1F)(GLuint  program, GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FEXT)(GLuint  program, GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1I)(GLuint  program, GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IEXT)(GLuint  program, GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UI)(GLuint  program, GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIEXT)(GLuint  program, GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM1UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FEXT)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2I)(GLuint  program, GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IEXT)(GLuint  program, GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIEXT)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM2UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FEXT)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IEXT)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIEXT)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM3UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4D)(GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2, GLdouble  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4DV)(GLuint  program, GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4F)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FEXT)(GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FV)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4FVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4I)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IEXT)(GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IV)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4IVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UI)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIEXT)(GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIV)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORM4UIVEXT)(GLuint  program, GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64ARB)(GLuint  program, GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMHANDLEUI64VARB)(GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X3FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX2X4FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X2FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX3X4FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X2FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3DV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3FV)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROGRAMUNIFORMMATRIX4X3FVEXT)(GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPPROVOKINGVERTEX)(GLenum  mode);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUP)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPPUSHDEBUGGROUPKHR)(GLenum  source, GLuint  id, GLsizei  length, const GLchar * message);
// typedef void  (APIENTRYP GPPUSHGROUPMARKEREXT)(GLsizei  length, const GLchar * marker);
// typedef void  (APIENTRYP GPQUERYCOUNTER)(GLuint  id, GLenum  target);
// typedef GLbitfield  (APIENTRYP GPQUERYMATRIXXOES)(GLfixed * mantissa, GLint * exponent);
// typedef void  (APIENTRYP GPRASTERPOS2XOES)(GLfixed  x, GLfixed  y);
// typedef void  (APIENTRYP GPRASTERPOS2XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPRASTERPOS3XOES)(GLfixed  x, GLfixed  y, GLfixed  z);
// typedef void  (APIENTRYP GPRASTERPOS3XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPRASTERPOS4XOES)(GLfixed  x, GLfixed  y, GLfixed  z, GLfixed  w);
// typedef void  (APIENTRYP GPRASTERPOS4XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPREADBUFFER)(GLenum  src);
// typedef void  (APIENTRYP GPREADPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels);
// typedef void  (APIENTRYP GPREADNPIXELS)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSARB)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPREADNPIXELSKHR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data);
// typedef void  (APIENTRYP GPRECTXOES)(GLfixed  x1, GLfixed  y1, GLfixed  x2, GLfixed  y2);
// typedef void  (APIENTRYP GPRECTXVOES)(const GLfixed * v1, const GLfixed * v2);
// typedef void  (APIENTRYP GPRELEASESHADERCOMPILER)();
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGE)(GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRENDERBUFFERSTORAGEMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPRESUMETRANSFORMFEEDBACK)();
// typedef void  (APIENTRYP GPROTATEXOES)(GLfixed  angle, GLfixed  x, GLfixed  y, GLfixed  z);
// typedef void  (APIENTRYP GPSAMPLECOVERAGE)(GLfloat  value, GLboolean  invert);
// typedef void  (APIENTRYP GPSAMPLECOVERAGEOES)(GLfixed  value, GLboolean  invert);
// typedef void  (APIENTRYP GPSAMPLECOVERAGEXOES)(GLclampx  value, GLboolean  invert);
// typedef void  (APIENTRYP GPSAMPLEMASKI)(GLuint  maskNumber, GLbitfield  mask);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIIV)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIUIV)(GLuint  sampler, GLenum  pname, const GLuint * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERF)(GLuint  sampler, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERFV)(GLuint  sampler, GLenum  pname, const GLfloat * param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERI)(GLuint  sampler, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPSAMPLERPARAMETERIV)(GLuint  sampler, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPSCALEXOES)(GLfixed  x, GLfixed  y, GLfixed  z);
// typedef void  (APIENTRYP GPSCISSOR)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORARRAYV)(GLuint  first, GLsizei  count, const GLint * v);
// typedef void  (APIENTRYP GPSCISSORINDEXED)(GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPSCISSORINDEXEDV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPSELECTPERFMONITORCOUNTERSAMD)(GLuint  monitor, GLboolean  enable, GLuint  group, GLint  numCounters, GLuint * counterList);
// typedef void  (APIENTRYP GPSETFENCENV)(GLuint  fence, GLenum  condition);
// typedef void  (APIENTRYP GPSHADERBINARY)(GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length);
// typedef void  (APIENTRYP GPSHADERSOURCE)(GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length);
// typedef void  (APIENTRYP GPSHADERSTORAGEBLOCKBINDING)(GLuint  program, GLuint  storageBlockIndex, GLuint  storageBlockBinding);
// typedef void  (APIENTRYP GPSTENCILFUNC)(GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILFUNCSEPARATE)(GLenum  face, GLenum  func, GLint  ref, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASK)(GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILMASKSEPARATE)(GLenum  face, GLuint  mask);
// typedef void  (APIENTRYP GPSTENCILOP)(GLenum  fail, GLenum  zfail, GLenum  zpass);
// typedef void  (APIENTRYP GPSTENCILOPSEPARATE)(GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass);
// typedef GLboolean  (APIENTRYP GPTESTFENCENV)(GLuint  fence);
// typedef void  (APIENTRYP GPTEXBUFFER)(GLenum  target, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXBUFFERRANGE)(GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size);
// typedef void  (APIENTRYP GPTEXCOORD1BOES)(GLbyte  s);
// typedef void  (APIENTRYP GPTEXCOORD1BVOES)(const GLbyte * coords);
// typedef void  (APIENTRYP GPTEXCOORD1XOES)(GLfixed  s);
// typedef void  (APIENTRYP GPTEXCOORD1XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPTEXCOORD2BOES)(GLbyte  s, GLbyte  t);
// typedef void  (APIENTRYP GPTEXCOORD2BVOES)(const GLbyte * coords);
// typedef void  (APIENTRYP GPTEXCOORD2XOES)(GLfixed  s, GLfixed  t);
// typedef void  (APIENTRYP GPTEXCOORD2XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPTEXCOORD3BOES)(GLbyte  s, GLbyte  t, GLbyte  r);
// typedef void  (APIENTRYP GPTEXCOORD3BVOES)(const GLbyte * coords);
// typedef void  (APIENTRYP GPTEXCOORD3XOES)(GLfixed  s, GLfixed  t, GLfixed  r);
// typedef void  (APIENTRYP GPTEXCOORD3XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPTEXCOORD4BOES)(GLbyte  s, GLbyte  t, GLbyte  r, GLbyte  q);
// typedef void  (APIENTRYP GPTEXCOORD4BVOES)(const GLbyte * coords);
// typedef void  (APIENTRYP GPTEXCOORD4XOES)(GLfixed  s, GLfixed  t, GLfixed  r, GLfixed  q);
// typedef void  (APIENTRYP GPTEXCOORD4XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPTEXENVXOES)(GLenum  target, GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPTEXENVXVOES)(GLenum  target, GLenum  pname, const GLfixed * params);
// typedef void  (APIENTRYP GPTEXGENXOES)(GLenum  coord, GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPTEXGENXVOES)(GLenum  coord, GLenum  pname, const GLfixed * params);
// typedef void  (APIENTRYP GPTEXIMAGE1D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE2D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE2DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXIMAGE3D)(GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXIMAGE3DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXPAGECOMMITMENTARB)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  resident);
// typedef void  (APIENTRYP GPTEXPARAMETERIIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERIUIV)(GLenum  target, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERF)(GLenum  target, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXPARAMETERFV)(GLenum  target, GLenum  pname, const GLfloat * params);
// typedef void  (APIENTRYP GPTEXPARAMETERI)(GLenum  target, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXPARAMETERIV)(GLenum  target, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXPARAMETERXOES)(GLenum  target, GLenum  pname, GLfixed  param);
// typedef void  (APIENTRYP GPTEXPARAMETERXVOES)(GLenum  target, GLenum  pname, const GLfixed * params);
// typedef void  (APIENTRYP GPTEXSTORAGE1D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXSTORAGE2D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXSTORAGE2DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSTORAGE3D)(GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXSTORAGE3DMULTISAMPLE)(GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXSUBIMAGE1D)(GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE2D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXSUBIMAGE3D)(GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTUREBARRIER)();
// typedef void  (APIENTRYP GPTEXTUREBUFFER)(GLuint  texture, GLenum  internalformat, GLuint  buffer);
// typedef void  (APIENTRYP GPTEXTUREBUFFERRANGE)(GLuint  texture, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizei  size);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIIV)(GLuint  texture, GLenum  pname, const GLint * params);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIUIV)(GLuint  texture, GLenum  pname, const GLuint * params);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERF)(GLuint  texture, GLenum  pname, GLfloat  param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERFV)(GLuint  texture, GLenum  pname, const GLfloat * param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERI)(GLuint  texture, GLenum  pname, GLint  param);
// typedef void  (APIENTRYP GPTEXTUREPARAMETERIV)(GLuint  texture, GLenum  pname, const GLint * param);
// typedef void  (APIENTRYP GPTEXTURESTORAGE1D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width);
// typedef void  (APIENTRYP GPTEXTURESTORAGE2D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPTEXTURESTORAGE2DMULTISAMPLE)(GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXTURESTORAGE3D)(GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth);
// typedef void  (APIENTRYP GPTEXTURESTORAGE3DMULTISAMPLE)(GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE1D)(GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE2D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTURESUBIMAGE3D)(GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels);
// typedef void  (APIENTRYP GPTEXTUREVIEW)(GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKBUFFERBASE)(GLuint  xfb, GLuint  index, GLuint  buffer);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKBUFFERRANGE)(GLuint  xfb, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizei  size);
// typedef void  (APIENTRYP GPTRANSFORMFEEDBACKVARYINGS)(GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode);
// typedef void  (APIENTRYP GPTRANSLATEXOES)(GLfixed  x, GLfixed  y, GLfixed  z);
// typedef void  (APIENTRYP GPUNIFORM1D)(GLint  location, GLdouble  x);
// typedef void  (APIENTRYP GPUNIFORM1DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM1F)(GLint  location, GLfloat  v0);
// typedef void  (APIENTRYP GPUNIFORM1FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM1I)(GLint  location, GLint  v0);
// typedef void  (APIENTRYP GPUNIFORM1IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM1UI)(GLint  location, GLuint  v0);
// typedef void  (APIENTRYP GPUNIFORM1UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM2D)(GLint  location, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPUNIFORM2DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM2F)(GLint  location, GLfloat  v0, GLfloat  v1);
// typedef void  (APIENTRYP GPUNIFORM2FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM2I)(GLint  location, GLint  v0, GLint  v1);
// typedef void  (APIENTRYP GPUNIFORM2IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM2UI)(GLint  location, GLuint  v0, GLuint  v1);
// typedef void  (APIENTRYP GPUNIFORM2UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM3D)(GLint  location, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPUNIFORM3DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM3F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2);
// typedef void  (APIENTRYP GPUNIFORM3FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM3I)(GLint  location, GLint  v0, GLint  v1, GLint  v2);
// typedef void  (APIENTRYP GPUNIFORM3IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM3UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2);
// typedef void  (APIENTRYP GPUNIFORM3UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORM4D)(GLint  location, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPUNIFORM4DV)(GLint  location, GLsizei  count, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORM4F)(GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3);
// typedef void  (APIENTRYP GPUNIFORM4FV)(GLint  location, GLsizei  count, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORM4I)(GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3);
// typedef void  (APIENTRYP GPUNIFORM4IV)(GLint  location, GLsizei  count, const GLint * value);
// typedef void  (APIENTRYP GPUNIFORM4UI)(GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3);
// typedef void  (APIENTRYP GPUNIFORM4UIV)(GLint  location, GLsizei  count, const GLuint * value);
// typedef void  (APIENTRYP GPUNIFORMBLOCKBINDING)(GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64ARB)(GLint  location, GLuint64  value);
// typedef void  (APIENTRYP GPUNIFORMHANDLEUI64VARB)(GLint  location, GLsizei  count, const GLuint64 * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX2X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX3X4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X2FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3DV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value);
// typedef void  (APIENTRYP GPUNIFORMMATRIX4X3FV)(GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value);
// typedef void  (APIENTRYP GPUNIFORMSUBROUTINESUIV)(GLenum  shadertype, GLsizei  count, const GLuint * indices);
// typedef GLboolean  (APIENTRYP GPUNMAPBUFFER)(GLenum  target);
// typedef GLboolean  (APIENTRYP GPUNMAPNAMEDBUFFER)(GLuint  buffer);
// typedef void  (APIENTRYP GPUSEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPUSEPROGRAMSTAGES)(GLuint  pipeline, GLbitfield  stages, GLuint  program);
// typedef void  (APIENTRYP GPUSEPROGRAMSTAGESEXT)(GLuint  pipeline, GLbitfield  stages, GLuint  program);
// typedef void  (APIENTRYP GPUSESHADERPROGRAMEXT)(GLenum  type, GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAM)(GLuint  program);
// typedef void  (APIENTRYP GPVALIDATEPROGRAMPIPELINE)(GLuint  pipeline);
// typedef void  (APIENTRYP GPVALIDATEPROGRAMPIPELINEEXT)(GLuint  pipeline);
// typedef void  (APIENTRYP GPVERTEX2BOES)(GLbyte  x, GLbyte  y);
// typedef void  (APIENTRYP GPVERTEX2BVOES)(const GLbyte * coords);
// typedef void  (APIENTRYP GPVERTEX2XOES)(GLfixed  x);
// typedef void  (APIENTRYP GPVERTEX2XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPVERTEX3BOES)(GLbyte  x, GLbyte  y, GLbyte  z);
// typedef void  (APIENTRYP GPVERTEX3BVOES)(const GLbyte * coords);
// typedef void  (APIENTRYP GPVERTEX3XOES)(GLfixed  x, GLfixed  y);
// typedef void  (APIENTRYP GPVERTEX3XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPVERTEX4BOES)(GLbyte  x, GLbyte  y, GLbyte  z, GLbyte  w);
// typedef void  (APIENTRYP GPVERTEX4BVOES)(const GLbyte * coords);
// typedef void  (APIENTRYP GPVERTEX4XOES)(GLfixed  x, GLfixed  y, GLfixed  z);
// typedef void  (APIENTRYP GPVERTEX4XVOES)(const GLfixed * coords);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBBINDING)(GLuint  vaobj, GLuint  attribindex, GLuint  bindingindex);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBIFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYATTRIBLFORMAT)(GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXARRAYBINDINGDIVISOR)(GLuint  vaobj, GLuint  bindingindex, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXARRAYELEMENTBUFFER)(GLuint  vaobj, GLuint  buffer);
// typedef void  (APIENTRYP GPVERTEXARRAYVERTEXBUFFER)(GLuint  vaobj, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride);
// typedef void  (APIENTRYP GPVERTEXARRAYVERTEXBUFFERS)(GLuint  vaobj, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides);
// typedef void  (APIENTRYP GPVERTEXATTRIB1D)(GLuint  index, GLdouble  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB1F)(GLuint  index, GLfloat  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB1S)(GLuint  index, GLshort  x);
// typedef void  (APIENTRYP GPVERTEXATTRIB1SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2D)(GLuint  index, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2F)(GLuint  index, GLfloat  x, GLfloat  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB2S)(GLuint  index, GLshort  x, GLshort  y);
// typedef void  (APIENTRYP GPVERTEXATTRIB2SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB3S)(GLuint  index, GLshort  x, GLshort  y, GLshort  z);
// typedef void  (APIENTRYP GPVERTEXATTRIB3SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NBV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NIV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NSV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUB)(GLuint  index, GLubyte  x, GLubyte  y, GLubyte  z, GLubyte  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4NUSV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4BV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4F)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4FV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4S)(GLuint  index, GLshort  x, GLshort  y, GLshort  z, GLshort  w);
// typedef void  (APIENTRYP GPVERTEXATTRIB4SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4UBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIB4USV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBBINDING)(GLuint  attribindex, GLuint  bindingindex);
// typedef void  (APIENTRYP GPVERTEXATTRIBDIVISOR)(GLuint  index, GLuint  divisor);
// typedef void  (APIENTRYP GPVERTEXATTRIBFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1I)(GLuint  index, GLint  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1UI)(GLuint  index, GLuint  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBI1UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2I)(GLuint  index, GLint  x, GLint  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2UI)(GLuint  index, GLuint  x, GLuint  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBI2UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3I)(GLuint  index, GLint  x, GLint  y, GLint  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3UI)(GLuint  index, GLuint  x, GLuint  y, GLuint  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBI3UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4BV)(GLuint  index, const GLbyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4I)(GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4IV)(GLuint  index, const GLint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4SV)(GLuint  index, const GLshort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UBV)(GLuint  index, const GLubyte * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UI)(GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4UIV)(GLuint  index, const GLuint * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBI4USV)(GLuint  index, const GLushort * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBIFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBIPOINTER)(GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1D)(GLuint  index, GLdouble  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1UI64ARB)(GLuint  index, GLuint64EXT  x);
// typedef void  (APIENTRYP GPVERTEXATTRIBL1UI64VARB)(GLuint  index, const GLuint64EXT * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL2D)(GLuint  index, GLdouble  x, GLdouble  y);
// typedef void  (APIENTRYP GPVERTEXATTRIBL2DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL3D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z);
// typedef void  (APIENTRYP GPVERTEXATTRIBL3DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBL4D)(GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w);
// typedef void  (APIENTRYP GPVERTEXATTRIBL4DV)(GLuint  index, const GLdouble * v);
// typedef void  (APIENTRYP GPVERTEXATTRIBLFORMAT)(GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset);
// typedef void  (APIENTRYP GPVERTEXATTRIBLPOINTER)(GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXATTRIBP1UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP1UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP2UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP2UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP3UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP3UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP4UI)(GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value);
// typedef void  (APIENTRYP GPVERTEXATTRIBP4UIV)(GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value);
// typedef void  (APIENTRYP GPVERTEXATTRIBPOINTER)(GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer);
// typedef void  (APIENTRYP GPVERTEXBINDINGDIVISOR)(GLuint  bindingindex, GLuint  divisor);
// typedef void  (APIENTRYP GPVIEWPORT)(GLint  x, GLint  y, GLsizei  width, GLsizei  height);
// typedef void  (APIENTRYP GPVIEWPORTARRAYV)(GLuint  first, GLsizei  count, const GLfloat * v);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDF)(GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h);
// typedef void  (APIENTRYP GPVIEWPORTINDEXEDFV)(GLuint  index, const GLfloat * v);
// typedef void  (APIENTRYP GPWAITSYNC)(GLsync  sync, GLbitfield  flags, GLuint64  timeout);
// static void  glowAccumxOES(GPACCUMXOES fnptr, GLenum  op, GLfixed  value) {
//   (*fnptr)(op, value);
// }
// static void  glowActiveProgramEXT(GPACTIVEPROGRAMEXT fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowActiveShaderProgram(GPACTIVESHADERPROGRAM fnptr, GLuint  pipeline, GLuint  program) {
//   (*fnptr)(pipeline, program);
// }
// static void  glowActiveShaderProgramEXT(GPACTIVESHADERPROGRAMEXT fnptr, GLuint  pipeline, GLuint  program) {
//   (*fnptr)(pipeline, program);
// }
// static void  glowActiveTexture(GPACTIVETEXTURE fnptr, GLenum  texture) {
//   (*fnptr)(texture);
// }
// static void  glowAlphaFuncxOES(GPALPHAFUNCXOES fnptr, GLenum  func, GLfixed  ref) {
//   (*fnptr)(func, ref);
// }
// static void  glowAttachShader(GPATTACHSHADER fnptr, GLuint  program, GLuint  shader) {
//   (*fnptr)(program, shader);
// }
// static void  glowBeginConditionalRender(GPBEGINCONDITIONALRENDER fnptr, GLuint  id, GLenum  mode) {
//   (*fnptr)(id, mode);
// }
// static void  glowBeginPerfMonitorAMD(GPBEGINPERFMONITORAMD fnptr, GLuint  monitor) {
//   (*fnptr)(monitor);
// }
// static void  glowBeginPerfQueryINTEL(GPBEGINPERFQUERYINTEL fnptr, GLuint  queryHandle) {
//   (*fnptr)(queryHandle);
// }
// static void  glowBeginQuery(GPBEGINQUERY fnptr, GLenum  target, GLuint  id) {
//   (*fnptr)(target, id);
// }
// static void  glowBeginQueryIndexed(GPBEGINQUERYINDEXED fnptr, GLenum  target, GLuint  index, GLuint  id) {
//   (*fnptr)(target, index, id);
// }
// static void  glowBeginTransformFeedback(GPBEGINTRANSFORMFEEDBACK fnptr, GLenum  primitiveMode) {
//   (*fnptr)(primitiveMode);
// }
// static void  glowBindAttribLocation(GPBINDATTRIBLOCATION fnptr, GLuint  program, GLuint  index, const GLchar * name) {
//   (*fnptr)(program, index, name);
// }
// static void  glowBindBuffer(GPBINDBUFFER fnptr, GLenum  target, GLuint  buffer) {
//   (*fnptr)(target, buffer);
// }
// static void  glowBindBufferBase(GPBINDBUFFERBASE fnptr, GLenum  target, GLuint  index, GLuint  buffer) {
//   (*fnptr)(target, index, buffer);
// }
// static void  glowBindBufferRange(GPBINDBUFFERRANGE fnptr, GLenum  target, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {
//   (*fnptr)(target, index, buffer, offset, size);
// }
// static void  glowBindBuffersBase(GPBINDBUFFERSBASE fnptr, GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers) {
//   (*fnptr)(target, first, count, buffers);
// }
// static void  glowBindBuffersRange(GPBINDBUFFERSRANGE fnptr, GLenum  target, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizeiptr * sizes) {
//   (*fnptr)(target, first, count, buffers, offsets, sizes);
// }
// static void  glowBindFragDataLocation(GPBINDFRAGDATALOCATION fnptr, GLuint  program, GLuint  color, const GLchar * name) {
//   (*fnptr)(program, color, name);
// }
// static void  glowBindFragDataLocationIndexed(GPBINDFRAGDATALOCATIONINDEXED fnptr, GLuint  program, GLuint  colorNumber, GLuint  index, const GLchar * name) {
//   (*fnptr)(program, colorNumber, index, name);
// }
// static void  glowBindFramebuffer(GPBINDFRAMEBUFFER fnptr, GLenum  target, GLuint  framebuffer) {
//   (*fnptr)(target, framebuffer);
// }
// static void  glowBindImageTexture(GPBINDIMAGETEXTURE fnptr, GLuint  unit, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  access, GLenum  format) {
//   (*fnptr)(unit, texture, level, layered, layer, access, format);
// }
// static void  glowBindImageTextures(GPBINDIMAGETEXTURES fnptr, GLuint  first, GLsizei  count, const GLuint * textures) {
//   (*fnptr)(first, count, textures);
// }
// static void  glowBindProgramPipeline(GPBINDPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowBindProgramPipelineEXT(GPBINDPROGRAMPIPELINEEXT fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowBindRenderbuffer(GPBINDRENDERBUFFER fnptr, GLenum  target, GLuint  renderbuffer) {
//   (*fnptr)(target, renderbuffer);
// }
// static void  glowBindSampler(GPBINDSAMPLER fnptr, GLuint  unit, GLuint  sampler) {
//   (*fnptr)(unit, sampler);
// }
// static void  glowBindSamplers(GPBINDSAMPLERS fnptr, GLuint  first, GLsizei  count, const GLuint * samplers) {
//   (*fnptr)(first, count, samplers);
// }
// static void  glowBindTexture(GPBINDTEXTURE fnptr, GLenum  target, GLuint  texture) {
//   (*fnptr)(target, texture);
// }
// static void  glowBindTextureUnit(GPBINDTEXTUREUNIT fnptr, GLuint  unit, GLuint  texture) {
//   (*fnptr)(unit, texture);
// }
// static void  glowBindTextures(GPBINDTEXTURES fnptr, GLuint  first, GLsizei  count, const GLuint * textures) {
//   (*fnptr)(first, count, textures);
// }
// static void  glowBindTransformFeedback(GPBINDTRANSFORMFEEDBACK fnptr, GLenum  target, GLuint  id) {
//   (*fnptr)(target, id);
// }
// static void  glowBindVertexArray(GPBINDVERTEXARRAY fnptr, GLuint  array) {
//   (*fnptr)(array);
// }
// static void  glowBindVertexBuffer(GPBINDVERTEXBUFFER fnptr, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride) {
//   (*fnptr)(bindingindex, buffer, offset, stride);
// }
// static void  glowBindVertexBuffers(GPBINDVERTEXBUFFERS fnptr, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides) {
//   (*fnptr)(first, count, buffers, offsets, strides);
// }
// static void  glowBitmapxOES(GPBITMAPXOES fnptr, GLsizei  width, GLsizei  height, GLfixed  xorig, GLfixed  yorig, GLfixed  xmove, GLfixed  ymove, const GLubyte * bitmap) {
//   (*fnptr)(width, height, xorig, yorig, xmove, ymove, bitmap);
// }
// static void  glowBlendBarrierKHR(GPBLENDBARRIERKHR fnptr) {
//   (*fnptr)();
// }
// static void  glowBlendBarrierNV(GPBLENDBARRIERNV fnptr) {
//   (*fnptr)();
// }
// static void  glowBlendColor(GPBLENDCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowBlendColorxOES(GPBLENDCOLORXOES fnptr, GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowBlendEquation(GPBLENDEQUATION fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowBlendEquationEXT(GPBLENDEQUATIONEXT fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowBlendEquationSeparate(GPBLENDEQUATIONSEPARATE fnptr, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(modeRGB, modeAlpha);
// }
// static void  glowBlendEquationSeparatei(GPBLENDEQUATIONSEPARATEI fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(buf, modeRGB, modeAlpha);
// }
// static void  glowBlendEquationSeparateiARB(GPBLENDEQUATIONSEPARATEIARB fnptr, GLuint  buf, GLenum  modeRGB, GLenum  modeAlpha) {
//   (*fnptr)(buf, modeRGB, modeAlpha);
// }
// static void  glowBlendEquationi(GPBLENDEQUATIONI fnptr, GLuint  buf, GLenum  mode) {
//   (*fnptr)(buf, mode);
// }
// static void  glowBlendEquationiARB(GPBLENDEQUATIONIARB fnptr, GLuint  buf, GLenum  mode) {
//   (*fnptr)(buf, mode);
// }
// static void  glowBlendFunc(GPBLENDFUNC fnptr, GLenum  sfactor, GLenum  dfactor) {
//   (*fnptr)(sfactor, dfactor);
// }
// static void  glowBlendFuncSeparate(GPBLENDFUNCSEPARATE fnptr, GLenum  sfactorRGB, GLenum  dfactorRGB, GLenum  sfactorAlpha, GLenum  dfactorAlpha) {
//   (*fnptr)(sfactorRGB, dfactorRGB, sfactorAlpha, dfactorAlpha);
// }
// static void  glowBlendFuncSeparatei(GPBLENDFUNCSEPARATEI fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {
//   (*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// static void  glowBlendFuncSeparateiARB(GPBLENDFUNCSEPARATEIARB fnptr, GLuint  buf, GLenum  srcRGB, GLenum  dstRGB, GLenum  srcAlpha, GLenum  dstAlpha) {
//   (*fnptr)(buf, srcRGB, dstRGB, srcAlpha, dstAlpha);
// }
// static void  glowBlendFunci(GPBLENDFUNCI fnptr, GLuint  buf, GLenum  src, GLenum  dst) {
//   (*fnptr)(buf, src, dst);
// }
// static void  glowBlendFunciARB(GPBLENDFUNCIARB fnptr, GLuint  buf, GLenum  src, GLenum  dst) {
//   (*fnptr)(buf, src, dst);
// }
// static void  glowBlendParameteriNV(GPBLENDPARAMETERINV fnptr, GLenum  pname, GLint  value) {
//   (*fnptr)(pname, value);
// }
// static void  glowBlitFramebuffer(GPBLITFRAMEBUFFER fnptr, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   (*fnptr)(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static void  glowBlitNamedFramebuffer(GPBLITNAMEDFRAMEBUFFER fnptr, GLuint  readFramebuffer, GLuint  drawFramebuffer, GLint  srcX0, GLint  srcY0, GLint  srcX1, GLint  srcY1, GLint  dstX0, GLint  dstY0, GLint  dstX1, GLint  dstY1, GLbitfield  mask, GLenum  filter) {
//   (*fnptr)(readFramebuffer, drawFramebuffer, srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter);
// }
// static void  glowBufferData(GPBUFFERDATA fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLenum  usage) {
//   (*fnptr)(target, size, data, usage);
// }
// static void  glowBufferPageCommitmentARB(GPBUFFERPAGECOMMITMENTARB fnptr, GLenum  target, GLintptr  offset, GLsizei  size, GLboolean  commit) {
//   (*fnptr)(target, offset, size, commit);
// }
// static void  glowBufferStorage(GPBUFFERSTORAGE fnptr, GLenum  target, GLsizeiptr  size, const void * data, GLbitfield  flags) {
//   (*fnptr)(target, size, data, flags);
// }
// static void  glowBufferSubData(GPBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, const void * data) {
//   (*fnptr)(target, offset, size, data);
// }
// static GLenum  glowCheckFramebufferStatus(GPCHECKFRAMEBUFFERSTATUS fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static GLenum  glowCheckNamedFramebufferStatus(GPCHECKNAMEDFRAMEBUFFERSTATUS fnptr, GLuint  framebuffer, GLenum  target) {
//   return (*fnptr)(framebuffer, target);
// }
// static void  glowClampColor(GPCLAMPCOLOR fnptr, GLenum  target, GLenum  clamp) {
//   (*fnptr)(target, clamp);
// }
// static void  glowClear(GPCLEAR fnptr, GLbitfield  mask) {
//   (*fnptr)(mask);
// }
// static void  glowClearAccumxOES(GPCLEARACCUMXOES fnptr, GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowClearBufferData(GPCLEARBUFFERDATA fnptr, GLenum  target, GLenum  internalformat, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(target, internalformat, format, type, data);
// }
// static void  glowClearBufferSubData(GPCLEARBUFFERSUBDATA fnptr, GLenum  target, GLenum  internalformat, GLintptr  offset, GLsizeiptr  size, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(target, internalformat, offset, size, format, type, data);
// }
// static void  glowClearBufferfi(GPCLEARBUFFERFI fnptr, GLenum  buffer, GLint  drawbuffer, GLfloat  depth, GLint  stencil) {
//   (*fnptr)(buffer, drawbuffer, depth, stencil);
// }
// static void  glowClearBufferfv(GPCLEARBUFFERFV fnptr, GLenum  buffer, GLint  drawbuffer, const GLfloat * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearBufferiv(GPCLEARBUFFERIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLint * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearBufferuiv(GPCLEARBUFFERUIV fnptr, GLenum  buffer, GLint  drawbuffer, const GLuint * value) {
//   (*fnptr)(buffer, drawbuffer, value);
// }
// static void  glowClearColor(GPCLEARCOLOR fnptr, GLfloat  red, GLfloat  green, GLfloat  blue, GLfloat  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowClearColorxOES(GPCLEARCOLORXOES fnptr, GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowClearDepth(GPCLEARDEPTH fnptr, GLdouble  depth) {
//   (*fnptr)(depth);
// }
// static void  glowClearDepthf(GPCLEARDEPTHF fnptr, GLfloat  d) {
//   (*fnptr)(d);
// }
// static void  glowClearDepthfOES(GPCLEARDEPTHFOES fnptr, GLclampf  depth) {
//   (*fnptr)(depth);
// }
// static void  glowClearDepthxOES(GPCLEARDEPTHXOES fnptr, GLfixed  depth) {
//   (*fnptr)(depth);
// }
// static void  glowClearNamedBufferData(GPCLEARNAMEDBUFFERDATA fnptr, GLuint  buffer, GLenum  internalformat, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(buffer, internalformat, format, type, data);
// }
// static void  glowClearNamedBufferSubData(GPCLEARNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLenum  internalformat, GLintptr  offset, GLsizei  size, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(buffer, internalformat, offset, size, format, type, data);
// }
// static void  glowClearNamedFramebufferfi(GPCLEARNAMEDFRAMEBUFFERFI fnptr, GLuint  framebuffer, GLenum  buffer, const GLfloat  depth, GLint  stencil) {
//   (*fnptr)(framebuffer, buffer, depth, stencil);
// }
// static void  glowClearNamedFramebufferfv(GPCLEARNAMEDFRAMEBUFFERFV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLfloat * value) {
//   (*fnptr)(framebuffer, buffer, drawbuffer, value);
// }
// static void  glowClearNamedFramebufferiv(GPCLEARNAMEDFRAMEBUFFERIV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLint * value) {
//   (*fnptr)(framebuffer, buffer, drawbuffer, value);
// }
// static void  glowClearNamedFramebufferuiv(GPCLEARNAMEDFRAMEBUFFERUIV fnptr, GLuint  framebuffer, GLenum  buffer, GLint  drawbuffer, const GLuint * value) {
//   (*fnptr)(framebuffer, buffer, drawbuffer, value);
// }
// static void  glowClearStencil(GPCLEARSTENCIL fnptr, GLint  s) {
//   (*fnptr)(s);
// }
// static void  glowClearTexImage(GPCLEARTEXIMAGE fnptr, GLuint  texture, GLint  level, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(texture, level, format, type, data);
// }
// static void  glowClearTexSubImage(GPCLEARTEXSUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * data) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, data);
// }
// static GLenum  glowClientWaitSync(GPCLIENTWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   return (*fnptr)(sync, flags, timeout);
// }
// static void  glowClipControl(GPCLIPCONTROL fnptr, GLenum  origin, GLenum  depth) {
//   (*fnptr)(origin, depth);
// }
// static void  glowClipPlanefOES(GPCLIPPLANEFOES fnptr, GLenum  plane, const GLfloat * equation) {
//   (*fnptr)(plane, equation);
// }
// static void  glowClipPlanexOES(GPCLIPPLANEXOES fnptr, GLenum  plane, const GLfixed * equation) {
//   (*fnptr)(plane, equation);
// }
// static void  glowColor3xOES(GPCOLOR3XOES fnptr, GLfixed  red, GLfixed  green, GLfixed  blue) {
//   (*fnptr)(red, green, blue);
// }
// static void  glowColor3xvOES(GPCOLOR3XVOES fnptr, const GLfixed * components) {
//   (*fnptr)(components);
// }
// static void  glowColor4xOES(GPCOLOR4XOES fnptr, GLfixed  red, GLfixed  green, GLfixed  blue, GLfixed  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColor4xvOES(GPCOLOR4XVOES fnptr, const GLfixed * components) {
//   (*fnptr)(components);
// }
// static void  glowColorMask(GPCOLORMASK fnptr, GLboolean  red, GLboolean  green, GLboolean  blue, GLboolean  alpha) {
//   (*fnptr)(red, green, blue, alpha);
// }
// static void  glowColorMaski(GPCOLORMASKI fnptr, GLuint  index, GLboolean  r, GLboolean  g, GLboolean  b, GLboolean  a) {
//   (*fnptr)(index, r, g, b, a);
// }
// static void  glowCompileShader(GPCOMPILESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static void  glowCompileShaderIncludeARB(GPCOMPILESHADERINCLUDEARB fnptr, GLuint  shader, GLsizei  count, const GLchar *const* path, const GLint * length) {
//   (*fnptr)(shader, count, path, length);
// }
// static void  glowCompressedTexImage1D(GPCOMPRESSEDTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, border, imageSize, data);
// }
// static void  glowCompressedTexImage2D(GPCOMPRESSEDTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, height, border, imageSize, data);
// }
// static void  glowCompressedTexImage3D(GPCOMPRESSEDTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, imageSize, data);
// }
// static void  glowCompressedTexSubImage1D(GPCOMPRESSEDTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, width, format, imageSize, data);
// }
// static void  glowCompressedTexSubImage2D(GPCOMPRESSEDTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, yoffset, width, height, format, imageSize, data);
// }
// static void  glowCompressedTexSubImage3D(GPCOMPRESSEDTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// static void  glowCompressedTextureSubImage1D(GPCOMPRESSEDTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(texture, level, xoffset, width, format, imageSize, data);
// }
// static void  glowCompressedTextureSubImage2D(GPCOMPRESSEDTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(texture, level, xoffset, yoffset, width, height, format, imageSize, data);
// }
// static void  glowCompressedTextureSubImage3D(GPCOMPRESSEDTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLsizei  imageSize, const void * data) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, imageSize, data);
// }
// static void  glowConvolutionParameterxOES(GPCONVOLUTIONPARAMETERXOES fnptr, GLenum  target, GLenum  pname, GLfixed  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowConvolutionParameterxvOES(GPCONVOLUTIONPARAMETERXVOES fnptr, GLenum  target, GLenum  pname, const GLfixed * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowCopyBufferSubData(GPCOPYBUFFERSUBDATA fnptr, GLenum  readTarget, GLenum  writeTarget, GLintptr  readOffset, GLintptr  writeOffset, GLsizeiptr  size) {
//   (*fnptr)(readTarget, writeTarget, readOffset, writeOffset, size);
// }
// static void  glowCopyImageSubData(GPCOPYIMAGESUBDATA fnptr, GLuint  srcName, GLenum  srcTarget, GLint  srcLevel, GLint  srcX, GLint  srcY, GLint  srcZ, GLuint  dstName, GLenum  dstTarget, GLint  dstLevel, GLint  dstX, GLint  dstY, GLint  dstZ, GLsizei  srcWidth, GLsizei  srcHeight, GLsizei  srcDepth) {
//   (*fnptr)(srcName, srcTarget, srcLevel, srcX, srcY, srcZ, dstName, dstTarget, dstLevel, dstX, dstY, dstZ, srcWidth, srcHeight, srcDepth);
// }
// static void  glowCopyNamedBufferSubData(GPCOPYNAMEDBUFFERSUBDATA fnptr, GLuint  readBuffer, GLuint  writeBuffer, GLintptr  readOffset, GLintptr  writeOffset, GLsizei  size) {
//   (*fnptr)(readBuffer, writeBuffer, readOffset, writeOffset, size);
// }
// static void  glowCopyTexImage1D(GPCOPYTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLint  border) {
//   (*fnptr)(target, level, internalformat, x, y, width, border);
// }
// static void  glowCopyTexImage2D(GPCOPYTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLenum  internalformat, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLint  border) {
//   (*fnptr)(target, level, internalformat, x, y, width, height, border);
// }
// static void  glowCopyTexSubImage1D(GPCOPYTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width) {
//   (*fnptr)(target, level, xoffset, x, y, width);
// }
// static void  glowCopyTexSubImage2D(GPCOPYTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, level, xoffset, yoffset, x, y, width, height);
// }
// static void  glowCopyTexSubImage3D(GPCOPYTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// static void  glowCopyTextureSubImage1D(GPCOPYTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  x, GLint  y, GLsizei  width) {
//   (*fnptr)(texture, level, xoffset, x, y, width);
// }
// static void  glowCopyTextureSubImage2D(GPCOPYTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(texture, level, xoffset, yoffset, x, y, width, height);
// }
// static void  glowCopyTextureSubImage3D(GPCOPYTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, x, y, width, height);
// }
// static void  glowCreateBuffers(GPCREATEBUFFERS fnptr, GLsizei  n, GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowCreateFramebuffers(GPCREATEFRAMEBUFFERS fnptr, GLsizei  n, GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static void  glowCreatePerfQueryINTEL(GPCREATEPERFQUERYINTEL fnptr, GLuint  queryId, GLuint * queryHandle) {
//   (*fnptr)(queryId, queryHandle);
// }
// static GLuint  glowCreateProgram(GPCREATEPROGRAM fnptr) {
//   return (*fnptr)();
// }
// static void  glowCreateProgramPipelines(GPCREATEPROGRAMPIPELINES fnptr, GLsizei  n, GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowCreateQueries(GPCREATEQUERIES fnptr, GLenum  target, GLsizei  n, GLuint * ids) {
//   (*fnptr)(target, n, ids);
// }
// static void  glowCreateRenderbuffers(GPCREATERENDERBUFFERS fnptr, GLsizei  n, GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowCreateSamplers(GPCREATESAMPLERS fnptr, GLsizei  n, GLuint * samplers) {
//   (*fnptr)(n, samplers);
// }
// static GLuint  glowCreateShader(GPCREATESHADER fnptr, GLenum  type) {
//   return (*fnptr)(type);
// }
// static GLuint  glowCreateShaderProgramEXT(GPCREATESHADERPROGRAMEXT fnptr, GLenum  type, const GLchar * string) {
//   return (*fnptr)(type, string);
// }
// static GLuint  glowCreateShaderProgramv(GPCREATESHADERPROGRAMV fnptr, GLenum  type, GLsizei  count, const GLchar *const* strings) {
//   return (*fnptr)(type, count, strings);
// }
// static GLuint  glowCreateShaderProgramvEXT(GPCREATESHADERPROGRAMVEXT fnptr, GLenum  type, GLsizei  count, const GLchar ** strings) {
//   return (*fnptr)(type, count, strings);
// }
// static GLsync  glowCreateSyncFromCLeventARB(GPCREATESYNCFROMCLEVENTARB fnptr, struct _cl_context * context, struct _cl_event * event, GLbitfield  flags) {
//   return (*fnptr)(context, event, flags);
// }
// static void  glowCreateTextures(GPCREATETEXTURES fnptr, GLenum  target, GLsizei  n, GLuint * textures) {
//   (*fnptr)(target, n, textures);
// }
// static void  glowCreateTransformFeedbacks(GPCREATETRANSFORMFEEDBACKS fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowCreateVertexArrays(GPCREATEVERTEXARRAYS fnptr, GLsizei  n, GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowCullFace(GPCULLFACE fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowDebugMessageCallback(GPDEBUGMESSAGECALLBACK fnptr, GLDEBUGPROC  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageCallbackARB(GPDEBUGMESSAGECALLBACKARB fnptr, GLDEBUGPROCARB  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageCallbackKHR(GPDEBUGMESSAGECALLBACKKHR fnptr, GLDEBUGPROCKHR  callback, const void * userParam) {
//   (*fnptr)(glowCDebugCallback, userParam);
// }
// static void  glowDebugMessageControl(GPDEBUGMESSAGECONTROL fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageControlARB(GPDEBUGMESSAGECONTROLARB fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageControlKHR(GPDEBUGMESSAGECONTROLKHR fnptr, GLenum  source, GLenum  type, GLenum  severity, GLsizei  count, const GLuint * ids, GLboolean  enabled) {
//   (*fnptr)(source, type, severity, count, ids, enabled);
// }
// static void  glowDebugMessageInsert(GPDEBUGMESSAGEINSERT fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDebugMessageInsertARB(GPDEBUGMESSAGEINSERTARB fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDebugMessageInsertKHR(GPDEBUGMESSAGEINSERTKHR fnptr, GLenum  source, GLenum  type, GLuint  id, GLenum  severity, GLsizei  length, const GLchar * buf) {
//   (*fnptr)(source, type, id, severity, length, buf);
// }
// static void  glowDeleteBuffers(GPDELETEBUFFERS fnptr, GLsizei  n, const GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowDeleteFencesNV(GPDELETEFENCESNV fnptr, GLsizei  n, const GLuint * fences) {
//   (*fnptr)(n, fences);
// }
// static void  glowDeleteFramebuffers(GPDELETEFRAMEBUFFERS fnptr, GLsizei  n, const GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static void  glowDeleteNamedStringARB(GPDELETENAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name) {
//   (*fnptr)(namelen, name);
// }
// static void  glowDeletePerfMonitorsAMD(GPDELETEPERFMONITORSAMD fnptr, GLsizei  n, GLuint * monitors) {
//   (*fnptr)(n, monitors);
// }
// static void  glowDeletePerfQueryINTEL(GPDELETEPERFQUERYINTEL fnptr, GLuint  queryHandle) {
//   (*fnptr)(queryHandle);
// }
// static void  glowDeleteProgram(GPDELETEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowDeleteProgramPipelines(GPDELETEPROGRAMPIPELINES fnptr, GLsizei  n, const GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowDeleteProgramPipelinesEXT(GPDELETEPROGRAMPIPELINESEXT fnptr, GLsizei  n, const GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowDeleteQueries(GPDELETEQUERIES fnptr, GLsizei  n, const GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowDeleteRenderbuffers(GPDELETERENDERBUFFERS fnptr, GLsizei  n, const GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowDeleteSamplers(GPDELETESAMPLERS fnptr, GLsizei  count, const GLuint * samplers) {
//   (*fnptr)(count, samplers);
// }
// static void  glowDeleteShader(GPDELETESHADER fnptr, GLuint  shader) {
//   (*fnptr)(shader);
// }
// static void  glowDeleteSync(GPDELETESYNC fnptr, GLsync  sync) {
//   (*fnptr)(sync);
// }
// static void  glowDeleteTextures(GPDELETETEXTURES fnptr, GLsizei  n, const GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static void  glowDeleteTransformFeedbacks(GPDELETETRANSFORMFEEDBACKS fnptr, GLsizei  n, const GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowDeleteVertexArrays(GPDELETEVERTEXARRAYS fnptr, GLsizei  n, const GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowDepthFunc(GPDEPTHFUNC fnptr, GLenum  func) {
//   (*fnptr)(func);
// }
// static void  glowDepthMask(GPDEPTHMASK fnptr, GLboolean  flag) {
//   (*fnptr)(flag);
// }
// static void  glowDepthRange(GPDEPTHRANGE fnptr, GLdouble  xnear, GLdouble  xfar) {
//   (*fnptr)(xnear, xfar);
// }
// static void  glowDepthRangeArrayv(GPDEPTHRANGEARRAYV fnptr, GLuint  first, GLsizei  count, const GLdouble * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowDepthRangeIndexed(GPDEPTHRANGEINDEXED fnptr, GLuint  index, GLdouble  n, GLdouble  f) {
//   (*fnptr)(index, n, f);
// }
// static void  glowDepthRangef(GPDEPTHRANGEF fnptr, GLfloat  n, GLfloat  f) {
//   (*fnptr)(n, f);
// }
// static void  glowDepthRangefOES(GPDEPTHRANGEFOES fnptr, GLclampf  n, GLclampf  f) {
//   (*fnptr)(n, f);
// }
// static void  glowDepthRangexOES(GPDEPTHRANGEXOES fnptr, GLfixed  n, GLfixed  f) {
//   (*fnptr)(n, f);
// }
// static void  glowDetachShader(GPDETACHSHADER fnptr, GLuint  program, GLuint  shader) {
//   (*fnptr)(program, shader);
// }
// static void  glowDisable(GPDISABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowDisableVertexArrayAttrib(GPDISABLEVERTEXARRAYATTRIB fnptr, GLuint  vaobj, GLuint  index) {
//   (*fnptr)(vaobj, index);
// }
// static void  glowDisableVertexAttribArray(GPDISABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowDisablei(GPDISABLEI fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowDispatchCompute(GPDISPATCHCOMPUTE fnptr, GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z) {
//   (*fnptr)(num_groups_x, num_groups_y, num_groups_z);
// }
// static void  glowDispatchComputeGroupSizeARB(GPDISPATCHCOMPUTEGROUPSIZEARB fnptr, GLuint  num_groups_x, GLuint  num_groups_y, GLuint  num_groups_z, GLuint  group_size_x, GLuint  group_size_y, GLuint  group_size_z) {
//   (*fnptr)(num_groups_x, num_groups_y, num_groups_z, group_size_x, group_size_y, group_size_z);
// }
// static void  glowDispatchComputeIndirect(GPDISPATCHCOMPUTEINDIRECT fnptr, GLintptr  indirect) {
//   (*fnptr)(indirect);
// }
// static void  glowDrawArrays(GPDRAWARRAYS fnptr, GLenum  mode, GLint  first, GLsizei  count) {
//   (*fnptr)(mode, first, count);
// }
// static void  glowDrawArraysIndirect(GPDRAWARRAYSINDIRECT fnptr, GLenum  mode, const void * indirect) {
//   (*fnptr)(mode, indirect);
// }
// static void  glowDrawArraysInstanced(GPDRAWARRAYSINSTANCED fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount) {
//   (*fnptr)(mode, first, count, instancecount);
// }
// static void  glowDrawArraysInstancedBaseInstance(GPDRAWARRAYSINSTANCEDBASEINSTANCE fnptr, GLenum  mode, GLint  first, GLsizei  count, GLsizei  instancecount, GLuint  baseinstance) {
//   (*fnptr)(mode, first, count, instancecount, baseinstance);
// }
// static void  glowDrawArraysInstancedEXT(GPDRAWARRAYSINSTANCEDEXT fnptr, GLenum  mode, GLint  start, GLsizei  count, GLsizei  primcount) {
//   (*fnptr)(mode, start, count, primcount);
// }
// static void  glowDrawBuffer(GPDRAWBUFFER fnptr, GLenum  buf) {
//   (*fnptr)(buf);
// }
// static void  glowDrawBuffers(GPDRAWBUFFERS fnptr, GLsizei  n, const GLenum * bufs) {
//   (*fnptr)(n, bufs);
// }
// static void  glowDrawElements(GPDRAWELEMENTS fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices) {
//   (*fnptr)(mode, count, type, indices);
// }
// static void  glowDrawElementsBaseVertex(GPDRAWELEMENTSBASEVERTEX fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, basevertex);
// }
// static void  glowDrawElementsIndirect(GPDRAWELEMENTSINDIRECT fnptr, GLenum  mode, GLenum  type, const void * indirect) {
//   (*fnptr)(mode, type, indirect);
// }
// static void  glowDrawElementsInstanced(GPDRAWELEMENTSINSTANCED fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount) {
//   (*fnptr)(mode, count, type, indices, instancecount);
// }
// static void  glowDrawElementsInstancedBaseInstance(GPDRAWELEMENTSINSTANCEDBASEINSTANCE fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLuint  baseinstance) {
//   (*fnptr)(mode, count, type, indices, instancecount, baseinstance);
// }
// static void  glowDrawElementsInstancedBaseVertex(GPDRAWELEMENTSINSTANCEDBASEVERTEX fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex) {
//   (*fnptr)(mode, count, type, indices, instancecount, basevertex);
// }
// static void  glowDrawElementsInstancedBaseVertexBaseInstance(GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  instancecount, GLint  basevertex, GLuint  baseinstance) {
//   (*fnptr)(mode, count, type, indices, instancecount, basevertex, baseinstance);
// }
// static void  glowDrawElementsInstancedEXT(GPDRAWELEMENTSINSTANCEDEXT fnptr, GLenum  mode, GLsizei  count, GLenum  type, const void * indices, GLsizei  primcount) {
//   (*fnptr)(mode, count, type, indices, primcount);
// }
// static void  glowDrawRangeElements(GPDRAWRANGEELEMENTS fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices) {
//   (*fnptr)(mode, start, end, count, type, indices);
// }
// static void  glowDrawRangeElementsBaseVertex(GPDRAWRANGEELEMENTSBASEVERTEX fnptr, GLenum  mode, GLuint  start, GLuint  end, GLsizei  count, GLenum  type, const void * indices, GLint  basevertex) {
//   (*fnptr)(mode, start, end, count, type, indices, basevertex);
// }
// static void  glowDrawTransformFeedback(GPDRAWTRANSFORMFEEDBACK fnptr, GLenum  mode, GLuint  id) {
//   (*fnptr)(mode, id);
// }
// static void  glowDrawTransformFeedbackInstanced(GPDRAWTRANSFORMFEEDBACKINSTANCED fnptr, GLenum  mode, GLuint  id, GLsizei  instancecount) {
//   (*fnptr)(mode, id, instancecount);
// }
// static void  glowDrawTransformFeedbackStream(GPDRAWTRANSFORMFEEDBACKSTREAM fnptr, GLenum  mode, GLuint  id, GLuint  stream) {
//   (*fnptr)(mode, id, stream);
// }
// static void  glowDrawTransformFeedbackStreamInstanced(GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED fnptr, GLenum  mode, GLuint  id, GLuint  stream, GLsizei  instancecount) {
//   (*fnptr)(mode, id, stream, instancecount);
// }
// static void  glowEnable(GPENABLE fnptr, GLenum  cap) {
//   (*fnptr)(cap);
// }
// static void  glowEnableVertexArrayAttrib(GPENABLEVERTEXARRAYATTRIB fnptr, GLuint  vaobj, GLuint  index) {
//   (*fnptr)(vaobj, index);
// }
// static void  glowEnableVertexAttribArray(GPENABLEVERTEXATTRIBARRAY fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowEnablei(GPENABLEI fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowEndConditionalRender(GPENDCONDITIONALRENDER fnptr) {
//   (*fnptr)();
// }
// static void  glowEndPerfMonitorAMD(GPENDPERFMONITORAMD fnptr, GLuint  monitor) {
//   (*fnptr)(monitor);
// }
// static void  glowEndPerfQueryINTEL(GPENDPERFQUERYINTEL fnptr, GLuint  queryHandle) {
//   (*fnptr)(queryHandle);
// }
// static void  glowEndQuery(GPENDQUERY fnptr, GLenum  target) {
//   (*fnptr)(target);
// }
// static void  glowEndQueryIndexed(GPENDQUERYINDEXED fnptr, GLenum  target, GLuint  index) {
//   (*fnptr)(target, index);
// }
// static void  glowEndTransformFeedback(GPENDTRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowEvalCoord1xOES(GPEVALCOORD1XOES fnptr, GLfixed  u) {
//   (*fnptr)(u);
// }
// static void  glowEvalCoord1xvOES(GPEVALCOORD1XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowEvalCoord2xOES(GPEVALCOORD2XOES fnptr, GLfixed  u, GLfixed  v) {
//   (*fnptr)(u, v);
// }
// static void  glowEvalCoord2xvOES(GPEVALCOORD2XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowFeedbackBufferxOES(GPFEEDBACKBUFFERXOES fnptr, GLsizei  n, GLenum  type, const GLfixed * buffer) {
//   (*fnptr)(n, type, buffer);
// }
// static GLsync  glowFenceSync(GPFENCESYNC fnptr, GLenum  condition, GLbitfield  flags) {
//   return (*fnptr)(condition, flags);
// }
// static void  glowFinish(GPFINISH fnptr) {
//   (*fnptr)();
// }
// static void  glowFinishFenceNV(GPFINISHFENCENV fnptr, GLuint  fence) {
//   (*fnptr)(fence);
// }
// static void  glowFlush(GPFLUSH fnptr) {
//   (*fnptr)();
// }
// static void  glowFlushMappedBufferRange(GPFLUSHMAPPEDBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length) {
//   (*fnptr)(target, offset, length);
// }
// static void  glowFlushMappedNamedBufferRange(GPFLUSHMAPPEDNAMEDBUFFERRANGE fnptr, GLuint  buffer, GLintptr  offset, GLsizei  length) {
//   (*fnptr)(buffer, offset, length);
// }
// static void  glowFogxOES(GPFOGXOES fnptr, GLenum  pname, GLfixed  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowFogxvOES(GPFOGXVOES fnptr, GLenum  pname, const GLfixed * param) {
//   (*fnptr)(pname, param);
// }
// static void  glowFramebufferParameteri(GPFRAMEBUFFERPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowFramebufferRenderbuffer(GPFRAMEBUFFERRENDERBUFFER fnptr, GLenum  target, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {
//   (*fnptr)(target, attachment, renderbuffertarget, renderbuffer);
// }
// static void  glowFramebufferTexture(GPFRAMEBUFFERTEXTURE fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, texture, level);
// }
// static void  glowFramebufferTexture1D(GPFRAMEBUFFERTEXTURE1D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, textarget, texture, level);
// }
// static void  glowFramebufferTexture2D(GPFRAMEBUFFERTEXTURE2D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level) {
//   (*fnptr)(target, attachment, textarget, texture, level);
// }
// static void  glowFramebufferTexture3D(GPFRAMEBUFFERTEXTURE3D fnptr, GLenum  target, GLenum  attachment, GLenum  textarget, GLuint  texture, GLint  level, GLint  zoffset) {
//   (*fnptr)(target, attachment, textarget, texture, level, zoffset);
// }
// static void  glowFramebufferTextureLayer(GPFRAMEBUFFERTEXTURELAYER fnptr, GLenum  target, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer) {
//   (*fnptr)(target, attachment, texture, level, layer);
// }
// static void  glowFrontFace(GPFRONTFACE fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowFrustumfOES(GPFRUSTUMFOES fnptr, GLfloat  l, GLfloat  r, GLfloat  b, GLfloat  t, GLfloat  n, GLfloat  f) {
//   (*fnptr)(l, r, b, t, n, f);
// }
// static void  glowFrustumxOES(GPFRUSTUMXOES fnptr, GLfixed  l, GLfixed  r, GLfixed  b, GLfixed  t, GLfixed  n, GLfixed  f) {
//   (*fnptr)(l, r, b, t, n, f);
// }
// static void  glowGenBuffers(GPGENBUFFERS fnptr, GLsizei  n, GLuint * buffers) {
//   (*fnptr)(n, buffers);
// }
// static void  glowGenFencesNV(GPGENFENCESNV fnptr, GLsizei  n, GLuint * fences) {
//   (*fnptr)(n, fences);
// }
// static void  glowGenFramebuffers(GPGENFRAMEBUFFERS fnptr, GLsizei  n, GLuint * framebuffers) {
//   (*fnptr)(n, framebuffers);
// }
// static void  glowGenPerfMonitorsAMD(GPGENPERFMONITORSAMD fnptr, GLsizei  n, GLuint * monitors) {
//   (*fnptr)(n, monitors);
// }
// static void  glowGenProgramPipelines(GPGENPROGRAMPIPELINES fnptr, GLsizei  n, GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowGenProgramPipelinesEXT(GPGENPROGRAMPIPELINESEXT fnptr, GLsizei  n, GLuint * pipelines) {
//   (*fnptr)(n, pipelines);
// }
// static void  glowGenQueries(GPGENQUERIES fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowGenRenderbuffers(GPGENRENDERBUFFERS fnptr, GLsizei  n, GLuint * renderbuffers) {
//   (*fnptr)(n, renderbuffers);
// }
// static void  glowGenSamplers(GPGENSAMPLERS fnptr, GLsizei  count, GLuint * samplers) {
//   (*fnptr)(count, samplers);
// }
// static void  glowGenTextures(GPGENTEXTURES fnptr, GLsizei  n, GLuint * textures) {
//   (*fnptr)(n, textures);
// }
// static void  glowGenTransformFeedbacks(GPGENTRANSFORMFEEDBACKS fnptr, GLsizei  n, GLuint * ids) {
//   (*fnptr)(n, ids);
// }
// static void  glowGenVertexArrays(GPGENVERTEXARRAYS fnptr, GLsizei  n, GLuint * arrays) {
//   (*fnptr)(n, arrays);
// }
// static void  glowGenerateMipmap(GPGENERATEMIPMAP fnptr, GLenum  target) {
//   (*fnptr)(target);
// }
// static void  glowGenerateTextureMipmap(GPGENERATETEXTUREMIPMAP fnptr, GLuint  texture) {
//   (*fnptr)(texture);
// }
// static void  glowGetActiveAtomicCounterBufferiv(GPGETACTIVEATOMICCOUNTERBUFFERIV fnptr, GLuint  program, GLuint  bufferIndex, GLenum  pname, GLint * params) {
//   (*fnptr)(program, bufferIndex, pname, params);
// }
// static void  glowGetActiveAttrib(GPGETACTIVEATTRIB fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetActiveSubroutineName(GPGETACTIVESUBROUTINENAME fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name) {
//   (*fnptr)(program, shadertype, index, bufsize, length, name);
// }
// static void  glowGetActiveSubroutineUniformName(GPGETACTIVESUBROUTINEUNIFORMNAME fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLsizei  bufsize, GLsizei * length, GLchar * name) {
//   (*fnptr)(program, shadertype, index, bufsize, length, name);
// }
// static void  glowGetActiveSubroutineUniformiv(GPGETACTIVESUBROUTINEUNIFORMIV fnptr, GLuint  program, GLenum  shadertype, GLuint  index, GLenum  pname, GLint * values) {
//   (*fnptr)(program, shadertype, index, pname, values);
// }
// static void  glowGetActiveUniform(GPGETACTIVEUNIFORM fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLint * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetActiveUniformBlockName(GPGETACTIVEUNIFORMBLOCKNAME fnptr, GLuint  program, GLuint  uniformBlockIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformBlockName) {
//   (*fnptr)(program, uniformBlockIndex, bufSize, length, uniformBlockName);
// }
// static void  glowGetActiveUniformBlockiv(GPGETACTIVEUNIFORMBLOCKIV fnptr, GLuint  program, GLuint  uniformBlockIndex, GLenum  pname, GLint * params) {
//   (*fnptr)(program, uniformBlockIndex, pname, params);
// }
// static void  glowGetActiveUniformName(GPGETACTIVEUNIFORMNAME fnptr, GLuint  program, GLuint  uniformIndex, GLsizei  bufSize, GLsizei * length, GLchar * uniformName) {
//   (*fnptr)(program, uniformIndex, bufSize, length, uniformName);
// }
// static void  glowGetActiveUniformsiv(GPGETACTIVEUNIFORMSIV fnptr, GLuint  program, GLsizei  uniformCount, const GLuint * uniformIndices, GLenum  pname, GLint * params) {
//   (*fnptr)(program, uniformCount, uniformIndices, pname, params);
// }
// static void  glowGetAttachedShaders(GPGETATTACHEDSHADERS fnptr, GLuint  program, GLsizei  maxCount, GLsizei * count, GLuint * shaders) {
//   (*fnptr)(program, maxCount, count, shaders);
// }
// static GLint  glowGetAttribLocation(GPGETATTRIBLOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetBooleani_v(GPGETBOOLEANI_V fnptr, GLenum  target, GLuint  index, GLboolean * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetBooleanv(GPGETBOOLEANV fnptr, GLenum  pname, GLboolean * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetBufferParameteri64v(GPGETBUFFERPARAMETERI64V fnptr, GLenum  target, GLenum  pname, GLint64 * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferParameteriv(GPGETBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferPointerv(GPGETBUFFERPOINTERV fnptr, GLenum  target, GLenum  pname, void ** params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetBufferSubData(GPGETBUFFERSUBDATA fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  size, void * data) {
//   (*fnptr)(target, offset, size, data);
// }
// static void  glowGetClipPlanefOES(GPGETCLIPPLANEFOES fnptr, GLenum  plane, GLfloat * equation) {
//   (*fnptr)(plane, equation);
// }
// static void  glowGetClipPlanexOES(GPGETCLIPPLANEXOES fnptr, GLenum  plane, GLfixed * equation) {
//   (*fnptr)(plane, equation);
// }
// static void  glowGetCompressedTexImage(GPGETCOMPRESSEDTEXIMAGE fnptr, GLenum  target, GLint  level, void * img) {
//   (*fnptr)(target, level, img);
// }
// static void  glowGetCompressedTextureImage(GPGETCOMPRESSEDTEXTUREIMAGE fnptr, GLuint  texture, GLint  level, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, bufSize, pixels);
// }
// static void  glowGetCompressedTextureSubImage(GPGETCOMPRESSEDTEXTURESUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, bufSize, pixels);
// }
// static void  glowGetConvolutionParameterxvOES(GPGETCONVOLUTIONPARAMETERXVOES fnptr, GLenum  target, GLenum  pname, GLfixed * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLuint  glowGetDebugMessageLog(GPGETDEBUGMESSAGELOG fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static GLuint  glowGetDebugMessageLogARB(GPGETDEBUGMESSAGELOGARB fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static GLuint  glowGetDebugMessageLogKHR(GPGETDEBUGMESSAGELOGKHR fnptr, GLuint  count, GLsizei  bufSize, GLenum * sources, GLenum * types, GLuint * ids, GLenum * severities, GLsizei * lengths, GLchar * messageLog) {
//   return (*fnptr)(count, bufSize, sources, types, ids, severities, lengths, messageLog);
// }
// static void  glowGetDoublei_v(GPGETDOUBLEI_V fnptr, GLenum  target, GLuint  index, GLdouble * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetDoublev(GPGETDOUBLEV fnptr, GLenum  pname, GLdouble * data) {
//   (*fnptr)(pname, data);
// }
// static GLenum  glowGetError(GPGETERROR fnptr) {
//   return (*fnptr)();
// }
// static void  glowGetFenceivNV(GPGETFENCEIVNV fnptr, GLuint  fence, GLenum  pname, GLint * params) {
//   (*fnptr)(fence, pname, params);
// }
// static void  glowGetFirstPerfQueryIdINTEL(GPGETFIRSTPERFQUERYIDINTEL fnptr, GLuint * queryId) {
//   (*fnptr)(queryId);
// }
// static void  glowGetFixedvOES(GPGETFIXEDVOES fnptr, GLenum  pname, GLfixed * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetFloati_v(GPGETFLOATI_V fnptr, GLenum  target, GLuint  index, GLfloat * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetFloatv(GPGETFLOATV fnptr, GLenum  pname, GLfloat * data) {
//   (*fnptr)(pname, data);
// }
// static GLint  glowGetFragDataIndex(GPGETFRAGDATAINDEX fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static GLint  glowGetFragDataLocation(GPGETFRAGDATALOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetFramebufferAttachmentParameteriv(GPGETFRAMEBUFFERATTACHMENTPARAMETERIV fnptr, GLenum  target, GLenum  attachment, GLenum  pname, GLint * params) {
//   (*fnptr)(target, attachment, pname, params);
// }
// static void  glowGetFramebufferParameteriv(GPGETFRAMEBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLenum  glowGetGraphicsResetStatus(GPGETGRAPHICSRESETSTATUS fnptr) {
//   return (*fnptr)();
// }
// static GLenum  glowGetGraphicsResetStatusARB(GPGETGRAPHICSRESETSTATUSARB fnptr) {
//   return (*fnptr)();
// }
// static GLenum  glowGetGraphicsResetStatusKHR(GPGETGRAPHICSRESETSTATUSKHR fnptr) {
//   return (*fnptr)();
// }
// static void  glowGetHistogramParameterxvOES(GPGETHISTOGRAMPARAMETERXVOES fnptr, GLenum  target, GLenum  pname, GLfixed * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLuint64  glowGetImageHandleARB(GPGETIMAGEHANDLEARB fnptr, GLuint  texture, GLint  level, GLboolean  layered, GLint  layer, GLenum  format) {
//   return (*fnptr)(texture, level, layered, layer, format);
// }
// static void  glowGetInteger64i_v(GPGETINTEGER64I_V fnptr, GLenum  target, GLuint  index, GLint64 * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetInteger64v(GPGETINTEGER64V fnptr, GLenum  pname, GLint64 * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetIntegeri_v(GPGETINTEGERI_V fnptr, GLenum  target, GLuint  index, GLint * data) {
//   (*fnptr)(target, index, data);
// }
// static void  glowGetIntegerv(GPGETINTEGERV fnptr, GLenum  pname, GLint * data) {
//   (*fnptr)(pname, data);
// }
// static void  glowGetInternalformati64v(GPGETINTERNALFORMATI64V fnptr, GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint64 * params) {
//   (*fnptr)(target, internalformat, pname, bufSize, params);
// }
// static void  glowGetInternalformativ(GPGETINTERNALFORMATIV fnptr, GLenum  target, GLenum  internalformat, GLenum  pname, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(target, internalformat, pname, bufSize, params);
// }
// static void  glowGetLightxOES(GPGETLIGHTXOES fnptr, GLenum  light, GLenum  pname, GLfixed * params) {
//   (*fnptr)(light, pname, params);
// }
// static void  glowGetLightxvOES(GPGETLIGHTXVOES fnptr, GLenum  light, GLenum  pname, GLfixed * params) {
//   (*fnptr)(light, pname, params);
// }
// static void  glowGetMapxvOES(GPGETMAPXVOES fnptr, GLenum  target, GLenum  query, GLfixed * v) {
//   (*fnptr)(target, query, v);
// }
// static void  glowGetMaterialxOES(GPGETMATERIALXOES fnptr, GLenum  face, GLenum  pname, GLfixed  param) {
//   (*fnptr)(face, pname, param);
// }
// static void  glowGetMaterialxvOES(GPGETMATERIALXVOES fnptr, GLenum  face, GLenum  pname, GLfixed * params) {
//   (*fnptr)(face, pname, params);
// }
// static void  glowGetMultisamplefv(GPGETMULTISAMPLEFV fnptr, GLenum  pname, GLuint  index, GLfloat * val) {
//   (*fnptr)(pname, index, val);
// }
// static void  glowGetNamedBufferParameteri64v(GPGETNAMEDBUFFERPARAMETERI64V fnptr, GLuint  buffer, GLenum  pname, GLint64 * params) {
//   (*fnptr)(buffer, pname, params);
// }
// static void  glowGetNamedBufferParameteriv(GPGETNAMEDBUFFERPARAMETERIV fnptr, GLuint  buffer, GLenum  pname, GLint * params) {
//   (*fnptr)(buffer, pname, params);
// }
// static void  glowGetNamedBufferPointerv(GPGETNAMEDBUFFERPOINTERV fnptr, GLuint  buffer, GLenum  pname, void ** params) {
//   (*fnptr)(buffer, pname, params);
// }
// static void  glowGetNamedBufferSubData(GPGETNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, void * data) {
//   (*fnptr)(buffer, offset, size, data);
// }
// static void  glowGetNamedFramebufferAttachmentParameteriv(GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV fnptr, GLuint  framebuffer, GLenum  attachment, GLenum  pname, GLint * params) {
//   (*fnptr)(framebuffer, attachment, pname, params);
// }
// static void  glowGetNamedFramebufferParameteriv(GPGETNAMEDFRAMEBUFFERPARAMETERIV fnptr, GLuint  framebuffer, GLenum  pname, GLint * param) {
//   (*fnptr)(framebuffer, pname, param);
// }
// static void  glowGetNamedRenderbufferParameteriv(GPGETNAMEDRENDERBUFFERPARAMETERIV fnptr, GLuint  renderbuffer, GLenum  pname, GLint * params) {
//   (*fnptr)(renderbuffer, pname, params);
// }
// static void  glowGetNamedStringARB(GPGETNAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name, GLsizei  bufSize, GLint * stringlen, GLchar * string) {
//   (*fnptr)(namelen, name, bufSize, stringlen, string);
// }
// static void  glowGetNamedStringivARB(GPGETNAMEDSTRINGIVARB fnptr, GLint  namelen, const GLchar * name, GLenum  pname, GLint * params) {
//   (*fnptr)(namelen, name, pname, params);
// }
// static void  glowGetNextPerfQueryIdINTEL(GPGETNEXTPERFQUERYIDINTEL fnptr, GLuint  queryId, GLuint * nextQueryId) {
//   (*fnptr)(queryId, nextQueryId);
// }
// static void  glowGetObjectLabel(GPGETOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(identifier, name, bufSize, length, label);
// }
// static void  glowGetObjectLabelEXT(GPGETOBJECTLABELEXT fnptr, GLenum  type, GLuint  object, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(type, object, bufSize, length, label);
// }
// static void  glowGetObjectLabelKHR(GPGETOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(identifier, name, bufSize, length, label);
// }
// static void  glowGetObjectPtrLabel(GPGETOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(ptr, bufSize, length, label);
// }
// static void  glowGetObjectPtrLabelKHR(GPGETOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  bufSize, GLsizei * length, GLchar * label) {
//   (*fnptr)(ptr, bufSize, length, label);
// }
// static void  glowGetPerfCounterInfoINTEL(GPGETPERFCOUNTERINFOINTEL fnptr, GLuint  queryId, GLuint  counterId, GLuint  counterNameLength, GLchar * counterName, GLuint  counterDescLength, GLchar * counterDesc, GLuint * counterOffset, GLuint * counterDataSize, GLuint * counterTypeEnum, GLuint * counterDataTypeEnum, GLuint64 * rawCounterMaxValue) {
//   (*fnptr)(queryId, counterId, counterNameLength, counterName, counterDescLength, counterDesc, counterOffset, counterDataSize, counterTypeEnum, counterDataTypeEnum, rawCounterMaxValue);
// }
// static void  glowGetPerfMonitorCounterDataAMD(GPGETPERFMONITORCOUNTERDATAAMD fnptr, GLuint  monitor, GLenum  pname, GLsizei  dataSize, GLuint * data, GLint * bytesWritten) {
//   (*fnptr)(monitor, pname, dataSize, data, bytesWritten);
// }
// static void  glowGetPerfMonitorCounterInfoAMD(GPGETPERFMONITORCOUNTERINFOAMD fnptr, GLuint  group, GLuint  counter, GLenum  pname, void * data) {
//   (*fnptr)(group, counter, pname, data);
// }
// static void  glowGetPerfMonitorCounterStringAMD(GPGETPERFMONITORCOUNTERSTRINGAMD fnptr, GLuint  group, GLuint  counter, GLsizei  bufSize, GLsizei * length, GLchar * counterString) {
//   (*fnptr)(group, counter, bufSize, length, counterString);
// }
// static void  glowGetPerfMonitorCountersAMD(GPGETPERFMONITORCOUNTERSAMD fnptr, GLuint  group, GLint * numCounters, GLint * maxActiveCounters, GLsizei  counterSize, GLuint * counters) {
//   (*fnptr)(group, numCounters, maxActiveCounters, counterSize, counters);
// }
// static void  glowGetPerfMonitorGroupStringAMD(GPGETPERFMONITORGROUPSTRINGAMD fnptr, GLuint  group, GLsizei  bufSize, GLsizei * length, GLchar * groupString) {
//   (*fnptr)(group, bufSize, length, groupString);
// }
// static void  glowGetPerfMonitorGroupsAMD(GPGETPERFMONITORGROUPSAMD fnptr, GLint * numGroups, GLsizei  groupsSize, GLuint * groups) {
//   (*fnptr)(numGroups, groupsSize, groups);
// }
// static void  glowGetPerfQueryDataINTEL(GPGETPERFQUERYDATAINTEL fnptr, GLuint  queryHandle, GLuint  flags, GLsizei  dataSize, GLvoid * data, GLuint * bytesWritten) {
//   (*fnptr)(queryHandle, flags, dataSize, data, bytesWritten);
// }
// static void  glowGetPerfQueryIdByNameINTEL(GPGETPERFQUERYIDBYNAMEINTEL fnptr, GLchar * queryName, GLuint * queryId) {
//   (*fnptr)(queryName, queryId);
// }
// static void  glowGetPerfQueryInfoINTEL(GPGETPERFQUERYINFOINTEL fnptr, GLuint  queryId, GLuint  queryNameLength, GLchar * queryName, GLuint * dataSize, GLuint * noCounters, GLuint * noInstances, GLuint * capsMask) {
//   (*fnptr)(queryId, queryNameLength, queryName, dataSize, noCounters, noInstances, capsMask);
// }
// static void  glowGetPixelMapxv(GPGETPIXELMAPXV fnptr, GLenum  map, GLint  size, GLfixed * values) {
//   (*fnptr)(map, size, values);
// }
// static void  glowGetPointerv(GPGETPOINTERV fnptr, GLenum  pname, void ** params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetPointervKHR(GPGETPOINTERVKHR fnptr, GLenum  pname, void ** params) {
//   (*fnptr)(pname, params);
// }
// static void  glowGetProgramBinary(GPGETPROGRAMBINARY fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLenum * binaryFormat, void * binary) {
//   (*fnptr)(program, bufSize, length, binaryFormat, binary);
// }
// static void  glowGetProgramInfoLog(GPGETPROGRAMINFOLOG fnptr, GLuint  program, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(program, bufSize, length, infoLog);
// }
// static void  glowGetProgramInterfaceiv(GPGETPROGRAMINTERFACEIV fnptr, GLuint  program, GLenum  programInterface, GLenum  pname, GLint * params) {
//   (*fnptr)(program, programInterface, pname, params);
// }
// static void  glowGetProgramPipelineInfoLog(GPGETPROGRAMPIPELINEINFOLOG fnptr, GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(pipeline, bufSize, length, infoLog);
// }
// static void  glowGetProgramPipelineInfoLogEXT(GPGETPROGRAMPIPELINEINFOLOGEXT fnptr, GLuint  pipeline, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(pipeline, bufSize, length, infoLog);
// }
// static void  glowGetProgramPipelineiv(GPGETPROGRAMPIPELINEIV fnptr, GLuint  pipeline, GLenum  pname, GLint * params) {
//   (*fnptr)(pipeline, pname, params);
// }
// static void  glowGetProgramPipelineivEXT(GPGETPROGRAMPIPELINEIVEXT fnptr, GLuint  pipeline, GLenum  pname, GLint * params) {
//   (*fnptr)(pipeline, pname, params);
// }
// static GLuint  glowGetProgramResourceIndex(GPGETPROGRAMRESOURCEINDEX fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static GLint  glowGetProgramResourceLocation(GPGETPROGRAMRESOURCELOCATION fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static GLint  glowGetProgramResourceLocationIndex(GPGETPROGRAMRESOURCELOCATIONINDEX fnptr, GLuint  program, GLenum  programInterface, const GLchar * name) {
//   return (*fnptr)(program, programInterface, name);
// }
// static void  glowGetProgramResourceName(GPGETPROGRAMRESOURCENAME fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  bufSize, GLsizei * length, GLchar * name) {
//   (*fnptr)(program, programInterface, index, bufSize, length, name);
// }
// static void  glowGetProgramResourceiv(GPGETPROGRAMRESOURCEIV fnptr, GLuint  program, GLenum  programInterface, GLuint  index, GLsizei  propCount, const GLenum * props, GLsizei  bufSize, GLsizei * length, GLint * params) {
//   (*fnptr)(program, programInterface, index, propCount, props, bufSize, length, params);
// }
// static void  glowGetProgramStageiv(GPGETPROGRAMSTAGEIV fnptr, GLuint  program, GLenum  shadertype, GLenum  pname, GLint * values) {
//   (*fnptr)(program, shadertype, pname, values);
// }
// static void  glowGetProgramiv(GPGETPROGRAMIV fnptr, GLuint  program, GLenum  pname, GLint * params) {
//   (*fnptr)(program, pname, params);
// }
// static void  glowGetQueryIndexediv(GPGETQUERYINDEXEDIV fnptr, GLenum  target, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(target, index, pname, params);
// }
// static void  glowGetQueryObjecti64v(GPGETQUERYOBJECTI64V fnptr, GLuint  id, GLenum  pname, GLint64 * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectiv(GPGETQUERYOBJECTIV fnptr, GLuint  id, GLenum  pname, GLint * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectui64v(GPGETQUERYOBJECTUI64V fnptr, GLuint  id, GLenum  pname, GLuint64 * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryObjectuiv(GPGETQUERYOBJECTUIV fnptr, GLuint  id, GLenum  pname, GLuint * params) {
//   (*fnptr)(id, pname, params);
// }
// static void  glowGetQueryiv(GPGETQUERYIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetRenderbufferParameteriv(GPGETRENDERBUFFERPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetSamplerParameterIiv(GPGETSAMPLERPARAMETERIIV fnptr, GLuint  sampler, GLenum  pname, GLint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterIuiv(GPGETSAMPLERPARAMETERIUIV fnptr, GLuint  sampler, GLenum  pname, GLuint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameterfv(GPGETSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, GLfloat * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetSamplerParameteriv(GPGETSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, GLint * params) {
//   (*fnptr)(sampler, pname, params);
// }
// static void  glowGetShaderInfoLog(GPGETSHADERINFOLOG fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * infoLog) {
//   (*fnptr)(shader, bufSize, length, infoLog);
// }
// static void  glowGetShaderPrecisionFormat(GPGETSHADERPRECISIONFORMAT fnptr, GLenum  shadertype, GLenum  precisiontype, GLint * range, GLint * precision) {
//   (*fnptr)(shadertype, precisiontype, range, precision);
// }
// static void  glowGetShaderSource(GPGETSHADERSOURCE fnptr, GLuint  shader, GLsizei  bufSize, GLsizei * length, GLchar * source) {
//   (*fnptr)(shader, bufSize, length, source);
// }
// static void  glowGetShaderiv(GPGETSHADERIV fnptr, GLuint  shader, GLenum  pname, GLint * params) {
//   (*fnptr)(shader, pname, params);
// }
// static const GLubyte * glowGetString(GPGETSTRING fnptr, GLenum  name) {
//   return (*fnptr)(name);
// }
// static const GLubyte * glowGetStringi(GPGETSTRINGI fnptr, GLenum  name, GLuint  index) {
//   return (*fnptr)(name, index);
// }
// static GLuint  glowGetSubroutineIndex(GPGETSUBROUTINEINDEX fnptr, GLuint  program, GLenum  shadertype, const GLchar * name) {
//   return (*fnptr)(program, shadertype, name);
// }
// static GLint  glowGetSubroutineUniformLocation(GPGETSUBROUTINEUNIFORMLOCATION fnptr, GLuint  program, GLenum  shadertype, const GLchar * name) {
//   return (*fnptr)(program, shadertype, name);
// }
// static void  glowGetSynciv(GPGETSYNCIV fnptr, GLsync  sync, GLenum  pname, GLsizei  bufSize, GLsizei * length, GLint * values) {
//   (*fnptr)(sync, pname, bufSize, length, values);
// }
// static void  glowGetTexEnvxvOES(GPGETTEXENVXVOES fnptr, GLenum  target, GLenum  pname, GLfixed * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexGenxvOES(GPGETTEXGENXVOES fnptr, GLenum  coord, GLenum  pname, GLfixed * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowGetTexImage(GPGETTEXIMAGE fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, void * pixels) {
//   (*fnptr)(target, level, format, type, pixels);
// }
// static void  glowGetTexLevelParameterfv(GPGETTEXLEVELPARAMETERFV fnptr, GLenum  target, GLint  level, GLenum  pname, GLfloat * params) {
//   (*fnptr)(target, level, pname, params);
// }
// static void  glowGetTexLevelParameteriv(GPGETTEXLEVELPARAMETERIV fnptr, GLenum  target, GLint  level, GLenum  pname, GLint * params) {
//   (*fnptr)(target, level, pname, params);
// }
// static void  glowGetTexLevelParameterxvOES(GPGETTEXLEVELPARAMETERXVOES fnptr, GLenum  target, GLint  level, GLenum  pname, GLfixed * params) {
//   (*fnptr)(target, level, pname, params);
// }
// static void  glowGetTexParameterIiv(GPGETTEXPARAMETERIIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterIuiv(GPGETTEXPARAMETERIUIV fnptr, GLenum  target, GLenum  pname, GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterfv(GPGETTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameteriv(GPGETTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowGetTexParameterxvOES(GPGETTEXPARAMETERXVOES fnptr, GLenum  target, GLenum  pname, GLfixed * params) {
//   (*fnptr)(target, pname, params);
// }
// static GLuint64  glowGetTextureHandleARB(GPGETTEXTUREHANDLEARB fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static void  glowGetTextureImage(GPGETTEXTUREIMAGE fnptr, GLuint  texture, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, format, type, bufSize, pixels);
// }
// static void  glowGetTextureLevelParameterfv(GPGETTEXTURELEVELPARAMETERFV fnptr, GLuint  texture, GLint  level, GLenum  pname, GLfloat * params) {
//   (*fnptr)(texture, level, pname, params);
// }
// static void  glowGetTextureLevelParameteriv(GPGETTEXTURELEVELPARAMETERIV fnptr, GLuint  texture, GLint  level, GLenum  pname, GLint * params) {
//   (*fnptr)(texture, level, pname, params);
// }
// static void  glowGetTextureParameterIiv(GPGETTEXTUREPARAMETERIIV fnptr, GLuint  texture, GLenum  pname, GLint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowGetTextureParameterIuiv(GPGETTEXTUREPARAMETERIUIV fnptr, GLuint  texture, GLenum  pname, GLuint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowGetTextureParameterfv(GPGETTEXTUREPARAMETERFV fnptr, GLuint  texture, GLenum  pname, GLfloat * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowGetTextureParameteriv(GPGETTEXTUREPARAMETERIV fnptr, GLuint  texture, GLenum  pname, GLint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static GLuint64  glowGetTextureSamplerHandleARB(GPGETTEXTURESAMPLERHANDLEARB fnptr, GLuint  texture, GLuint  sampler) {
//   return (*fnptr)(texture, sampler);
// }
// static void  glowGetTextureSubImage(GPGETTEXTURESUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, GLsizei  bufSize, void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, bufSize, pixels);
// }
// static void  glowGetTransformFeedbackVarying(GPGETTRANSFORMFEEDBACKVARYING fnptr, GLuint  program, GLuint  index, GLsizei  bufSize, GLsizei * length, GLsizei * size, GLenum * type, GLchar * name) {
//   (*fnptr)(program, index, bufSize, length, size, type, name);
// }
// static void  glowGetTransformFeedbacki64_v(GPGETTRANSFORMFEEDBACKI64_V fnptr, GLuint  xfb, GLenum  pname, GLuint  index, GLint64 * param) {
//   (*fnptr)(xfb, pname, index, param);
// }
// static void  glowGetTransformFeedbacki_v(GPGETTRANSFORMFEEDBACKI_V fnptr, GLuint  xfb, GLenum  pname, GLuint  index, GLint * param) {
//   (*fnptr)(xfb, pname, index, param);
// }
// static void  glowGetTransformFeedbackiv(GPGETTRANSFORMFEEDBACKIV fnptr, GLuint  xfb, GLenum  pname, GLint * param) {
//   (*fnptr)(xfb, pname, param);
// }
// static GLuint  glowGetUniformBlockIndex(GPGETUNIFORMBLOCKINDEX fnptr, GLuint  program, const GLchar * uniformBlockName) {
//   return (*fnptr)(program, uniformBlockName);
// }
// static void  glowGetUniformIndices(GPGETUNIFORMINDICES fnptr, GLuint  program, GLsizei  uniformCount, const GLchar *const* uniformNames, GLuint * uniformIndices) {
//   (*fnptr)(program, uniformCount, uniformNames, uniformIndices);
// }
// static GLint  glowGetUniformLocation(GPGETUNIFORMLOCATION fnptr, GLuint  program, const GLchar * name) {
//   return (*fnptr)(program, name);
// }
// static void  glowGetUniformSubroutineuiv(GPGETUNIFORMSUBROUTINEUIV fnptr, GLenum  shadertype, GLint  location, GLuint * params) {
//   (*fnptr)(shadertype, location, params);
// }
// static void  glowGetUniformdv(GPGETUNIFORMDV fnptr, GLuint  program, GLint  location, GLdouble * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformfv(GPGETUNIFORMFV fnptr, GLuint  program, GLint  location, GLfloat * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformiv(GPGETUNIFORMIV fnptr, GLuint  program, GLint  location, GLint * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetUniformuiv(GPGETUNIFORMUIV fnptr, GLuint  program, GLint  location, GLuint * params) {
//   (*fnptr)(program, location, params);
// }
// static void  glowGetVertexArrayIndexed64iv(GPGETVERTEXARRAYINDEXED64IV fnptr, GLuint  vaobj, GLuint  index, GLenum  pname, GLint64 * param) {
//   (*fnptr)(vaobj, index, pname, param);
// }
// static void  glowGetVertexArrayIndexediv(GPGETVERTEXARRAYINDEXEDIV fnptr, GLuint  vaobj, GLuint  index, GLenum  pname, GLint * param) {
//   (*fnptr)(vaobj, index, pname, param);
// }
// static void  glowGetVertexArrayiv(GPGETVERTEXARRAYIV fnptr, GLuint  vaobj, GLenum  pname, GLint * param) {
//   (*fnptr)(vaobj, pname, param);
// }
// static void  glowGetVertexAttribIiv(GPGETVERTEXATTRIBIIV fnptr, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribIuiv(GPGETVERTEXATTRIBIUIV fnptr, GLuint  index, GLenum  pname, GLuint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribLdv(GPGETVERTEXATTRIBLDV fnptr, GLuint  index, GLenum  pname, GLdouble * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribLui64vARB(GPGETVERTEXATTRIBLUI64VARB fnptr, GLuint  index, GLenum  pname, GLuint64EXT * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribPointerv(GPGETVERTEXATTRIBPOINTERV fnptr, GLuint  index, GLenum  pname, void ** pointer) {
//   (*fnptr)(index, pname, pointer);
// }
// static void  glowGetVertexAttribdv(GPGETVERTEXATTRIBDV fnptr, GLuint  index, GLenum  pname, GLdouble * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribfv(GPGETVERTEXATTRIBFV fnptr, GLuint  index, GLenum  pname, GLfloat * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetVertexAttribiv(GPGETVERTEXATTRIBIV fnptr, GLuint  index, GLenum  pname, GLint * params) {
//   (*fnptr)(index, pname, params);
// }
// static void  glowGetnCompressedTexImageARB(GPGETNCOMPRESSEDTEXIMAGEARB fnptr, GLenum  target, GLint  lod, GLsizei  bufSize, void * img) {
//   (*fnptr)(target, lod, bufSize, img);
// }
// static void  glowGetnTexImageARB(GPGETNTEXIMAGEARB fnptr, GLenum  target, GLint  level, GLenum  format, GLenum  type, GLsizei  bufSize, void * img) {
//   (*fnptr)(target, level, format, type, bufSize, img);
// }
// static void  glowGetnUniformdvARB(GPGETNUNIFORMDVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLdouble * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfv(GPGETNUNIFORMFV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfvARB(GPGETNUNIFORMFVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformfvKHR(GPGETNUNIFORMFVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLfloat * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformiv(GPGETNUNIFORMIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformivARB(GPGETNUNIFORMIVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformivKHR(GPGETNUNIFORMIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuiv(GPGETNUNIFORMUIV fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuivARB(GPGETNUNIFORMUIVARB fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowGetnUniformuivKHR(GPGETNUNIFORMUIVKHR fnptr, GLuint  program, GLint  location, GLsizei  bufSize, GLuint * params) {
//   (*fnptr)(program, location, bufSize, params);
// }
// static void  glowHint(GPHINT fnptr, GLenum  target, GLenum  mode) {
//   (*fnptr)(target, mode);
// }
// static void  glowIndexxOES(GPINDEXXOES fnptr, GLfixed  component) {
//   (*fnptr)(component);
// }
// static void  glowIndexxvOES(GPINDEXXVOES fnptr, const GLfixed * component) {
//   (*fnptr)(component);
// }
// static void  glowInsertEventMarkerEXT(GPINSERTEVENTMARKEREXT fnptr, GLsizei  length, const GLchar * marker) {
//   (*fnptr)(length, marker);
// }
// static void  glowInvalidateBufferData(GPINVALIDATEBUFFERDATA fnptr, GLuint  buffer) {
//   (*fnptr)(buffer);
// }
// static void  glowInvalidateBufferSubData(GPINVALIDATEBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizeiptr  length) {
//   (*fnptr)(buffer, offset, length);
// }
// static void  glowInvalidateFramebuffer(GPINVALIDATEFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments) {
//   (*fnptr)(target, numAttachments, attachments);
// }
// static void  glowInvalidateNamedFramebufferData(GPINVALIDATENAMEDFRAMEBUFFERDATA fnptr, GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments) {
//   (*fnptr)(framebuffer, numAttachments, attachments);
// }
// static void  glowInvalidateNamedFramebufferSubData(GPINVALIDATENAMEDFRAMEBUFFERSUBDATA fnptr, GLuint  framebuffer, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(framebuffer, numAttachments, attachments, x, y, width, height);
// }
// static void  glowInvalidateSubFramebuffer(GPINVALIDATESUBFRAMEBUFFER fnptr, GLenum  target, GLsizei  numAttachments, const GLenum * attachments, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, numAttachments, attachments, x, y, width, height);
// }
// static void  glowInvalidateTexImage(GPINVALIDATETEXIMAGE fnptr, GLuint  texture, GLint  level) {
//   (*fnptr)(texture, level);
// }
// static void  glowInvalidateTexSubImage(GPINVALIDATETEXSUBIMAGE fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth);
// }
// static GLboolean  glowIsBuffer(GPISBUFFER fnptr, GLuint  buffer) {
//   return (*fnptr)(buffer);
// }
// static GLboolean  glowIsEnabled(GPISENABLED fnptr, GLenum  cap) {
//   return (*fnptr)(cap);
// }
// static GLboolean  glowIsEnabledi(GPISENABLEDI fnptr, GLenum  target, GLuint  index) {
//   return (*fnptr)(target, index);
// }
// static GLboolean  glowIsFenceNV(GPISFENCENV fnptr, GLuint  fence) {
//   return (*fnptr)(fence);
// }
// static GLboolean  glowIsFramebuffer(GPISFRAMEBUFFER fnptr, GLuint  framebuffer) {
//   return (*fnptr)(framebuffer);
// }
// static GLboolean  glowIsImageHandleResidentARB(GPISIMAGEHANDLERESIDENTARB fnptr, GLuint64  handle) {
//   return (*fnptr)(handle);
// }
// static GLboolean  glowIsNamedStringARB(GPISNAMEDSTRINGARB fnptr, GLint  namelen, const GLchar * name) {
//   return (*fnptr)(namelen, name);
// }
// static GLboolean  glowIsProgram(GPISPROGRAM fnptr, GLuint  program) {
//   return (*fnptr)(program);
// }
// static GLboolean  glowIsProgramPipeline(GPISPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   return (*fnptr)(pipeline);
// }
// static GLboolean  glowIsProgramPipelineEXT(GPISPROGRAMPIPELINEEXT fnptr, GLuint  pipeline) {
//   return (*fnptr)(pipeline);
// }
// static GLboolean  glowIsQuery(GPISQUERY fnptr, GLuint  id) {
//   return (*fnptr)(id);
// }
// static GLboolean  glowIsRenderbuffer(GPISRENDERBUFFER fnptr, GLuint  renderbuffer) {
//   return (*fnptr)(renderbuffer);
// }
// static GLboolean  glowIsSampler(GPISSAMPLER fnptr, GLuint  sampler) {
//   return (*fnptr)(sampler);
// }
// static GLboolean  glowIsShader(GPISSHADER fnptr, GLuint  shader) {
//   return (*fnptr)(shader);
// }
// static GLboolean  glowIsSync(GPISSYNC fnptr, GLsync  sync) {
//   return (*fnptr)(sync);
// }
// static GLboolean  glowIsTexture(GPISTEXTURE fnptr, GLuint  texture) {
//   return (*fnptr)(texture);
// }
// static GLboolean  glowIsTextureHandleResidentARB(GPISTEXTUREHANDLERESIDENTARB fnptr, GLuint64  handle) {
//   return (*fnptr)(handle);
// }
// static GLboolean  glowIsTransformFeedback(GPISTRANSFORMFEEDBACK fnptr, GLuint  id) {
//   return (*fnptr)(id);
// }
// static GLboolean  glowIsVertexArray(GPISVERTEXARRAY fnptr, GLuint  array) {
//   return (*fnptr)(array);
// }
// static void  glowLabelObjectEXT(GPLABELOBJECTEXT fnptr, GLenum  type, GLuint  object, GLsizei  length, const GLchar * label) {
//   (*fnptr)(type, object, length, label);
// }
// static void  glowLightModelxOES(GPLIGHTMODELXOES fnptr, GLenum  pname, GLfixed  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowLightModelxvOES(GPLIGHTMODELXVOES fnptr, GLenum  pname, const GLfixed * param) {
//   (*fnptr)(pname, param);
// }
// static void  glowLightxOES(GPLIGHTXOES fnptr, GLenum  light, GLenum  pname, GLfixed  param) {
//   (*fnptr)(light, pname, param);
// }
// static void  glowLightxvOES(GPLIGHTXVOES fnptr, GLenum  light, GLenum  pname, const GLfixed * params) {
//   (*fnptr)(light, pname, params);
// }
// static void  glowLineWidth(GPLINEWIDTH fnptr, GLfloat  width) {
//   (*fnptr)(width);
// }
// static void  glowLineWidthxOES(GPLINEWIDTHXOES fnptr, GLfixed  width) {
//   (*fnptr)(width);
// }
// static void  glowLinkProgram(GPLINKPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowLoadMatrixxOES(GPLOADMATRIXXOES fnptr, const GLfixed * m) {
//   (*fnptr)(m);
// }
// static void  glowLoadTransposeMatrixxOES(GPLOADTRANSPOSEMATRIXXOES fnptr, const GLfixed * m) {
//   (*fnptr)(m);
// }
// static void  glowLogicOp(GPLOGICOP fnptr, GLenum  opcode) {
//   (*fnptr)(opcode);
// }
// static void  glowMakeImageHandleNonResidentARB(GPMAKEIMAGEHANDLENONRESIDENTARB fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMakeImageHandleResidentARB(GPMAKEIMAGEHANDLERESIDENTARB fnptr, GLuint64  handle, GLenum  access) {
//   (*fnptr)(handle, access);
// }
// static void  glowMakeTextureHandleNonResidentARB(GPMAKETEXTUREHANDLENONRESIDENTARB fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMakeTextureHandleResidentARB(GPMAKETEXTUREHANDLERESIDENTARB fnptr, GLuint64  handle) {
//   (*fnptr)(handle);
// }
// static void  glowMap1xOES(GPMAP1XOES fnptr, GLenum  target, GLfixed  u1, GLfixed  u2, GLint  stride, GLint  order, GLfixed  points) {
//   (*fnptr)(target, u1, u2, stride, order, points);
// }
// static void  glowMap2xOES(GPMAP2XOES fnptr, GLenum  target, GLfixed  u1, GLfixed  u2, GLint  ustride, GLint  uorder, GLfixed  v1, GLfixed  v2, GLint  vstride, GLint  vorder, GLfixed  points) {
//   (*fnptr)(target, u1, u2, ustride, uorder, v1, v2, vstride, vorder, points);
// }
// static void * glowMapBuffer(GPMAPBUFFER fnptr, GLenum  target, GLenum  access) {
//   return (*fnptr)(target, access);
// }
// static void * glowMapBufferRange(GPMAPBUFFERRANGE fnptr, GLenum  target, GLintptr  offset, GLsizeiptr  length, GLbitfield  access) {
//   return (*fnptr)(target, offset, length, access);
// }
// static void  glowMapGrid1xOES(GPMAPGRID1XOES fnptr, GLint  n, GLfixed  u1, GLfixed  u2) {
//   (*fnptr)(n, u1, u2);
// }
// static void  glowMapGrid2xOES(GPMAPGRID2XOES fnptr, GLint  n, GLfixed  u1, GLfixed  u2, GLfixed  v1, GLfixed  v2) {
//   (*fnptr)(n, u1, u2, v1, v2);
// }
// static void * glowMapNamedBuffer(GPMAPNAMEDBUFFER fnptr, GLuint  buffer, GLenum  access) {
//   return (*fnptr)(buffer, access);
// }
// static void * glowMapNamedBufferRange(GPMAPNAMEDBUFFERRANGE fnptr, GLuint  buffer, GLintptr  offset, GLsizei  length, GLbitfield  access) {
//   return (*fnptr)(buffer, offset, length, access);
// }
// static void  glowMaterialxOES(GPMATERIALXOES fnptr, GLenum  face, GLenum  pname, GLfixed  param) {
//   (*fnptr)(face, pname, param);
// }
// static void  glowMaterialxvOES(GPMATERIALXVOES fnptr, GLenum  face, GLenum  pname, const GLfixed * param) {
//   (*fnptr)(face, pname, param);
// }
// static void  glowMemoryBarrier(GPMEMORYBARRIER fnptr, GLbitfield  barriers) {
//   (*fnptr)(barriers);
// }
// static void  glowMemoryBarrierByRegion(GPMEMORYBARRIERBYREGION fnptr, GLbitfield  barriers) {
//   (*fnptr)(barriers);
// }
// static void  glowMinSampleShading(GPMINSAMPLESHADING fnptr, GLfloat  value) {
//   (*fnptr)(value);
// }
// static void  glowMinSampleShadingARB(GPMINSAMPLESHADINGARB fnptr, GLfloat  value) {
//   (*fnptr)(value);
// }
// static void  glowMultMatrixxOES(GPMULTMATRIXXOES fnptr, const GLfixed * m) {
//   (*fnptr)(m);
// }
// static void  glowMultTransposeMatrixxOES(GPMULTTRANSPOSEMATRIXXOES fnptr, const GLfixed * m) {
//   (*fnptr)(m);
// }
// static void  glowMultiDrawArrays(GPMULTIDRAWARRAYS fnptr, GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  drawcount) {
//   (*fnptr)(mode, first, count, drawcount);
// }
// static void  glowMultiDrawArraysEXT(GPMULTIDRAWARRAYSEXT fnptr, GLenum  mode, const GLint * first, const GLsizei * count, GLsizei  primcount) {
//   (*fnptr)(mode, first, count, primcount);
// }
// static void  glowMultiDrawArraysIndirect(GPMULTIDRAWARRAYSINDIRECT fnptr, GLenum  mode, const void * indirect, GLsizei  drawcount, GLsizei  stride) {
//   (*fnptr)(mode, indirect, drawcount, stride);
// }
// static void  glowMultiDrawArraysIndirectCountARB(GPMULTIDRAWARRAYSINDIRECTCOUNTARB fnptr, GLenum  mode, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride) {
//   (*fnptr)(mode, indirect, drawcount, maxdrawcount, stride);
// }
// static void  glowMultiDrawElements(GPMULTIDRAWELEMENTS fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount) {
//   (*fnptr)(mode, count, type, indices, drawcount);
// }
// static void  glowMultiDrawElementsBaseVertex(GPMULTIDRAWELEMENTSBASEVERTEX fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  drawcount, const GLint * basevertex) {
//   (*fnptr)(mode, count, type, indices, drawcount, basevertex);
// }
// static void  glowMultiDrawElementsEXT(GPMULTIDRAWELEMENTSEXT fnptr, GLenum  mode, const GLsizei * count, GLenum  type, const void *const* indices, GLsizei  primcount) {
//   (*fnptr)(mode, count, type, indices, primcount);
// }
// static void  glowMultiDrawElementsIndirect(GPMULTIDRAWELEMENTSINDIRECT fnptr, GLenum  mode, GLenum  type, const void * indirect, GLsizei  drawcount, GLsizei  stride) {
//   (*fnptr)(mode, type, indirect, drawcount, stride);
// }
// static void  glowMultiDrawElementsIndirectCountARB(GPMULTIDRAWELEMENTSINDIRECTCOUNTARB fnptr, GLenum  mode, GLenum  type, GLintptr  indirect, GLintptr  drawcount, GLsizei  maxdrawcount, GLsizei  stride) {
//   (*fnptr)(mode, type, indirect, drawcount, maxdrawcount, stride);
// }
// static void  glowMultiTexCoord1bOES(GPMULTITEXCOORD1BOES fnptr, GLenum  texture, GLbyte  s) {
//   (*fnptr)(texture, s);
// }
// static void  glowMultiTexCoord1bvOES(GPMULTITEXCOORD1BVOES fnptr, GLenum  texture, const GLbyte * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowMultiTexCoord1xOES(GPMULTITEXCOORD1XOES fnptr, GLenum  texture, GLfixed  s) {
//   (*fnptr)(texture, s);
// }
// static void  glowMultiTexCoord1xvOES(GPMULTITEXCOORD1XVOES fnptr, GLenum  texture, const GLfixed * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowMultiTexCoord2bOES(GPMULTITEXCOORD2BOES fnptr, GLenum  texture, GLbyte  s, GLbyte  t) {
//   (*fnptr)(texture, s, t);
// }
// static void  glowMultiTexCoord2bvOES(GPMULTITEXCOORD2BVOES fnptr, GLenum  texture, const GLbyte * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowMultiTexCoord2xOES(GPMULTITEXCOORD2XOES fnptr, GLenum  texture, GLfixed  s, GLfixed  t) {
//   (*fnptr)(texture, s, t);
// }
// static void  glowMultiTexCoord2xvOES(GPMULTITEXCOORD2XVOES fnptr, GLenum  texture, const GLfixed * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowMultiTexCoord3bOES(GPMULTITEXCOORD3BOES fnptr, GLenum  texture, GLbyte  s, GLbyte  t, GLbyte  r) {
//   (*fnptr)(texture, s, t, r);
// }
// static void  glowMultiTexCoord3bvOES(GPMULTITEXCOORD3BVOES fnptr, GLenum  texture, const GLbyte * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowMultiTexCoord3xOES(GPMULTITEXCOORD3XOES fnptr, GLenum  texture, GLfixed  s, GLfixed  t, GLfixed  r) {
//   (*fnptr)(texture, s, t, r);
// }
// static void  glowMultiTexCoord3xvOES(GPMULTITEXCOORD3XVOES fnptr, GLenum  texture, const GLfixed * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowMultiTexCoord4bOES(GPMULTITEXCOORD4BOES fnptr, GLenum  texture, GLbyte  s, GLbyte  t, GLbyte  r, GLbyte  q) {
//   (*fnptr)(texture, s, t, r, q);
// }
// static void  glowMultiTexCoord4bvOES(GPMULTITEXCOORD4BVOES fnptr, GLenum  texture, const GLbyte * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowMultiTexCoord4xOES(GPMULTITEXCOORD4XOES fnptr, GLenum  texture, GLfixed  s, GLfixed  t, GLfixed  r, GLfixed  q) {
//   (*fnptr)(texture, s, t, r, q);
// }
// static void  glowMultiTexCoord4xvOES(GPMULTITEXCOORD4XVOES fnptr, GLenum  texture, const GLfixed * coords) {
//   (*fnptr)(texture, coords);
// }
// static void  glowNamedBufferData(GPNAMEDBUFFERDATA fnptr, GLuint  buffer, GLsizei  size, const void * data, GLenum  usage) {
//   (*fnptr)(buffer, size, data, usage);
// }
// static void  glowNamedBufferPageCommitmentARB(GPNAMEDBUFFERPAGECOMMITMENTARB fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit) {
//   (*fnptr)(buffer, offset, size, commit);
// }
// static void  glowNamedBufferPageCommitmentEXT(GPNAMEDBUFFERPAGECOMMITMENTEXT fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, GLboolean  commit) {
//   (*fnptr)(buffer, offset, size, commit);
// }
// static void  glowNamedBufferStorage(GPNAMEDBUFFERSTORAGE fnptr, GLuint  buffer, GLsizei  size, const void * data, GLbitfield  flags) {
//   (*fnptr)(buffer, size, data, flags);
// }
// static void  glowNamedBufferSubData(GPNAMEDBUFFERSUBDATA fnptr, GLuint  buffer, GLintptr  offset, GLsizei  size, const void * data) {
//   (*fnptr)(buffer, offset, size, data);
// }
// static void  glowNamedFramebufferDrawBuffer(GPNAMEDFRAMEBUFFERDRAWBUFFER fnptr, GLuint  framebuffer, GLenum  buf) {
//   (*fnptr)(framebuffer, buf);
// }
// static void  glowNamedFramebufferDrawBuffers(GPNAMEDFRAMEBUFFERDRAWBUFFERS fnptr, GLuint  framebuffer, GLsizei  n, const GLenum * bufs) {
//   (*fnptr)(framebuffer, n, bufs);
// }
// static void  glowNamedFramebufferParameteri(GPNAMEDFRAMEBUFFERPARAMETERI fnptr, GLuint  framebuffer, GLenum  pname, GLint  param) {
//   (*fnptr)(framebuffer, pname, param);
// }
// static void  glowNamedFramebufferReadBuffer(GPNAMEDFRAMEBUFFERREADBUFFER fnptr, GLuint  framebuffer, GLenum  src) {
//   (*fnptr)(framebuffer, src);
// }
// static void  glowNamedFramebufferRenderbuffer(GPNAMEDFRAMEBUFFERRENDERBUFFER fnptr, GLuint  framebuffer, GLenum  attachment, GLenum  renderbuffertarget, GLuint  renderbuffer) {
//   (*fnptr)(framebuffer, attachment, renderbuffertarget, renderbuffer);
// }
// static void  glowNamedFramebufferTexture(GPNAMEDFRAMEBUFFERTEXTURE fnptr, GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level) {
//   (*fnptr)(framebuffer, attachment, texture, level);
// }
// static void  glowNamedFramebufferTextureLayer(GPNAMEDFRAMEBUFFERTEXTURELAYER fnptr, GLuint  framebuffer, GLenum  attachment, GLuint  texture, GLint  level, GLint  layer) {
//   (*fnptr)(framebuffer, attachment, texture, level, layer);
// }
// static void  glowNamedRenderbufferStorage(GPNAMEDRENDERBUFFERSTORAGE fnptr, GLuint  renderbuffer, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(renderbuffer, internalformat, width, height);
// }
// static void  glowNamedRenderbufferStorageMultisample(GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE fnptr, GLuint  renderbuffer, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(renderbuffer, samples, internalformat, width, height);
// }
// static void  glowNamedStringARB(GPNAMEDSTRINGARB fnptr, GLenum  type, GLint  namelen, const GLchar * name, GLint  stringlen, const GLchar * string) {
//   (*fnptr)(type, namelen, name, stringlen, string);
// }
// static void  glowNormal3xOES(GPNORMAL3XOES fnptr, GLfixed  nx, GLfixed  ny, GLfixed  nz) {
//   (*fnptr)(nx, ny, nz);
// }
// static void  glowNormal3xvOES(GPNORMAL3XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowObjectLabel(GPOBJECTLABEL fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {
//   (*fnptr)(identifier, name, length, label);
// }
// static void  glowObjectLabelKHR(GPOBJECTLABELKHR fnptr, GLenum  identifier, GLuint  name, GLsizei  length, const GLchar * label) {
//   (*fnptr)(identifier, name, length, label);
// }
// static void  glowObjectPtrLabel(GPOBJECTPTRLABEL fnptr, const void * ptr, GLsizei  length, const GLchar * label) {
//   (*fnptr)(ptr, length, label);
// }
// static void  glowObjectPtrLabelKHR(GPOBJECTPTRLABELKHR fnptr, const void * ptr, GLsizei  length, const GLchar * label) {
//   (*fnptr)(ptr, length, label);
// }
// static void  glowOrthofOES(GPORTHOFOES fnptr, GLfloat  l, GLfloat  r, GLfloat  b, GLfloat  t, GLfloat  n, GLfloat  f) {
//   (*fnptr)(l, r, b, t, n, f);
// }
// static void  glowOrthoxOES(GPORTHOXOES fnptr, GLfixed  l, GLfixed  r, GLfixed  b, GLfixed  t, GLfixed  n, GLfixed  f) {
//   (*fnptr)(l, r, b, t, n, f);
// }
// static void  glowPassThroughxOES(GPPASSTHROUGHXOES fnptr, GLfixed  token) {
//   (*fnptr)(token);
// }
// static void  glowPatchParameterfv(GPPATCHPARAMETERFV fnptr, GLenum  pname, const GLfloat * values) {
//   (*fnptr)(pname, values);
// }
// static void  glowPatchParameteri(GPPATCHPARAMETERI fnptr, GLenum  pname, GLint  value) {
//   (*fnptr)(pname, value);
// }
// static void  glowPauseTransformFeedback(GPPAUSETRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowPixelMapx(GPPIXELMAPX fnptr, GLenum  map, GLint  size, const GLfixed * values) {
//   (*fnptr)(map, size, values);
// }
// static void  glowPixelStoref(GPPIXELSTOREF fnptr, GLenum  pname, GLfloat  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelStorei(GPPIXELSTOREI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelStorex(GPPIXELSTOREX fnptr, GLenum  pname, GLfixed  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelTransferxOES(GPPIXELTRANSFERXOES fnptr, GLenum  pname, GLfixed  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPixelZoomxOES(GPPIXELZOOMXOES fnptr, GLfixed  xfactor, GLfixed  yfactor) {
//   (*fnptr)(xfactor, yfactor);
// }
// static void  glowPointParameterf(GPPOINTPARAMETERF fnptr, GLenum  pname, GLfloat  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPointParameterfv(GPPOINTPARAMETERFV fnptr, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowPointParameteri(GPPOINTPARAMETERI fnptr, GLenum  pname, GLint  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPointParameteriv(GPPOINTPARAMETERIV fnptr, GLenum  pname, const GLint * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowPointParameterxOES(GPPOINTPARAMETERXOES fnptr, GLenum  pname, GLfixed  param) {
//   (*fnptr)(pname, param);
// }
// static void  glowPointParameterxvOES(GPPOINTPARAMETERXVOES fnptr, GLenum  pname, const GLfixed * params) {
//   (*fnptr)(pname, params);
// }
// static void  glowPointSize(GPPOINTSIZE fnptr, GLfloat  size) {
//   (*fnptr)(size);
// }
// static void  glowPointSizexOES(GPPOINTSIZEXOES fnptr, GLfixed  size) {
//   (*fnptr)(size);
// }
// static void  glowPolygonMode(GPPOLYGONMODE fnptr, GLenum  face, GLenum  mode) {
//   (*fnptr)(face, mode);
// }
// static void  glowPolygonOffset(GPPOLYGONOFFSET fnptr, GLfloat  factor, GLfloat  units) {
//   (*fnptr)(factor, units);
// }
// static void  glowPolygonOffsetxOES(GPPOLYGONOFFSETXOES fnptr, GLfixed  factor, GLfixed  units) {
//   (*fnptr)(factor, units);
// }
// static void  glowPopDebugGroup(GPPOPDEBUGGROUP fnptr) {
//   (*fnptr)();
// }
// static void  glowPopDebugGroupKHR(GPPOPDEBUGGROUPKHR fnptr) {
//   (*fnptr)();
// }
// static void  glowPopGroupMarkerEXT(GPPOPGROUPMARKEREXT fnptr) {
//   (*fnptr)();
// }
// static void  glowPrimitiveRestartIndex(GPPRIMITIVERESTARTINDEX fnptr, GLuint  index) {
//   (*fnptr)(index);
// }
// static void  glowPrioritizeTexturesxOES(GPPRIORITIZETEXTURESXOES fnptr, GLsizei  n, const GLuint * textures, const GLfixed * priorities) {
//   (*fnptr)(n, textures, priorities);
// }
// static void  glowProgramBinary(GPPROGRAMBINARY fnptr, GLuint  program, GLenum  binaryFormat, const void * binary, GLsizei  length) {
//   (*fnptr)(program, binaryFormat, binary, length);
// }
// static void  glowProgramParameteri(GPPROGRAMPARAMETERI fnptr, GLuint  program, GLenum  pname, GLint  value) {
//   (*fnptr)(program, pname, value);
// }
// static void  glowProgramParameteriEXT(GPPROGRAMPARAMETERIEXT fnptr, GLuint  program, GLenum  pname, GLint  value) {
//   (*fnptr)(program, pname, value);
// }
// static void  glowProgramUniform1d(GPPROGRAMUNIFORM1D fnptr, GLuint  program, GLint  location, GLdouble  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1dv(GPPROGRAMUNIFORM1DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1f(GPPROGRAMUNIFORM1F fnptr, GLuint  program, GLint  location, GLfloat  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1fEXT(GPPROGRAMUNIFORM1FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1fv(GPPROGRAMUNIFORM1FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1fvEXT(GPPROGRAMUNIFORM1FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1i(GPPROGRAMUNIFORM1I fnptr, GLuint  program, GLint  location, GLint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1iEXT(GPPROGRAMUNIFORM1IEXT fnptr, GLuint  program, GLint  location, GLint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1iv(GPPROGRAMUNIFORM1IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1ivEXT(GPPROGRAMUNIFORM1IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1ui(GPPROGRAMUNIFORM1UI fnptr, GLuint  program, GLint  location, GLuint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1uiEXT(GPPROGRAMUNIFORM1UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0) {
//   (*fnptr)(program, location, v0);
// }
// static void  glowProgramUniform1uiv(GPPROGRAMUNIFORM1UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform1uivEXT(GPPROGRAMUNIFORM1UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2d(GPPROGRAMUNIFORM2D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2dv(GPPROGRAMUNIFORM2DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2f(GPPROGRAMUNIFORM2F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2fEXT(GPPROGRAMUNIFORM2FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2fv(GPPROGRAMUNIFORM2FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2fvEXT(GPPROGRAMUNIFORM2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2i(GPPROGRAMUNIFORM2I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2iEXT(GPPROGRAMUNIFORM2IEXT fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2iv(GPPROGRAMUNIFORM2IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2ivEXT(GPPROGRAMUNIFORM2IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2ui(GPPROGRAMUNIFORM2UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2uiEXT(GPPROGRAMUNIFORM2UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(program, location, v0, v1);
// }
// static void  glowProgramUniform2uiv(GPPROGRAMUNIFORM2UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform2uivEXT(GPPROGRAMUNIFORM2UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3d(GPPROGRAMUNIFORM3D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3dv(GPPROGRAMUNIFORM3DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3f(GPPROGRAMUNIFORM3F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3fEXT(GPPROGRAMUNIFORM3FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3fv(GPPROGRAMUNIFORM3FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3fvEXT(GPPROGRAMUNIFORM3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3i(GPPROGRAMUNIFORM3I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3iEXT(GPPROGRAMUNIFORM3IEXT fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3iv(GPPROGRAMUNIFORM3IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3ivEXT(GPPROGRAMUNIFORM3IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3ui(GPPROGRAMUNIFORM3UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3uiEXT(GPPROGRAMUNIFORM3UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(program, location, v0, v1, v2);
// }
// static void  glowProgramUniform3uiv(GPPROGRAMUNIFORM3UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform3uivEXT(GPPROGRAMUNIFORM3UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4d(GPPROGRAMUNIFORM4D fnptr, GLuint  program, GLint  location, GLdouble  v0, GLdouble  v1, GLdouble  v2, GLdouble  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4dv(GPPROGRAMUNIFORM4DV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4f(GPPROGRAMUNIFORM4F fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4fEXT(GPPROGRAMUNIFORM4FEXT fnptr, GLuint  program, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4fv(GPPROGRAMUNIFORM4FV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4fvEXT(GPPROGRAMUNIFORM4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4i(GPPROGRAMUNIFORM4I fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4iEXT(GPPROGRAMUNIFORM4IEXT fnptr, GLuint  program, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4iv(GPPROGRAMUNIFORM4IV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4ivEXT(GPPROGRAMUNIFORM4IVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4ui(GPPROGRAMUNIFORM4UI fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4uiEXT(GPPROGRAMUNIFORM4UIEXT fnptr, GLuint  program, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(program, location, v0, v1, v2, v3);
// }
// static void  glowProgramUniform4uiv(GPPROGRAMUNIFORM4UIV fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniform4uivEXT(GPPROGRAMUNIFORM4UIVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(program, location, count, value);
// }
// static void  glowProgramUniformHandleui64ARB(GPPROGRAMUNIFORMHANDLEUI64ARB fnptr, GLuint  program, GLint  location, GLuint64  value) {
//   (*fnptr)(program, location, value);
// }
// static void  glowProgramUniformHandleui64vARB(GPPROGRAMUNIFORMHANDLEUI64VARB fnptr, GLuint  program, GLint  location, GLsizei  count, const GLuint64 * values) {
//   (*fnptr)(program, location, count, values);
// }
// static void  glowProgramUniformMatrix2dv(GPPROGRAMUNIFORMMATRIX2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2fv(GPPROGRAMUNIFORMMATRIX2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2fvEXT(GPPROGRAMUNIFORMMATRIX2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x3dv(GPPROGRAMUNIFORMMATRIX2X3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x3fv(GPPROGRAMUNIFORMMATRIX2X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x3fvEXT(GPPROGRAMUNIFORMMATRIX2X3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x4dv(GPPROGRAMUNIFORMMATRIX2X4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x4fv(GPPROGRAMUNIFORMMATRIX2X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix2x4fvEXT(GPPROGRAMUNIFORMMATRIX2X4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3dv(GPPROGRAMUNIFORMMATRIX3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3fv(GPPROGRAMUNIFORMMATRIX3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3fvEXT(GPPROGRAMUNIFORMMATRIX3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x2dv(GPPROGRAMUNIFORMMATRIX3X2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x2fv(GPPROGRAMUNIFORMMATRIX3X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x2fvEXT(GPPROGRAMUNIFORMMATRIX3X2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x4dv(GPPROGRAMUNIFORMMATRIX3X4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x4fv(GPPROGRAMUNIFORMMATRIX3X4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix3x4fvEXT(GPPROGRAMUNIFORMMATRIX3X4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4dv(GPPROGRAMUNIFORMMATRIX4DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4fv(GPPROGRAMUNIFORMMATRIX4FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4fvEXT(GPPROGRAMUNIFORMMATRIX4FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x2dv(GPPROGRAMUNIFORMMATRIX4X2DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x2fv(GPPROGRAMUNIFORMMATRIX4X2FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x2fvEXT(GPPROGRAMUNIFORMMATRIX4X2FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x3dv(GPPROGRAMUNIFORMMATRIX4X3DV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x3fv(GPPROGRAMUNIFORMMATRIX4X3FV fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProgramUniformMatrix4x3fvEXT(GPPROGRAMUNIFORMMATRIX4X3FVEXT fnptr, GLuint  program, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(program, location, count, transpose, value);
// }
// static void  glowProvokingVertex(GPPROVOKINGVERTEX fnptr, GLenum  mode) {
//   (*fnptr)(mode);
// }
// static void  glowPushDebugGroup(GPPUSHDEBUGGROUP fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {
//   (*fnptr)(source, id, length, message);
// }
// static void  glowPushDebugGroupKHR(GPPUSHDEBUGGROUPKHR fnptr, GLenum  source, GLuint  id, GLsizei  length, const GLchar * message) {
//   (*fnptr)(source, id, length, message);
// }
// static void  glowPushGroupMarkerEXT(GPPUSHGROUPMARKEREXT fnptr, GLsizei  length, const GLchar * marker) {
//   (*fnptr)(length, marker);
// }
// static void  glowQueryCounter(GPQUERYCOUNTER fnptr, GLuint  id, GLenum  target) {
//   (*fnptr)(id, target);
// }
// static GLbitfield  glowQueryMatrixxOES(GPQUERYMATRIXXOES fnptr, GLfixed * mantissa, GLint * exponent) {
//   return (*fnptr)(mantissa, exponent);
// }
// static void  glowRasterPos2xOES(GPRASTERPOS2XOES fnptr, GLfixed  x, GLfixed  y) {
//   (*fnptr)(x, y);
// }
// static void  glowRasterPos2xvOES(GPRASTERPOS2XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowRasterPos3xOES(GPRASTERPOS3XOES fnptr, GLfixed  x, GLfixed  y, GLfixed  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowRasterPos3xvOES(GPRASTERPOS3XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowRasterPos4xOES(GPRASTERPOS4XOES fnptr, GLfixed  x, GLfixed  y, GLfixed  z, GLfixed  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowRasterPos4xvOES(GPRASTERPOS4XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowReadBuffer(GPREADBUFFER fnptr, GLenum  src) {
//   (*fnptr)(src);
// }
// static void  glowReadPixels(GPREADPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, void * pixels) {
//   (*fnptr)(x, y, width, height, format, type, pixels);
// }
// static void  glowReadnPixels(GPREADNPIXELS fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowReadnPixelsARB(GPREADNPIXELSARB fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowReadnPixelsKHR(GPREADNPIXELSKHR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, GLsizei  bufSize, void * data) {
//   (*fnptr)(x, y, width, height, format, type, bufSize, data);
// }
// static void  glowRectxOES(GPRECTXOES fnptr, GLfixed  x1, GLfixed  y1, GLfixed  x2, GLfixed  y2) {
//   (*fnptr)(x1, y1, x2, y2);
// }
// static void  glowRectxvOES(GPRECTXVOES fnptr, const GLfixed * v1, const GLfixed * v2) {
//   (*fnptr)(v1, v2);
// }
// static void  glowReleaseShaderCompiler(GPRELEASESHADERCOMPILER fnptr) {
//   (*fnptr)();
// }
// static void  glowRenderbufferStorage(GPRENDERBUFFERSTORAGE fnptr, GLenum  target, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, internalformat, width, height);
// }
// static void  glowRenderbufferStorageMultisample(GPRENDERBUFFERSTORAGEMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, samples, internalformat, width, height);
// }
// static void  glowResumeTransformFeedback(GPRESUMETRANSFORMFEEDBACK fnptr) {
//   (*fnptr)();
// }
// static void  glowRotatexOES(GPROTATEXOES fnptr, GLfixed  angle, GLfixed  x, GLfixed  y, GLfixed  z) {
//   (*fnptr)(angle, x, y, z);
// }
// static void  glowSampleCoverage(GPSAMPLECOVERAGE fnptr, GLfloat  value, GLboolean  invert) {
//   (*fnptr)(value, invert);
// }
// static void  glowSampleCoverageOES(GPSAMPLECOVERAGEOES fnptr, GLfixed  value, GLboolean  invert) {
//   (*fnptr)(value, invert);
// }
// static void  glowSampleCoveragexOES(GPSAMPLECOVERAGEXOES fnptr, GLclampx  value, GLboolean  invert) {
//   (*fnptr)(value, invert);
// }
// static void  glowSampleMaski(GPSAMPLEMASKI fnptr, GLuint  maskNumber, GLbitfield  mask) {
//   (*fnptr)(maskNumber, mask);
// }
// static void  glowSamplerParameterIiv(GPSAMPLERPARAMETERIIV fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterIuiv(GPSAMPLERPARAMETERIUIV fnptr, GLuint  sampler, GLenum  pname, const GLuint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterf(GPSAMPLERPARAMETERF fnptr, GLuint  sampler, GLenum  pname, GLfloat  param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameterfv(GPSAMPLERPARAMETERFV fnptr, GLuint  sampler, GLenum  pname, const GLfloat * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameteri(GPSAMPLERPARAMETERI fnptr, GLuint  sampler, GLenum  pname, GLint  param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowSamplerParameteriv(GPSAMPLERPARAMETERIV fnptr, GLuint  sampler, GLenum  pname, const GLint * param) {
//   (*fnptr)(sampler, pname, param);
// }
// static void  glowScalexOES(GPSCALEXOES fnptr, GLfixed  x, GLfixed  y, GLfixed  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowScissor(GPSCISSOR fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
// static void  glowScissorArrayv(GPSCISSORARRAYV fnptr, GLuint  first, GLsizei  count, const GLint * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowScissorIndexed(GPSCISSORINDEXED fnptr, GLuint  index, GLint  left, GLint  bottom, GLsizei  width, GLsizei  height) {
//   (*fnptr)(index, left, bottom, width, height);
// }
// static void  glowScissorIndexedv(GPSCISSORINDEXEDV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowSelectPerfMonitorCountersAMD(GPSELECTPERFMONITORCOUNTERSAMD fnptr, GLuint  monitor, GLboolean  enable, GLuint  group, GLint  numCounters, GLuint * counterList) {
//   (*fnptr)(monitor, enable, group, numCounters, counterList);
// }
// static void  glowSetFenceNV(GPSETFENCENV fnptr, GLuint  fence, GLenum  condition) {
//   (*fnptr)(fence, condition);
// }
// static void  glowShaderBinary(GPSHADERBINARY fnptr, GLsizei  count, const GLuint * shaders, GLenum  binaryformat, const void * binary, GLsizei  length) {
//   (*fnptr)(count, shaders, binaryformat, binary, length);
// }
// static void  glowShaderSource(GPSHADERSOURCE fnptr, GLuint  shader, GLsizei  count, const GLchar *const* string, const GLint * length) {
//   (*fnptr)(shader, count, string, length);
// }
// static void  glowShaderStorageBlockBinding(GPSHADERSTORAGEBLOCKBINDING fnptr, GLuint  program, GLuint  storageBlockIndex, GLuint  storageBlockBinding) {
//   (*fnptr)(program, storageBlockIndex, storageBlockBinding);
// }
// static void  glowStencilFunc(GPSTENCILFUNC fnptr, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(func, ref, mask);
// }
// static void  glowStencilFuncSeparate(GPSTENCILFUNCSEPARATE fnptr, GLenum  face, GLenum  func, GLint  ref, GLuint  mask) {
//   (*fnptr)(face, func, ref, mask);
// }
// static void  glowStencilMask(GPSTENCILMASK fnptr, GLuint  mask) {
//   (*fnptr)(mask);
// }
// static void  glowStencilMaskSeparate(GPSTENCILMASKSEPARATE fnptr, GLenum  face, GLuint  mask) {
//   (*fnptr)(face, mask);
// }
// static void  glowStencilOp(GPSTENCILOP fnptr, GLenum  fail, GLenum  zfail, GLenum  zpass) {
//   (*fnptr)(fail, zfail, zpass);
// }
// static void  glowStencilOpSeparate(GPSTENCILOPSEPARATE fnptr, GLenum  face, GLenum  sfail, GLenum  dpfail, GLenum  dppass) {
//   (*fnptr)(face, sfail, dpfail, dppass);
// }
// static GLboolean  glowTestFenceNV(GPTESTFENCENV fnptr, GLuint  fence) {
//   return (*fnptr)(fence);
// }
// static void  glowTexBuffer(GPTEXBUFFER fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer) {
//   (*fnptr)(target, internalformat, buffer);
// }
// static void  glowTexBufferRange(GPTEXBUFFERRANGE fnptr, GLenum  target, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizeiptr  size) {
//   (*fnptr)(target, internalformat, buffer, offset, size);
// }
// static void  glowTexCoord1bOES(GPTEXCOORD1BOES fnptr, GLbyte  s) {
//   (*fnptr)(s);
// }
// static void  glowTexCoord1bvOES(GPTEXCOORD1BVOES fnptr, const GLbyte * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexCoord1xOES(GPTEXCOORD1XOES fnptr, GLfixed  s) {
//   (*fnptr)(s);
// }
// static void  glowTexCoord1xvOES(GPTEXCOORD1XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexCoord2bOES(GPTEXCOORD2BOES fnptr, GLbyte  s, GLbyte  t) {
//   (*fnptr)(s, t);
// }
// static void  glowTexCoord2bvOES(GPTEXCOORD2BVOES fnptr, const GLbyte * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexCoord2xOES(GPTEXCOORD2XOES fnptr, GLfixed  s, GLfixed  t) {
//   (*fnptr)(s, t);
// }
// static void  glowTexCoord2xvOES(GPTEXCOORD2XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexCoord3bOES(GPTEXCOORD3BOES fnptr, GLbyte  s, GLbyte  t, GLbyte  r) {
//   (*fnptr)(s, t, r);
// }
// static void  glowTexCoord3bvOES(GPTEXCOORD3BVOES fnptr, const GLbyte * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexCoord3xOES(GPTEXCOORD3XOES fnptr, GLfixed  s, GLfixed  t, GLfixed  r) {
//   (*fnptr)(s, t, r);
// }
// static void  glowTexCoord3xvOES(GPTEXCOORD3XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexCoord4bOES(GPTEXCOORD4BOES fnptr, GLbyte  s, GLbyte  t, GLbyte  r, GLbyte  q) {
//   (*fnptr)(s, t, r, q);
// }
// static void  glowTexCoord4bvOES(GPTEXCOORD4BVOES fnptr, const GLbyte * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexCoord4xOES(GPTEXCOORD4XOES fnptr, GLfixed  s, GLfixed  t, GLfixed  r, GLfixed  q) {
//   (*fnptr)(s, t, r, q);
// }
// static void  glowTexCoord4xvOES(GPTEXCOORD4XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowTexEnvxOES(GPTEXENVXOES fnptr, GLenum  target, GLenum  pname, GLfixed  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexEnvxvOES(GPTEXENVXVOES fnptr, GLenum  target, GLenum  pname, const GLfixed * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexGenxOES(GPTEXGENXOES fnptr, GLenum  coord, GLenum  pname, GLfixed  param) {
//   (*fnptr)(coord, pname, param);
// }
// static void  glowTexGenxvOES(GPTEXGENXVOES fnptr, GLenum  coord, GLenum  pname, const GLfixed * params) {
//   (*fnptr)(coord, pname, params);
// }
// static void  glowTexImage1D(GPTEXIMAGE1D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, border, format, type, pixels);
// }
// static void  glowTexImage2D(GPTEXIMAGE2D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, border, format, type, pixels);
// }
// static void  glowTexImage2DMultisample(GPTEXIMAGE2DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, fixedsamplelocations);
// }
// static void  glowTexImage3D(GPTEXIMAGE3D fnptr, GLenum  target, GLint  level, GLint  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLint  border, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, internalformat, width, height, depth, border, format, type, pixels);
// }
// static void  glowTexImage3DMultisample(GPTEXIMAGE3DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// static void  glowTexPageCommitmentARB(GPTEXPAGECOMMITMENTARB fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  resident) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, resident);
// }
// static void  glowTexParameterIiv(GPTEXPARAMETERIIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterIuiv(GPTEXPARAMETERIUIV fnptr, GLenum  target, GLenum  pname, const GLuint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterf(GPTEXPARAMETERF fnptr, GLenum  target, GLenum  pname, GLfloat  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexParameterfv(GPTEXPARAMETERFV fnptr, GLenum  target, GLenum  pname, const GLfloat * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameteri(GPTEXPARAMETERI fnptr, GLenum  target, GLenum  pname, GLint  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexParameteriv(GPTEXPARAMETERIV fnptr, GLenum  target, GLenum  pname, const GLint * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexParameterxOES(GPTEXPARAMETERXOES fnptr, GLenum  target, GLenum  pname, GLfixed  param) {
//   (*fnptr)(target, pname, param);
// }
// static void  glowTexParameterxvOES(GPTEXPARAMETERXVOES fnptr, GLenum  target, GLenum  pname, const GLfixed * params) {
//   (*fnptr)(target, pname, params);
// }
// static void  glowTexStorage1D(GPTEXSTORAGE1D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width) {
//   (*fnptr)(target, levels, internalformat, width);
// }
// static void  glowTexStorage2D(GPTEXSTORAGE2D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(target, levels, internalformat, width, height);
// }
// static void  glowTexStorage2DMultisample(GPTEXSTORAGE2DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, fixedsamplelocations);
// }
// static void  glowTexStorage3D(GPTEXSTORAGE3D fnptr, GLenum  target, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(target, levels, internalformat, width, height, depth);
// }
// static void  glowTexStorage3DMultisample(GPTEXSTORAGE3DMULTISAMPLE fnptr, GLenum  target, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {
//   (*fnptr)(target, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// static void  glowTexSubImage1D(GPTEXSUBIMAGE1D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, width, format, type, pixels);
// }
// static void  glowTexSubImage2D(GPTEXSUBIMAGE2D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, width, height, format, type, pixels);
// }
// static void  glowTexSubImage3D(GPTEXSUBIMAGE3D fnptr, GLenum  target, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(target, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
// }
// static void  glowTextureBarrier(GPTEXTUREBARRIER fnptr) {
//   (*fnptr)();
// }
// static void  glowTextureBuffer(GPTEXTUREBUFFER fnptr, GLuint  texture, GLenum  internalformat, GLuint  buffer) {
//   (*fnptr)(texture, internalformat, buffer);
// }
// static void  glowTextureBufferRange(GPTEXTUREBUFFERRANGE fnptr, GLuint  texture, GLenum  internalformat, GLuint  buffer, GLintptr  offset, GLsizei  size) {
//   (*fnptr)(texture, internalformat, buffer, offset, size);
// }
// static void  glowTextureParameterIiv(GPTEXTUREPARAMETERIIV fnptr, GLuint  texture, GLenum  pname, const GLint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowTextureParameterIuiv(GPTEXTUREPARAMETERIUIV fnptr, GLuint  texture, GLenum  pname, const GLuint * params) {
//   (*fnptr)(texture, pname, params);
// }
// static void  glowTextureParameterf(GPTEXTUREPARAMETERF fnptr, GLuint  texture, GLenum  pname, GLfloat  param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureParameterfv(GPTEXTUREPARAMETERFV fnptr, GLuint  texture, GLenum  pname, const GLfloat * param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureParameteri(GPTEXTUREPARAMETERI fnptr, GLuint  texture, GLenum  pname, GLint  param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureParameteriv(GPTEXTUREPARAMETERIV fnptr, GLuint  texture, GLenum  pname, const GLint * param) {
//   (*fnptr)(texture, pname, param);
// }
// static void  glowTextureStorage1D(GPTEXTURESTORAGE1D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width) {
//   (*fnptr)(texture, levels, internalformat, width);
// }
// static void  glowTextureStorage2D(GPTEXTURESTORAGE2D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height) {
//   (*fnptr)(texture, levels, internalformat, width, height);
// }
// static void  glowTextureStorage2DMultisample(GPTEXTURESTORAGE2DMULTISAMPLE fnptr, GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLboolean  fixedsamplelocations) {
//   (*fnptr)(texture, samples, internalformat, width, height, fixedsamplelocations);
// }
// static void  glowTextureStorage3D(GPTEXTURESTORAGE3D fnptr, GLuint  texture, GLsizei  levels, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth) {
//   (*fnptr)(texture, levels, internalformat, width, height, depth);
// }
// static void  glowTextureStorage3DMultisample(GPTEXTURESTORAGE3DMULTISAMPLE fnptr, GLuint  texture, GLsizei  samples, GLenum  internalformat, GLsizei  width, GLsizei  height, GLsizei  depth, GLboolean  fixedsamplelocations) {
//   (*fnptr)(texture, samples, internalformat, width, height, depth, fixedsamplelocations);
// }
// static void  glowTextureSubImage1D(GPTEXTURESUBIMAGE1D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLsizei  width, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(texture, level, xoffset, width, format, type, pixels);
// }
// static void  glowTextureSubImage2D(GPTEXTURESUBIMAGE2D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLsizei  width, GLsizei  height, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, width, height, format, type, pixels);
// }
// static void  glowTextureSubImage3D(GPTEXTURESUBIMAGE3D fnptr, GLuint  texture, GLint  level, GLint  xoffset, GLint  yoffset, GLint  zoffset, GLsizei  width, GLsizei  height, GLsizei  depth, GLenum  format, GLenum  type, const void * pixels) {
//   (*fnptr)(texture, level, xoffset, yoffset, zoffset, width, height, depth, format, type, pixels);
// }
// static void  glowTextureView(GPTEXTUREVIEW fnptr, GLuint  texture, GLenum  target, GLuint  origtexture, GLenum  internalformat, GLuint  minlevel, GLuint  numlevels, GLuint  minlayer, GLuint  numlayers) {
//   (*fnptr)(texture, target, origtexture, internalformat, minlevel, numlevels, minlayer, numlayers);
// }
// static void  glowTransformFeedbackBufferBase(GPTRANSFORMFEEDBACKBUFFERBASE fnptr, GLuint  xfb, GLuint  index, GLuint  buffer) {
//   (*fnptr)(xfb, index, buffer);
// }
// static void  glowTransformFeedbackBufferRange(GPTRANSFORMFEEDBACKBUFFERRANGE fnptr, GLuint  xfb, GLuint  index, GLuint  buffer, GLintptr  offset, GLsizei  size) {
//   (*fnptr)(xfb, index, buffer, offset, size);
// }
// static void  glowTransformFeedbackVaryings(GPTRANSFORMFEEDBACKVARYINGS fnptr, GLuint  program, GLsizei  count, const GLchar *const* varyings, GLenum  bufferMode) {
//   (*fnptr)(program, count, varyings, bufferMode);
// }
// static void  glowTranslatexOES(GPTRANSLATEXOES fnptr, GLfixed  x, GLfixed  y, GLfixed  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowUniform1d(GPUNIFORM1D fnptr, GLint  location, GLdouble  x) {
//   (*fnptr)(location, x);
// }
// static void  glowUniform1dv(GPUNIFORM1DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1f(GPUNIFORM1F fnptr, GLint  location, GLfloat  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1fv(GPUNIFORM1FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1i(GPUNIFORM1I fnptr, GLint  location, GLint  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1iv(GPUNIFORM1IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform1ui(GPUNIFORM1UI fnptr, GLint  location, GLuint  v0) {
//   (*fnptr)(location, v0);
// }
// static void  glowUniform1uiv(GPUNIFORM1UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2d(GPUNIFORM2D fnptr, GLint  location, GLdouble  x, GLdouble  y) {
//   (*fnptr)(location, x, y);
// }
// static void  glowUniform2dv(GPUNIFORM2DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2f(GPUNIFORM2F fnptr, GLint  location, GLfloat  v0, GLfloat  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2fv(GPUNIFORM2FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2i(GPUNIFORM2I fnptr, GLint  location, GLint  v0, GLint  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2iv(GPUNIFORM2IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform2ui(GPUNIFORM2UI fnptr, GLint  location, GLuint  v0, GLuint  v1) {
//   (*fnptr)(location, v0, v1);
// }
// static void  glowUniform2uiv(GPUNIFORM2UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3d(GPUNIFORM3D fnptr, GLint  location, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(location, x, y, z);
// }
// static void  glowUniform3dv(GPUNIFORM3DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3f(GPUNIFORM3F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3fv(GPUNIFORM3FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3i(GPUNIFORM3I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3iv(GPUNIFORM3IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform3ui(GPUNIFORM3UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2) {
//   (*fnptr)(location, v0, v1, v2);
// }
// static void  glowUniform3uiv(GPUNIFORM3UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4d(GPUNIFORM4D fnptr, GLint  location, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(location, x, y, z, w);
// }
// static void  glowUniform4dv(GPUNIFORM4DV fnptr, GLint  location, GLsizei  count, const GLdouble * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4f(GPUNIFORM4F fnptr, GLint  location, GLfloat  v0, GLfloat  v1, GLfloat  v2, GLfloat  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4fv(GPUNIFORM4FV fnptr, GLint  location, GLsizei  count, const GLfloat * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4i(GPUNIFORM4I fnptr, GLint  location, GLint  v0, GLint  v1, GLint  v2, GLint  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4iv(GPUNIFORM4IV fnptr, GLint  location, GLsizei  count, const GLint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniform4ui(GPUNIFORM4UI fnptr, GLint  location, GLuint  v0, GLuint  v1, GLuint  v2, GLuint  v3) {
//   (*fnptr)(location, v0, v1, v2, v3);
// }
// static void  glowUniform4uiv(GPUNIFORM4UIV fnptr, GLint  location, GLsizei  count, const GLuint * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformBlockBinding(GPUNIFORMBLOCKBINDING fnptr, GLuint  program, GLuint  uniformBlockIndex, GLuint  uniformBlockBinding) {
//   (*fnptr)(program, uniformBlockIndex, uniformBlockBinding);
// }
// static void  glowUniformHandleui64ARB(GPUNIFORMHANDLEUI64ARB fnptr, GLint  location, GLuint64  value) {
//   (*fnptr)(location, value);
// }
// static void  glowUniformHandleui64vARB(GPUNIFORMHANDLEUI64VARB fnptr, GLint  location, GLsizei  count, const GLuint64 * value) {
//   (*fnptr)(location, count, value);
// }
// static void  glowUniformMatrix2dv(GPUNIFORMMATRIX2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2fv(GPUNIFORMMATRIX2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x3dv(GPUNIFORMMATRIX2X3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x3fv(GPUNIFORMMATRIX2X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x4dv(GPUNIFORMMATRIX2X4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix2x4fv(GPUNIFORMMATRIX2X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3dv(GPUNIFORMMATRIX3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3fv(GPUNIFORMMATRIX3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x2dv(GPUNIFORMMATRIX3X2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x2fv(GPUNIFORMMATRIX3X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x4dv(GPUNIFORMMATRIX3X4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix3x4fv(GPUNIFORMMATRIX3X4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4dv(GPUNIFORMMATRIX4DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4fv(GPUNIFORMMATRIX4FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x2dv(GPUNIFORMMATRIX4X2DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x2fv(GPUNIFORMMATRIX4X2FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x3dv(GPUNIFORMMATRIX4X3DV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLdouble * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformMatrix4x3fv(GPUNIFORMMATRIX4X3FV fnptr, GLint  location, GLsizei  count, GLboolean  transpose, const GLfloat * value) {
//   (*fnptr)(location, count, transpose, value);
// }
// static void  glowUniformSubroutinesuiv(GPUNIFORMSUBROUTINESUIV fnptr, GLenum  shadertype, GLsizei  count, const GLuint * indices) {
//   (*fnptr)(shadertype, count, indices);
// }
// static GLboolean  glowUnmapBuffer(GPUNMAPBUFFER fnptr, GLenum  target) {
//   return (*fnptr)(target);
// }
// static GLboolean  glowUnmapNamedBuffer(GPUNMAPNAMEDBUFFER fnptr, GLuint  buffer) {
//   return (*fnptr)(buffer);
// }
// static void  glowUseProgram(GPUSEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowUseProgramStages(GPUSEPROGRAMSTAGES fnptr, GLuint  pipeline, GLbitfield  stages, GLuint  program) {
//   (*fnptr)(pipeline, stages, program);
// }
// static void  glowUseProgramStagesEXT(GPUSEPROGRAMSTAGESEXT fnptr, GLuint  pipeline, GLbitfield  stages, GLuint  program) {
//   (*fnptr)(pipeline, stages, program);
// }
// static void  glowUseShaderProgramEXT(GPUSESHADERPROGRAMEXT fnptr, GLenum  type, GLuint  program) {
//   (*fnptr)(type, program);
// }
// static void  glowValidateProgram(GPVALIDATEPROGRAM fnptr, GLuint  program) {
//   (*fnptr)(program);
// }
// static void  glowValidateProgramPipeline(GPVALIDATEPROGRAMPIPELINE fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowValidateProgramPipelineEXT(GPVALIDATEPROGRAMPIPELINEEXT fnptr, GLuint  pipeline) {
//   (*fnptr)(pipeline);
// }
// static void  glowVertex2bOES(GPVERTEX2BOES fnptr, GLbyte  x, GLbyte  y) {
//   (*fnptr)(x, y);
// }
// static void  glowVertex2bvOES(GPVERTEX2BVOES fnptr, const GLbyte * coords) {
//   (*fnptr)(coords);
// }
// static void  glowVertex2xOES(GPVERTEX2XOES fnptr, GLfixed  x) {
//   (*fnptr)(x);
// }
// static void  glowVertex2xvOES(GPVERTEX2XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowVertex3bOES(GPVERTEX3BOES fnptr, GLbyte  x, GLbyte  y, GLbyte  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowVertex3bvOES(GPVERTEX3BVOES fnptr, const GLbyte * coords) {
//   (*fnptr)(coords);
// }
// static void  glowVertex3xOES(GPVERTEX3XOES fnptr, GLfixed  x, GLfixed  y) {
//   (*fnptr)(x, y);
// }
// static void  glowVertex3xvOES(GPVERTEX3XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowVertex4bOES(GPVERTEX4BOES fnptr, GLbyte  x, GLbyte  y, GLbyte  z, GLbyte  w) {
//   (*fnptr)(x, y, z, w);
// }
// static void  glowVertex4bvOES(GPVERTEX4BVOES fnptr, const GLbyte * coords) {
//   (*fnptr)(coords);
// }
// static void  glowVertex4xOES(GPVERTEX4XOES fnptr, GLfixed  x, GLfixed  y, GLfixed  z) {
//   (*fnptr)(x, y, z);
// }
// static void  glowVertex4xvOES(GPVERTEX4XVOES fnptr, const GLfixed * coords) {
//   (*fnptr)(coords);
// }
// static void  glowVertexArrayAttribBinding(GPVERTEXARRAYATTRIBBINDING fnptr, GLuint  vaobj, GLuint  attribindex, GLuint  bindingindex) {
//   (*fnptr)(vaobj, attribindex, bindingindex);
// }
// static void  glowVertexArrayAttribFormat(GPVERTEXARRAYATTRIBFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset) {
//   (*fnptr)(vaobj, attribindex, size, type, normalized, relativeoffset);
// }
// static void  glowVertexArrayAttribIFormat(GPVERTEXARRAYATTRIBIFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(vaobj, attribindex, size, type, relativeoffset);
// }
// static void  glowVertexArrayAttribLFormat(GPVERTEXARRAYATTRIBLFORMAT fnptr, GLuint  vaobj, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(vaobj, attribindex, size, type, relativeoffset);
// }
// static void  glowVertexArrayBindingDivisor(GPVERTEXARRAYBINDINGDIVISOR fnptr, GLuint  vaobj, GLuint  bindingindex, GLuint  divisor) {
//   (*fnptr)(vaobj, bindingindex, divisor);
// }
// static void  glowVertexArrayElementBuffer(GPVERTEXARRAYELEMENTBUFFER fnptr, GLuint  vaobj, GLuint  buffer) {
//   (*fnptr)(vaobj, buffer);
// }
// static void  glowVertexArrayVertexBuffer(GPVERTEXARRAYVERTEXBUFFER fnptr, GLuint  vaobj, GLuint  bindingindex, GLuint  buffer, GLintptr  offset, GLsizei  stride) {
//   (*fnptr)(vaobj, bindingindex, buffer, offset, stride);
// }
// static void  glowVertexArrayVertexBuffers(GPVERTEXARRAYVERTEXBUFFERS fnptr, GLuint  vaobj, GLuint  first, GLsizei  count, const GLuint * buffers, const GLintptr * offsets, const GLsizei * strides) {
//   (*fnptr)(vaobj, first, count, buffers, offsets, strides);
// }
// static void  glowVertexAttrib1d(GPVERTEXATTRIB1D fnptr, GLuint  index, GLdouble  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttrib1dv(GPVERTEXATTRIB1DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib1f(GPVERTEXATTRIB1F fnptr, GLuint  index, GLfloat  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttrib1fv(GPVERTEXATTRIB1FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib1s(GPVERTEXATTRIB1S fnptr, GLuint  index, GLshort  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttrib1sv(GPVERTEXATTRIB1SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib2d(GPVERTEXATTRIB2D fnptr, GLuint  index, GLdouble  x, GLdouble  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttrib2dv(GPVERTEXATTRIB2DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib2f(GPVERTEXATTRIB2F fnptr, GLuint  index, GLfloat  x, GLfloat  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttrib2fv(GPVERTEXATTRIB2FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib2s(GPVERTEXATTRIB2S fnptr, GLuint  index, GLshort  x, GLshort  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttrib2sv(GPVERTEXATTRIB2SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib3d(GPVERTEXATTRIB3D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttrib3dv(GPVERTEXATTRIB3DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib3f(GPVERTEXATTRIB3F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttrib3fv(GPVERTEXATTRIB3FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib3s(GPVERTEXATTRIB3S fnptr, GLuint  index, GLshort  x, GLshort  y, GLshort  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttrib3sv(GPVERTEXATTRIB3SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nbv(GPVERTEXATTRIB4NBV fnptr, GLuint  index, const GLbyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Niv(GPVERTEXATTRIB4NIV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nsv(GPVERTEXATTRIB4NSV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nub(GPVERTEXATTRIB4NUB fnptr, GLuint  index, GLubyte  x, GLubyte  y, GLubyte  z, GLubyte  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4Nubv(GPVERTEXATTRIB4NUBV fnptr, GLuint  index, const GLubyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nuiv(GPVERTEXATTRIB4NUIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4Nusv(GPVERTEXATTRIB4NUSV fnptr, GLuint  index, const GLushort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4bv(GPVERTEXATTRIB4BV fnptr, GLuint  index, const GLbyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4d(GPVERTEXATTRIB4D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4dv(GPVERTEXATTRIB4DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4f(GPVERTEXATTRIB4F fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  z, GLfloat  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4fv(GPVERTEXATTRIB4FV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4iv(GPVERTEXATTRIB4IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4s(GPVERTEXATTRIB4S fnptr, GLuint  index, GLshort  x, GLshort  y, GLshort  z, GLshort  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttrib4sv(GPVERTEXATTRIB4SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4ubv(GPVERTEXATTRIB4UBV fnptr, GLuint  index, const GLubyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4uiv(GPVERTEXATTRIB4UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttrib4usv(GPVERTEXATTRIB4USV fnptr, GLuint  index, const GLushort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribBinding(GPVERTEXATTRIBBINDING fnptr, GLuint  attribindex, GLuint  bindingindex) {
//   (*fnptr)(attribindex, bindingindex);
// }
// static void  glowVertexAttribDivisor(GPVERTEXATTRIBDIVISOR fnptr, GLuint  index, GLuint  divisor) {
//   (*fnptr)(index, divisor);
// }
// static void  glowVertexAttribFormat(GPVERTEXATTRIBFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLboolean  normalized, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, normalized, relativeoffset);
// }
// static void  glowVertexAttribI1i(GPVERTEXATTRIBI1I fnptr, GLuint  index, GLint  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribI1iv(GPVERTEXATTRIBI1IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI1ui(GPVERTEXATTRIBI1UI fnptr, GLuint  index, GLuint  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribI1uiv(GPVERTEXATTRIBI1UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI2i(GPVERTEXATTRIBI2I fnptr, GLuint  index, GLint  x, GLint  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttribI2iv(GPVERTEXATTRIBI2IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI2ui(GPVERTEXATTRIBI2UI fnptr, GLuint  index, GLuint  x, GLuint  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttribI2uiv(GPVERTEXATTRIBI2UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI3i(GPVERTEXATTRIBI3I fnptr, GLuint  index, GLint  x, GLint  y, GLint  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttribI3iv(GPVERTEXATTRIBI3IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI3ui(GPVERTEXATTRIBI3UI fnptr, GLuint  index, GLuint  x, GLuint  y, GLuint  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttribI3uiv(GPVERTEXATTRIBI3UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4bv(GPVERTEXATTRIBI4BV fnptr, GLuint  index, const GLbyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4i(GPVERTEXATTRIBI4I fnptr, GLuint  index, GLint  x, GLint  y, GLint  z, GLint  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribI4iv(GPVERTEXATTRIBI4IV fnptr, GLuint  index, const GLint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4sv(GPVERTEXATTRIBI4SV fnptr, GLuint  index, const GLshort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4ubv(GPVERTEXATTRIBI4UBV fnptr, GLuint  index, const GLubyte * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4ui(GPVERTEXATTRIBI4UI fnptr, GLuint  index, GLuint  x, GLuint  y, GLuint  z, GLuint  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribI4uiv(GPVERTEXATTRIBI4UIV fnptr, GLuint  index, const GLuint * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribI4usv(GPVERTEXATTRIBI4USV fnptr, GLuint  index, const GLushort * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribIFormat(GPVERTEXATTRIBIFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, relativeoffset);
// }
// static void  glowVertexAttribIPointer(GPVERTEXATTRIBIPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, stride, pointer);
// }
// static void  glowVertexAttribL1d(GPVERTEXATTRIBL1D fnptr, GLuint  index, GLdouble  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribL1dv(GPVERTEXATTRIBL1DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL1ui64ARB(GPVERTEXATTRIBL1UI64ARB fnptr, GLuint  index, GLuint64EXT  x) {
//   (*fnptr)(index, x);
// }
// static void  glowVertexAttribL1ui64vARB(GPVERTEXATTRIBL1UI64VARB fnptr, GLuint  index, const GLuint64EXT * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL2d(GPVERTEXATTRIBL2D fnptr, GLuint  index, GLdouble  x, GLdouble  y) {
//   (*fnptr)(index, x, y);
// }
// static void  glowVertexAttribL2dv(GPVERTEXATTRIBL2DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL3d(GPVERTEXATTRIBL3D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z) {
//   (*fnptr)(index, x, y, z);
// }
// static void  glowVertexAttribL3dv(GPVERTEXATTRIBL3DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribL4d(GPVERTEXATTRIBL4D fnptr, GLuint  index, GLdouble  x, GLdouble  y, GLdouble  z, GLdouble  w) {
//   (*fnptr)(index, x, y, z, w);
// }
// static void  glowVertexAttribL4dv(GPVERTEXATTRIBL4DV fnptr, GLuint  index, const GLdouble * v) {
//   (*fnptr)(index, v);
// }
// static void  glowVertexAttribLFormat(GPVERTEXATTRIBLFORMAT fnptr, GLuint  attribindex, GLint  size, GLenum  type, GLuint  relativeoffset) {
//   (*fnptr)(attribindex, size, type, relativeoffset);
// }
// static void  glowVertexAttribLPointer(GPVERTEXATTRIBLPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, stride, pointer);
// }
// static void  glowVertexAttribP1ui(GPVERTEXATTRIBP1UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP1uiv(GPVERTEXATTRIBP1UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP2ui(GPVERTEXATTRIBP2UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP2uiv(GPVERTEXATTRIBP2UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP3ui(GPVERTEXATTRIBP3UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP3uiv(GPVERTEXATTRIBP3UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP4ui(GPVERTEXATTRIBP4UI fnptr, GLuint  index, GLenum  type, GLboolean  normalized, GLuint  value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribP4uiv(GPVERTEXATTRIBP4UIV fnptr, GLuint  index, GLenum  type, GLboolean  normalized, const GLuint * value) {
//   (*fnptr)(index, type, normalized, value);
// }
// static void  glowVertexAttribPointer(GPVERTEXATTRIBPOINTER fnptr, GLuint  index, GLint  size, GLenum  type, GLboolean  normalized, GLsizei  stride, const void * pointer) {
//   (*fnptr)(index, size, type, normalized, stride, pointer);
// }
// static void  glowVertexBindingDivisor(GPVERTEXBINDINGDIVISOR fnptr, GLuint  bindingindex, GLuint  divisor) {
//   (*fnptr)(bindingindex, divisor);
// }
// static void  glowViewport(GPVIEWPORT fnptr, GLint  x, GLint  y, GLsizei  width, GLsizei  height) {
//   (*fnptr)(x, y, width, height);
// }
// static void  glowViewportArrayv(GPVIEWPORTARRAYV fnptr, GLuint  first, GLsizei  count, const GLfloat * v) {
//   (*fnptr)(first, count, v);
// }
// static void  glowViewportIndexedf(GPVIEWPORTINDEXEDF fnptr, GLuint  index, GLfloat  x, GLfloat  y, GLfloat  w, GLfloat  h) {
//   (*fnptr)(index, x, y, w, h);
// }
// static void  glowViewportIndexedfv(GPVIEWPORTINDEXEDFV fnptr, GLuint  index, const GLfloat * v) {
//   (*fnptr)(index, v);
// }
// static void  glowWaitSync(GPWAITSYNC fnptr, GLsync  sync, GLbitfield  flags, GLuint64  timeout) {
//   (*fnptr)(sync, flags, timeout);
// }
import "C"
import (
	"errors"
	"github.com/go-gl/gl/procaddr"
	"github.com/go-gl/gl/procaddr/auto"
	"unsafe"
)

const (
	ACTIVE_ATOMIC_COUNTER_BUFFERS                              = 0x92D9
	ACTIVE_ATTRIBUTES                                          = 0x8B89
	ACTIVE_ATTRIBUTE_MAX_LENGTH                                = 0x8B8A
	ACTIVE_PROGRAM                                             = 0x8259
	ACTIVE_PROGRAM_EXT                                         = 0x8B8D
	ACTIVE_RESOURCES                                           = 0x92F5
	ACTIVE_SUBROUTINES                                         = 0x8DE5
	ACTIVE_SUBROUTINE_MAX_LENGTH                               = 0x8E48
	ACTIVE_SUBROUTINE_UNIFORMS                                 = 0x8DE6
	ACTIVE_SUBROUTINE_UNIFORM_LOCATIONS                        = 0x8E47
	ACTIVE_SUBROUTINE_UNIFORM_MAX_LENGTH                       = 0x8E49
	ACTIVE_TEXTURE                                             = 0x84E0
	ACTIVE_UNIFORMS                                            = 0x8B86
	ACTIVE_UNIFORM_BLOCKS                                      = 0x8A36
	ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH                       = 0x8A35
	ACTIVE_UNIFORM_MAX_LENGTH                                  = 0x8B87
	ACTIVE_VARIABLES                                           = 0x9305
	ALIASED_LINE_WIDTH_RANGE                                   = 0x846E
	ALL_BARRIER_BITS                                           = 0xFFFFFFFF
	ALL_COMPLETED_NV                                           = 0x84F2
	ALL_SHADER_BITS                                            = 0xFFFFFFFF
	ALL_SHADER_BITS_EXT                                        = 0xFFFFFFFF
	ALPHA                                                      = 0x1906
	ALREADY_SIGNALED                                           = 0x911A
	ALWAYS                                                     = 0x0207
	AND                                                        = 0x1501
	AND_INVERTED                                               = 0x1504
	AND_REVERSE                                                = 0x1502
	ANY_SAMPLES_PASSED                                         = 0x8C2F
	ANY_SAMPLES_PASSED_CONSERVATIVE                            = 0x8D6A
	ARRAY_BUFFER                                               = 0x8892
	ARRAY_BUFFER_BINDING                                       = 0x8894
	ARRAY_SIZE                                                 = 0x92FB
	ARRAY_STRIDE                                               = 0x92FE
	ATOMIC_COUNTER_BARRIER_BIT                                 = 0x00001000
	ATOMIC_COUNTER_BUFFER                                      = 0x92C0
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTERS               = 0x92C5
	ATOMIC_COUNTER_BUFFER_ACTIVE_ATOMIC_COUNTER_INDICES        = 0x92C6
	ATOMIC_COUNTER_BUFFER_BINDING                              = 0x92C1
	ATOMIC_COUNTER_BUFFER_DATA_SIZE                            = 0x92C4
	ATOMIC_COUNTER_BUFFER_INDEX                                = 0x9301
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_COMPUTE_SHADER         = 0x90ED
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_FRAGMENT_SHADER        = 0x92CB
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_GEOMETRY_SHADER        = 0x92CA
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_CONTROL_SHADER    = 0x92C8
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_TESS_EVALUATION_SHADER = 0x92C9
	ATOMIC_COUNTER_BUFFER_REFERENCED_BY_VERTEX_SHADER          = 0x92C7
	ATOMIC_COUNTER_BUFFER_SIZE                                 = 0x92C3
	ATOMIC_COUNTER_BUFFER_START                                = 0x92C2
	ATTACHED_SHADERS                                           = 0x8B85
	AUTO_GENERATE_MIPMAP                                       = 0x8295
	BACK                                                       = 0x0405
	BACK_LEFT                                                  = 0x0402
	BACK_RIGHT                                                 = 0x0403
	BGR                                                        = 0x80E0
	BGRA                                                       = 0x80E1
	BGRA_INTEGER                                               = 0x8D9B
	BGR_INTEGER                                                = 0x8D9A
	BLEND                                                      = 0x0BE2
	BLEND_ADVANCED_COHERENT_KHR                                = 0x9285
	BLEND_ADVANCED_COHERENT_NV                                 = 0x9285
	BLEND_COLOR                                                = 0x8005
	BLEND_DST                                                  = 0x0BE0
	BLEND_DST_ALPHA                                            = 0x80CA
	BLEND_DST_RGB                                              = 0x80C8
	BLEND_EQUATION                                             = 0x8009
	BLEND_EQUATION_ALPHA                                       = 0x883D
	BLEND_EQUATION_EXT                                         = 0x8009
	BLEND_EQUATION_RGB                                         = 0x8009
	BLEND_OVERLAP_NV                                           = 0x9281
	BLEND_PREMULTIPLIED_SRC_NV                                 = 0x9280
	BLEND_SRC                                                  = 0x0BE1
	BLEND_SRC_ALPHA                                            = 0x80CB
	BLEND_SRC_RGB                                              = 0x80C9
	BLOCK_INDEX                                                = 0x92FD
	BLUE                                                       = 0x1905
	BLUE_INTEGER                                               = 0x8D96
	BLUE_NV                                                    = 0x1905
	BOOL                                                       = 0x8B56
	BOOL_VEC2                                                  = 0x8B57
	BOOL_VEC3                                                  = 0x8B58
	BOOL_VEC4                                                  = 0x8B59
	BUFFER                                                     = 0x82E0
	BUFFER_ACCESS                                              = 0x88BB
	BUFFER_ACCESS_FLAGS                                        = 0x911F
	BUFFER_BINDING                                             = 0x9302
	BUFFER_DATA_SIZE                                           = 0x9303
	BUFFER_IMMUTABLE_STORAGE                                   = 0x821F
	BUFFER_KHR                                                 = 0x82E0
	BUFFER_MAPPED                                              = 0x88BC
	BUFFER_MAP_LENGTH                                          = 0x9120
	BUFFER_MAP_OFFSET                                          = 0x9121
	BUFFER_MAP_POINTER                                         = 0x88BD
	BUFFER_OBJECT_EXT                                          = 0x9151
	BUFFER_SIZE                                                = 0x8764
	BUFFER_STORAGE_FLAGS                                       = 0x8220
	BUFFER_UPDATE_BARRIER_BIT                                  = 0x00000200
	BUFFER_USAGE                                               = 0x8765
	BUFFER_VARIABLE                                            = 0x92E5
	BYTE                                                       = 0x1400
	CAVEAT_SUPPORT                                             = 0x82B8
	CCW                                                        = 0x0901
	CLAMP_READ_COLOR                                           = 0x891C
	CLAMP_TO_BORDER                                            = 0x812D
	CLAMP_TO_EDGE                                              = 0x812F
	CLEAR                                                      = 0x1500
	CLEAR_BUFFER                                               = 0x82B4
	CLEAR_TEXTURE                                              = 0x9365
	CLIENT_MAPPED_BUFFER_BARRIER_BIT                           = 0x00004000
	CLIENT_STORAGE_BIT                                         = 0x0200
	CLIPPING_INPUT_PRIMITIVES_ARB                              = 0x82F6
	CLIPPING_OUTPUT_PRIMITIVES_ARB                             = 0x82F7
	CLIP_DEPTH_MODE                                            = 0x935D
	CLIP_DISTANCE0                                             = 0x3000
	CLIP_DISTANCE1                                             = 0x3001
	CLIP_DISTANCE2                                             = 0x3002
	CLIP_DISTANCE3                                             = 0x3003
	CLIP_DISTANCE4                                             = 0x3004
	CLIP_DISTANCE5                                             = 0x3005
	CLIP_DISTANCE6                                             = 0x3006
	CLIP_DISTANCE7                                             = 0x3007
	CLIP_ORIGIN                                                = 0x935C
	COLOR                                                      = 0x1800
	COLORBURN_KHR                                              = 0x929A
	COLORBURN_NV                                               = 0x929A
	COLORDODGE_KHR                                             = 0x9299
	COLORDODGE_NV                                              = 0x9299
	COLOR_ATTACHMENT0                                          = 0x8CE0
	COLOR_ATTACHMENT1                                          = 0x8CE1
	COLOR_ATTACHMENT10                                         = 0x8CEA
	COLOR_ATTACHMENT11                                         = 0x8CEB
	COLOR_ATTACHMENT12                                         = 0x8CEC
	COLOR_ATTACHMENT13                                         = 0x8CED
	COLOR_ATTACHMENT14                                         = 0x8CEE
	COLOR_ATTACHMENT15                                         = 0x8CEF
	COLOR_ATTACHMENT2                                          = 0x8CE2
	COLOR_ATTACHMENT3                                          = 0x8CE3
	COLOR_ATTACHMENT4                                          = 0x8CE4
	COLOR_ATTACHMENT5                                          = 0x8CE5
	COLOR_ATTACHMENT6                                          = 0x8CE6
	COLOR_ATTACHMENT7                                          = 0x8CE7
	COLOR_ATTACHMENT8                                          = 0x8CE8
	COLOR_ATTACHMENT9                                          = 0x8CE9
	COLOR_BUFFER_BIT                                           = 0x00004000
	COLOR_CLEAR_VALUE                                          = 0x0C22
	COLOR_COMPONENTS                                           = 0x8283
	COLOR_ENCODING                                             = 0x8296
	COLOR_LOGIC_OP                                             = 0x0BF2
	COLOR_RENDERABLE                                           = 0x8286
	COLOR_WRITEMASK                                            = 0x0C23
	COMMAND_BARRIER_BIT                                        = 0x00000040
	COMPARE_REF_TO_TEXTURE                                     = 0x884E
	COMPATIBLE_SUBROUTINES                                     = 0x8E4B
	COMPILE_STATUS                                             = 0x8B81
	COMPRESSED_R11_EAC                                         = 0x9270
	COMPRESSED_RED                                             = 0x8225
	COMPRESSED_RED_RGTC1                                       = 0x8DBB
	COMPRESSED_RG                                              = 0x8226
	COMPRESSED_RG11_EAC                                        = 0x9272
	COMPRESSED_RGB                                             = 0x84ED
	COMPRESSED_RGB8_ETC2                                       = 0x9274
	COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2                   = 0x9276
	COMPRESSED_RGBA                                            = 0x84EE
	COMPRESSED_RGBA8_ETC2_EAC                                  = 0x9278
	COMPRESSED_RGBA_ASTC_10x10_KHR                             = 0x93BB
	COMPRESSED_RGBA_ASTC_10x5_KHR                              = 0x93B8
	COMPRESSED_RGBA_ASTC_10x6_KHR                              = 0x93B9
	COMPRESSED_RGBA_ASTC_10x8_KHR                              = 0x93BA
	COMPRESSED_RGBA_ASTC_12x10_KHR                             = 0x93BC
	COMPRESSED_RGBA_ASTC_12x12_KHR                             = 0x93BD
	COMPRESSED_RGBA_ASTC_4x4_KHR                               = 0x93B0
	COMPRESSED_RGBA_ASTC_5x4_KHR                               = 0x93B1
	COMPRESSED_RGBA_ASTC_5x5_KHR                               = 0x93B2
	COMPRESSED_RGBA_ASTC_6x5_KHR                               = 0x93B3
	COMPRESSED_RGBA_ASTC_6x6_KHR                               = 0x93B4
	COMPRESSED_RGBA_ASTC_8x5_KHR                               = 0x93B5
	COMPRESSED_RGBA_ASTC_8x6_KHR                               = 0x93B6
	COMPRESSED_RGBA_ASTC_8x8_KHR                               = 0x93B7
	COMPRESSED_RGBA_BPTC_UNORM_ARB                             = 0x8E8C
	COMPRESSED_RGBA_S3TC_DXT1_EXT                              = 0x83F1
	COMPRESSED_RGBA_S3TC_DXT3_EXT                              = 0x83F2
	COMPRESSED_RGBA_S3TC_DXT5_EXT                              = 0x83F3
	COMPRESSED_RGB_BPTC_SIGNED_FLOAT_ARB                       = 0x8E8E
	COMPRESSED_RGB_BPTC_UNSIGNED_FLOAT_ARB                     = 0x8E8F
	COMPRESSED_RGB_S3TC_DXT1_EXT                               = 0x83F0
	COMPRESSED_RG_RGTC2                                        = 0x8DBD
	COMPRESSED_SIGNED_R11_EAC                                  = 0x9271
	COMPRESSED_SIGNED_RED_RGTC1                                = 0x8DBC
	COMPRESSED_SIGNED_RG11_EAC                                 = 0x9273
	COMPRESSED_SIGNED_RG_RGTC2                                 = 0x8DBE
	COMPRESSED_SRGB                                            = 0x8C48
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR                     = 0x93DB
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR                      = 0x93D8
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR                      = 0x93D9
	COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR                      = 0x93DA
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR                     = 0x93DC
	COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR                     = 0x93DD
	COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR                       = 0x93D0
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR                       = 0x93D1
	COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR                       = 0x93D2
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR                       = 0x93D3
	COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR                       = 0x93D4
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR                       = 0x93D5
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR                       = 0x93D6
	COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR                       = 0x93D7
	COMPRESSED_SRGB8_ALPHA8_ETC2_EAC                           = 0x9279
	COMPRESSED_SRGB8_ETC2                                      = 0x9275
	COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2                  = 0x9277
	COMPRESSED_SRGB_ALPHA                                      = 0x8C49
	COMPRESSED_SRGB_ALPHA_BPTC_UNORM_ARB                       = 0x8E8D
	COMPRESSED_TEXTURE_FORMATS                                 = 0x86A3
	COMPUTE_SHADER                                             = 0x91B9
	COMPUTE_SHADER_BIT                                         = 0x00000020
	COMPUTE_SHADER_INVOCATIONS_ARB                             = 0x82F5
	COMPUTE_SUBROUTINE                                         = 0x92ED
	COMPUTE_SUBROUTINE_UNIFORM                                 = 0x92F3
	COMPUTE_TEXTURE                                            = 0x82A0
	COMPUTE_WORK_GROUP_SIZE                                    = 0x8267
	CONDITION_SATISFIED                                        = 0x911C
	CONJOINT_NV                                                = 0x9284
	CONSTANT_ALPHA                                             = 0x8003
	CONSTANT_COLOR                                             = 0x8001
	CONTEXT_COMPATIBILITY_PROFILE_BIT                          = 0x00000002
	CONTEXT_CORE_PROFILE_BIT                                   = 0x00000001
	CONTEXT_FLAGS                                              = 0x821E
	CONTEXT_FLAG_DEBUG_BIT                                     = 0x00000002
	CONTEXT_FLAG_DEBUG_BIT_KHR                                 = 0x00000002
	CONTEXT_FLAG_FORWARD_COMPATIBLE_BIT                        = 0x00000001
	CONTEXT_FLAG_ROBUST_ACCESS_BIT_ARB                         = 0x00000004
	CONTEXT_LOST                                               = 0x0507
	CONTEXT_LOST_KHR                                           = 0x0507
	CONTEXT_PROFILE_MASK                                       = 0x9126
	CONTEXT_RELEASE_BEHAVIOR                                   = 0x82FB
	CONTEXT_RELEASE_BEHAVIOR_FLUSH                             = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_FLUSH_KHR                         = 0x82FC
	CONTEXT_RELEASE_BEHAVIOR_KHR                               = 0x82FB
	CONTEXT_ROBUST_ACCESS                                      = 0x90F3
	CONTEXT_ROBUST_ACCESS_KHR                                  = 0x90F3
	CONTRAST_NV                                                = 0x92A1
	COPY                                                       = 0x1503
	COPY_INVERTED                                              = 0x150C
	COPY_READ_BUFFER                                           = 0x8F36
	COPY_READ_BUFFER_BINDING                                   = 0x8F36
	COPY_WRITE_BUFFER                                          = 0x8F37
	COPY_WRITE_BUFFER_BINDING                                  = 0x8F37
	COUNTER_RANGE_AMD                                          = 0x8BC1
	COUNTER_TYPE_AMD                                           = 0x8BC0
	CULL_FACE                                                  = 0x0B44
	CULL_FACE_MODE                                             = 0x0B45
	CURRENT_PROGRAM                                            = 0x8B8D
	CURRENT_QUERY                                              = 0x8865
	CURRENT_VERTEX_ATTRIB                                      = 0x8626
	CW                                                         = 0x0900
	DARKEN_KHR                                                 = 0x9297
	DARKEN_NV                                                  = 0x9297
	DEBUG_CALLBACK_FUNCTION                                    = 0x8244
	DEBUG_CALLBACK_FUNCTION_ARB                                = 0x8244
	DEBUG_CALLBACK_FUNCTION_KHR                                = 0x8244
	DEBUG_CALLBACK_USER_PARAM                                  = 0x8245
	DEBUG_CALLBACK_USER_PARAM_ARB                              = 0x8245
	DEBUG_CALLBACK_USER_PARAM_KHR                              = 0x8245
	DEBUG_GROUP_STACK_DEPTH                                    = 0x826D
	DEBUG_GROUP_STACK_DEPTH_KHR                                = 0x826D
	DEBUG_LOGGED_MESSAGES                                      = 0x9145
	DEBUG_LOGGED_MESSAGES_ARB                                  = 0x9145
	DEBUG_LOGGED_MESSAGES_KHR                                  = 0x9145
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH                           = 0x8243
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_ARB                       = 0x8243
	DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR                       = 0x8243
	DEBUG_OUTPUT                                               = 0x92E0
	DEBUG_OUTPUT_KHR                                           = 0x92E0
	DEBUG_OUTPUT_SYNCHRONOUS                                   = 0x8242
	DEBUG_OUTPUT_SYNCHRONOUS_ARB                               = 0x8242
	DEBUG_OUTPUT_SYNCHRONOUS_KHR                               = 0x8242
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	DEBUG_SEVERITY_HIGH_ARB                                    = 0x9146
	DEBUG_SEVERITY_HIGH_KHR                                    = 0x9146
	DEBUG_SEVERITY_LOW                                         = 0x9148
	DEBUG_SEVERITY_LOW_ARB                                     = 0x9148
	DEBUG_SEVERITY_LOW_KHR                                     = 0x9148
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	DEBUG_SEVERITY_MEDIUM_ARB                                  = 0x9147
	DEBUG_SEVERITY_MEDIUM_KHR                                  = 0x9147
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	DEBUG_SEVERITY_NOTIFICATION_KHR                            = 0x826B
	DEBUG_SOURCE_API                                           = 0x8246
	DEBUG_SOURCE_API_ARB                                       = 0x8246
	DEBUG_SOURCE_API_KHR                                       = 0x8246
	DEBUG_SOURCE_APPLICATION                                   = 0x824A
	DEBUG_SOURCE_APPLICATION_ARB                               = 0x824A
	DEBUG_SOURCE_APPLICATION_KHR                               = 0x824A
	DEBUG_SOURCE_OTHER                                         = 0x824B
	DEBUG_SOURCE_OTHER_ARB                                     = 0x824B
	DEBUG_SOURCE_OTHER_KHR                                     = 0x824B
	DEBUG_SOURCE_SHADER_COMPILER                               = 0x8248
	DEBUG_SOURCE_SHADER_COMPILER_ARB                           = 0x8248
	DEBUG_SOURCE_SHADER_COMPILER_KHR                           = 0x8248
	DEBUG_SOURCE_THIRD_PARTY                                   = 0x8249
	DEBUG_SOURCE_THIRD_PARTY_ARB                               = 0x8249
	DEBUG_SOURCE_THIRD_PARTY_KHR                               = 0x8249
	DEBUG_SOURCE_WINDOW_SYSTEM                                 = 0x8247
	DEBUG_SOURCE_WINDOW_SYSTEM_ARB                             = 0x8247
	DEBUG_SOURCE_WINDOW_SYSTEM_KHR                             = 0x8247
	DEBUG_TYPE_DEPRECATED_BEHAVIOR                             = 0x824D
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_ARB                         = 0x824D
	DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR                         = 0x824D
	DEBUG_TYPE_ERROR                                           = 0x824C
	DEBUG_TYPE_ERROR_ARB                                       = 0x824C
	DEBUG_TYPE_ERROR_KHR                                       = 0x824C
	DEBUG_TYPE_MARKER                                          = 0x8268
	DEBUG_TYPE_MARKER_KHR                                      = 0x8268
	DEBUG_TYPE_OTHER                                           = 0x8251
	DEBUG_TYPE_OTHER_ARB                                       = 0x8251
	DEBUG_TYPE_OTHER_KHR                                       = 0x8251
	DEBUG_TYPE_PERFORMANCE                                     = 0x8250
	DEBUG_TYPE_PERFORMANCE_ARB                                 = 0x8250
	DEBUG_TYPE_PERFORMANCE_KHR                                 = 0x8250
	DEBUG_TYPE_POP_GROUP                                       = 0x826A
	DEBUG_TYPE_POP_GROUP_KHR                                   = 0x826A
	DEBUG_TYPE_PORTABILITY                                     = 0x824F
	DEBUG_TYPE_PORTABILITY_ARB                                 = 0x824F
	DEBUG_TYPE_PORTABILITY_KHR                                 = 0x824F
	DEBUG_TYPE_PUSH_GROUP                                      = 0x8269
	DEBUG_TYPE_PUSH_GROUP_KHR                                  = 0x8269
	DEBUG_TYPE_UNDEFINED_BEHAVIOR                              = 0x824E
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_ARB                          = 0x824E
	DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR                          = 0x824E
	DECODE_EXT                                                 = 0x8A49
	DECR                                                       = 0x1E03
	DECR_WRAP                                                  = 0x8508
	DELETE_STATUS                                              = 0x8B80
	DEPTH                                                      = 0x1801
	DEPTH24_STENCIL8                                           = 0x88F0
	DEPTH32F_STENCIL8                                          = 0x8CAD
	DEPTH_ATTACHMENT                                           = 0x8D00
	DEPTH_BUFFER_BIT                                           = 0x00000100
	DEPTH_CLAMP                                                = 0x864F
	DEPTH_CLEAR_VALUE                                          = 0x0B73
	DEPTH_COMPONENT                                            = 0x1902
	DEPTH_COMPONENT16                                          = 0x81A5
	DEPTH_COMPONENT24                                          = 0x81A6
	DEPTH_COMPONENT32                                          = 0x81A7
	DEPTH_COMPONENT32F                                         = 0x8CAC
	DEPTH_COMPONENTS                                           = 0x8284
	DEPTH_FUNC                                                 = 0x0B74
	DEPTH_RANGE                                                = 0x0B70
	DEPTH_RENDERABLE                                           = 0x8287
	DEPTH_STENCIL                                              = 0x84F9
	DEPTH_STENCIL_ATTACHMENT                                   = 0x821A
	DEPTH_STENCIL_TEXTURE_MODE                                 = 0x90EA
	DEPTH_TEST                                                 = 0x0B71
	DEPTH_WRITEMASK                                            = 0x0B72
	DIFFERENCE_KHR                                             = 0x929E
	DIFFERENCE_NV                                              = 0x929E
	DISJOINT_NV                                                = 0x9283
	DISPATCH_INDIRECT_BUFFER                                   = 0x90EE
	DISPATCH_INDIRECT_BUFFER_BINDING                           = 0x90EF
	DITHER                                                     = 0x0BD0
	DONT_CARE                                                  = 0x1100
	DOUBLE                                                     = 0x140A
	DOUBLEBUFFER                                               = 0x0C32
	DOUBLE_MAT2                                                = 0x8F46
	DOUBLE_MAT2x3                                              = 0x8F49
	DOUBLE_MAT2x4                                              = 0x8F4A
	DOUBLE_MAT3                                                = 0x8F47
	DOUBLE_MAT3x2                                              = 0x8F4B
	DOUBLE_MAT3x4                                              = 0x8F4C
	DOUBLE_MAT4                                                = 0x8F48
	DOUBLE_MAT4x2                                              = 0x8F4D
	DOUBLE_MAT4x3                                              = 0x8F4E
	DOUBLE_VEC2                                                = 0x8FFC
	DOUBLE_VEC3                                                = 0x8FFD
	DOUBLE_VEC4                                                = 0x8FFE
	DRAW_BUFFER                                                = 0x0C01
	DRAW_BUFFER0                                               = 0x8825
	DRAW_BUFFER1                                               = 0x8826
	DRAW_BUFFER10                                              = 0x882F
	DRAW_BUFFER11                                              = 0x8830
	DRAW_BUFFER12                                              = 0x8831
	DRAW_BUFFER13                                              = 0x8832
	DRAW_BUFFER14                                              = 0x8833
	DRAW_BUFFER15                                              = 0x8834
	DRAW_BUFFER2                                               = 0x8827
	DRAW_BUFFER3                                               = 0x8828
	DRAW_BUFFER4                                               = 0x8829
	DRAW_BUFFER5                                               = 0x882A
	DRAW_BUFFER6                                               = 0x882B
	DRAW_BUFFER7                                               = 0x882C
	DRAW_BUFFER8                                               = 0x882D
	DRAW_BUFFER9                                               = 0x882E
	DRAW_FRAMEBUFFER                                           = 0x8CA9
	DRAW_FRAMEBUFFER_BINDING                                   = 0x8CA6
	DRAW_INDIRECT_BUFFER                                       = 0x8F3F
	DRAW_INDIRECT_BUFFER_BINDING                               = 0x8F43
	DST_ALPHA                                                  = 0x0304
	DST_ATOP_NV                                                = 0x928F
	DST_COLOR                                                  = 0x0306
	DST_IN_NV                                                  = 0x928B
	DST_NV                                                     = 0x9287
	DST_OUT_NV                                                 = 0x928D
	DST_OVER_NV                                                = 0x9289
	DYNAMIC_COPY                                               = 0x88EA
	DYNAMIC_DRAW                                               = 0x88E8
	DYNAMIC_READ                                               = 0x88E9
	DYNAMIC_STORAGE_BIT                                        = 0x0100
	ELEMENT_ARRAY_BARRIER_BIT                                  = 0x00000002
	ELEMENT_ARRAY_BUFFER                                       = 0x8893
	ELEMENT_ARRAY_BUFFER_BINDING                               = 0x8895
	EQUAL                                                      = 0x0202
	EQUIV                                                      = 0x1509
	EXCLUSION_KHR                                              = 0x92A0
	EXCLUSION_NV                                               = 0x92A0
	EXTENSIONS                                                 = 0x1F03
	FALSE                                                      = 0
	FASTEST                                                    = 0x1101
	FENCE_CONDITION_NV                                         = 0x84F4
	FENCE_STATUS_NV                                            = 0x84F3
	FILL                                                       = 0x1B02
	FILTER                                                     = 0x829A
	FIRST_VERTEX_CONVENTION                                    = 0x8E4D
	FIXED                                                      = 0x140C
	FIXED_OES                                                  = 0x140C
	FIXED_ONLY                                                 = 0x891D
	FLOAT                                                      = 0x1406
	FLOAT_32_UNSIGNED_INT_24_8_REV                             = 0x8DAD
	FLOAT_MAT2                                                 = 0x8B5A
	FLOAT_MAT2x3                                               = 0x8B65
	FLOAT_MAT2x4                                               = 0x8B66
	FLOAT_MAT3                                                 = 0x8B5B
	FLOAT_MAT3x2                                               = 0x8B67
	FLOAT_MAT3x4                                               = 0x8B68
	FLOAT_MAT4                                                 = 0x8B5C
	FLOAT_MAT4x2                                               = 0x8B69
	FLOAT_MAT4x3                                               = 0x8B6A
	FLOAT_VEC2                                                 = 0x8B50
	FLOAT_VEC3                                                 = 0x8B51
	FLOAT_VEC4                                                 = 0x8B52
	FRACTIONAL_EVEN                                            = 0x8E7C
	FRACTIONAL_ODD                                             = 0x8E7B
	FRAGMENT_INTERPOLATION_OFFSET_BITS                         = 0x8E5D
	FRAGMENT_SHADER                                            = 0x8B30
	FRAGMENT_SHADER_BIT                                        = 0x00000002
	FRAGMENT_SHADER_BIT_EXT                                    = 0x00000002
	FRAGMENT_SHADER_DERIVATIVE_HINT                            = 0x8B8B
	FRAGMENT_SHADER_INVOCATIONS_ARB                            = 0x82F4
	FRAGMENT_SUBROUTINE                                        = 0x92EC
	FRAGMENT_SUBROUTINE_UNIFORM                                = 0x92F2
	FRAGMENT_TEXTURE                                           = 0x829F
	FRAMEBUFFER                                                = 0x8D40
	FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE                          = 0x8215
	FRAMEBUFFER_ATTACHMENT_BLUE_SIZE                           = 0x8214
	FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING                      = 0x8210
	FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE                      = 0x8211
	FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE                          = 0x8216
	FRAMEBUFFER_ATTACHMENT_GREEN_SIZE                          = 0x8213
	FRAMEBUFFER_ATTACHMENT_LAYERED                             = 0x8DA7
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME                         = 0x8CD1
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE                         = 0x8CD0
	FRAMEBUFFER_ATTACHMENT_RED_SIZE                            = 0x8212
	FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE                        = 0x8217
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE               = 0x8CD3
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER                       = 0x8CD4
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL                       = 0x8CD2
	FRAMEBUFFER_BARRIER_BIT                                    = 0x00000400
	FRAMEBUFFER_BINDING                                        = 0x8CA6
	FRAMEBUFFER_BLEND                                          = 0x828B
	FRAMEBUFFER_COMPLETE                                       = 0x8CD5
	FRAMEBUFFER_DEFAULT                                        = 0x8218
	FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS                 = 0x9314
	FRAMEBUFFER_DEFAULT_HEIGHT                                 = 0x9311
	FRAMEBUFFER_DEFAULT_LAYERS                                 = 0x9312
	FRAMEBUFFER_DEFAULT_SAMPLES                                = 0x9313
	FRAMEBUFFER_DEFAULT_WIDTH                                  = 0x9310
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT                          = 0x8CD6
	FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER                         = 0x8CDB
	FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS                       = 0x8DA8
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT                  = 0x8CD7
	FRAMEBUFFER_INCOMPLETE_MULTISAMPLE                         = 0x8D56
	FRAMEBUFFER_INCOMPLETE_READ_BUFFER                         = 0x8CDC
	FRAMEBUFFER_RENDERABLE                                     = 0x8289
	FRAMEBUFFER_RENDERABLE_LAYERED                             = 0x828A
	FRAMEBUFFER_SRGB                                           = 0x8DB9
	FRAMEBUFFER_UNDEFINED                                      = 0x8219
	FRAMEBUFFER_UNSUPPORTED                                    = 0x8CDD
	FRONT                                                      = 0x0404
	FRONT_AND_BACK                                             = 0x0408
	FRONT_FACE                                                 = 0x0B46
	FRONT_LEFT                                                 = 0x0400
	FRONT_RIGHT                                                = 0x0401
	FULL_SUPPORT                                               = 0x82B7
	FUNC_ADD                                                   = 0x8006
	FUNC_ADD_EXT                                               = 0x8006
	FUNC_REVERSE_SUBTRACT                                      = 0x800B
	FUNC_SUBTRACT                                              = 0x800A
	GEOMETRY_INPUT_TYPE                                        = 0x8917
	GEOMETRY_OUTPUT_TYPE                                       = 0x8918
	GEOMETRY_SHADER                                            = 0x8DD9
	GEOMETRY_SHADER_BIT                                        = 0x00000004
	GEOMETRY_SHADER_INVOCATIONS                                = 0x887F
	GEOMETRY_SHADER_PRIMITIVES_EMITTED_ARB                     = 0x82F3
	GEOMETRY_SUBROUTINE                                        = 0x92EB
	GEOMETRY_SUBROUTINE_UNIFORM                                = 0x92F1
	GEOMETRY_TEXTURE                                           = 0x829E
	GEOMETRY_VERTICES_OUT                                      = 0x8916
	GEQUAL                                                     = 0x0206
	GET_TEXTURE_IMAGE_FORMAT                                   = 0x8291
	GET_TEXTURE_IMAGE_TYPE                                     = 0x8292
	GREATER                                                    = 0x0204
	GREEN                                                      = 0x1904
	GREEN_INTEGER                                              = 0x8D95
	GREEN_NV                                                   = 0x1904
	GUILTY_CONTEXT_RESET                                       = 0x8253
	GUILTY_CONTEXT_RESET_ARB                                   = 0x8253
	GUILTY_CONTEXT_RESET_KHR                                   = 0x8253
	HALF_FLOAT                                                 = 0x140B
	HARDLIGHT_KHR                                              = 0x929B
	HARDLIGHT_NV                                               = 0x929B
	HARDMIX_NV                                                 = 0x92A9
	HIGH_FLOAT                                                 = 0x8DF2
	HIGH_INT                                                   = 0x8DF5
	HSL_COLOR_KHR                                              = 0x92AF
	HSL_COLOR_NV                                               = 0x92AF
	HSL_HUE_KHR                                                = 0x92AD
	HSL_HUE_NV                                                 = 0x92AD
	HSL_LUMINOSITY_KHR                                         = 0x92B0
	HSL_LUMINOSITY_NV                                          = 0x92B0
	HSL_SATURATION_KHR                                         = 0x92AE
	HSL_SATURATION_NV                                          = 0x92AE
	IMAGE_1D                                                   = 0x904C
	IMAGE_1D_ARRAY                                             = 0x9052
	IMAGE_2D                                                   = 0x904D
	IMAGE_2D_ARRAY                                             = 0x9053
	IMAGE_2D_MULTISAMPLE                                       = 0x9055
	IMAGE_2D_MULTISAMPLE_ARRAY                                 = 0x9056
	IMAGE_2D_RECT                                              = 0x904F
	IMAGE_3D                                                   = 0x904E
	IMAGE_BINDING_ACCESS                                       = 0x8F3E
	IMAGE_BINDING_FORMAT                                       = 0x906E
	IMAGE_BINDING_LAYER                                        = 0x8F3D
	IMAGE_BINDING_LAYERED                                      = 0x8F3C
	IMAGE_BINDING_LEVEL                                        = 0x8F3B
	IMAGE_BINDING_NAME                                         = 0x8F3A
	IMAGE_BUFFER                                               = 0x9051
	IMAGE_CLASS_10_10_10_2                                     = 0x82C3
	IMAGE_CLASS_11_11_10                                       = 0x82C2
	IMAGE_CLASS_1_X_16                                         = 0x82BE
	IMAGE_CLASS_1_X_32                                         = 0x82BB
	IMAGE_CLASS_1_X_8                                          = 0x82C1
	IMAGE_CLASS_2_X_16                                         = 0x82BD
	IMAGE_CLASS_2_X_32                                         = 0x82BA
	IMAGE_CLASS_2_X_8                                          = 0x82C0
	IMAGE_CLASS_4_X_16                                         = 0x82BC
	IMAGE_CLASS_4_X_32                                         = 0x82B9
	IMAGE_CLASS_4_X_8                                          = 0x82BF
	IMAGE_COMPATIBILITY_CLASS                                  = 0x82A8
	IMAGE_CUBE                                                 = 0x9050
	IMAGE_CUBE_MAP_ARRAY                                       = 0x9054
	IMAGE_FORMAT_COMPATIBILITY_BY_CLASS                        = 0x90C9
	IMAGE_FORMAT_COMPATIBILITY_BY_SIZE                         = 0x90C8
	IMAGE_FORMAT_COMPATIBILITY_TYPE                            = 0x90C7
	IMAGE_PIXEL_FORMAT                                         = 0x82A9
	IMAGE_PIXEL_TYPE                                           = 0x82AA
	IMAGE_TEXEL_SIZE                                           = 0x82A7
	IMPLEMENTATION_COLOR_READ_FORMAT                           = 0x8B9B
	IMPLEMENTATION_COLOR_READ_FORMAT_OES                       = 0x8B9B
	IMPLEMENTATION_COLOR_READ_TYPE                             = 0x8B9A
	IMPLEMENTATION_COLOR_READ_TYPE_OES                         = 0x8B9A
	INCR                                                       = 0x1E02
	INCR_WRAP                                                  = 0x8507
	INFO_LOG_LENGTH                                            = 0x8B84
	INNOCENT_CONTEXT_RESET                                     = 0x8254
	INNOCENT_CONTEXT_RESET_ARB                                 = 0x8254
	INNOCENT_CONTEXT_RESET_KHR                                 = 0x8254
	INT                                                        = 0x1404
	INTERLEAVED_ATTRIBS                                        = 0x8C8C
	INTERNALFORMAT_ALPHA_SIZE                                  = 0x8274
	INTERNALFORMAT_ALPHA_TYPE                                  = 0x827B
	INTERNALFORMAT_BLUE_SIZE                                   = 0x8273
	INTERNALFORMAT_BLUE_TYPE                                   = 0x827A
	INTERNALFORMAT_DEPTH_SIZE                                  = 0x8275
	INTERNALFORMAT_DEPTH_TYPE                                  = 0x827C
	INTERNALFORMAT_GREEN_SIZE                                  = 0x8272
	INTERNALFORMAT_GREEN_TYPE                                  = 0x8279
	INTERNALFORMAT_PREFERRED                                   = 0x8270
	INTERNALFORMAT_RED_SIZE                                    = 0x8271
	INTERNALFORMAT_RED_TYPE                                    = 0x8278
	INTERNALFORMAT_SHARED_SIZE                                 = 0x8277
	INTERNALFORMAT_STENCIL_SIZE                                = 0x8276
	INTERNALFORMAT_STENCIL_TYPE                                = 0x827D
	INTERNALFORMAT_SUPPORTED                                   = 0x826F
	INT_2_10_10_10_REV                                         = 0x8D9F
	INT_IMAGE_1D                                               = 0x9057
	INT_IMAGE_1D_ARRAY                                         = 0x905D
	INT_IMAGE_2D                                               = 0x9058
	INT_IMAGE_2D_ARRAY                                         = 0x905E
	INT_IMAGE_2D_MULTISAMPLE                                   = 0x9060
	INT_IMAGE_2D_MULTISAMPLE_ARRAY                             = 0x9061
	INT_IMAGE_2D_RECT                                          = 0x905A
	INT_IMAGE_3D                                               = 0x9059
	INT_IMAGE_BUFFER                                           = 0x905C
	INT_IMAGE_CUBE                                             = 0x905B
	INT_IMAGE_CUBE_MAP_ARRAY                                   = 0x905F
	INT_SAMPLER_1D                                             = 0x8DC9
	INT_SAMPLER_1D_ARRAY                                       = 0x8DCE
	INT_SAMPLER_2D                                             = 0x8DCA
	INT_SAMPLER_2D_ARRAY                                       = 0x8DCF
	INT_SAMPLER_2D_MULTISAMPLE                                 = 0x9109
	INT_SAMPLER_2D_MULTISAMPLE_ARRAY                           = 0x910C
	INT_SAMPLER_2D_RECT                                        = 0x8DCD
	INT_SAMPLER_3D                                             = 0x8DCB
	INT_SAMPLER_BUFFER                                         = 0x8DD0
	INT_SAMPLER_CUBE                                           = 0x8DCC
	INT_SAMPLER_CUBE_MAP_ARRAY                                 = 0x900E
	INT_SAMPLER_CUBE_MAP_ARRAY_ARB                             = 0x900E
	INT_VEC2                                                   = 0x8B53
	INT_VEC3                                                   = 0x8B54
	INT_VEC4                                                   = 0x8B55
	INVALID_ENUM                                               = 0x0500
	INVALID_FRAMEBUFFER_OPERATION                              = 0x0506
	INVALID_INDEX                                              = 0xFFFFFFFF
	INVALID_OPERATION                                          = 0x0502
	INVALID_VALUE                                              = 0x0501
	INVERT                                                     = 0x150A
	INVERT_OVG_NV                                              = 0x92B4
	INVERT_RGB_NV                                              = 0x92A3
	ISOLINES                                                   = 0x8E7A
	IS_PER_PATCH                                               = 0x92E7
	IS_ROW_MAJOR                                               = 0x9300
	KEEP                                                       = 0x1E00
	LAST_VERTEX_CONVENTION                                     = 0x8E4E
	LAYER_PROVOKING_VERTEX                                     = 0x825E
	LEFT                                                       = 0x0406
	LEQUAL                                                     = 0x0203
	LESS                                                       = 0x0201
	LIGHTEN_KHR                                                = 0x9298
	LIGHTEN_NV                                                 = 0x9298
	LINE                                                       = 0x1B01
	LINEAR                                                     = 0x2601
	LINEARBURN_NV                                              = 0x92A5
	LINEARDODGE_NV                                             = 0x92A4
	LINEARLIGHT_NV                                             = 0x92A7
	LINEAR_MIPMAP_LINEAR                                       = 0x2703
	LINEAR_MIPMAP_NEAREST                                      = 0x2701
	LINES                                                      = 0x0001
	LINES_ADJACENCY                                            = 0x000A
	LINE_LOOP                                                  = 0x0002
	LINE_SMOOTH                                                = 0x0B20
	LINE_SMOOTH_HINT                                           = 0x0C52
	LINE_STRIP                                                 = 0x0003
	LINE_STRIP_ADJACENCY                                       = 0x000B
	LINE_WIDTH                                                 = 0x0B21
	LINE_WIDTH_GRANULARITY                                     = 0x0B23
	LINE_WIDTH_RANGE                                           = 0x0B22
	LINK_STATUS                                                = 0x8B82
	LOCATION                                                   = 0x930E
	LOCATION_COMPONENT                                         = 0x934A
	LOCATION_INDEX                                             = 0x930F
	LOGIC_OP_MODE                                              = 0x0BF0
	LOSE_CONTEXT_ON_RESET                                      = 0x8252
	LOSE_CONTEXT_ON_RESET_ARB                                  = 0x8252
	LOSE_CONTEXT_ON_RESET_KHR                                  = 0x8252
	LOWER_LEFT                                                 = 0x8CA1
	LOW_FLOAT                                                  = 0x8DF0
	LOW_INT                                                    = 0x8DF3
	MAJOR_VERSION                                              = 0x821B
	MANUAL_GENERATE_MIPMAP                                     = 0x8294
	MAP_COHERENT_BIT                                           = 0x0080
	MAP_FLUSH_EXPLICIT_BIT                                     = 0x0010
	MAP_INVALIDATE_BUFFER_BIT                                  = 0x0008
	MAP_INVALIDATE_RANGE_BIT                                   = 0x0004
	MAP_PERSISTENT_BIT                                         = 0x0040
	MAP_READ_BIT                                               = 0x0001
	MAP_UNSYNCHRONIZED_BIT                                     = 0x0020
	MAP_WRITE_BIT                                              = 0x0002
	MATRIX_STRIDE                                              = 0x92FF
	MAX                                                        = 0x8008
	MAX_3D_TEXTURE_SIZE                                        = 0x8073
	MAX_ARRAY_TEXTURE_LAYERS                                   = 0x88FF
	MAX_ATOMIC_COUNTER_BUFFER_BINDINGS                         = 0x92DC
	MAX_ATOMIC_COUNTER_BUFFER_SIZE                             = 0x92D8
	MAX_CLIP_DISTANCES                                         = 0x0D32
	MAX_COLOR_ATTACHMENTS                                      = 0x8CDF
	MAX_COLOR_TEXTURE_SAMPLES                                  = 0x910E
	MAX_COMBINED_ATOMIC_COUNTERS                               = 0x92D7
	MAX_COMBINED_ATOMIC_COUNTER_BUFFERS                        = 0x92D1
	MAX_COMBINED_CLIP_AND_CULL_DISTANCES                       = 0x82FA
	MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS                    = 0x8266
	MAX_COMBINED_DIMENSIONS                                    = 0x8282
	MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS                   = 0x8A33
	MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS                   = 0x8A32
	MAX_COMBINED_IMAGE_UNIFORMS                                = 0x90CF
	MAX_COMBINED_IMAGE_UNITS_AND_FRAGMENT_OUTPUTS              = 0x8F39
	MAX_COMBINED_SHADER_OUTPUT_RESOURCES                       = 0x8F39
	MAX_COMBINED_SHADER_STORAGE_BLOCKS                         = 0x90DC
	MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS               = 0x8E1E
	MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS            = 0x8E1F
	MAX_COMBINED_TEXTURE_IMAGE_UNITS                           = 0x8B4D
	MAX_COMBINED_UNIFORM_BLOCKS                                = 0x8A2E
	MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS                     = 0x8A31
	MAX_COMPUTE_ATOMIC_COUNTERS                                = 0x8265
	MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS                         = 0x8264
	MAX_COMPUTE_FIXED_GROUP_INVOCATIONS_ARB                    = 0x90EB
	MAX_COMPUTE_FIXED_GROUP_SIZE_ARB                           = 0x91BF
	MAX_COMPUTE_IMAGE_UNIFORMS                                 = 0x91BD
	MAX_COMPUTE_SHADER_STORAGE_BLOCKS                          = 0x90DB
	MAX_COMPUTE_SHARED_MEMORY_SIZE                             = 0x8262
	MAX_COMPUTE_TEXTURE_IMAGE_UNITS                            = 0x91BC
	MAX_COMPUTE_UNIFORM_BLOCKS                                 = 0x91BB
	MAX_COMPUTE_UNIFORM_COMPONENTS                             = 0x8263
	MAX_COMPUTE_VARIABLE_GROUP_INVOCATIONS_ARB                 = 0x9344
	MAX_COMPUTE_VARIABLE_GROUP_SIZE_ARB                        = 0x9345
	MAX_COMPUTE_WORK_GROUP_COUNT                               = 0x91BE
	MAX_COMPUTE_WORK_GROUP_INVOCATIONS                         = 0x90EB
	MAX_COMPUTE_WORK_GROUP_SIZE                                = 0x91BF
	MAX_CUBE_MAP_TEXTURE_SIZE                                  = 0x851C
	MAX_CULL_DISTANCES                                         = 0x82F9
	MAX_DEBUG_GROUP_STACK_DEPTH                                = 0x826C
	MAX_DEBUG_GROUP_STACK_DEPTH_KHR                            = 0x826C
	MAX_DEBUG_LOGGED_MESSAGES                                  = 0x9144
	MAX_DEBUG_LOGGED_MESSAGES_ARB                              = 0x9144
	MAX_DEBUG_LOGGED_MESSAGES_KHR                              = 0x9144
	MAX_DEBUG_MESSAGE_LENGTH                                   = 0x9143
	MAX_DEBUG_MESSAGE_LENGTH_ARB                               = 0x9143
	MAX_DEBUG_MESSAGE_LENGTH_KHR                               = 0x9143
	MAX_DEPTH                                                  = 0x8280
	MAX_DEPTH_TEXTURE_SAMPLES                                  = 0x910F
	MAX_DRAW_BUFFERS                                           = 0x8824
	MAX_DUAL_SOURCE_DRAW_BUFFERS                               = 0x88FC
	MAX_ELEMENTS_INDICES                                       = 0x80E9
	MAX_ELEMENTS_VERTICES                                      = 0x80E8
	MAX_ELEMENT_INDEX                                          = 0x8D6B
	MAX_EXT                                                    = 0x8008
	MAX_FRAGMENT_ATOMIC_COUNTERS                               = 0x92D6
	MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS                        = 0x92D0
	MAX_FRAGMENT_IMAGE_UNIFORMS                                = 0x90CE
	MAX_FRAGMENT_INPUT_COMPONENTS                              = 0x9125
	MAX_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5C
	MAX_FRAGMENT_SHADER_STORAGE_BLOCKS                         = 0x90DA
	MAX_FRAGMENT_UNIFORM_BLOCKS                                = 0x8A2D
	MAX_FRAGMENT_UNIFORM_COMPONENTS                            = 0x8B49
	MAX_FRAGMENT_UNIFORM_VECTORS                               = 0x8DFD
	MAX_FRAMEBUFFER_HEIGHT                                     = 0x9316
	MAX_FRAMEBUFFER_LAYERS                                     = 0x9317
	MAX_FRAMEBUFFER_SAMPLES                                    = 0x9318
	MAX_FRAMEBUFFER_WIDTH                                      = 0x9315
	MAX_GEOMETRY_ATOMIC_COUNTERS                               = 0x92D5
	MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS                        = 0x92CF
	MAX_GEOMETRY_IMAGE_UNIFORMS                                = 0x90CD
	MAX_GEOMETRY_INPUT_COMPONENTS                              = 0x9123
	MAX_GEOMETRY_OUTPUT_COMPONENTS                             = 0x9124
	MAX_GEOMETRY_OUTPUT_VERTICES                               = 0x8DE0
	MAX_GEOMETRY_SHADER_INVOCATIONS                            = 0x8E5A
	MAX_GEOMETRY_SHADER_STORAGE_BLOCKS                         = 0x90D7
	MAX_GEOMETRY_TEXTURE_IMAGE_UNITS                           = 0x8C29
	MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS                       = 0x8DE1
	MAX_GEOMETRY_UNIFORM_BLOCKS                                = 0x8A2C
	MAX_GEOMETRY_UNIFORM_COMPONENTS                            = 0x8DDF
	MAX_HEIGHT                                                 = 0x827F
	MAX_IMAGE_SAMPLES                                          = 0x906D
	MAX_IMAGE_UNITS                                            = 0x8F38
	MAX_INTEGER_SAMPLES                                        = 0x9110
	MAX_LABEL_LENGTH                                           = 0x82E8
	MAX_LABEL_LENGTH_KHR                                       = 0x82E8
	MAX_LAYERS                                                 = 0x8281
	MAX_NAME_LENGTH                                            = 0x92F6
	MAX_NUM_ACTIVE_VARIABLES                                   = 0x92F7
	MAX_NUM_COMPATIBLE_SUBROUTINES                             = 0x92F8
	MAX_PATCH_VERTICES                                         = 0x8E7D
	MAX_PROGRAM_TEXEL_OFFSET                                   = 0x8905
	MAX_PROGRAM_TEXTURE_GATHER_COMPONENTS_ARB                  = 0x8F9F
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5F
	MAX_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5F
	MAX_RECTANGLE_TEXTURE_SIZE                                 = 0x84F8
	MAX_RENDERBUFFER_SIZE                                      = 0x84E8
	MAX_SAMPLES                                                = 0x8D57
	MAX_SAMPLE_MASK_WORDS                                      = 0x8E59
	MAX_SERVER_WAIT_TIMEOUT                                    = 0x9111
	MAX_SHADER_STORAGE_BLOCK_SIZE                              = 0x90DE
	MAX_SHADER_STORAGE_BUFFER_BINDINGS                         = 0x90DD
	MAX_SPARSE_3D_TEXTURE_SIZE_ARB                             = 0x9199
	MAX_SPARSE_ARRAY_TEXTURE_LAYERS_ARB                        = 0x919A
	MAX_SPARSE_TEXTURE_SIZE_ARB                                = 0x9198
	MAX_SUBROUTINES                                            = 0x8DE7
	MAX_SUBROUTINE_UNIFORM_LOCATIONS                           = 0x8DE8
	MAX_TESS_CONTROL_ATOMIC_COUNTERS                           = 0x92D3
	MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS                    = 0x92CD
	MAX_TESS_CONTROL_IMAGE_UNIFORMS                            = 0x90CB
	MAX_TESS_CONTROL_INPUT_COMPONENTS                          = 0x886C
	MAX_TESS_CONTROL_OUTPUT_COMPONENTS                         = 0x8E83
	MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS                     = 0x90D8
	MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS                       = 0x8E81
	MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS                   = 0x8E85
	MAX_TESS_CONTROL_UNIFORM_BLOCKS                            = 0x8E89
	MAX_TESS_CONTROL_UNIFORM_COMPONENTS                        = 0x8E7F
	MAX_TESS_EVALUATION_ATOMIC_COUNTERS                        = 0x92D4
	MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS                 = 0x92CE
	MAX_TESS_EVALUATION_IMAGE_UNIFORMS                         = 0x90CC
	MAX_TESS_EVALUATION_INPUT_COMPONENTS                       = 0x886D
	MAX_TESS_EVALUATION_OUTPUT_COMPONENTS                      = 0x8E86
	MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS                  = 0x90D9
	MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS                    = 0x8E82
	MAX_TESS_EVALUATION_UNIFORM_BLOCKS                         = 0x8E8A
	MAX_TESS_EVALUATION_UNIFORM_COMPONENTS                     = 0x8E80
	MAX_TESS_GEN_LEVEL                                         = 0x8E7E
	MAX_TESS_PATCH_COMPONENTS                                  = 0x8E84
	MAX_TEXTURE_BUFFER_SIZE                                    = 0x8C2B
	MAX_TEXTURE_IMAGE_UNITS                                    = 0x8872
	MAX_TEXTURE_LOD_BIAS                                       = 0x84FD
	MAX_TEXTURE_LOD_BIAS_EXT                                   = 0x84FD
	MAX_TEXTURE_MAX_ANISOTROPY_EXT                             = 0x84FF
	MAX_TEXTURE_SIZE                                           = 0x0D33
	MAX_TRANSFORM_FEEDBACK_BUFFERS                             = 0x8E70
	MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS              = 0x8C8A
	MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS                    = 0x8C8B
	MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS                 = 0x8C80
	MAX_UNIFORM_BLOCK_SIZE                                     = 0x8A30
	MAX_UNIFORM_BUFFER_BINDINGS                                = 0x8A2F
	MAX_UNIFORM_LOCATIONS                                      = 0x826E
	MAX_VARYING_COMPONENTS                                     = 0x8B4B
	MAX_VARYING_FLOATS                                         = 0x8B4B
	MAX_VARYING_VECTORS                                        = 0x8DFC
	MAX_VERTEX_ATOMIC_COUNTERS                                 = 0x92D2
	MAX_VERTEX_ATOMIC_COUNTER_BUFFERS                          = 0x92CC
	MAX_VERTEX_ATTRIBS                                         = 0x8869
	MAX_VERTEX_ATTRIB_BINDINGS                                 = 0x82DA
	MAX_VERTEX_ATTRIB_RELATIVE_OFFSET                          = 0x82D9
	MAX_VERTEX_IMAGE_UNIFORMS                                  = 0x90CA
	MAX_VERTEX_OUTPUT_COMPONENTS                               = 0x9122
	MAX_VERTEX_SHADER_STORAGE_BLOCKS                           = 0x90D6
	MAX_VERTEX_STREAMS                                         = 0x8E71
	MAX_VERTEX_TEXTURE_IMAGE_UNITS                             = 0x8B4C
	MAX_VERTEX_UNIFORM_BLOCKS                                  = 0x8A2B
	MAX_VERTEX_UNIFORM_COMPONENTS                              = 0x8B4A
	MAX_VERTEX_UNIFORM_VECTORS                                 = 0x8DFB
	MAX_VIEWPORTS                                              = 0x825B
	MAX_VIEWPORT_DIMS                                          = 0x0D3A
	MAX_WIDTH                                                  = 0x827E
	MEDIUM_FLOAT                                               = 0x8DF1
	MEDIUM_INT                                                 = 0x8DF4
	MIN                                                        = 0x8007
	MINOR_VERSION                                              = 0x821C
	MINUS_CLAMPED_NV                                           = 0x92B3
	MINUS_NV                                                   = 0x929F
	MIN_EXT                                                    = 0x8007
	MIN_FRAGMENT_INTERPOLATION_OFFSET                          = 0x8E5B
	MIN_MAP_BUFFER_ALIGNMENT                                   = 0x90BC
	MIN_PROGRAM_TEXEL_OFFSET                                   = 0x8904
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET                          = 0x8E5E
	MIN_PROGRAM_TEXTURE_GATHER_OFFSET_ARB                      = 0x8E5E
	MIN_SAMPLE_SHADING_VALUE                                   = 0x8C37
	MIN_SAMPLE_SHADING_VALUE_ARB                               = 0x8C37
	MIN_SPARSE_LEVEL_ARB                                       = 0x919B
	MIPMAP                                                     = 0x8293
	MIRRORED_REPEAT                                            = 0x8370
	MIRROR_CLAMP_TO_EDGE                                       = 0x8743
	MULTIPLY_KHR                                               = 0x9294
	MULTIPLY_NV                                                = 0x9294
	MULTISAMPLE                                                = 0x809D
	NAMED_STRING_LENGTH_ARB                                    = 0x8DE9
	NAMED_STRING_TYPE_ARB                                      = 0x8DEA
	NAME_LENGTH                                                = 0x92F9
	NAND                                                       = 0x150E
	NEAREST                                                    = 0x2600
	NEAREST_MIPMAP_LINEAR                                      = 0x2702
	NEAREST_MIPMAP_NEAREST                                     = 0x2700
	NEGATIVE_ONE_TO_ONE                                        = 0x935E
	NEVER                                                      = 0x0200
	NICEST                                                     = 0x1102
	NONE                                                       = 0
	NOOP                                                       = 0x1505
	NOR                                                        = 0x1508
	NOTEQUAL                                                   = 0x0205
	NO_ERROR                                                   = 0
	NO_RESET_NOTIFICATION                                      = 0x8261
	NO_RESET_NOTIFICATION_ARB                                  = 0x8261
	NO_RESET_NOTIFICATION_KHR                                  = 0x8261
	NUM_ACTIVE_VARIABLES                                       = 0x9304
	NUM_COMPATIBLE_SUBROUTINES                                 = 0x8E4A
	NUM_COMPRESSED_TEXTURE_FORMATS                             = 0x86A2
	NUM_EXTENSIONS                                             = 0x821D
	NUM_PROGRAM_BINARY_FORMATS                                 = 0x87FE
	NUM_SAMPLE_COUNTS                                          = 0x9380
	NUM_SHADER_BINARY_FORMATS                                  = 0x8DF9
	NUM_VIRTUAL_PAGE_SIZES_ARB                                 = 0x91A8
	OBJECT_TYPE                                                = 0x9112
	OFFSET                                                     = 0x92FC
	ONE                                                        = 1
	ONE_MINUS_CONSTANT_ALPHA                                   = 0x8004
	ONE_MINUS_CONSTANT_COLOR                                   = 0x8002
	ONE_MINUS_DST_ALPHA                                        = 0x0305
	ONE_MINUS_DST_COLOR                                        = 0x0307
	ONE_MINUS_SRC1_ALPHA                                       = 0x88FB
	ONE_MINUS_SRC1_COLOR                                       = 0x88FA
	ONE_MINUS_SRC_ALPHA                                        = 0x0303
	ONE_MINUS_SRC_COLOR                                        = 0x0301
	OR                                                         = 0x1507
	OR_INVERTED                                                = 0x150D
	OR_REVERSE                                                 = 0x150B
	OUT_OF_MEMORY                                              = 0x0505
	OVERLAY_KHR                                                = 0x9296
	OVERLAY_NV                                                 = 0x9296
	PACK_ALIGNMENT                                             = 0x0D05
	PACK_COMPRESSED_BLOCK_DEPTH                                = 0x912D
	PACK_COMPRESSED_BLOCK_HEIGHT                               = 0x912C
	PACK_COMPRESSED_BLOCK_SIZE                                 = 0x912E
	PACK_COMPRESSED_BLOCK_WIDTH                                = 0x912B
	PACK_IMAGE_HEIGHT                                          = 0x806C
	PACK_LSB_FIRST                                             = 0x0D01
	PACK_ROW_LENGTH                                            = 0x0D02
	PACK_SKIP_IMAGES                                           = 0x806B
	PACK_SKIP_PIXELS                                           = 0x0D04
	PACK_SKIP_ROWS                                             = 0x0D03
	PACK_SWAP_BYTES                                            = 0x0D00
	PALETTE4_R5_G6_B5_OES                                      = 0x8B92
	PALETTE4_RGB5_A1_OES                                       = 0x8B94
	PALETTE4_RGB8_OES                                          = 0x8B90
	PALETTE4_RGBA4_OES                                         = 0x8B93
	PALETTE4_RGBA8_OES                                         = 0x8B91
	PALETTE8_R5_G6_B5_OES                                      = 0x8B97
	PALETTE8_RGB5_A1_OES                                       = 0x8B99
	PALETTE8_RGB8_OES                                          = 0x8B95
	PALETTE8_RGBA4_OES                                         = 0x8B98
	PALETTE8_RGBA8_OES                                         = 0x8B96
	PARAMETER_BUFFER_ARB                                       = 0x80EE
	PARAMETER_BUFFER_BINDING_ARB                               = 0x80EF
	PATCHES                                                    = 0x000E
	PATCH_DEFAULT_INNER_LEVEL                                  = 0x8E73
	PATCH_DEFAULT_OUTER_LEVEL                                  = 0x8E74
	PATCH_VERTICES                                             = 0x8E72
	PERCENTAGE_AMD                                             = 0x8BC3
	PERFMON_RESULT_AMD                                         = 0x8BC6
	PERFMON_RESULT_AVAILABLE_AMD                               = 0x8BC4
	PERFMON_RESULT_SIZE_AMD                                    = 0x8BC5
	PERFQUERY_COUNTER_DATA_BOOL32_INTEL                        = 0x94FC
	PERFQUERY_COUNTER_DATA_DOUBLE_INTEL                        = 0x94FB
	PERFQUERY_COUNTER_DATA_FLOAT_INTEL                         = 0x94FA
	PERFQUERY_COUNTER_DATA_UINT32_INTEL                        = 0x94F8
	PERFQUERY_COUNTER_DATA_UINT64_INTEL                        = 0x94F9
	PERFQUERY_COUNTER_DESC_LENGTH_MAX_INTEL                    = 0x94FF
	PERFQUERY_COUNTER_DURATION_NORM_INTEL                      = 0x94F1
	PERFQUERY_COUNTER_DURATION_RAW_INTEL                       = 0x94F2
	PERFQUERY_COUNTER_EVENT_INTEL                              = 0x94F0
	PERFQUERY_COUNTER_NAME_LENGTH_MAX_INTEL                    = 0x94FE
	PERFQUERY_COUNTER_RAW_INTEL                                = 0x94F4
	PERFQUERY_COUNTER_THROUGHPUT_INTEL                         = 0x94F3
	PERFQUERY_COUNTER_TIMESTAMP_INTEL                          = 0x94F5
	PERFQUERY_DONOT_FLUSH_INTEL                                = 0x83F9
	PERFQUERY_FLUSH_INTEL                                      = 0x83FA
	PERFQUERY_GLOBAL_CONTEXT_INTEL                             = 0x00000001
	PERFQUERY_GPA_EXTENDED_COUNTERS_INTEL                      = 0x9500
	PERFQUERY_QUERY_NAME_LENGTH_MAX_INTEL                      = 0x94FD
	PERFQUERY_SINGLE_CONTEXT_INTEL                             = 0x00000000
	PERFQUERY_WAIT_INTEL                                       = 0x83FB
	PINLIGHT_NV                                                = 0x92A8
	PIXEL_BUFFER_BARRIER_BIT                                   = 0x00000080
	PIXEL_PACK_BUFFER                                          = 0x88EB
	PIXEL_PACK_BUFFER_BINDING                                  = 0x88ED
	PIXEL_UNPACK_BUFFER                                        = 0x88EC
	PIXEL_UNPACK_BUFFER_BINDING                                = 0x88EF
	PLUS_CLAMPED_ALPHA_NV                                      = 0x92B2
	PLUS_CLAMPED_NV                                            = 0x92B1
	PLUS_DARKER_NV                                             = 0x9292
	PLUS_NV                                                    = 0x9291
	POINT                                                      = 0x1B00
	POINTS                                                     = 0x0000
	POINT_FADE_THRESHOLD_SIZE                                  = 0x8128
	POINT_SIZE                                                 = 0x0B11
	POINT_SIZE_GRANULARITY                                     = 0x0B13
	POINT_SIZE_RANGE                                           = 0x0B12
	POINT_SPRITE_COORD_ORIGIN                                  = 0x8CA0
	POLYGON_MODE                                               = 0x0B40
	POLYGON_OFFSET_FACTOR                                      = 0x8038
	POLYGON_OFFSET_FILL                                        = 0x8037
	POLYGON_OFFSET_LINE                                        = 0x2A02
	POLYGON_OFFSET_POINT                                       = 0x2A01
	POLYGON_OFFSET_UNITS                                       = 0x2A00
	POLYGON_SMOOTH                                             = 0x0B41
	POLYGON_SMOOTH_HINT                                        = 0x0C53
	PRIMITIVES_GENERATED                                       = 0x8C87
	PRIMITIVES_SUBMITTED_ARB                                   = 0x82EF
	PRIMITIVE_RESTART                                          = 0x8F9D
	PRIMITIVE_RESTART_FIXED_INDEX                              = 0x8D69
	PRIMITIVE_RESTART_INDEX                                    = 0x8F9E
	PROGRAM                                                    = 0x82E2
	PROGRAM_BINARY_FORMATS                                     = 0x87FF
	PROGRAM_BINARY_LENGTH                                      = 0x8741
	PROGRAM_BINARY_RETRIEVABLE_HINT                            = 0x8257
	PROGRAM_INPUT                                              = 0x92E3
	PROGRAM_KHR                                                = 0x82E2
	PROGRAM_OBJECT_EXT                                         = 0x8B40
	PROGRAM_OUTPUT                                             = 0x92E4
	PROGRAM_PIPELINE                                           = 0x82E4
	PROGRAM_PIPELINE_BINDING                                   = 0x825A
	PROGRAM_PIPELINE_BINDING_EXT                               = 0x825A
	PROGRAM_PIPELINE_OBJECT_EXT                                = 0x8A4F
	PROGRAM_POINT_SIZE                                         = 0x8642
	PROGRAM_SEPARABLE                                          = 0x8258
	PROGRAM_SEPARABLE_EXT                                      = 0x8258
	PROVOKING_VERTEX                                           = 0x8E4F
	PROXY_TEXTURE_1D                                           = 0x8063
	PROXY_TEXTURE_1D_ARRAY                                     = 0x8C19
	PROXY_TEXTURE_2D                                           = 0x8064
	PROXY_TEXTURE_2D_ARRAY                                     = 0x8C1B
	PROXY_TEXTURE_2D_MULTISAMPLE                               = 0x9101
	PROXY_TEXTURE_2D_MULTISAMPLE_ARRAY                         = 0x9103
	PROXY_TEXTURE_3D                                           = 0x8070
	PROXY_TEXTURE_CUBE_MAP                                     = 0x851B
	PROXY_TEXTURE_CUBE_MAP_ARRAY                               = 0x900B
	PROXY_TEXTURE_CUBE_MAP_ARRAY_ARB                           = 0x900B
	PROXY_TEXTURE_RECTANGLE                                    = 0x84F7
	QUADS                                                      = 0x0007
	QUADS_FOLLOW_PROVOKING_VERTEX_CONVENTION                   = 0x8E4C
	QUERY                                                      = 0x82E3
	QUERY_BUFFER                                               = 0x9192
	QUERY_BUFFER_BARRIER_BIT                                   = 0x00008000
	QUERY_BUFFER_BINDING                                       = 0x9193
	QUERY_BY_REGION_NO_WAIT                                    = 0x8E16
	QUERY_BY_REGION_NO_WAIT_INVERTED                           = 0x8E1A
	QUERY_BY_REGION_WAIT                                       = 0x8E15
	QUERY_BY_REGION_WAIT_INVERTED                              = 0x8E19
	QUERY_COUNTER_BITS                                         = 0x8864
	QUERY_KHR                                                  = 0x82E3
	QUERY_NO_WAIT                                              = 0x8E14
	QUERY_NO_WAIT_INVERTED                                     = 0x8E18
	QUERY_OBJECT_EXT                                           = 0x9153
	QUERY_RESULT                                               = 0x8866
	QUERY_RESULT_AVAILABLE                                     = 0x8867
	QUERY_RESULT_NO_WAIT                                       = 0x9194
	QUERY_TARGET                                               = 0x82EA
	QUERY_WAIT                                                 = 0x8E13
	QUERY_WAIT_INVERTED                                        = 0x8E17
	R11F_G11F_B10F                                             = 0x8C3A
	R16                                                        = 0x822A
	R16F                                                       = 0x822D
	R16I                                                       = 0x8233
	R16UI                                                      = 0x8234
	R16_SNORM                                                  = 0x8F98
	R32F                                                       = 0x822E
	R32I                                                       = 0x8235
	R32UI                                                      = 0x8236
	R3_G3_B2                                                   = 0x2A10
	R8                                                         = 0x8229
	R8I                                                        = 0x8231
	R8UI                                                       = 0x8232
	R8_SNORM                                                   = 0x8F94
	RASTERIZER_DISCARD                                         = 0x8C89
	READ_BUFFER                                                = 0x0C02
	READ_FRAMEBUFFER                                           = 0x8CA8
	READ_FRAMEBUFFER_BINDING                                   = 0x8CAA
	READ_ONLY                                                  = 0x88B8
	READ_PIXELS                                                = 0x828C
	READ_PIXELS_FORMAT                                         = 0x828D
	READ_PIXELS_TYPE                                           = 0x828E
	READ_WRITE                                                 = 0x88BA
	RED                                                        = 0x1903
	RED_INTEGER                                                = 0x8D94
	RED_NV                                                     = 0x1903
	REFERENCED_BY_COMPUTE_SHADER                               = 0x930B
	REFERENCED_BY_FRAGMENT_SHADER                              = 0x930A
	REFERENCED_BY_GEOMETRY_SHADER                              = 0x9309
	REFERENCED_BY_TESS_CONTROL_SHADER                          = 0x9307
	REFERENCED_BY_TESS_EVALUATION_SHADER                       = 0x9308
	REFERENCED_BY_VERTEX_SHADER                                = 0x9306
	RENDERBUFFER                                               = 0x8D41
	RENDERBUFFER_ALPHA_SIZE                                    = 0x8D53
	RENDERBUFFER_BINDING                                       = 0x8CA7
	RENDERBUFFER_BLUE_SIZE                                     = 0x8D52
	RENDERBUFFER_DEPTH_SIZE                                    = 0x8D54
	RENDERBUFFER_GREEN_SIZE                                    = 0x8D51
	RENDERBUFFER_HEIGHT                                        = 0x8D43
	RENDERBUFFER_INTERNAL_FORMAT                               = 0x8D44
	RENDERBUFFER_RED_SIZE                                      = 0x8D50
	RENDERBUFFER_SAMPLES                                       = 0x8CAB
	RENDERBUFFER_STENCIL_SIZE                                  = 0x8D55
	RENDERBUFFER_WIDTH                                         = 0x8D42
	RENDERER                                                   = 0x1F01
	REPEAT                                                     = 0x2901
	REPLACE                                                    = 0x1E01
	RESET_NOTIFICATION_STRATEGY                                = 0x8256
	RESET_NOTIFICATION_STRATEGY_ARB                            = 0x8256
	RESET_NOTIFICATION_STRATEGY_KHR                            = 0x8256
	RG                                                         = 0x8227
	RG16                                                       = 0x822C
	RG16F                                                      = 0x822F
	RG16I                                                      = 0x8239
	RG16UI                                                     = 0x823A
	RG16_SNORM                                                 = 0x8F99
	RG32F                                                      = 0x8230
	RG32I                                                      = 0x823B
	RG32UI                                                     = 0x823C
	RG8                                                        = 0x822B
	RG8I                                                       = 0x8237
	RG8UI                                                      = 0x8238
	RG8_SNORM                                                  = 0x8F95
	RGB                                                        = 0x1907
	RGB10                                                      = 0x8052
	RGB10_A2                                                   = 0x8059
	RGB10_A2UI                                                 = 0x906F
	RGB12                                                      = 0x8053
	RGB16                                                      = 0x8054
	RGB16F                                                     = 0x881B
	RGB16I                                                     = 0x8D89
	RGB16UI                                                    = 0x8D77
	RGB16_SNORM                                                = 0x8F9A
	RGB32F                                                     = 0x8815
	RGB32I                                                     = 0x8D83
	RGB32UI                                                    = 0x8D71
	RGB4                                                       = 0x804F
	RGB5                                                       = 0x8050
	RGB565                                                     = 0x8D62
	RGB5_A1                                                    = 0x8057
	RGB8                                                       = 0x8051
	RGB8I                                                      = 0x8D8F
	RGB8UI                                                     = 0x8D7D
	RGB8_SNORM                                                 = 0x8F96
	RGB9_E5                                                    = 0x8C3D
	RGBA                                                       = 0x1908
	RGBA12                                                     = 0x805A
	RGBA16                                                     = 0x805B
	RGBA16F                                                    = 0x881A
	RGBA16I                                                    = 0x8D88
	RGBA16UI                                                   = 0x8D76
	RGBA16_SNORM                                               = 0x8F9B
	RGBA2                                                      = 0x8055
	RGBA32F                                                    = 0x8814
	RGBA32I                                                    = 0x8D82
	RGBA32UI                                                   = 0x8D70
	RGBA4                                                      = 0x8056
	RGBA8                                                      = 0x8058
	RGBA8I                                                     = 0x8D8E
	RGBA8UI                                                    = 0x8D7C
	RGBA8_SNORM                                                = 0x8F97
	RGBA_INTEGER                                               = 0x8D99
	RGB_422_APPLE                                              = 0x8A1F
	RGB_INTEGER                                                = 0x8D98
	RGB_RAW_422_APPLE                                          = 0x8A51
	RG_INTEGER                                                 = 0x8228
	RIGHT                                                      = 0x0407
	SAMPLER                                                    = 0x82E6
	SAMPLER_1D                                                 = 0x8B5D
	SAMPLER_1D_ARRAY                                           = 0x8DC0
	SAMPLER_1D_ARRAY_SHADOW                                    = 0x8DC3
	SAMPLER_1D_SHADOW                                          = 0x8B61
	SAMPLER_2D                                                 = 0x8B5E
	SAMPLER_2D_ARRAY                                           = 0x8DC1
	SAMPLER_2D_ARRAY_SHADOW                                    = 0x8DC4
	SAMPLER_2D_MULTISAMPLE                                     = 0x9108
	SAMPLER_2D_MULTISAMPLE_ARRAY                               = 0x910B
	SAMPLER_2D_RECT                                            = 0x8B63
	SAMPLER_2D_RECT_SHADOW                                     = 0x8B64
	SAMPLER_2D_SHADOW                                          = 0x8B62
	SAMPLER_3D                                                 = 0x8B5F
	SAMPLER_BINDING                                            = 0x8919
	SAMPLER_BUFFER                                             = 0x8DC2
	SAMPLER_CUBE                                               = 0x8B60
	SAMPLER_CUBE_MAP_ARRAY                                     = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_ARB                                 = 0x900C
	SAMPLER_CUBE_MAP_ARRAY_SHADOW                              = 0x900D
	SAMPLER_CUBE_MAP_ARRAY_SHADOW_ARB                          = 0x900D
	SAMPLER_CUBE_SHADOW                                        = 0x8DC5
	SAMPLER_KHR                                                = 0x82E6
	SAMPLES                                                    = 0x80A9
	SAMPLES_PASSED                                             = 0x8914
	SAMPLE_ALPHA_TO_COVERAGE                                   = 0x809E
	SAMPLE_ALPHA_TO_ONE                                        = 0x809F
	SAMPLE_BUFFERS                                             = 0x80A8
	SAMPLE_COVERAGE                                            = 0x80A0
	SAMPLE_COVERAGE_INVERT                                     = 0x80AB
	SAMPLE_COVERAGE_VALUE                                      = 0x80AA
	SAMPLE_MASK                                                = 0x8E51
	SAMPLE_MASK_VALUE                                          = 0x8E52
	SAMPLE_POSITION                                            = 0x8E50
	SAMPLE_SHADING                                             = 0x8C36
	SAMPLE_SHADING_ARB                                         = 0x8C36
	SCISSOR_BOX                                                = 0x0C10
	SCISSOR_TEST                                               = 0x0C11
	SCREEN_KHR                                                 = 0x9295
	SCREEN_NV                                                  = 0x9295
	SEPARATE_ATTRIBS                                           = 0x8C8D
	SET                                                        = 0x150F
	SHADER                                                     = 0x82E1
	SHADER_BINARY_FORMATS                                      = 0x8DF8
	SHADER_COMPILER                                            = 0x8DFA
	SHADER_IMAGE_ACCESS_BARRIER_BIT                            = 0x00000020
	SHADER_IMAGE_ATOMIC                                        = 0x82A6
	SHADER_IMAGE_LOAD                                          = 0x82A4
	SHADER_IMAGE_STORE                                         = 0x82A5
	SHADER_INCLUDE_ARB                                         = 0x8DAE
	SHADER_KHR                                                 = 0x82E1
	SHADER_OBJECT_EXT                                          = 0x8B48
	SHADER_SOURCE_LENGTH                                       = 0x8B88
	SHADER_STORAGE_BARRIER_BIT                                 = 0x00002000
	SHADER_STORAGE_BLOCK                                       = 0x92E6
	SHADER_STORAGE_BUFFER                                      = 0x90D2
	SHADER_STORAGE_BUFFER_BINDING                              = 0x90D3
	SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT                     = 0x90DF
	SHADER_STORAGE_BUFFER_SIZE                                 = 0x90D5
	SHADER_STORAGE_BUFFER_START                                = 0x90D4
	SHADER_TYPE                                                = 0x8B4F
	SHADING_LANGUAGE_VERSION                                   = 0x8B8C
	SHORT                                                      = 0x1402
	SIGNALED                                                   = 0x9119
	SIGNED_NORMALIZED                                          = 0x8F9C
	SIMULTANEOUS_TEXTURE_AND_DEPTH_TEST                        = 0x82AC
	SIMULTANEOUS_TEXTURE_AND_DEPTH_WRITE                       = 0x82AE
	SIMULTANEOUS_TEXTURE_AND_STENCIL_TEST                      = 0x82AD
	SIMULTANEOUS_TEXTURE_AND_STENCIL_WRITE                     = 0x82AF
	SKIP_DECODE_EXT                                            = 0x8A4A
	SMOOTH_LINE_WIDTH_GRANULARITY                              = 0x0B23
	SMOOTH_LINE_WIDTH_RANGE                                    = 0x0B22
	SMOOTH_POINT_SIZE_GRANULARITY                              = 0x0B13
	SMOOTH_POINT_SIZE_RANGE                                    = 0x0B12
	SOFTLIGHT_KHR                                              = 0x929C
	SOFTLIGHT_NV                                               = 0x929C
	SPARSE_BUFFER_PAGE_SIZE_ARB                                = 0x82F8
	SPARSE_STORAGE_BIT_ARB                                     = 0x0400
	SPARSE_TEXTURE_FULL_ARRAY_CUBE_MIPMAPS_ARB                 = 0x91A9
	SRC1_ALPHA                                                 = 0x8589
	SRC1_COLOR                                                 = 0x88F9
	SRC_ALPHA                                                  = 0x0302
	SRC_ALPHA_SATURATE                                         = 0x0308
	SRC_ATOP_NV                                                = 0x928E
	SRC_COLOR                                                  = 0x0300
	SRC_IN_NV                                                  = 0x928A
	SRC_NV                                                     = 0x9286
	SRC_OUT_NV                                                 = 0x928C
	SRC_OVER_NV                                                = 0x9288
	SRGB                                                       = 0x8C40
	SRGB8                                                      = 0x8C41
	SRGB8_ALPHA8                                               = 0x8C43
	SRGB_ALPHA                                                 = 0x8C42
	SRGB_DECODE_ARB                                            = 0x8299
	SRGB_READ                                                  = 0x8297
	SRGB_WRITE                                                 = 0x8298
	STACK_OVERFLOW                                             = 0x0503
	STACK_OVERFLOW_KHR                                         = 0x0503
	STACK_UNDERFLOW                                            = 0x0504
	STACK_UNDERFLOW_KHR                                        = 0x0504
	STATIC_COPY                                                = 0x88E6
	STATIC_DRAW                                                = 0x88E4
	STATIC_READ                                                = 0x88E5
	STENCIL                                                    = 0x1802
	STENCIL_ATTACHMENT                                         = 0x8D20
	STENCIL_BACK_FAIL                                          = 0x8801
	STENCIL_BACK_FUNC                                          = 0x8800
	STENCIL_BACK_PASS_DEPTH_FAIL                               = 0x8802
	STENCIL_BACK_PASS_DEPTH_PASS                               = 0x8803
	STENCIL_BACK_REF                                           = 0x8CA3
	STENCIL_BACK_VALUE_MASK                                    = 0x8CA4
	STENCIL_BACK_WRITEMASK                                     = 0x8CA5
	STENCIL_BUFFER_BIT                                         = 0x00000400
	STENCIL_CLEAR_VALUE                                        = 0x0B91
	STENCIL_COMPONENTS                                         = 0x8285
	STENCIL_FAIL                                               = 0x0B94
	STENCIL_FUNC                                               = 0x0B92
	STENCIL_INDEX                                              = 0x1901
	STENCIL_INDEX1                                             = 0x8D46
	STENCIL_INDEX16                                            = 0x8D49
	STENCIL_INDEX4                                             = 0x8D47
	STENCIL_INDEX8                                             = 0x8D48
	STENCIL_PASS_DEPTH_FAIL                                    = 0x0B95
	STENCIL_PASS_DEPTH_PASS                                    = 0x0B96
	STENCIL_REF                                                = 0x0B97
	STENCIL_RENDERABLE                                         = 0x8288
	STENCIL_TEST                                               = 0x0B90
	STENCIL_VALUE_MASK                                         = 0x0B93
	STENCIL_WRITEMASK                                          = 0x0B98
	STEREO                                                     = 0x0C33
	STREAM_COPY                                                = 0x88E2
	STREAM_DRAW                                                = 0x88E0
	STREAM_READ                                                = 0x88E1
	SUBPIXEL_BITS                                              = 0x0D50
	SYNC_CL_EVENT_ARB                                          = 0x8240
	SYNC_CL_EVENT_COMPLETE_ARB                                 = 0x8241
	SYNC_CONDITION                                             = 0x9113
	SYNC_FENCE                                                 = 0x9116
	SYNC_FLAGS                                                 = 0x9115
	SYNC_FLUSH_COMMANDS_BIT                                    = 0x00000001
	SYNC_GPU_COMMANDS_COMPLETE                                 = 0x9117
	SYNC_STATUS                                                = 0x9114
	TESS_CONTROL_OUTPUT_VERTICES                               = 0x8E75
	TESS_CONTROL_SHADER                                        = 0x8E88
	TESS_CONTROL_SHADER_BIT                                    = 0x00000008
	TESS_CONTROL_SHADER_PATCHES_ARB                            = 0x82F1
	TESS_CONTROL_SUBROUTINE                                    = 0x92E9
	TESS_CONTROL_SUBROUTINE_UNIFORM                            = 0x92EF
	TESS_CONTROL_TEXTURE                                       = 0x829C
	TESS_EVALUATION_SHADER                                     = 0x8E87
	TESS_EVALUATION_SHADER_BIT                                 = 0x00000010
	TESS_EVALUATION_SHADER_INVOCATIONS_ARB                     = 0x82F2
	TESS_EVALUATION_SUBROUTINE                                 = 0x92EA
	TESS_EVALUATION_SUBROUTINE_UNIFORM                         = 0x92F0
	TESS_EVALUATION_TEXTURE                                    = 0x829D
	TESS_GEN_MODE                                              = 0x8E76
	TESS_GEN_POINT_MODE                                        = 0x8E79
	TESS_GEN_SPACING                                           = 0x8E77
	TESS_GEN_VERTEX_ORDER                                      = 0x8E78
	TEXTURE                                                    = 0x1702
	TEXTURE0                                                   = 0x84C0
	TEXTURE1                                                   = 0x84C1
	TEXTURE10                                                  = 0x84CA
	TEXTURE11                                                  = 0x84CB
	TEXTURE12                                                  = 0x84CC
	TEXTURE13                                                  = 0x84CD
	TEXTURE14                                                  = 0x84CE
	TEXTURE15                                                  = 0x84CF
	TEXTURE16                                                  = 0x84D0
	TEXTURE17                                                  = 0x84D1
	TEXTURE18                                                  = 0x84D2
	TEXTURE19                                                  = 0x84D3
	TEXTURE2                                                   = 0x84C2
	TEXTURE20                                                  = 0x84D4
	TEXTURE21                                                  = 0x84D5
	TEXTURE22                                                  = 0x84D6
	TEXTURE23                                                  = 0x84D7
	TEXTURE24                                                  = 0x84D8
	TEXTURE25                                                  = 0x84D9
	TEXTURE26                                                  = 0x84DA
	TEXTURE27                                                  = 0x84DB
	TEXTURE28                                                  = 0x84DC
	TEXTURE29                                                  = 0x84DD
	TEXTURE3                                                   = 0x84C3
	TEXTURE30                                                  = 0x84DE
	TEXTURE31                                                  = 0x84DF
	TEXTURE4                                                   = 0x84C4
	TEXTURE5                                                   = 0x84C5
	TEXTURE6                                                   = 0x84C6
	TEXTURE7                                                   = 0x84C7
	TEXTURE8                                                   = 0x84C8
	TEXTURE9                                                   = 0x84C9
	TEXTURE_1D                                                 = 0x0DE0
	TEXTURE_1D_ARRAY                                           = 0x8C18
	TEXTURE_2D                                                 = 0x0DE1
	TEXTURE_2D_ARRAY                                           = 0x8C1A
	TEXTURE_2D_MULTISAMPLE                                     = 0x9100
	TEXTURE_2D_MULTISAMPLE_ARRAY                               = 0x9102
	TEXTURE_3D                                                 = 0x806F
	TEXTURE_ALPHA_SIZE                                         = 0x805F
	TEXTURE_ALPHA_TYPE                                         = 0x8C13
	TEXTURE_BASE_LEVEL                                         = 0x813C
	TEXTURE_BINDING                                            = 0x82EB
	TEXTURE_BINDING_1D                                         = 0x8068
	TEXTURE_BINDING_1D_ARRAY                                   = 0x8C1C
	TEXTURE_BINDING_2D                                         = 0x8069
	TEXTURE_BINDING_2D_ARRAY                                   = 0x8C1D
	TEXTURE_BINDING_2D_MULTISAMPLE                             = 0x9104
	TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY                       = 0x9105
	TEXTURE_BINDING_3D                                         = 0x806A
	TEXTURE_BINDING_BUFFER                                     = 0x8C2C
	TEXTURE_BINDING_CUBE_MAP                                   = 0x8514
	TEXTURE_BINDING_CUBE_MAP_ARRAY                             = 0x900A
	TEXTURE_BINDING_CUBE_MAP_ARRAY_ARB                         = 0x900A
	TEXTURE_BINDING_RECTANGLE                                  = 0x84F6
	TEXTURE_BLUE_SIZE                                          = 0x805E
	TEXTURE_BLUE_TYPE                                          = 0x8C12
	TEXTURE_BORDER_COLOR                                       = 0x1004
	TEXTURE_BUFFER                                             = 0x8C2A
	TEXTURE_BUFFER_DATA_STORE_BINDING                          = 0x8C2D
	TEXTURE_BUFFER_OFFSET                                      = 0x919D
	TEXTURE_BUFFER_OFFSET_ALIGNMENT                            = 0x919F
	TEXTURE_BUFFER_SIZE                                        = 0x919E
	TEXTURE_COMPARE_FUNC                                       = 0x884D
	TEXTURE_COMPARE_MODE                                       = 0x884C
	TEXTURE_COMPRESSED                                         = 0x86A1
	TEXTURE_COMPRESSED_BLOCK_HEIGHT                            = 0x82B2
	TEXTURE_COMPRESSED_BLOCK_SIZE                              = 0x82B3
	TEXTURE_COMPRESSED_BLOCK_WIDTH                             = 0x82B1
	TEXTURE_COMPRESSED_IMAGE_SIZE                              = 0x86A0
	TEXTURE_COMPRESSION_HINT                                   = 0x84EF
	TEXTURE_CUBE_MAP                                           = 0x8513
	TEXTURE_CUBE_MAP_ARRAY                                     = 0x9009
	TEXTURE_CUBE_MAP_ARRAY_ARB                                 = 0x9009
	TEXTURE_CUBE_MAP_NEGATIVE_X                                = 0x8516
	TEXTURE_CUBE_MAP_NEGATIVE_Y                                = 0x8518
	TEXTURE_CUBE_MAP_NEGATIVE_Z                                = 0x851A
	TEXTURE_CUBE_MAP_POSITIVE_X                                = 0x8515
	TEXTURE_CUBE_MAP_POSITIVE_Y                                = 0x8517
	TEXTURE_CUBE_MAP_POSITIVE_Z                                = 0x8519
	TEXTURE_CUBE_MAP_SEAMLESS                                  = 0x884F
	TEXTURE_DEPTH                                              = 0x8071
	TEXTURE_DEPTH_SIZE                                         = 0x884A
	TEXTURE_DEPTH_TYPE                                         = 0x8C16
	TEXTURE_FETCH_BARRIER_BIT                                  = 0x00000008
	TEXTURE_FILTER_CONTROL_EXT                                 = 0x8500
	TEXTURE_FIXED_SAMPLE_LOCATIONS                             = 0x9107
	TEXTURE_GATHER                                             = 0x82A2
	TEXTURE_GATHER_SHADOW                                      = 0x82A3
	TEXTURE_GREEN_SIZE                                         = 0x805D
	TEXTURE_GREEN_TYPE                                         = 0x8C11
	TEXTURE_HEIGHT                                             = 0x1001
	TEXTURE_IMAGE_FORMAT                                       = 0x828F
	TEXTURE_IMAGE_TYPE                                         = 0x8290
	TEXTURE_IMMUTABLE_FORMAT                                   = 0x912F
	TEXTURE_IMMUTABLE_LEVELS                                   = 0x82DF
	TEXTURE_INTERNAL_FORMAT                                    = 0x1003
	TEXTURE_LOD_BIAS                                           = 0x8501
	TEXTURE_LOD_BIAS_EXT                                       = 0x8501
	TEXTURE_MAG_FILTER                                         = 0x2800
	TEXTURE_MAX_ANISOTROPY_EXT                                 = 0x84FE
	TEXTURE_MAX_LEVEL                                          = 0x813D
	TEXTURE_MAX_LOD                                            = 0x813B
	TEXTURE_MIN_FILTER                                         = 0x2801
	TEXTURE_MIN_LOD                                            = 0x813A
	TEXTURE_RECTANGLE                                          = 0x84F5
	TEXTURE_RED_SIZE                                           = 0x805C
	TEXTURE_RED_TYPE                                           = 0x8C10
	TEXTURE_SAMPLES                                            = 0x9106
	TEXTURE_SHADOW                                             = 0x82A1
	TEXTURE_SHARED_SIZE                                        = 0x8C3F
	TEXTURE_SPARSE_ARB                                         = 0x91A6
	TEXTURE_SRGB_DECODE_EXT                                    = 0x8A48
	TEXTURE_STENCIL_SIZE                                       = 0x88F1
	TEXTURE_SWIZZLE_A                                          = 0x8E45
	TEXTURE_SWIZZLE_B                                          = 0x8E44
	TEXTURE_SWIZZLE_G                                          = 0x8E43
	TEXTURE_SWIZZLE_R                                          = 0x8E42
	TEXTURE_SWIZZLE_RGBA                                       = 0x8E46
	TEXTURE_TARGET                                             = 0x1006
	TEXTURE_UPDATE_BARRIER_BIT                                 = 0x00000100
	TEXTURE_VIEW                                               = 0x82B5
	TEXTURE_VIEW_MIN_LAYER                                     = 0x82DD
	TEXTURE_VIEW_MIN_LEVEL                                     = 0x82DB
	TEXTURE_VIEW_NUM_LAYERS                                    = 0x82DE
	TEXTURE_VIEW_NUM_LEVELS                                    = 0x82DC
	TEXTURE_WIDTH                                              = 0x1000
	TEXTURE_WRAP_R                                             = 0x8072
	TEXTURE_WRAP_S                                             = 0x2802
	TEXTURE_WRAP_T                                             = 0x2803
	TIMEOUT_EXPIRED                                            = 0x911B
	TIMEOUT_IGNORED                                            = 0xFFFFFFFFFFFFFFFF
	TIMESTAMP                                                  = 0x8E28
	TIME_ELAPSED                                               = 0x88BF
	TOP_LEVEL_ARRAY_SIZE                                       = 0x930C
	TOP_LEVEL_ARRAY_STRIDE                                     = 0x930D
	TRANSFORM_FEEDBACK                                         = 0x8E22
	TRANSFORM_FEEDBACK_ACTIVE                                  = 0x8E24
	TRANSFORM_FEEDBACK_BARRIER_BIT                             = 0x00000800
	TRANSFORM_FEEDBACK_BINDING                                 = 0x8E25
	TRANSFORM_FEEDBACK_BUFFER                                  = 0x8C8E
	TRANSFORM_FEEDBACK_BUFFER_ACTIVE                           = 0x8E24
	TRANSFORM_FEEDBACK_BUFFER_BINDING                          = 0x8C8F
	TRANSFORM_FEEDBACK_BUFFER_INDEX                            = 0x934B
	TRANSFORM_FEEDBACK_BUFFER_MODE                             = 0x8C7F
	TRANSFORM_FEEDBACK_BUFFER_PAUSED                           = 0x8E23
	TRANSFORM_FEEDBACK_BUFFER_SIZE                             = 0x8C85
	TRANSFORM_FEEDBACK_BUFFER_START                            = 0x8C84
	TRANSFORM_FEEDBACK_BUFFER_STRIDE                           = 0x934C
	TRANSFORM_FEEDBACK_OVERFLOW_ARB                            = 0x82EC
	TRANSFORM_FEEDBACK_PAUSED                                  = 0x8E23
	TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN                      = 0x8C88
	TRANSFORM_FEEDBACK_STREAM_OVERFLOW_ARB                     = 0x82ED
	TRANSFORM_FEEDBACK_VARYING                                 = 0x92F4
	TRANSFORM_FEEDBACK_VARYINGS                                = 0x8C83
	TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH                      = 0x8C76
	TRIANGLES                                                  = 0x0004
	TRIANGLES_ADJACENCY                                        = 0x000C
	TRIANGLE_FAN                                               = 0x0006
	TRIANGLE_STRIP                                             = 0x0005
	TRIANGLE_STRIP_ADJACENCY                                   = 0x000D
	TRUE                                                       = 1
	TYPE                                                       = 0x92FA
	UNCORRELATED_NV                                            = 0x9282
	UNDEFINED_VERTEX                                           = 0x8260
	UNIFORM                                                    = 0x92E1
	UNIFORM_ARRAY_STRIDE                                       = 0x8A3C
	UNIFORM_ATOMIC_COUNTER_BUFFER_INDEX                        = 0x92DA
	UNIFORM_BARRIER_BIT                                        = 0x00000004
	UNIFORM_BLOCK                                              = 0x92E2
	UNIFORM_BLOCK_ACTIVE_UNIFORMS                              = 0x8A42
	UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES                       = 0x8A43
	UNIFORM_BLOCK_BINDING                                      = 0x8A3F
	UNIFORM_BLOCK_DATA_SIZE                                    = 0x8A40
	UNIFORM_BLOCK_INDEX                                        = 0x8A3A
	UNIFORM_BLOCK_NAME_LENGTH                                  = 0x8A41
	UNIFORM_BLOCK_REFERENCED_BY_COMPUTE_SHADER                 = 0x90EC
	UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER                = 0x8A46
	UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER                = 0x8A45
	UNIFORM_BLOCK_REFERENCED_BY_TESS_CONTROL_SHADER            = 0x84F0
	UNIFORM_BLOCK_REFERENCED_BY_TESS_EVALUATION_SHADER         = 0x84F1
	UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER                  = 0x8A44
	UNIFORM_BUFFER                                             = 0x8A11
	UNIFORM_BUFFER_BINDING                                     = 0x8A28
	UNIFORM_BUFFER_OFFSET_ALIGNMENT                            = 0x8A34
	UNIFORM_BUFFER_SIZE                                        = 0x8A2A
	UNIFORM_BUFFER_START                                       = 0x8A29
	UNIFORM_IS_ROW_MAJOR                                       = 0x8A3E
	UNIFORM_MATRIX_STRIDE                                      = 0x8A3D
	UNIFORM_NAME_LENGTH                                        = 0x8A39
	UNIFORM_OFFSET                                             = 0x8A3B
	UNIFORM_SIZE                                               = 0x8A38
	UNIFORM_TYPE                                               = 0x8A37
	UNKNOWN_CONTEXT_RESET                                      = 0x8255
	UNKNOWN_CONTEXT_RESET_ARB                                  = 0x8255
	UNKNOWN_CONTEXT_RESET_KHR                                  = 0x8255
	UNPACK_ALIGNMENT                                           = 0x0CF5
	UNPACK_COMPRESSED_BLOCK_DEPTH                              = 0x9129
	UNPACK_COMPRESSED_BLOCK_HEIGHT                             = 0x9128
	UNPACK_COMPRESSED_BLOCK_SIZE                               = 0x912A
	UNPACK_COMPRESSED_BLOCK_WIDTH                              = 0x9127
	UNPACK_IMAGE_HEIGHT                                        = 0x806E
	UNPACK_LSB_FIRST                                           = 0x0CF1
	UNPACK_ROW_LENGTH                                          = 0x0CF2
	UNPACK_SKIP_IMAGES                                         = 0x806D
	UNPACK_SKIP_PIXELS                                         = 0x0CF4
	UNPACK_SKIP_ROWS                                           = 0x0CF3
	UNPACK_SWAP_BYTES                                          = 0x0CF0
	UNSIGNALED                                                 = 0x9118
	UNSIGNED_BYTE                                              = 0x1401
	UNSIGNED_BYTE_2_3_3_REV                                    = 0x8362
	UNSIGNED_BYTE_3_3_2                                        = 0x8032
	UNSIGNED_INT                                               = 0x1405
	UNSIGNED_INT64_AMD                                         = 0x8BC2
	UNSIGNED_INT64_ARB                                         = 0x140F
	UNSIGNED_INT_10F_11F_11F_REV                               = 0x8C3B
	UNSIGNED_INT_10_10_10_2                                    = 0x8036
	UNSIGNED_INT_24_8                                          = 0x84FA
	UNSIGNED_INT_2_10_10_10_REV                                = 0x8368
	UNSIGNED_INT_5_9_9_9_REV                                   = 0x8C3E
	UNSIGNED_INT_8_8_8_8                                       = 0x8035
	UNSIGNED_INT_8_8_8_8_REV                                   = 0x8367
	UNSIGNED_INT_ATOMIC_COUNTER                                = 0x92DB
	UNSIGNED_INT_IMAGE_1D                                      = 0x9062
	UNSIGNED_INT_IMAGE_1D_ARRAY                                = 0x9068
	UNSIGNED_INT_IMAGE_2D                                      = 0x9063
	UNSIGNED_INT_IMAGE_2D_ARRAY                                = 0x9069
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE                          = 0x906B
	UNSIGNED_INT_IMAGE_2D_MULTISAMPLE_ARRAY                    = 0x906C
	UNSIGNED_INT_IMAGE_2D_RECT                                 = 0x9065
	UNSIGNED_INT_IMAGE_3D                                      = 0x9064
	UNSIGNED_INT_IMAGE_BUFFER                                  = 0x9067
	UNSIGNED_INT_IMAGE_CUBE                                    = 0x9066
	UNSIGNED_INT_IMAGE_CUBE_MAP_ARRAY                          = 0x906A
	UNSIGNED_INT_SAMPLER_1D                                    = 0x8DD1
	UNSIGNED_INT_SAMPLER_1D_ARRAY                              = 0x8DD6
	UNSIGNED_INT_SAMPLER_2D                                    = 0x8DD2
	UNSIGNED_INT_SAMPLER_2D_ARRAY                              = 0x8DD7
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE                        = 0x910A
	UNSIGNED_INT_SAMPLER_2D_MULTISAMPLE_ARRAY                  = 0x910D
	UNSIGNED_INT_SAMPLER_2D_RECT                               = 0x8DD5
	UNSIGNED_INT_SAMPLER_3D                                    = 0x8DD3
	UNSIGNED_INT_SAMPLER_BUFFER                                = 0x8DD8
	UNSIGNED_INT_SAMPLER_CUBE                                  = 0x8DD4
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY                        = 0x900F
	UNSIGNED_INT_SAMPLER_CUBE_MAP_ARRAY_ARB                    = 0x900F
	UNSIGNED_INT_VEC2                                          = 0x8DC6
	UNSIGNED_INT_VEC3                                          = 0x8DC7
	UNSIGNED_INT_VEC4                                          = 0x8DC8
	UNSIGNED_NORMALIZED                                        = 0x8C17
	UNSIGNED_SHORT                                             = 0x1403
	UNSIGNED_SHORT_1_5_5_5_REV                                 = 0x8366
	UNSIGNED_SHORT_4_4_4_4                                     = 0x8033
	UNSIGNED_SHORT_4_4_4_4_REV                                 = 0x8365
	UNSIGNED_SHORT_5_5_5_1                                     = 0x8034
	UNSIGNED_SHORT_5_6_5                                       = 0x8363
	UNSIGNED_SHORT_5_6_5_REV                                   = 0x8364
	UNSIGNED_SHORT_8_8_APPLE                                   = 0x85BA
	UNSIGNED_SHORT_8_8_REV_APPLE                               = 0x85BB
	UPPER_LEFT                                                 = 0x8CA2
	VALIDATE_STATUS                                            = 0x8B83
	VENDOR                                                     = 0x1F00
	VERSION                                                    = 0x1F02
	VERTEX_ARRAY                                               = 0x8074
	VERTEX_ARRAY_BINDING                                       = 0x85B5
	VERTEX_ARRAY_KHR                                           = 0x8074
	VERTEX_ARRAY_OBJECT_EXT                                    = 0x9154
	VERTEX_ATTRIB_ARRAY_BARRIER_BIT                            = 0x00000001
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING                         = 0x889F
	VERTEX_ATTRIB_ARRAY_DIVISOR                                = 0x88FE
	VERTEX_ATTRIB_ARRAY_ENABLED                                = 0x8622
	VERTEX_ATTRIB_ARRAY_INTEGER                                = 0x88FD
	VERTEX_ATTRIB_ARRAY_NORMALIZED                             = 0x886A
	VERTEX_ATTRIB_ARRAY_POINTER                                = 0x8645
	VERTEX_ATTRIB_ARRAY_SIZE                                   = 0x8623
	VERTEX_ATTRIB_ARRAY_STRIDE                                 = 0x8624
	VERTEX_ATTRIB_ARRAY_TYPE                                   = 0x8625
	VERTEX_ATTRIB_BINDING                                      = 0x82D4
	VERTEX_ATTRIB_RELATIVE_OFFSET                              = 0x82D5
	VERTEX_BINDING_DIVISOR                                     = 0x82D6
	VERTEX_BINDING_OFFSET                                      = 0x82D7
	VERTEX_BINDING_STRIDE                                      = 0x82D8
	VERTEX_PROGRAM_POINT_SIZE                                  = 0x8642
	VERTEX_SHADER                                              = 0x8B31
	VERTEX_SHADER_BIT                                          = 0x00000001
	VERTEX_SHADER_BIT_EXT                                      = 0x00000001
	VERTEX_SHADER_INVOCATIONS_ARB                              = 0x82F0
	VERTEX_SUBROUTINE                                          = 0x92E8
	VERTEX_SUBROUTINE_UNIFORM                                  = 0x92EE
	VERTEX_TEXTURE                                             = 0x829B
	VERTICES_SUBMITTED_ARB                                     = 0x82EE
	VIEWPORT                                                   = 0x0BA2
	VIEWPORT_BOUNDS_RANGE                                      = 0x825D
	VIEWPORT_INDEX_PROVOKING_VERTEX                            = 0x825F
	VIEWPORT_SUBPIXEL_BITS                                     = 0x825C
	VIEW_CLASS_128_BITS                                        = 0x82C4
	VIEW_CLASS_16_BITS                                         = 0x82CA
	VIEW_CLASS_24_BITS                                         = 0x82C9
	VIEW_CLASS_32_BITS                                         = 0x82C8
	VIEW_CLASS_48_BITS                                         = 0x82C7
	VIEW_CLASS_64_BITS                                         = 0x82C6
	VIEW_CLASS_8_BITS                                          = 0x82CB
	VIEW_CLASS_96_BITS                                         = 0x82C5
	VIEW_CLASS_BPTC_FLOAT                                      = 0x82D3
	VIEW_CLASS_BPTC_UNORM                                      = 0x82D2
	VIEW_CLASS_RGTC1_RED                                       = 0x82D0
	VIEW_CLASS_RGTC2_RG                                        = 0x82D1
	VIEW_CLASS_S3TC_DXT1_RGB                                   = 0x82CC
	VIEW_CLASS_S3TC_DXT1_RGBA                                  = 0x82CD
	VIEW_CLASS_S3TC_DXT3_RGBA                                  = 0x82CE
	VIEW_CLASS_S3TC_DXT5_RGBA                                  = 0x82CF
	VIEW_COMPATIBILITY_CLASS                                   = 0x82B6
	VIRTUAL_PAGE_SIZE_INDEX_ARB                                = 0x91A7
	VIRTUAL_PAGE_SIZE_X_ARB                                    = 0x9195
	VIRTUAL_PAGE_SIZE_Y_ARB                                    = 0x9196
	VIRTUAL_PAGE_SIZE_Z_ARB                                    = 0x9197
	VIVIDLIGHT_NV                                              = 0x92A6
	WAIT_FAILED                                                = 0x911D
	WRITE_ONLY                                                 = 0x88B9
	XOR                                                        = 0x1506
	XOR_NV                                                     = 0x1506
	ZERO                                                       = 0
	ZERO_TO_ONE                                                = 0x935F
)

var (
	gpAccumxOES                                   C.GPACCUMXOES
	gpActiveProgramEXT                            C.GPACTIVEPROGRAMEXT
	gpActiveShaderProgram                         C.GPACTIVESHADERPROGRAM
	gpActiveShaderProgramEXT                      C.GPACTIVESHADERPROGRAMEXT
	gpActiveTexture                               C.GPACTIVETEXTURE
	gpAlphaFuncxOES                               C.GPALPHAFUNCXOES
	gpAttachShader                                C.GPATTACHSHADER
	gpBeginConditionalRender                      C.GPBEGINCONDITIONALRENDER
	gpBeginPerfMonitorAMD                         C.GPBEGINPERFMONITORAMD
	gpBeginPerfQueryINTEL                         C.GPBEGINPERFQUERYINTEL
	gpBeginQuery                                  C.GPBEGINQUERY
	gpBeginQueryIndexed                           C.GPBEGINQUERYINDEXED
	gpBeginTransformFeedback                      C.GPBEGINTRANSFORMFEEDBACK
	gpBindAttribLocation                          C.GPBINDATTRIBLOCATION
	gpBindBuffer                                  C.GPBINDBUFFER
	gpBindBufferBase                              C.GPBINDBUFFERBASE
	gpBindBufferRange                             C.GPBINDBUFFERRANGE
	gpBindBuffersBase                             C.GPBINDBUFFERSBASE
	gpBindBuffersRange                            C.GPBINDBUFFERSRANGE
	gpBindFragDataLocation                        C.GPBINDFRAGDATALOCATION
	gpBindFragDataLocationIndexed                 C.GPBINDFRAGDATALOCATIONINDEXED
	gpBindFramebuffer                             C.GPBINDFRAMEBUFFER
	gpBindImageTexture                            C.GPBINDIMAGETEXTURE
	gpBindImageTextures                           C.GPBINDIMAGETEXTURES
	gpBindProgramPipeline                         C.GPBINDPROGRAMPIPELINE
	gpBindProgramPipelineEXT                      C.GPBINDPROGRAMPIPELINEEXT
	gpBindRenderbuffer                            C.GPBINDRENDERBUFFER
	gpBindSampler                                 C.GPBINDSAMPLER
	gpBindSamplers                                C.GPBINDSAMPLERS
	gpBindTexture                                 C.GPBINDTEXTURE
	gpBindTextureUnit                             C.GPBINDTEXTUREUNIT
	gpBindTextures                                C.GPBINDTEXTURES
	gpBindTransformFeedback                       C.GPBINDTRANSFORMFEEDBACK
	gpBindVertexArray                             C.GPBINDVERTEXARRAY
	gpBindVertexBuffer                            C.GPBINDVERTEXBUFFER
	gpBindVertexBuffers                           C.GPBINDVERTEXBUFFERS
	gpBitmapxOES                                  C.GPBITMAPXOES
	gpBlendBarrierKHR                             C.GPBLENDBARRIERKHR
	gpBlendBarrierNV                              C.GPBLENDBARRIERNV
	gpBlendColor                                  C.GPBLENDCOLOR
	gpBlendColorxOES                              C.GPBLENDCOLORXOES
	gpBlendEquation                               C.GPBLENDEQUATION
	gpBlendEquationEXT                            C.GPBLENDEQUATIONEXT
	gpBlendEquationSeparate                       C.GPBLENDEQUATIONSEPARATE
	gpBlendEquationSeparatei                      C.GPBLENDEQUATIONSEPARATEI
	gpBlendEquationSeparateiARB                   C.GPBLENDEQUATIONSEPARATEIARB
	gpBlendEquationi                              C.GPBLENDEQUATIONI
	gpBlendEquationiARB                           C.GPBLENDEQUATIONIARB
	gpBlendFunc                                   C.GPBLENDFUNC
	gpBlendFuncSeparate                           C.GPBLENDFUNCSEPARATE
	gpBlendFuncSeparatei                          C.GPBLENDFUNCSEPARATEI
	gpBlendFuncSeparateiARB                       C.GPBLENDFUNCSEPARATEIARB
	gpBlendFunci                                  C.GPBLENDFUNCI
	gpBlendFunciARB                               C.GPBLENDFUNCIARB
	gpBlendParameteriNV                           C.GPBLENDPARAMETERINV
	gpBlitFramebuffer                             C.GPBLITFRAMEBUFFER
	gpBlitNamedFramebuffer                        C.GPBLITNAMEDFRAMEBUFFER
	gpBufferData                                  C.GPBUFFERDATA
	gpBufferPageCommitmentARB                     C.GPBUFFERPAGECOMMITMENTARB
	gpBufferStorage                               C.GPBUFFERSTORAGE
	gpBufferSubData                               C.GPBUFFERSUBDATA
	gpCheckFramebufferStatus                      C.GPCHECKFRAMEBUFFERSTATUS
	gpCheckNamedFramebufferStatus                 C.GPCHECKNAMEDFRAMEBUFFERSTATUS
	gpClampColor                                  C.GPCLAMPCOLOR
	gpClear                                       C.GPCLEAR
	gpClearAccumxOES                              C.GPCLEARACCUMXOES
	gpClearBufferData                             C.GPCLEARBUFFERDATA
	gpClearBufferSubData                          C.GPCLEARBUFFERSUBDATA
	gpClearBufferfi                               C.GPCLEARBUFFERFI
	gpClearBufferfv                               C.GPCLEARBUFFERFV
	gpClearBufferiv                               C.GPCLEARBUFFERIV
	gpClearBufferuiv                              C.GPCLEARBUFFERUIV
	gpClearColor                                  C.GPCLEARCOLOR
	gpClearColorxOES                              C.GPCLEARCOLORXOES
	gpClearDepth                                  C.GPCLEARDEPTH
	gpClearDepthf                                 C.GPCLEARDEPTHF
	gpClearDepthfOES                              C.GPCLEARDEPTHFOES
	gpClearDepthxOES                              C.GPCLEARDEPTHXOES
	gpClearNamedBufferData                        C.GPCLEARNAMEDBUFFERDATA
	gpClearNamedBufferSubData                     C.GPCLEARNAMEDBUFFERSUBDATA
	gpClearNamedFramebufferfi                     C.GPCLEARNAMEDFRAMEBUFFERFI
	gpClearNamedFramebufferfv                     C.GPCLEARNAMEDFRAMEBUFFERFV
	gpClearNamedFramebufferiv                     C.GPCLEARNAMEDFRAMEBUFFERIV
	gpClearNamedFramebufferuiv                    C.GPCLEARNAMEDFRAMEBUFFERUIV
	gpClearStencil                                C.GPCLEARSTENCIL
	gpClearTexImage                               C.GPCLEARTEXIMAGE
	gpClearTexSubImage                            C.GPCLEARTEXSUBIMAGE
	gpClientWaitSync                              C.GPCLIENTWAITSYNC
	gpClipControl                                 C.GPCLIPCONTROL
	gpClipPlanefOES                               C.GPCLIPPLANEFOES
	gpClipPlanexOES                               C.GPCLIPPLANEXOES
	gpColor3xOES                                  C.GPCOLOR3XOES
	gpColor3xvOES                                 C.GPCOLOR3XVOES
	gpColor4xOES                                  C.GPCOLOR4XOES
	gpColor4xvOES                                 C.GPCOLOR4XVOES
	gpColorMask                                   C.GPCOLORMASK
	gpColorMaski                                  C.GPCOLORMASKI
	gpCompileShader                               C.GPCOMPILESHADER
	gpCompileShaderIncludeARB                     C.GPCOMPILESHADERINCLUDEARB
	gpCompressedTexImage1D                        C.GPCOMPRESSEDTEXIMAGE1D
	gpCompressedTexImage2D                        C.GPCOMPRESSEDTEXIMAGE2D
	gpCompressedTexImage3D                        C.GPCOMPRESSEDTEXIMAGE3D
	gpCompressedTexSubImage1D                     C.GPCOMPRESSEDTEXSUBIMAGE1D
	gpCompressedTexSubImage2D                     C.GPCOMPRESSEDTEXSUBIMAGE2D
	gpCompressedTexSubImage3D                     C.GPCOMPRESSEDTEXSUBIMAGE3D
	gpCompressedTextureSubImage1D                 C.GPCOMPRESSEDTEXTURESUBIMAGE1D
	gpCompressedTextureSubImage2D                 C.GPCOMPRESSEDTEXTURESUBIMAGE2D
	gpCompressedTextureSubImage3D                 C.GPCOMPRESSEDTEXTURESUBIMAGE3D
	gpConvolutionParameterxOES                    C.GPCONVOLUTIONPARAMETERXOES
	gpConvolutionParameterxvOES                   C.GPCONVOLUTIONPARAMETERXVOES
	gpCopyBufferSubData                           C.GPCOPYBUFFERSUBDATA
	gpCopyImageSubData                            C.GPCOPYIMAGESUBDATA
	gpCopyNamedBufferSubData                      C.GPCOPYNAMEDBUFFERSUBDATA
	gpCopyTexImage1D                              C.GPCOPYTEXIMAGE1D
	gpCopyTexImage2D                              C.GPCOPYTEXIMAGE2D
	gpCopyTexSubImage1D                           C.GPCOPYTEXSUBIMAGE1D
	gpCopyTexSubImage2D                           C.GPCOPYTEXSUBIMAGE2D
	gpCopyTexSubImage3D                           C.GPCOPYTEXSUBIMAGE3D
	gpCopyTextureSubImage1D                       C.GPCOPYTEXTURESUBIMAGE1D
	gpCopyTextureSubImage2D                       C.GPCOPYTEXTURESUBIMAGE2D
	gpCopyTextureSubImage3D                       C.GPCOPYTEXTURESUBIMAGE3D
	gpCreateBuffers                               C.GPCREATEBUFFERS
	gpCreateFramebuffers                          C.GPCREATEFRAMEBUFFERS
	gpCreatePerfQueryINTEL                        C.GPCREATEPERFQUERYINTEL
	gpCreateProgram                               C.GPCREATEPROGRAM
	gpCreateProgramPipelines                      C.GPCREATEPROGRAMPIPELINES
	gpCreateQueries                               C.GPCREATEQUERIES
	gpCreateRenderbuffers                         C.GPCREATERENDERBUFFERS
	gpCreateSamplers                              C.GPCREATESAMPLERS
	gpCreateShader                                C.GPCREATESHADER
	gpCreateShaderProgramEXT                      C.GPCREATESHADERPROGRAMEXT
	gpCreateShaderProgramv                        C.GPCREATESHADERPROGRAMV
	gpCreateShaderProgramvEXT                     C.GPCREATESHADERPROGRAMVEXT
	gpCreateSyncFromCLeventARB                    C.GPCREATESYNCFROMCLEVENTARB
	gpCreateTextures                              C.GPCREATETEXTURES
	gpCreateTransformFeedbacks                    C.GPCREATETRANSFORMFEEDBACKS
	gpCreateVertexArrays                          C.GPCREATEVERTEXARRAYS
	gpCullFace                                    C.GPCULLFACE
	gpDebugMessageCallback                        C.GPDEBUGMESSAGECALLBACK
	gpDebugMessageCallbackARB                     C.GPDEBUGMESSAGECALLBACKARB
	gpDebugMessageCallbackKHR                     C.GPDEBUGMESSAGECALLBACKKHR
	gpDebugMessageControl                         C.GPDEBUGMESSAGECONTROL
	gpDebugMessageControlARB                      C.GPDEBUGMESSAGECONTROLARB
	gpDebugMessageControlKHR                      C.GPDEBUGMESSAGECONTROLKHR
	gpDebugMessageInsert                          C.GPDEBUGMESSAGEINSERT
	gpDebugMessageInsertARB                       C.GPDEBUGMESSAGEINSERTARB
	gpDebugMessageInsertKHR                       C.GPDEBUGMESSAGEINSERTKHR
	gpDeleteBuffers                               C.GPDELETEBUFFERS
	gpDeleteFencesNV                              C.GPDELETEFENCESNV
	gpDeleteFramebuffers                          C.GPDELETEFRAMEBUFFERS
	gpDeleteNamedStringARB                        C.GPDELETENAMEDSTRINGARB
	gpDeletePerfMonitorsAMD                       C.GPDELETEPERFMONITORSAMD
	gpDeletePerfQueryINTEL                        C.GPDELETEPERFQUERYINTEL
	gpDeleteProgram                               C.GPDELETEPROGRAM
	gpDeleteProgramPipelines                      C.GPDELETEPROGRAMPIPELINES
	gpDeleteProgramPipelinesEXT                   C.GPDELETEPROGRAMPIPELINESEXT
	gpDeleteQueries                               C.GPDELETEQUERIES
	gpDeleteRenderbuffers                         C.GPDELETERENDERBUFFERS
	gpDeleteSamplers                              C.GPDELETESAMPLERS
	gpDeleteShader                                C.GPDELETESHADER
	gpDeleteSync                                  C.GPDELETESYNC
	gpDeleteTextures                              C.GPDELETETEXTURES
	gpDeleteTransformFeedbacks                    C.GPDELETETRANSFORMFEEDBACKS
	gpDeleteVertexArrays                          C.GPDELETEVERTEXARRAYS
	gpDepthFunc                                   C.GPDEPTHFUNC
	gpDepthMask                                   C.GPDEPTHMASK
	gpDepthRange                                  C.GPDEPTHRANGE
	gpDepthRangeArrayv                            C.GPDEPTHRANGEARRAYV
	gpDepthRangeIndexed                           C.GPDEPTHRANGEINDEXED
	gpDepthRangef                                 C.GPDEPTHRANGEF
	gpDepthRangefOES                              C.GPDEPTHRANGEFOES
	gpDepthRangexOES                              C.GPDEPTHRANGEXOES
	gpDetachShader                                C.GPDETACHSHADER
	gpDisable                                     C.GPDISABLE
	gpDisableVertexArrayAttrib                    C.GPDISABLEVERTEXARRAYATTRIB
	gpDisableVertexAttribArray                    C.GPDISABLEVERTEXATTRIBARRAY
	gpDisablei                                    C.GPDISABLEI
	gpDispatchCompute                             C.GPDISPATCHCOMPUTE
	gpDispatchComputeGroupSizeARB                 C.GPDISPATCHCOMPUTEGROUPSIZEARB
	gpDispatchComputeIndirect                     C.GPDISPATCHCOMPUTEINDIRECT
	gpDrawArrays                                  C.GPDRAWARRAYS
	gpDrawArraysIndirect                          C.GPDRAWARRAYSINDIRECT
	gpDrawArraysInstanced                         C.GPDRAWARRAYSINSTANCED
	gpDrawArraysInstancedBaseInstance             C.GPDRAWARRAYSINSTANCEDBASEINSTANCE
	gpDrawArraysInstancedEXT                      C.GPDRAWARRAYSINSTANCEDEXT
	gpDrawBuffer                                  C.GPDRAWBUFFER
	gpDrawBuffers                                 C.GPDRAWBUFFERS
	gpDrawElements                                C.GPDRAWELEMENTS
	gpDrawElementsBaseVertex                      C.GPDRAWELEMENTSBASEVERTEX
	gpDrawElementsIndirect                        C.GPDRAWELEMENTSINDIRECT
	gpDrawElementsInstanced                       C.GPDRAWELEMENTSINSTANCED
	gpDrawElementsInstancedBaseInstance           C.GPDRAWELEMENTSINSTANCEDBASEINSTANCE
	gpDrawElementsInstancedBaseVertex             C.GPDRAWELEMENTSINSTANCEDBASEVERTEX
	gpDrawElementsInstancedBaseVertexBaseInstance C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE
	gpDrawElementsInstancedEXT                    C.GPDRAWELEMENTSINSTANCEDEXT
	gpDrawRangeElements                           C.GPDRAWRANGEELEMENTS
	gpDrawRangeElementsBaseVertex                 C.GPDRAWRANGEELEMENTSBASEVERTEX
	gpDrawTransformFeedback                       C.GPDRAWTRANSFORMFEEDBACK
	gpDrawTransformFeedbackInstanced              C.GPDRAWTRANSFORMFEEDBACKINSTANCED
	gpDrawTransformFeedbackStream                 C.GPDRAWTRANSFORMFEEDBACKSTREAM
	gpDrawTransformFeedbackStreamInstanced        C.GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED
	gpEnable                                      C.GPENABLE
	gpEnableVertexArrayAttrib                     C.GPENABLEVERTEXARRAYATTRIB
	gpEnableVertexAttribArray                     C.GPENABLEVERTEXATTRIBARRAY
	gpEnablei                                     C.GPENABLEI
	gpEndConditionalRender                        C.GPENDCONDITIONALRENDER
	gpEndPerfMonitorAMD                           C.GPENDPERFMONITORAMD
	gpEndPerfQueryINTEL                           C.GPENDPERFQUERYINTEL
	gpEndQuery                                    C.GPENDQUERY
	gpEndQueryIndexed                             C.GPENDQUERYINDEXED
	gpEndTransformFeedback                        C.GPENDTRANSFORMFEEDBACK
	gpEvalCoord1xOES                              C.GPEVALCOORD1XOES
	gpEvalCoord1xvOES                             C.GPEVALCOORD1XVOES
	gpEvalCoord2xOES                              C.GPEVALCOORD2XOES
	gpEvalCoord2xvOES                             C.GPEVALCOORD2XVOES
	gpFeedbackBufferxOES                          C.GPFEEDBACKBUFFERXOES
	gpFenceSync                                   C.GPFENCESYNC
	gpFinish                                      C.GPFINISH
	gpFinishFenceNV                               C.GPFINISHFENCENV
	gpFlush                                       C.GPFLUSH
	gpFlushMappedBufferRange                      C.GPFLUSHMAPPEDBUFFERRANGE
	gpFlushMappedNamedBufferRange                 C.GPFLUSHMAPPEDNAMEDBUFFERRANGE
	gpFogxOES                                     C.GPFOGXOES
	gpFogxvOES                                    C.GPFOGXVOES
	gpFramebufferParameteri                       C.GPFRAMEBUFFERPARAMETERI
	gpFramebufferRenderbuffer                     C.GPFRAMEBUFFERRENDERBUFFER
	gpFramebufferTexture                          C.GPFRAMEBUFFERTEXTURE
	gpFramebufferTexture1D                        C.GPFRAMEBUFFERTEXTURE1D
	gpFramebufferTexture2D                        C.GPFRAMEBUFFERTEXTURE2D
	gpFramebufferTexture3D                        C.GPFRAMEBUFFERTEXTURE3D
	gpFramebufferTextureLayer                     C.GPFRAMEBUFFERTEXTURELAYER
	gpFrontFace                                   C.GPFRONTFACE
	gpFrustumfOES                                 C.GPFRUSTUMFOES
	gpFrustumxOES                                 C.GPFRUSTUMXOES
	gpGenBuffers                                  C.GPGENBUFFERS
	gpGenFencesNV                                 C.GPGENFENCESNV
	gpGenFramebuffers                             C.GPGENFRAMEBUFFERS
	gpGenPerfMonitorsAMD                          C.GPGENPERFMONITORSAMD
	gpGenProgramPipelines                         C.GPGENPROGRAMPIPELINES
	gpGenProgramPipelinesEXT                      C.GPGENPROGRAMPIPELINESEXT
	gpGenQueries                                  C.GPGENQUERIES
	gpGenRenderbuffers                            C.GPGENRENDERBUFFERS
	gpGenSamplers                                 C.GPGENSAMPLERS
	gpGenTextures                                 C.GPGENTEXTURES
	gpGenTransformFeedbacks                       C.GPGENTRANSFORMFEEDBACKS
	gpGenVertexArrays                             C.GPGENVERTEXARRAYS
	gpGenerateMipmap                              C.GPGENERATEMIPMAP
	gpGenerateTextureMipmap                       C.GPGENERATETEXTUREMIPMAP
	gpGetActiveAtomicCounterBufferiv              C.GPGETACTIVEATOMICCOUNTERBUFFERIV
	gpGetActiveAttrib                             C.GPGETACTIVEATTRIB
	gpGetActiveSubroutineName                     C.GPGETACTIVESUBROUTINENAME
	gpGetActiveSubroutineUniformName              C.GPGETACTIVESUBROUTINEUNIFORMNAME
	gpGetActiveSubroutineUniformiv                C.GPGETACTIVESUBROUTINEUNIFORMIV
	gpGetActiveUniform                            C.GPGETACTIVEUNIFORM
	gpGetActiveUniformBlockName                   C.GPGETACTIVEUNIFORMBLOCKNAME
	gpGetActiveUniformBlockiv                     C.GPGETACTIVEUNIFORMBLOCKIV
	gpGetActiveUniformName                        C.GPGETACTIVEUNIFORMNAME
	gpGetActiveUniformsiv                         C.GPGETACTIVEUNIFORMSIV
	gpGetAttachedShaders                          C.GPGETATTACHEDSHADERS
	gpGetAttribLocation                           C.GPGETATTRIBLOCATION
	gpGetBooleani_v                               C.GPGETBOOLEANI_V
	gpGetBooleanv                                 C.GPGETBOOLEANV
	gpGetBufferParameteri64v                      C.GPGETBUFFERPARAMETERI64V
	gpGetBufferParameteriv                        C.GPGETBUFFERPARAMETERIV
	gpGetBufferPointerv                           C.GPGETBUFFERPOINTERV
	gpGetBufferSubData                            C.GPGETBUFFERSUBDATA
	gpGetClipPlanefOES                            C.GPGETCLIPPLANEFOES
	gpGetClipPlanexOES                            C.GPGETCLIPPLANEXOES
	gpGetCompressedTexImage                       C.GPGETCOMPRESSEDTEXIMAGE
	gpGetCompressedTextureImage                   C.GPGETCOMPRESSEDTEXTUREIMAGE
	gpGetCompressedTextureSubImage                C.GPGETCOMPRESSEDTEXTURESUBIMAGE
	gpGetConvolutionParameterxvOES                C.GPGETCONVOLUTIONPARAMETERXVOES
	gpGetDebugMessageLog                          C.GPGETDEBUGMESSAGELOG
	gpGetDebugMessageLogARB                       C.GPGETDEBUGMESSAGELOGARB
	gpGetDebugMessageLogKHR                       C.GPGETDEBUGMESSAGELOGKHR
	gpGetDoublei_v                                C.GPGETDOUBLEI_V
	gpGetDoublev                                  C.GPGETDOUBLEV
	gpGetError                                    C.GPGETERROR
	gpGetFenceivNV                                C.GPGETFENCEIVNV
	gpGetFirstPerfQueryIdINTEL                    C.GPGETFIRSTPERFQUERYIDINTEL
	gpGetFixedvOES                                C.GPGETFIXEDVOES
	gpGetFloati_v                                 C.GPGETFLOATI_V
	gpGetFloatv                                   C.GPGETFLOATV
	gpGetFragDataIndex                            C.GPGETFRAGDATAINDEX
	gpGetFragDataLocation                         C.GPGETFRAGDATALOCATION
	gpGetFramebufferAttachmentParameteriv         C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV
	gpGetFramebufferParameteriv                   C.GPGETFRAMEBUFFERPARAMETERIV
	gpGetGraphicsResetStatus                      C.GPGETGRAPHICSRESETSTATUS
	gpGetGraphicsResetStatusARB                   C.GPGETGRAPHICSRESETSTATUSARB
	gpGetGraphicsResetStatusKHR                   C.GPGETGRAPHICSRESETSTATUSKHR
	gpGetHistogramParameterxvOES                  C.GPGETHISTOGRAMPARAMETERXVOES
	gpGetImageHandleARB                           C.GPGETIMAGEHANDLEARB
	gpGetInteger64i_v                             C.GPGETINTEGER64I_V
	gpGetInteger64v                               C.GPGETINTEGER64V
	gpGetIntegeri_v                               C.GPGETINTEGERI_V
	gpGetIntegerv                                 C.GPGETINTEGERV
	gpGetInternalformati64v                       C.GPGETINTERNALFORMATI64V
	gpGetInternalformativ                         C.GPGETINTERNALFORMATIV
	gpGetLightxOES                                C.GPGETLIGHTXOES
	gpGetLightxvOES                               C.GPGETLIGHTXVOES
	gpGetMapxvOES                                 C.GPGETMAPXVOES
	gpGetMaterialxOES                             C.GPGETMATERIALXOES
	gpGetMaterialxvOES                            C.GPGETMATERIALXVOES
	gpGetMultisamplefv                            C.GPGETMULTISAMPLEFV
	gpGetNamedBufferParameteri64v                 C.GPGETNAMEDBUFFERPARAMETERI64V
	gpGetNamedBufferParameteriv                   C.GPGETNAMEDBUFFERPARAMETERIV
	gpGetNamedBufferPointerv                      C.GPGETNAMEDBUFFERPOINTERV
	gpGetNamedBufferSubData                       C.GPGETNAMEDBUFFERSUBDATA
	gpGetNamedFramebufferAttachmentParameteriv    C.GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV
	gpGetNamedFramebufferParameteriv              C.GPGETNAMEDFRAMEBUFFERPARAMETERIV
	gpGetNamedRenderbufferParameteriv             C.GPGETNAMEDRENDERBUFFERPARAMETERIV
	gpGetNamedStringARB                           C.GPGETNAMEDSTRINGARB
	gpGetNamedStringivARB                         C.GPGETNAMEDSTRINGIVARB
	gpGetNextPerfQueryIdINTEL                     C.GPGETNEXTPERFQUERYIDINTEL
	gpGetObjectLabel                              C.GPGETOBJECTLABEL
	gpGetObjectLabelEXT                           C.GPGETOBJECTLABELEXT
	gpGetObjectLabelKHR                           C.GPGETOBJECTLABELKHR
	gpGetObjectPtrLabel                           C.GPGETOBJECTPTRLABEL
	gpGetObjectPtrLabelKHR                        C.GPGETOBJECTPTRLABELKHR
	gpGetPerfCounterInfoINTEL                     C.GPGETPERFCOUNTERINFOINTEL
	gpGetPerfMonitorCounterDataAMD                C.GPGETPERFMONITORCOUNTERDATAAMD
	gpGetPerfMonitorCounterInfoAMD                C.GPGETPERFMONITORCOUNTERINFOAMD
	gpGetPerfMonitorCounterStringAMD              C.GPGETPERFMONITORCOUNTERSTRINGAMD
	gpGetPerfMonitorCountersAMD                   C.GPGETPERFMONITORCOUNTERSAMD
	gpGetPerfMonitorGroupStringAMD                C.GPGETPERFMONITORGROUPSTRINGAMD
	gpGetPerfMonitorGroupsAMD                     C.GPGETPERFMONITORGROUPSAMD
	gpGetPerfQueryDataINTEL                       C.GPGETPERFQUERYDATAINTEL
	gpGetPerfQueryIdByNameINTEL                   C.GPGETPERFQUERYIDBYNAMEINTEL
	gpGetPerfQueryInfoINTEL                       C.GPGETPERFQUERYINFOINTEL
	gpGetPixelMapxv                               C.GPGETPIXELMAPXV
	gpGetPointerv                                 C.GPGETPOINTERV
	gpGetPointervKHR                              C.GPGETPOINTERVKHR
	gpGetProgramBinary                            C.GPGETPROGRAMBINARY
	gpGetProgramInfoLog                           C.GPGETPROGRAMINFOLOG
	gpGetProgramInterfaceiv                       C.GPGETPROGRAMINTERFACEIV
	gpGetProgramPipelineInfoLog                   C.GPGETPROGRAMPIPELINEINFOLOG
	gpGetProgramPipelineInfoLogEXT                C.GPGETPROGRAMPIPELINEINFOLOGEXT
	gpGetProgramPipelineiv                        C.GPGETPROGRAMPIPELINEIV
	gpGetProgramPipelineivEXT                     C.GPGETPROGRAMPIPELINEIVEXT
	gpGetProgramResourceIndex                     C.GPGETPROGRAMRESOURCEINDEX
	gpGetProgramResourceLocation                  C.GPGETPROGRAMRESOURCELOCATION
	gpGetProgramResourceLocationIndex             C.GPGETPROGRAMRESOURCELOCATIONINDEX
	gpGetProgramResourceName                      C.GPGETPROGRAMRESOURCENAME
	gpGetProgramResourceiv                        C.GPGETPROGRAMRESOURCEIV
	gpGetProgramStageiv                           C.GPGETPROGRAMSTAGEIV
	gpGetProgramiv                                C.GPGETPROGRAMIV
	gpGetQueryIndexediv                           C.GPGETQUERYINDEXEDIV
	gpGetQueryObjecti64v                          C.GPGETQUERYOBJECTI64V
	gpGetQueryObjectiv                            C.GPGETQUERYOBJECTIV
	gpGetQueryObjectui64v                         C.GPGETQUERYOBJECTUI64V
	gpGetQueryObjectuiv                           C.GPGETQUERYOBJECTUIV
	gpGetQueryiv                                  C.GPGETQUERYIV
	gpGetRenderbufferParameteriv                  C.GPGETRENDERBUFFERPARAMETERIV
	gpGetSamplerParameterIiv                      C.GPGETSAMPLERPARAMETERIIV
	gpGetSamplerParameterIuiv                     C.GPGETSAMPLERPARAMETERIUIV
	gpGetSamplerParameterfv                       C.GPGETSAMPLERPARAMETERFV
	gpGetSamplerParameteriv                       C.GPGETSAMPLERPARAMETERIV
	gpGetShaderInfoLog                            C.GPGETSHADERINFOLOG
	gpGetShaderPrecisionFormat                    C.GPGETSHADERPRECISIONFORMAT
	gpGetShaderSource                             C.GPGETSHADERSOURCE
	gpGetShaderiv                                 C.GPGETSHADERIV
	gpGetString                                   C.GPGETSTRING
	gpGetStringi                                  C.GPGETSTRINGI
	gpGetSubroutineIndex                          C.GPGETSUBROUTINEINDEX
	gpGetSubroutineUniformLocation                C.GPGETSUBROUTINEUNIFORMLOCATION
	gpGetSynciv                                   C.GPGETSYNCIV
	gpGetTexEnvxvOES                              C.GPGETTEXENVXVOES
	gpGetTexGenxvOES                              C.GPGETTEXGENXVOES
	gpGetTexImage                                 C.GPGETTEXIMAGE
	gpGetTexLevelParameterfv                      C.GPGETTEXLEVELPARAMETERFV
	gpGetTexLevelParameteriv                      C.GPGETTEXLEVELPARAMETERIV
	gpGetTexLevelParameterxvOES                   C.GPGETTEXLEVELPARAMETERXVOES
	gpGetTexParameterIiv                          C.GPGETTEXPARAMETERIIV
	gpGetTexParameterIuiv                         C.GPGETTEXPARAMETERIUIV
	gpGetTexParameterfv                           C.GPGETTEXPARAMETERFV
	gpGetTexParameteriv                           C.GPGETTEXPARAMETERIV
	gpGetTexParameterxvOES                        C.GPGETTEXPARAMETERXVOES
	gpGetTextureHandleARB                         C.GPGETTEXTUREHANDLEARB
	gpGetTextureImage                             C.GPGETTEXTUREIMAGE
	gpGetTextureLevelParameterfv                  C.GPGETTEXTURELEVELPARAMETERFV
	gpGetTextureLevelParameteriv                  C.GPGETTEXTURELEVELPARAMETERIV
	gpGetTextureParameterIiv                      C.GPGETTEXTUREPARAMETERIIV
	gpGetTextureParameterIuiv                     C.GPGETTEXTUREPARAMETERIUIV
	gpGetTextureParameterfv                       C.GPGETTEXTUREPARAMETERFV
	gpGetTextureParameteriv                       C.GPGETTEXTUREPARAMETERIV
	gpGetTextureSamplerHandleARB                  C.GPGETTEXTURESAMPLERHANDLEARB
	gpGetTextureSubImage                          C.GPGETTEXTURESUBIMAGE
	gpGetTransformFeedbackVarying                 C.GPGETTRANSFORMFEEDBACKVARYING
	gpGetTransformFeedbacki64_v                   C.GPGETTRANSFORMFEEDBACKI64_V
	gpGetTransformFeedbacki_v                     C.GPGETTRANSFORMFEEDBACKI_V
	gpGetTransformFeedbackiv                      C.GPGETTRANSFORMFEEDBACKIV
	gpGetUniformBlockIndex                        C.GPGETUNIFORMBLOCKINDEX
	gpGetUniformIndices                           C.GPGETUNIFORMINDICES
	gpGetUniformLocation                          C.GPGETUNIFORMLOCATION
	gpGetUniformSubroutineuiv                     C.GPGETUNIFORMSUBROUTINEUIV
	gpGetUniformdv                                C.GPGETUNIFORMDV
	gpGetUniformfv                                C.GPGETUNIFORMFV
	gpGetUniformiv                                C.GPGETUNIFORMIV
	gpGetUniformuiv                               C.GPGETUNIFORMUIV
	gpGetVertexArrayIndexed64iv                   C.GPGETVERTEXARRAYINDEXED64IV
	gpGetVertexArrayIndexediv                     C.GPGETVERTEXARRAYINDEXEDIV
	gpGetVertexArrayiv                            C.GPGETVERTEXARRAYIV
	gpGetVertexAttribIiv                          C.GPGETVERTEXATTRIBIIV
	gpGetVertexAttribIuiv                         C.GPGETVERTEXATTRIBIUIV
	gpGetVertexAttribLdv                          C.GPGETVERTEXATTRIBLDV
	gpGetVertexAttribLui64vARB                    C.GPGETVERTEXATTRIBLUI64VARB
	gpGetVertexAttribPointerv                     C.GPGETVERTEXATTRIBPOINTERV
	gpGetVertexAttribdv                           C.GPGETVERTEXATTRIBDV
	gpGetVertexAttribfv                           C.GPGETVERTEXATTRIBFV
	gpGetVertexAttribiv                           C.GPGETVERTEXATTRIBIV
	gpGetnCompressedTexImageARB                   C.GPGETNCOMPRESSEDTEXIMAGEARB
	gpGetnTexImageARB                             C.GPGETNTEXIMAGEARB
	gpGetnUniformdvARB                            C.GPGETNUNIFORMDVARB
	gpGetnUniformfv                               C.GPGETNUNIFORMFV
	gpGetnUniformfvARB                            C.GPGETNUNIFORMFVARB
	gpGetnUniformfvKHR                            C.GPGETNUNIFORMFVKHR
	gpGetnUniformiv                               C.GPGETNUNIFORMIV
	gpGetnUniformivARB                            C.GPGETNUNIFORMIVARB
	gpGetnUniformivKHR                            C.GPGETNUNIFORMIVKHR
	gpGetnUniformuiv                              C.GPGETNUNIFORMUIV
	gpGetnUniformuivARB                           C.GPGETNUNIFORMUIVARB
	gpGetnUniformuivKHR                           C.GPGETNUNIFORMUIVKHR
	gpHint                                        C.GPHINT
	gpIndexxOES                                   C.GPINDEXXOES
	gpIndexxvOES                                  C.GPINDEXXVOES
	gpInsertEventMarkerEXT                        C.GPINSERTEVENTMARKEREXT
	gpInvalidateBufferData                        C.GPINVALIDATEBUFFERDATA
	gpInvalidateBufferSubData                     C.GPINVALIDATEBUFFERSUBDATA
	gpInvalidateFramebuffer                       C.GPINVALIDATEFRAMEBUFFER
	gpInvalidateNamedFramebufferData              C.GPINVALIDATENAMEDFRAMEBUFFERDATA
	gpInvalidateNamedFramebufferSubData           C.GPINVALIDATENAMEDFRAMEBUFFERSUBDATA
	gpInvalidateSubFramebuffer                    C.GPINVALIDATESUBFRAMEBUFFER
	gpInvalidateTexImage                          C.GPINVALIDATETEXIMAGE
	gpInvalidateTexSubImage                       C.GPINVALIDATETEXSUBIMAGE
	gpIsBuffer                                    C.GPISBUFFER
	gpIsEnabled                                   C.GPISENABLED
	gpIsEnabledi                                  C.GPISENABLEDI
	gpIsFenceNV                                   C.GPISFENCENV
	gpIsFramebuffer                               C.GPISFRAMEBUFFER
	gpIsImageHandleResidentARB                    C.GPISIMAGEHANDLERESIDENTARB
	gpIsNamedStringARB                            C.GPISNAMEDSTRINGARB
	gpIsProgram                                   C.GPISPROGRAM
	gpIsProgramPipeline                           C.GPISPROGRAMPIPELINE
	gpIsProgramPipelineEXT                        C.GPISPROGRAMPIPELINEEXT
	gpIsQuery                                     C.GPISQUERY
	gpIsRenderbuffer                              C.GPISRENDERBUFFER
	gpIsSampler                                   C.GPISSAMPLER
	gpIsShader                                    C.GPISSHADER
	gpIsSync                                      C.GPISSYNC
	gpIsTexture                                   C.GPISTEXTURE
	gpIsTextureHandleResidentARB                  C.GPISTEXTUREHANDLERESIDENTARB
	gpIsTransformFeedback                         C.GPISTRANSFORMFEEDBACK
	gpIsVertexArray                               C.GPISVERTEXARRAY
	gpLabelObjectEXT                              C.GPLABELOBJECTEXT
	gpLightModelxOES                              C.GPLIGHTMODELXOES
	gpLightModelxvOES                             C.GPLIGHTMODELXVOES
	gpLightxOES                                   C.GPLIGHTXOES
	gpLightxvOES                                  C.GPLIGHTXVOES
	gpLineWidth                                   C.GPLINEWIDTH
	gpLineWidthxOES                               C.GPLINEWIDTHXOES
	gpLinkProgram                                 C.GPLINKPROGRAM
	gpLoadMatrixxOES                              C.GPLOADMATRIXXOES
	gpLoadTransposeMatrixxOES                     C.GPLOADTRANSPOSEMATRIXXOES
	gpLogicOp                                     C.GPLOGICOP
	gpMakeImageHandleNonResidentARB               C.GPMAKEIMAGEHANDLENONRESIDENTARB
	gpMakeImageHandleResidentARB                  C.GPMAKEIMAGEHANDLERESIDENTARB
	gpMakeTextureHandleNonResidentARB             C.GPMAKETEXTUREHANDLENONRESIDENTARB
	gpMakeTextureHandleResidentARB                C.GPMAKETEXTUREHANDLERESIDENTARB
	gpMap1xOES                                    C.GPMAP1XOES
	gpMap2xOES                                    C.GPMAP2XOES
	gpMapBuffer                                   C.GPMAPBUFFER
	gpMapBufferRange                              C.GPMAPBUFFERRANGE
	gpMapGrid1xOES                                C.GPMAPGRID1XOES
	gpMapGrid2xOES                                C.GPMAPGRID2XOES
	gpMapNamedBuffer                              C.GPMAPNAMEDBUFFER
	gpMapNamedBufferRange                         C.GPMAPNAMEDBUFFERRANGE
	gpMaterialxOES                                C.GPMATERIALXOES
	gpMaterialxvOES                               C.GPMATERIALXVOES
	gpMemoryBarrier                               C.GPMEMORYBARRIER
	gpMemoryBarrierByRegion                       C.GPMEMORYBARRIERBYREGION
	gpMinSampleShading                            C.GPMINSAMPLESHADING
	gpMinSampleShadingARB                         C.GPMINSAMPLESHADINGARB
	gpMultMatrixxOES                              C.GPMULTMATRIXXOES
	gpMultTransposeMatrixxOES                     C.GPMULTTRANSPOSEMATRIXXOES
	gpMultiDrawArrays                             C.GPMULTIDRAWARRAYS
	gpMultiDrawArraysEXT                          C.GPMULTIDRAWARRAYSEXT
	gpMultiDrawArraysIndirect                     C.GPMULTIDRAWARRAYSINDIRECT
	gpMultiDrawArraysIndirectCountARB             C.GPMULTIDRAWARRAYSINDIRECTCOUNTARB
	gpMultiDrawElements                           C.GPMULTIDRAWELEMENTS
	gpMultiDrawElementsBaseVertex                 C.GPMULTIDRAWELEMENTSBASEVERTEX
	gpMultiDrawElementsEXT                        C.GPMULTIDRAWELEMENTSEXT
	gpMultiDrawElementsIndirect                   C.GPMULTIDRAWELEMENTSINDIRECT
	gpMultiDrawElementsIndirectCountARB           C.GPMULTIDRAWELEMENTSINDIRECTCOUNTARB
	gpMultiTexCoord1bOES                          C.GPMULTITEXCOORD1BOES
	gpMultiTexCoord1bvOES                         C.GPMULTITEXCOORD1BVOES
	gpMultiTexCoord1xOES                          C.GPMULTITEXCOORD1XOES
	gpMultiTexCoord1xvOES                         C.GPMULTITEXCOORD1XVOES
	gpMultiTexCoord2bOES                          C.GPMULTITEXCOORD2BOES
	gpMultiTexCoord2bvOES                         C.GPMULTITEXCOORD2BVOES
	gpMultiTexCoord2xOES                          C.GPMULTITEXCOORD2XOES
	gpMultiTexCoord2xvOES                         C.GPMULTITEXCOORD2XVOES
	gpMultiTexCoord3bOES                          C.GPMULTITEXCOORD3BOES
	gpMultiTexCoord3bvOES                         C.GPMULTITEXCOORD3BVOES
	gpMultiTexCoord3xOES                          C.GPMULTITEXCOORD3XOES
	gpMultiTexCoord3xvOES                         C.GPMULTITEXCOORD3XVOES
	gpMultiTexCoord4bOES                          C.GPMULTITEXCOORD4BOES
	gpMultiTexCoord4bvOES                         C.GPMULTITEXCOORD4BVOES
	gpMultiTexCoord4xOES                          C.GPMULTITEXCOORD4XOES
	gpMultiTexCoord4xvOES                         C.GPMULTITEXCOORD4XVOES
	gpNamedBufferData                             C.GPNAMEDBUFFERDATA
	gpNamedBufferPageCommitmentARB                C.GPNAMEDBUFFERPAGECOMMITMENTARB
	gpNamedBufferPageCommitmentEXT                C.GPNAMEDBUFFERPAGECOMMITMENTEXT
	gpNamedBufferStorage                          C.GPNAMEDBUFFERSTORAGE
	gpNamedBufferSubData                          C.GPNAMEDBUFFERSUBDATA
	gpNamedFramebufferDrawBuffer                  C.GPNAMEDFRAMEBUFFERDRAWBUFFER
	gpNamedFramebufferDrawBuffers                 C.GPNAMEDFRAMEBUFFERDRAWBUFFERS
	gpNamedFramebufferParameteri                  C.GPNAMEDFRAMEBUFFERPARAMETERI
	gpNamedFramebufferReadBuffer                  C.GPNAMEDFRAMEBUFFERREADBUFFER
	gpNamedFramebufferRenderbuffer                C.GPNAMEDFRAMEBUFFERRENDERBUFFER
	gpNamedFramebufferTexture                     C.GPNAMEDFRAMEBUFFERTEXTURE
	gpNamedFramebufferTextureLayer                C.GPNAMEDFRAMEBUFFERTEXTURELAYER
	gpNamedRenderbufferStorage                    C.GPNAMEDRENDERBUFFERSTORAGE
	gpNamedRenderbufferStorageMultisample         C.GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE
	gpNamedStringARB                              C.GPNAMEDSTRINGARB
	gpNormal3xOES                                 C.GPNORMAL3XOES
	gpNormal3xvOES                                C.GPNORMAL3XVOES
	gpObjectLabel                                 C.GPOBJECTLABEL
	gpObjectLabelKHR                              C.GPOBJECTLABELKHR
	gpObjectPtrLabel                              C.GPOBJECTPTRLABEL
	gpObjectPtrLabelKHR                           C.GPOBJECTPTRLABELKHR
	gpOrthofOES                                   C.GPORTHOFOES
	gpOrthoxOES                                   C.GPORTHOXOES
	gpPassThroughxOES                             C.GPPASSTHROUGHXOES
	gpPatchParameterfv                            C.GPPATCHPARAMETERFV
	gpPatchParameteri                             C.GPPATCHPARAMETERI
	gpPauseTransformFeedback                      C.GPPAUSETRANSFORMFEEDBACK
	gpPixelMapx                                   C.GPPIXELMAPX
	gpPixelStoref                                 C.GPPIXELSTOREF
	gpPixelStorei                                 C.GPPIXELSTOREI
	gpPixelStorex                                 C.GPPIXELSTOREX
	gpPixelTransferxOES                           C.GPPIXELTRANSFERXOES
	gpPixelZoomxOES                               C.GPPIXELZOOMXOES
	gpPointParameterf                             C.GPPOINTPARAMETERF
	gpPointParameterfv                            C.GPPOINTPARAMETERFV
	gpPointParameteri                             C.GPPOINTPARAMETERI
	gpPointParameteriv                            C.GPPOINTPARAMETERIV
	gpPointParameterxOES                          C.GPPOINTPARAMETERXOES
	gpPointParameterxvOES                         C.GPPOINTPARAMETERXVOES
	gpPointSize                                   C.GPPOINTSIZE
	gpPointSizexOES                               C.GPPOINTSIZEXOES
	gpPolygonMode                                 C.GPPOLYGONMODE
	gpPolygonOffset                               C.GPPOLYGONOFFSET
	gpPolygonOffsetxOES                           C.GPPOLYGONOFFSETXOES
	gpPopDebugGroup                               C.GPPOPDEBUGGROUP
	gpPopDebugGroupKHR                            C.GPPOPDEBUGGROUPKHR
	gpPopGroupMarkerEXT                           C.GPPOPGROUPMARKEREXT
	gpPrimitiveRestartIndex                       C.GPPRIMITIVERESTARTINDEX
	gpPrioritizeTexturesxOES                      C.GPPRIORITIZETEXTURESXOES
	gpProgramBinary                               C.GPPROGRAMBINARY
	gpProgramParameteri                           C.GPPROGRAMPARAMETERI
	gpProgramParameteriEXT                        C.GPPROGRAMPARAMETERIEXT
	gpProgramUniform1d                            C.GPPROGRAMUNIFORM1D
	gpProgramUniform1dv                           C.GPPROGRAMUNIFORM1DV
	gpProgramUniform1f                            C.GPPROGRAMUNIFORM1F
	gpProgramUniform1fEXT                         C.GPPROGRAMUNIFORM1FEXT
	gpProgramUniform1fv                           C.GPPROGRAMUNIFORM1FV
	gpProgramUniform1fvEXT                        C.GPPROGRAMUNIFORM1FVEXT
	gpProgramUniform1i                            C.GPPROGRAMUNIFORM1I
	gpProgramUniform1iEXT                         C.GPPROGRAMUNIFORM1IEXT
	gpProgramUniform1iv                           C.GPPROGRAMUNIFORM1IV
	gpProgramUniform1ivEXT                        C.GPPROGRAMUNIFORM1IVEXT
	gpProgramUniform1ui                           C.GPPROGRAMUNIFORM1UI
	gpProgramUniform1uiEXT                        C.GPPROGRAMUNIFORM1UIEXT
	gpProgramUniform1uiv                          C.GPPROGRAMUNIFORM1UIV
	gpProgramUniform1uivEXT                       C.GPPROGRAMUNIFORM1UIVEXT
	gpProgramUniform2d                            C.GPPROGRAMUNIFORM2D
	gpProgramUniform2dv                           C.GPPROGRAMUNIFORM2DV
	gpProgramUniform2f                            C.GPPROGRAMUNIFORM2F
	gpProgramUniform2fEXT                         C.GPPROGRAMUNIFORM2FEXT
	gpProgramUniform2fv                           C.GPPROGRAMUNIFORM2FV
	gpProgramUniform2fvEXT                        C.GPPROGRAMUNIFORM2FVEXT
	gpProgramUniform2i                            C.GPPROGRAMUNIFORM2I
	gpProgramUniform2iEXT                         C.GPPROGRAMUNIFORM2IEXT
	gpProgramUniform2iv                           C.GPPROGRAMUNIFORM2IV
	gpProgramUniform2ivEXT                        C.GPPROGRAMUNIFORM2IVEXT
	gpProgramUniform2ui                           C.GPPROGRAMUNIFORM2UI
	gpProgramUniform2uiEXT                        C.GPPROGRAMUNIFORM2UIEXT
	gpProgramUniform2uiv                          C.GPPROGRAMUNIFORM2UIV
	gpProgramUniform2uivEXT                       C.GPPROGRAMUNIFORM2UIVEXT
	gpProgramUniform3d                            C.GPPROGRAMUNIFORM3D
	gpProgramUniform3dv                           C.GPPROGRAMUNIFORM3DV
	gpProgramUniform3f                            C.GPPROGRAMUNIFORM3F
	gpProgramUniform3fEXT                         C.GPPROGRAMUNIFORM3FEXT
	gpProgramUniform3fv                           C.GPPROGRAMUNIFORM3FV
	gpProgramUniform3fvEXT                        C.GPPROGRAMUNIFORM3FVEXT
	gpProgramUniform3i                            C.GPPROGRAMUNIFORM3I
	gpProgramUniform3iEXT                         C.GPPROGRAMUNIFORM3IEXT
	gpProgramUniform3iv                           C.GPPROGRAMUNIFORM3IV
	gpProgramUniform3ivEXT                        C.GPPROGRAMUNIFORM3IVEXT
	gpProgramUniform3ui                           C.GPPROGRAMUNIFORM3UI
	gpProgramUniform3uiEXT                        C.GPPROGRAMUNIFORM3UIEXT
	gpProgramUniform3uiv                          C.GPPROGRAMUNIFORM3UIV
	gpProgramUniform3uivEXT                       C.GPPROGRAMUNIFORM3UIVEXT
	gpProgramUniform4d                            C.GPPROGRAMUNIFORM4D
	gpProgramUniform4dv                           C.GPPROGRAMUNIFORM4DV
	gpProgramUniform4f                            C.GPPROGRAMUNIFORM4F
	gpProgramUniform4fEXT                         C.GPPROGRAMUNIFORM4FEXT
	gpProgramUniform4fv                           C.GPPROGRAMUNIFORM4FV
	gpProgramUniform4fvEXT                        C.GPPROGRAMUNIFORM4FVEXT
	gpProgramUniform4i                            C.GPPROGRAMUNIFORM4I
	gpProgramUniform4iEXT                         C.GPPROGRAMUNIFORM4IEXT
	gpProgramUniform4iv                           C.GPPROGRAMUNIFORM4IV
	gpProgramUniform4ivEXT                        C.GPPROGRAMUNIFORM4IVEXT
	gpProgramUniform4ui                           C.GPPROGRAMUNIFORM4UI
	gpProgramUniform4uiEXT                        C.GPPROGRAMUNIFORM4UIEXT
	gpProgramUniform4uiv                          C.GPPROGRAMUNIFORM4UIV
	gpProgramUniform4uivEXT                       C.GPPROGRAMUNIFORM4UIVEXT
	gpProgramUniformHandleui64ARB                 C.GPPROGRAMUNIFORMHANDLEUI64ARB
	gpProgramUniformHandleui64vARB                C.GPPROGRAMUNIFORMHANDLEUI64VARB
	gpProgramUniformMatrix2dv                     C.GPPROGRAMUNIFORMMATRIX2DV
	gpProgramUniformMatrix2fv                     C.GPPROGRAMUNIFORMMATRIX2FV
	gpProgramUniformMatrix2fvEXT                  C.GPPROGRAMUNIFORMMATRIX2FVEXT
	gpProgramUniformMatrix2x3dv                   C.GPPROGRAMUNIFORMMATRIX2X3DV
	gpProgramUniformMatrix2x3fv                   C.GPPROGRAMUNIFORMMATRIX2X3FV
	gpProgramUniformMatrix2x3fvEXT                C.GPPROGRAMUNIFORMMATRIX2X3FVEXT
	gpProgramUniformMatrix2x4dv                   C.GPPROGRAMUNIFORMMATRIX2X4DV
	gpProgramUniformMatrix2x4fv                   C.GPPROGRAMUNIFORMMATRIX2X4FV
	gpProgramUniformMatrix2x4fvEXT                C.GPPROGRAMUNIFORMMATRIX2X4FVEXT
	gpProgramUniformMatrix3dv                     C.GPPROGRAMUNIFORMMATRIX3DV
	gpProgramUniformMatrix3fv                     C.GPPROGRAMUNIFORMMATRIX3FV
	gpProgramUniformMatrix3fvEXT                  C.GPPROGRAMUNIFORMMATRIX3FVEXT
	gpProgramUniformMatrix3x2dv                   C.GPPROGRAMUNIFORMMATRIX3X2DV
	gpProgramUniformMatrix3x2fv                   C.GPPROGRAMUNIFORMMATRIX3X2FV
	gpProgramUniformMatrix3x2fvEXT                C.GPPROGRAMUNIFORMMATRIX3X2FVEXT
	gpProgramUniformMatrix3x4dv                   C.GPPROGRAMUNIFORMMATRIX3X4DV
	gpProgramUniformMatrix3x4fv                   C.GPPROGRAMUNIFORMMATRIX3X4FV
	gpProgramUniformMatrix3x4fvEXT                C.GPPROGRAMUNIFORMMATRIX3X4FVEXT
	gpProgramUniformMatrix4dv                     C.GPPROGRAMUNIFORMMATRIX4DV
	gpProgramUniformMatrix4fv                     C.GPPROGRAMUNIFORMMATRIX4FV
	gpProgramUniformMatrix4fvEXT                  C.GPPROGRAMUNIFORMMATRIX4FVEXT
	gpProgramUniformMatrix4x2dv                   C.GPPROGRAMUNIFORMMATRIX4X2DV
	gpProgramUniformMatrix4x2fv                   C.GPPROGRAMUNIFORMMATRIX4X2FV
	gpProgramUniformMatrix4x2fvEXT                C.GPPROGRAMUNIFORMMATRIX4X2FVEXT
	gpProgramUniformMatrix4x3dv                   C.GPPROGRAMUNIFORMMATRIX4X3DV
	gpProgramUniformMatrix4x3fv                   C.GPPROGRAMUNIFORMMATRIX4X3FV
	gpProgramUniformMatrix4x3fvEXT                C.GPPROGRAMUNIFORMMATRIX4X3FVEXT
	gpProvokingVertex                             C.GPPROVOKINGVERTEX
	gpPushDebugGroup                              C.GPPUSHDEBUGGROUP
	gpPushDebugGroupKHR                           C.GPPUSHDEBUGGROUPKHR
	gpPushGroupMarkerEXT                          C.GPPUSHGROUPMARKEREXT
	gpQueryCounter                                C.GPQUERYCOUNTER
	gpQueryMatrixxOES                             C.GPQUERYMATRIXXOES
	gpRasterPos2xOES                              C.GPRASTERPOS2XOES
	gpRasterPos2xvOES                             C.GPRASTERPOS2XVOES
	gpRasterPos3xOES                              C.GPRASTERPOS3XOES
	gpRasterPos3xvOES                             C.GPRASTERPOS3XVOES
	gpRasterPos4xOES                              C.GPRASTERPOS4XOES
	gpRasterPos4xvOES                             C.GPRASTERPOS4XVOES
	gpReadBuffer                                  C.GPREADBUFFER
	gpReadPixels                                  C.GPREADPIXELS
	gpReadnPixels                                 C.GPREADNPIXELS
	gpReadnPixelsARB                              C.GPREADNPIXELSARB
	gpReadnPixelsKHR                              C.GPREADNPIXELSKHR
	gpRectxOES                                    C.GPRECTXOES
	gpRectxvOES                                   C.GPRECTXVOES
	gpReleaseShaderCompiler                       C.GPRELEASESHADERCOMPILER
	gpRenderbufferStorage                         C.GPRENDERBUFFERSTORAGE
	gpRenderbufferStorageMultisample              C.GPRENDERBUFFERSTORAGEMULTISAMPLE
	gpResumeTransformFeedback                     C.GPRESUMETRANSFORMFEEDBACK
	gpRotatexOES                                  C.GPROTATEXOES
	gpSampleCoverage                              C.GPSAMPLECOVERAGE
	gpSampleCoverageOES                           C.GPSAMPLECOVERAGEOES
	gpSampleCoveragexOES                          C.GPSAMPLECOVERAGEXOES
	gpSampleMaski                                 C.GPSAMPLEMASKI
	gpSamplerParameterIiv                         C.GPSAMPLERPARAMETERIIV
	gpSamplerParameterIuiv                        C.GPSAMPLERPARAMETERIUIV
	gpSamplerParameterf                           C.GPSAMPLERPARAMETERF
	gpSamplerParameterfv                          C.GPSAMPLERPARAMETERFV
	gpSamplerParameteri                           C.GPSAMPLERPARAMETERI
	gpSamplerParameteriv                          C.GPSAMPLERPARAMETERIV
	gpScalexOES                                   C.GPSCALEXOES
	gpScissor                                     C.GPSCISSOR
	gpScissorArrayv                               C.GPSCISSORARRAYV
	gpScissorIndexed                              C.GPSCISSORINDEXED
	gpScissorIndexedv                             C.GPSCISSORINDEXEDV
	gpSelectPerfMonitorCountersAMD                C.GPSELECTPERFMONITORCOUNTERSAMD
	gpSetFenceNV                                  C.GPSETFENCENV
	gpShaderBinary                                C.GPSHADERBINARY
	gpShaderSource                                C.GPSHADERSOURCE
	gpShaderStorageBlockBinding                   C.GPSHADERSTORAGEBLOCKBINDING
	gpStencilFunc                                 C.GPSTENCILFUNC
	gpStencilFuncSeparate                         C.GPSTENCILFUNCSEPARATE
	gpStencilMask                                 C.GPSTENCILMASK
	gpStencilMaskSeparate                         C.GPSTENCILMASKSEPARATE
	gpStencilOp                                   C.GPSTENCILOP
	gpStencilOpSeparate                           C.GPSTENCILOPSEPARATE
	gpTestFenceNV                                 C.GPTESTFENCENV
	gpTexBuffer                                   C.GPTEXBUFFER
	gpTexBufferRange                              C.GPTEXBUFFERRANGE
	gpTexCoord1bOES                               C.GPTEXCOORD1BOES
	gpTexCoord1bvOES                              C.GPTEXCOORD1BVOES
	gpTexCoord1xOES                               C.GPTEXCOORD1XOES
	gpTexCoord1xvOES                              C.GPTEXCOORD1XVOES
	gpTexCoord2bOES                               C.GPTEXCOORD2BOES
	gpTexCoord2bvOES                              C.GPTEXCOORD2BVOES
	gpTexCoord2xOES                               C.GPTEXCOORD2XOES
	gpTexCoord2xvOES                              C.GPTEXCOORD2XVOES
	gpTexCoord3bOES                               C.GPTEXCOORD3BOES
	gpTexCoord3bvOES                              C.GPTEXCOORD3BVOES
	gpTexCoord3xOES                               C.GPTEXCOORD3XOES
	gpTexCoord3xvOES                              C.GPTEXCOORD3XVOES
	gpTexCoord4bOES                               C.GPTEXCOORD4BOES
	gpTexCoord4bvOES                              C.GPTEXCOORD4BVOES
	gpTexCoord4xOES                               C.GPTEXCOORD4XOES
	gpTexCoord4xvOES                              C.GPTEXCOORD4XVOES
	gpTexEnvxOES                                  C.GPTEXENVXOES
	gpTexEnvxvOES                                 C.GPTEXENVXVOES
	gpTexGenxOES                                  C.GPTEXGENXOES
	gpTexGenxvOES                                 C.GPTEXGENXVOES
	gpTexImage1D                                  C.GPTEXIMAGE1D
	gpTexImage2D                                  C.GPTEXIMAGE2D
	gpTexImage2DMultisample                       C.GPTEXIMAGE2DMULTISAMPLE
	gpTexImage3D                                  C.GPTEXIMAGE3D
	gpTexImage3DMultisample                       C.GPTEXIMAGE3DMULTISAMPLE
	gpTexPageCommitmentARB                        C.GPTEXPAGECOMMITMENTARB
	gpTexParameterIiv                             C.GPTEXPARAMETERIIV
	gpTexParameterIuiv                            C.GPTEXPARAMETERIUIV
	gpTexParameterf                               C.GPTEXPARAMETERF
	gpTexParameterfv                              C.GPTEXPARAMETERFV
	gpTexParameteri                               C.GPTEXPARAMETERI
	gpTexParameteriv                              C.GPTEXPARAMETERIV
	gpTexParameterxOES                            C.GPTEXPARAMETERXOES
	gpTexParameterxvOES                           C.GPTEXPARAMETERXVOES
	gpTexStorage1D                                C.GPTEXSTORAGE1D
	gpTexStorage2D                                C.GPTEXSTORAGE2D
	gpTexStorage2DMultisample                     C.GPTEXSTORAGE2DMULTISAMPLE
	gpTexStorage3D                                C.GPTEXSTORAGE3D
	gpTexStorage3DMultisample                     C.GPTEXSTORAGE3DMULTISAMPLE
	gpTexSubImage1D                               C.GPTEXSUBIMAGE1D
	gpTexSubImage2D                               C.GPTEXSUBIMAGE2D
	gpTexSubImage3D                               C.GPTEXSUBIMAGE3D
	gpTextureBarrier                              C.GPTEXTUREBARRIER
	gpTextureBuffer                               C.GPTEXTUREBUFFER
	gpTextureBufferRange                          C.GPTEXTUREBUFFERRANGE
	gpTextureParameterIiv                         C.GPTEXTUREPARAMETERIIV
	gpTextureParameterIuiv                        C.GPTEXTUREPARAMETERIUIV
	gpTextureParameterf                           C.GPTEXTUREPARAMETERF
	gpTextureParameterfv                          C.GPTEXTUREPARAMETERFV
	gpTextureParameteri                           C.GPTEXTUREPARAMETERI
	gpTextureParameteriv                          C.GPTEXTUREPARAMETERIV
	gpTextureStorage1D                            C.GPTEXTURESTORAGE1D
	gpTextureStorage2D                            C.GPTEXTURESTORAGE2D
	gpTextureStorage2DMultisample                 C.GPTEXTURESTORAGE2DMULTISAMPLE
	gpTextureStorage3D                            C.GPTEXTURESTORAGE3D
	gpTextureStorage3DMultisample                 C.GPTEXTURESTORAGE3DMULTISAMPLE
	gpTextureSubImage1D                           C.GPTEXTURESUBIMAGE1D
	gpTextureSubImage2D                           C.GPTEXTURESUBIMAGE2D
	gpTextureSubImage3D                           C.GPTEXTURESUBIMAGE3D
	gpTextureView                                 C.GPTEXTUREVIEW
	gpTransformFeedbackBufferBase                 C.GPTRANSFORMFEEDBACKBUFFERBASE
	gpTransformFeedbackBufferRange                C.GPTRANSFORMFEEDBACKBUFFERRANGE
	gpTransformFeedbackVaryings                   C.GPTRANSFORMFEEDBACKVARYINGS
	gpTranslatexOES                               C.GPTRANSLATEXOES
	gpUniform1d                                   C.GPUNIFORM1D
	gpUniform1dv                                  C.GPUNIFORM1DV
	gpUniform1f                                   C.GPUNIFORM1F
	gpUniform1fv                                  C.GPUNIFORM1FV
	gpUniform1i                                   C.GPUNIFORM1I
	gpUniform1iv                                  C.GPUNIFORM1IV
	gpUniform1ui                                  C.GPUNIFORM1UI
	gpUniform1uiv                                 C.GPUNIFORM1UIV
	gpUniform2d                                   C.GPUNIFORM2D
	gpUniform2dv                                  C.GPUNIFORM2DV
	gpUniform2f                                   C.GPUNIFORM2F
	gpUniform2fv                                  C.GPUNIFORM2FV
	gpUniform2i                                   C.GPUNIFORM2I
	gpUniform2iv                                  C.GPUNIFORM2IV
	gpUniform2ui                                  C.GPUNIFORM2UI
	gpUniform2uiv                                 C.GPUNIFORM2UIV
	gpUniform3d                                   C.GPUNIFORM3D
	gpUniform3dv                                  C.GPUNIFORM3DV
	gpUniform3f                                   C.GPUNIFORM3F
	gpUniform3fv                                  C.GPUNIFORM3FV
	gpUniform3i                                   C.GPUNIFORM3I
	gpUniform3iv                                  C.GPUNIFORM3IV
	gpUniform3ui                                  C.GPUNIFORM3UI
	gpUniform3uiv                                 C.GPUNIFORM3UIV
	gpUniform4d                                   C.GPUNIFORM4D
	gpUniform4dv                                  C.GPUNIFORM4DV
	gpUniform4f                                   C.GPUNIFORM4F
	gpUniform4fv                                  C.GPUNIFORM4FV
	gpUniform4i                                   C.GPUNIFORM4I
	gpUniform4iv                                  C.GPUNIFORM4IV
	gpUniform4ui                                  C.GPUNIFORM4UI
	gpUniform4uiv                                 C.GPUNIFORM4UIV
	gpUniformBlockBinding                         C.GPUNIFORMBLOCKBINDING
	gpUniformHandleui64ARB                        C.GPUNIFORMHANDLEUI64ARB
	gpUniformHandleui64vARB                       C.GPUNIFORMHANDLEUI64VARB
	gpUniformMatrix2dv                            C.GPUNIFORMMATRIX2DV
	gpUniformMatrix2fv                            C.GPUNIFORMMATRIX2FV
	gpUniformMatrix2x3dv                          C.GPUNIFORMMATRIX2X3DV
	gpUniformMatrix2x3fv                          C.GPUNIFORMMATRIX2X3FV
	gpUniformMatrix2x4dv                          C.GPUNIFORMMATRIX2X4DV
	gpUniformMatrix2x4fv                          C.GPUNIFORMMATRIX2X4FV
	gpUniformMatrix3dv                            C.GPUNIFORMMATRIX3DV
	gpUniformMatrix3fv                            C.GPUNIFORMMATRIX3FV
	gpUniformMatrix3x2dv                          C.GPUNIFORMMATRIX3X2DV
	gpUniformMatrix3x2fv                          C.GPUNIFORMMATRIX3X2FV
	gpUniformMatrix3x4dv                          C.GPUNIFORMMATRIX3X4DV
	gpUniformMatrix3x4fv                          C.GPUNIFORMMATRIX3X4FV
	gpUniformMatrix4dv                            C.GPUNIFORMMATRIX4DV
	gpUniformMatrix4fv                            C.GPUNIFORMMATRIX4FV
	gpUniformMatrix4x2dv                          C.GPUNIFORMMATRIX4X2DV
	gpUniformMatrix4x2fv                          C.GPUNIFORMMATRIX4X2FV
	gpUniformMatrix4x3dv                          C.GPUNIFORMMATRIX4X3DV
	gpUniformMatrix4x3fv                          C.GPUNIFORMMATRIX4X3FV
	gpUniformSubroutinesuiv                       C.GPUNIFORMSUBROUTINESUIV
	gpUnmapBuffer                                 C.GPUNMAPBUFFER
	gpUnmapNamedBuffer                            C.GPUNMAPNAMEDBUFFER
	gpUseProgram                                  C.GPUSEPROGRAM
	gpUseProgramStages                            C.GPUSEPROGRAMSTAGES
	gpUseProgramStagesEXT                         C.GPUSEPROGRAMSTAGESEXT
	gpUseShaderProgramEXT                         C.GPUSESHADERPROGRAMEXT
	gpValidateProgram                             C.GPVALIDATEPROGRAM
	gpValidateProgramPipeline                     C.GPVALIDATEPROGRAMPIPELINE
	gpValidateProgramPipelineEXT                  C.GPVALIDATEPROGRAMPIPELINEEXT
	gpVertex2bOES                                 C.GPVERTEX2BOES
	gpVertex2bvOES                                C.GPVERTEX2BVOES
	gpVertex2xOES                                 C.GPVERTEX2XOES
	gpVertex2xvOES                                C.GPVERTEX2XVOES
	gpVertex3bOES                                 C.GPVERTEX3BOES
	gpVertex3bvOES                                C.GPVERTEX3BVOES
	gpVertex3xOES                                 C.GPVERTEX3XOES
	gpVertex3xvOES                                C.GPVERTEX3XVOES
	gpVertex4bOES                                 C.GPVERTEX4BOES
	gpVertex4bvOES                                C.GPVERTEX4BVOES
	gpVertex4xOES                                 C.GPVERTEX4XOES
	gpVertex4xvOES                                C.GPVERTEX4XVOES
	gpVertexArrayAttribBinding                    C.GPVERTEXARRAYATTRIBBINDING
	gpVertexArrayAttribFormat                     C.GPVERTEXARRAYATTRIBFORMAT
	gpVertexArrayAttribIFormat                    C.GPVERTEXARRAYATTRIBIFORMAT
	gpVertexArrayAttribLFormat                    C.GPVERTEXARRAYATTRIBLFORMAT
	gpVertexArrayBindingDivisor                   C.GPVERTEXARRAYBINDINGDIVISOR
	gpVertexArrayElementBuffer                    C.GPVERTEXARRAYELEMENTBUFFER
	gpVertexArrayVertexBuffer                     C.GPVERTEXARRAYVERTEXBUFFER
	gpVertexArrayVertexBuffers                    C.GPVERTEXARRAYVERTEXBUFFERS
	gpVertexAttrib1d                              C.GPVERTEXATTRIB1D
	gpVertexAttrib1dv                             C.GPVERTEXATTRIB1DV
	gpVertexAttrib1f                              C.GPVERTEXATTRIB1F
	gpVertexAttrib1fv                             C.GPVERTEXATTRIB1FV
	gpVertexAttrib1s                              C.GPVERTEXATTRIB1S
	gpVertexAttrib1sv                             C.GPVERTEXATTRIB1SV
	gpVertexAttrib2d                              C.GPVERTEXATTRIB2D
	gpVertexAttrib2dv                             C.GPVERTEXATTRIB2DV
	gpVertexAttrib2f                              C.GPVERTEXATTRIB2F
	gpVertexAttrib2fv                             C.GPVERTEXATTRIB2FV
	gpVertexAttrib2s                              C.GPVERTEXATTRIB2S
	gpVertexAttrib2sv                             C.GPVERTEXATTRIB2SV
	gpVertexAttrib3d                              C.GPVERTEXATTRIB3D
	gpVertexAttrib3dv                             C.GPVERTEXATTRIB3DV
	gpVertexAttrib3f                              C.GPVERTEXATTRIB3F
	gpVertexAttrib3fv                             C.GPVERTEXATTRIB3FV
	gpVertexAttrib3s                              C.GPVERTEXATTRIB3S
	gpVertexAttrib3sv                             C.GPVERTEXATTRIB3SV
	gpVertexAttrib4Nbv                            C.GPVERTEXATTRIB4NBV
	gpVertexAttrib4Niv                            C.GPVERTEXATTRIB4NIV
	gpVertexAttrib4Nsv                            C.GPVERTEXATTRIB4NSV
	gpVertexAttrib4Nub                            C.GPVERTEXATTRIB4NUB
	gpVertexAttrib4Nubv                           C.GPVERTEXATTRIB4NUBV
	gpVertexAttrib4Nuiv                           C.GPVERTEXATTRIB4NUIV
	gpVertexAttrib4Nusv                           C.GPVERTEXATTRIB4NUSV
	gpVertexAttrib4bv                             C.GPVERTEXATTRIB4BV
	gpVertexAttrib4d                              C.GPVERTEXATTRIB4D
	gpVertexAttrib4dv                             C.GPVERTEXATTRIB4DV
	gpVertexAttrib4f                              C.GPVERTEXATTRIB4F
	gpVertexAttrib4fv                             C.GPVERTEXATTRIB4FV
	gpVertexAttrib4iv                             C.GPVERTEXATTRIB4IV
	gpVertexAttrib4s                              C.GPVERTEXATTRIB4S
	gpVertexAttrib4sv                             C.GPVERTEXATTRIB4SV
	gpVertexAttrib4ubv                            C.GPVERTEXATTRIB4UBV
	gpVertexAttrib4uiv                            C.GPVERTEXATTRIB4UIV
	gpVertexAttrib4usv                            C.GPVERTEXATTRIB4USV
	gpVertexAttribBinding                         C.GPVERTEXATTRIBBINDING
	gpVertexAttribDivisor                         C.GPVERTEXATTRIBDIVISOR
	gpVertexAttribFormat                          C.GPVERTEXATTRIBFORMAT
	gpVertexAttribI1i                             C.GPVERTEXATTRIBI1I
	gpVertexAttribI1iv                            C.GPVERTEXATTRIBI1IV
	gpVertexAttribI1ui                            C.GPVERTEXATTRIBI1UI
	gpVertexAttribI1uiv                           C.GPVERTEXATTRIBI1UIV
	gpVertexAttribI2i                             C.GPVERTEXATTRIBI2I
	gpVertexAttribI2iv                            C.GPVERTEXATTRIBI2IV
	gpVertexAttribI2ui                            C.GPVERTEXATTRIBI2UI
	gpVertexAttribI2uiv                           C.GPVERTEXATTRIBI2UIV
	gpVertexAttribI3i                             C.GPVERTEXATTRIBI3I
	gpVertexAttribI3iv                            C.GPVERTEXATTRIBI3IV
	gpVertexAttribI3ui                            C.GPVERTEXATTRIBI3UI
	gpVertexAttribI3uiv                           C.GPVERTEXATTRIBI3UIV
	gpVertexAttribI4bv                            C.GPVERTEXATTRIBI4BV
	gpVertexAttribI4i                             C.GPVERTEXATTRIBI4I
	gpVertexAttribI4iv                            C.GPVERTEXATTRIBI4IV
	gpVertexAttribI4sv                            C.GPVERTEXATTRIBI4SV
	gpVertexAttribI4ubv                           C.GPVERTEXATTRIBI4UBV
	gpVertexAttribI4ui                            C.GPVERTEXATTRIBI4UI
	gpVertexAttribI4uiv                           C.GPVERTEXATTRIBI4UIV
	gpVertexAttribI4usv                           C.GPVERTEXATTRIBI4USV
	gpVertexAttribIFormat                         C.GPVERTEXATTRIBIFORMAT
	gpVertexAttribIPointer                        C.GPVERTEXATTRIBIPOINTER
	gpVertexAttribL1d                             C.GPVERTEXATTRIBL1D
	gpVertexAttribL1dv                            C.GPVERTEXATTRIBL1DV
	gpVertexAttribL1ui64ARB                       C.GPVERTEXATTRIBL1UI64ARB
	gpVertexAttribL1ui64vARB                      C.GPVERTEXATTRIBL1UI64VARB
	gpVertexAttribL2d                             C.GPVERTEXATTRIBL2D
	gpVertexAttribL2dv                            C.GPVERTEXATTRIBL2DV
	gpVertexAttribL3d                             C.GPVERTEXATTRIBL3D
	gpVertexAttribL3dv                            C.GPVERTEXATTRIBL3DV
	gpVertexAttribL4d                             C.GPVERTEXATTRIBL4D
	gpVertexAttribL4dv                            C.GPVERTEXATTRIBL4DV
	gpVertexAttribLFormat                         C.GPVERTEXATTRIBLFORMAT
	gpVertexAttribLPointer                        C.GPVERTEXATTRIBLPOINTER
	gpVertexAttribP1ui                            C.GPVERTEXATTRIBP1UI
	gpVertexAttribP1uiv                           C.GPVERTEXATTRIBP1UIV
	gpVertexAttribP2ui                            C.GPVERTEXATTRIBP2UI
	gpVertexAttribP2uiv                           C.GPVERTEXATTRIBP2UIV
	gpVertexAttribP3ui                            C.GPVERTEXATTRIBP3UI
	gpVertexAttribP3uiv                           C.GPVERTEXATTRIBP3UIV
	gpVertexAttribP4ui                            C.GPVERTEXATTRIBP4UI
	gpVertexAttribP4uiv                           C.GPVERTEXATTRIBP4UIV
	gpVertexAttribPointer                         C.GPVERTEXATTRIBPOINTER
	gpVertexBindingDivisor                        C.GPVERTEXBINDINGDIVISOR
	gpViewport                                    C.GPVIEWPORT
	gpViewportArrayv                              C.GPVIEWPORTARRAYV
	gpViewportIndexedf                            C.GPVIEWPORTINDEXEDF
	gpViewportIndexedfv                           C.GPVIEWPORTINDEXEDFV
	gpWaitSync                                    C.GPWAITSYNC
)

// Helper functions
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
func AccumxOES(op uint32, value int32) {
	C.glowAccumxOES(gpAccumxOES, (C.GLenum)(op), (C.GLfixed)(value))
}
func ActiveProgramEXT(program uint32) {
	C.glowActiveProgramEXT(gpActiveProgramEXT, (C.GLuint)(program))
}

// set the active program object for a program pipeline object
func ActiveShaderProgram(pipeline uint32, program uint32) {
	C.glowActiveShaderProgram(gpActiveShaderProgram, (C.GLuint)(pipeline), (C.GLuint)(program))
}
func ActiveShaderProgramEXT(pipeline uint32, program uint32) {
	C.glowActiveShaderProgramEXT(gpActiveShaderProgramEXT, (C.GLuint)(pipeline), (C.GLuint)(program))
}

// select active texture unit
func ActiveTexture(texture uint32) {
	C.glowActiveTexture(gpActiveTexture, (C.GLenum)(texture))
}
func AlphaFuncxOES(xfunc uint32, ref int32) {
	C.glowAlphaFuncxOES(gpAlphaFuncxOES, (C.GLenum)(xfunc), (C.GLfixed)(ref))
}

// Attaches a shader object to a program object
func AttachShader(program uint32, shader uint32) {
	C.glowAttachShader(gpAttachShader, (C.GLuint)(program), (C.GLuint)(shader))
}

// start conditional rendering
func BeginConditionalRender(id uint32, mode uint32) {
	C.glowBeginConditionalRender(gpBeginConditionalRender, (C.GLuint)(id), (C.GLenum)(mode))
}
func BeginPerfMonitorAMD(monitor uint32) {
	C.glowBeginPerfMonitorAMD(gpBeginPerfMonitorAMD, (C.GLuint)(monitor))
}
func BeginPerfQueryINTEL(queryHandle uint32) {
	C.glowBeginPerfQueryINTEL(gpBeginPerfQueryINTEL, (C.GLuint)(queryHandle))
}

// delimit the boundaries of a query object
func BeginQuery(target uint32, id uint32) {
	C.glowBeginQuery(gpBeginQuery, (C.GLenum)(target), (C.GLuint)(id))
}
func BeginQueryIndexed(target uint32, index uint32, id uint32) {
	C.glowBeginQueryIndexed(gpBeginQueryIndexed, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(id))
}

// start transform feedback operation
func BeginTransformFeedback(primitiveMode uint32) {
	C.glowBeginTransformFeedback(gpBeginTransformFeedback, (C.GLenum)(primitiveMode))
}

// Associates a generic vertex attribute index with a named attribute variable
func BindAttribLocation(program uint32, index uint32, name *uint8) {
	C.glowBindAttribLocation(gpBindAttribLocation, (C.GLuint)(program), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a named buffer object
func BindBuffer(target uint32, buffer uint32) {
	C.glowBindBuffer(gpBindBuffer, (C.GLenum)(target), (C.GLuint)(buffer))
}

// bind a buffer object to an indexed buffer target
func BindBufferBase(target uint32, index uint32, buffer uint32) {
	C.glowBindBufferBase(gpBindBufferBase, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer))
}

// bind a range within a buffer object to an indexed buffer target
func BindBufferRange(target uint32, index uint32, buffer uint32, offset int, size int) {
	C.glowBindBufferRange(gpBindBufferRange, (C.GLenum)(target), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}

// bind one or more buffer objects to a sequence of indexed buffer targets
func BindBuffersBase(target uint32, first uint32, count int32, buffers *uint32) {
	C.glowBindBuffersBase(gpBindBuffersBase, (C.GLenum)(target), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)))
}

// bind ranges of one or more buffer objects to a sequence of indexed buffer targets
func BindBuffersRange(target uint32, first uint32, count int32, buffers *uint32, offsets *int, sizes *int) {
	C.glowBindBuffersRange(gpBindBuffersRange, (C.GLenum)(target), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizeiptr)(unsafe.Pointer(sizes)))
}

// bind a user-defined varying out variable to a fragment shader color number
func BindFragDataLocation(program uint32, color uint32, name *uint8) {
	C.glowBindFragDataLocation(gpBindFragDataLocation, (C.GLuint)(program), (C.GLuint)(color), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a user-defined varying out variable to a fragment shader color number and index
func BindFragDataLocationIndexed(program uint32, colorNumber uint32, index uint32, name *uint8) {
	C.glowBindFragDataLocationIndexed(gpBindFragDataLocationIndexed, (C.GLuint)(program), (C.GLuint)(colorNumber), (C.GLuint)(index), (*C.GLchar)(unsafe.Pointer(name)))
}

// bind a framebuffer to a framebuffer target
func BindFramebuffer(target uint32, framebuffer uint32) {
	C.glowBindFramebuffer(gpBindFramebuffer, (C.GLenum)(target), (C.GLuint)(framebuffer))
}

// bind a level of a texture to an image unit
func BindImageTexture(unit uint32, texture uint32, level int32, layered bool, layer int32, access uint32, format uint32) {
	C.glowBindImageTexture(gpBindImageTexture, (C.GLuint)(unit), (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(access), (C.GLenum)(format))
}

// bind one or more named texture images to a sequence of consecutive image units
func BindImageTextures(first uint32, count int32, textures *uint32) {
	C.glowBindImageTextures(gpBindImageTextures, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(textures)))
}

// bind a program pipeline to the current context
func BindProgramPipeline(pipeline uint32) {
	C.glowBindProgramPipeline(gpBindProgramPipeline, (C.GLuint)(pipeline))
}
func BindProgramPipelineEXT(pipeline uint32) {
	C.glowBindProgramPipelineEXT(gpBindProgramPipelineEXT, (C.GLuint)(pipeline))
}

// bind a renderbuffer to a renderbuffer target
func BindRenderbuffer(target uint32, renderbuffer uint32) {
	C.glowBindRenderbuffer(gpBindRenderbuffer, (C.GLenum)(target), (C.GLuint)(renderbuffer))
}

// bind a named sampler to a texturing target
func BindSampler(unit uint32, sampler uint32) {
	C.glowBindSampler(gpBindSampler, (C.GLuint)(unit), (C.GLuint)(sampler))
}

// bind one or more named sampler objects to a sequence of consecutive sampler units
func BindSamplers(first uint32, count int32, samplers *uint32) {
	C.glowBindSamplers(gpBindSamplers, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// bind a named texture to a texturing target
func BindTexture(target uint32, texture uint32) {
	C.glowBindTexture(gpBindTexture, (C.GLenum)(target), (C.GLuint)(texture))
}

// bind an existing texture object to the specified texture unit
func BindTextureUnit(unit uint32, texture uint32) {
	C.glowBindTextureUnit(gpBindTextureUnit, (C.GLuint)(unit), (C.GLuint)(texture))
}

// bind one or more named textures to a sequence of consecutive texture units
func BindTextures(first uint32, count int32, textures *uint32) {
	C.glowBindTextures(gpBindTextures, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(textures)))
}

// bind a transform feedback object
func BindTransformFeedback(target uint32, id uint32) {
	C.glowBindTransformFeedback(gpBindTransformFeedback, (C.GLenum)(target), (C.GLuint)(id))
}

// bind a vertex array object
func BindVertexArray(array uint32) {
	C.glowBindVertexArray(gpBindVertexArray, (C.GLuint)(array))
}

// bind a buffer to a vertex buffer bind point
func BindVertexBuffer(bindingindex uint32, buffer uint32, offset int, stride int32) {
	C.glowBindVertexBuffer(gpBindVertexBuffer, (C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}

// attach multiple buffer objects to a vertex array object
func BindVertexBuffers(first uint32, count int32, buffers *uint32, offsets *int, strides *int32) {
	C.glowBindVertexBuffers(gpBindVertexBuffers, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizei)(unsafe.Pointer(strides)))
}
func BitmapxOES(width int32, height int32, xorig int32, yorig int32, xmove int32, ymove int32, bitmap *uint8) {
	C.glowBitmapxOES(gpBitmapxOES, (C.GLsizei)(width), (C.GLsizei)(height), (C.GLfixed)(xorig), (C.GLfixed)(yorig), (C.GLfixed)(xmove), (C.GLfixed)(ymove), (*C.GLubyte)(unsafe.Pointer(bitmap)))
}
func BlendBarrierKHR() {
	C.glowBlendBarrierKHR(gpBlendBarrierKHR)
}
func BlendBarrierNV() {
	C.glowBlendBarrierNV(gpBlendBarrierNV)
}

// set the blend color
func BlendColor(red float32, green float32, blue float32, alpha float32) {
	C.glowBlendColor(gpBlendColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func BlendColorxOES(red int32, green int32, blue int32, alpha int32) {
	C.glowBlendColorxOES(gpBlendColorxOES, (C.GLfixed)(red), (C.GLfixed)(green), (C.GLfixed)(blue), (C.GLfixed)(alpha))
}

// specify the equation used for both the RGB blend equation and the Alpha blend equation
func BlendEquation(mode uint32) {
	C.glowBlendEquation(gpBlendEquation, (C.GLenum)(mode))
}
func BlendEquationEXT(mode uint32) {
	C.glowBlendEquationEXT(gpBlendEquationEXT, (C.GLenum)(mode))
}

// set the RGB blend equation and the alpha blend equation separately
func BlendEquationSeparate(modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparate(gpBlendEquationSeparate, (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparatei(buf uint32, modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparatei(gpBlendEquationSeparatei, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationSeparateiARB(buf uint32, modeRGB uint32, modeAlpha uint32) {
	C.glowBlendEquationSeparateiARB(gpBlendEquationSeparateiARB, (C.GLuint)(buf), (C.GLenum)(modeRGB), (C.GLenum)(modeAlpha))
}
func BlendEquationi(buf uint32, mode uint32) {
	C.glowBlendEquationi(gpBlendEquationi, (C.GLuint)(buf), (C.GLenum)(mode))
}
func BlendEquationiARB(buf uint32, mode uint32) {
	C.glowBlendEquationiARB(gpBlendEquationiARB, (C.GLuint)(buf), (C.GLenum)(mode))
}

// specify pixel arithmetic
func BlendFunc(sfactor uint32, dfactor uint32) {
	C.glowBlendFunc(gpBlendFunc, (C.GLenum)(sfactor), (C.GLenum)(dfactor))
}

// specify pixel arithmetic for RGB and alpha components separately
func BlendFuncSeparate(sfactorRGB uint32, dfactorRGB uint32, sfactorAlpha uint32, dfactorAlpha uint32) {
	C.glowBlendFuncSeparate(gpBlendFuncSeparate, (C.GLenum)(sfactorRGB), (C.GLenum)(dfactorRGB), (C.GLenum)(sfactorAlpha), (C.GLenum)(dfactorAlpha))
}
func BlendFuncSeparatei(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	C.glowBlendFuncSeparatei(gpBlendFuncSeparatei, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFuncSeparateiARB(buf uint32, srcRGB uint32, dstRGB uint32, srcAlpha uint32, dstAlpha uint32) {
	C.glowBlendFuncSeparateiARB(gpBlendFuncSeparateiARB, (C.GLuint)(buf), (C.GLenum)(srcRGB), (C.GLenum)(dstRGB), (C.GLenum)(srcAlpha), (C.GLenum)(dstAlpha))
}
func BlendFunci(buf uint32, src uint32, dst uint32) {
	C.glowBlendFunci(gpBlendFunci, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlendFunciARB(buf uint32, src uint32, dst uint32) {
	C.glowBlendFunciARB(gpBlendFunciARB, (C.GLuint)(buf), (C.GLenum)(src), (C.GLenum)(dst))
}
func BlendParameteriNV(pname uint32, value int32) {
	C.glowBlendParameteriNV(gpBlendParameteriNV, (C.GLenum)(pname), (C.GLint)(value))
}

// copy a block of pixels from one framebuffer object to another
func BlitFramebuffer(srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	C.glowBlitFramebuffer(gpBlitFramebuffer, (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}

// copy a block of pixels from one framebuffer object to another
func BlitNamedFramebuffer(readFramebuffer uint32, drawFramebuffer uint32, srcX0 int32, srcY0 int32, srcX1 int32, srcY1 int32, dstX0 int32, dstY0 int32, dstX1 int32, dstY1 int32, mask uint32, filter uint32) {
	C.glowBlitNamedFramebuffer(gpBlitNamedFramebuffer, (C.GLuint)(readFramebuffer), (C.GLuint)(drawFramebuffer), (C.GLint)(srcX0), (C.GLint)(srcY0), (C.GLint)(srcX1), (C.GLint)(srcY1), (C.GLint)(dstX0), (C.GLint)(dstY0), (C.GLint)(dstX1), (C.GLint)(dstY1), (C.GLbitfield)(mask), (C.GLenum)(filter))
}

// creates and initializes a buffer object's data     store
func BufferData(target uint32, size int, data unsafe.Pointer, usage uint32) {
	C.glowBufferData(gpBufferData, (C.GLenum)(target), (C.GLsizeiptr)(size), data, (C.GLenum)(usage))
}
func BufferPageCommitmentARB(target uint32, offset int, size int32, commit bool) {
	C.glowBufferPageCommitmentARB(gpBufferPageCommitmentARB, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}

// creates and initializes a buffer object's immutable data     store
func BufferStorage(target uint32, size int, data unsafe.Pointer, flags uint32) {
	C.glowBufferStorage(gpBufferStorage, (C.GLenum)(target), (C.GLsizeiptr)(size), data, (C.GLbitfield)(flags))
}

// updates a subset of a buffer object's data store
func BufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	C.glowBufferSubData(gpBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), data)
}

// check the completeness status of a framebuffer
func CheckFramebufferStatus(target uint32) uint32 {
	ret := C.glowCheckFramebufferStatus(gpCheckFramebufferStatus, (C.GLenum)(target))
	return (uint32)(ret)
}

// check the completeness status of a framebuffer
func CheckNamedFramebufferStatus(framebuffer uint32, target uint32) uint32 {
	ret := C.glowCheckNamedFramebufferStatus(gpCheckNamedFramebufferStatus, (C.GLuint)(framebuffer), (C.GLenum)(target))
	return (uint32)(ret)
}

// specify whether data read via  should be clamped
func ClampColor(target uint32, clamp uint32) {
	C.glowClampColor(gpClampColor, (C.GLenum)(target), (C.GLenum)(clamp))
}

// clear buffers to preset values
func Clear(mask uint32) {
	C.glowClear(gpClear, (C.GLbitfield)(mask))
}
func ClearAccumxOES(red int32, green int32, blue int32, alpha int32) {
	C.glowClearAccumxOES(gpClearAccumxOES, (C.GLfixed)(red), (C.GLfixed)(green), (C.GLfixed)(blue), (C.GLfixed)(alpha))
}

// fill a buffer object's data store with a fixed value
func ClearBufferData(target uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearBufferData(gpClearBufferData, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// fill all or part of buffer object's data store with a fixed value
func ClearBufferSubData(target uint32, internalformat uint32, offset int, size int, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearBufferSubData(gpClearBufferSubData, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizeiptr)(size), (C.GLenum)(format), (C.GLenum)(xtype), data)
}
func ClearBufferfi(buffer uint32, drawbuffer int32, depth float32, stencil int32) {
	C.glowClearBufferfi(gpClearBufferfi, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
func ClearBufferfv(buffer uint32, drawbuffer int32, value *float32) {
	C.glowClearBufferfv(gpClearBufferfv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ClearBufferiv(buffer uint32, drawbuffer int32, value *int32) {
	C.glowClearBufferiv(gpClearBufferiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(unsafe.Pointer(value)))
}
func ClearBufferuiv(buffer uint32, drawbuffer int32, value *uint32) {
	C.glowClearBufferuiv(gpClearBufferuiv, (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(unsafe.Pointer(value)))
}

// specify clear values for the color buffers
func ClearColor(red float32, green float32, blue float32, alpha float32) {
	C.glowClearColor(gpClearColor, (C.GLfloat)(red), (C.GLfloat)(green), (C.GLfloat)(blue), (C.GLfloat)(alpha))
}
func ClearColorxOES(red int32, green int32, blue int32, alpha int32) {
	C.glowClearColorxOES(gpClearColorxOES, (C.GLfixed)(red), (C.GLfixed)(green), (C.GLfixed)(blue), (C.GLfixed)(alpha))
}

// specify the clear value for the depth buffer
func ClearDepth(depth float64) {
	C.glowClearDepth(gpClearDepth, (C.GLdouble)(depth))
}
func ClearDepthf(d float32) {
	C.glowClearDepthf(gpClearDepthf, (C.GLfloat)(d))
}
func ClearDepthfOES(depth float32) {
	C.glowClearDepthfOES(gpClearDepthfOES, (C.GLclampf)(depth))
}
func ClearDepthxOES(depth int32) {
	C.glowClearDepthxOES(gpClearDepthxOES, (C.GLfixed)(depth))
}

// fill a buffer object's data store with a fixed value
func ClearNamedBufferData(buffer uint32, internalformat uint32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearNamedBufferData(gpClearNamedBufferData, (C.GLuint)(buffer), (C.GLenum)(internalformat), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// fill all or part of buffer object's data store with a fixed value
func ClearNamedBufferSubData(buffer uint32, internalformat uint32, offset int, size int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearNamedBufferSubData(gpClearNamedBufferSubData, (C.GLuint)(buffer), (C.GLenum)(internalformat), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLenum)(format), (C.GLenum)(xtype), data)
}
func ClearNamedFramebufferfi(framebuffer uint32, buffer uint32, depth float32, stencil int32) {
	C.glowClearNamedFramebufferfi(gpClearNamedFramebufferfi, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLfloat)(depth), (C.GLint)(stencil))
}
func ClearNamedFramebufferfv(framebuffer uint32, buffer uint32, drawbuffer int32, value *float32) {
	C.glowClearNamedFramebufferfv(gpClearNamedFramebufferfv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ClearNamedFramebufferiv(framebuffer uint32, buffer uint32, drawbuffer int32, value *int32) {
	C.glowClearNamedFramebufferiv(gpClearNamedFramebufferiv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLint)(unsafe.Pointer(value)))
}
func ClearNamedFramebufferuiv(framebuffer uint32, buffer uint32, drawbuffer int32, value *uint32) {
	C.glowClearNamedFramebufferuiv(gpClearNamedFramebufferuiv, (C.GLuint)(framebuffer), (C.GLenum)(buffer), (C.GLint)(drawbuffer), (*C.GLuint)(unsafe.Pointer(value)))
}

// specify the clear value for the stencil buffer
func ClearStencil(s int32) {
	C.glowClearStencil(gpClearStencil, (C.GLint)(s))
}

// fills all a texture image with a constant value
func ClearTexImage(texture uint32, level int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearTexImage(gpClearTexImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// fills all or part of a texture image with a constant value
func ClearTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, data unsafe.Pointer) {
	C.glowClearTexSubImage(gpClearTexSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), data)
}

// block and wait for a sync object to become signaled
func ClientWaitSync(sync unsafe.Pointer, flags uint32, timeout uint64) uint32 {
	ret := C.glowClientWaitSync(gpClientWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
	return (uint32)(ret)
}

// control clip coordinate to window coordinate behavior
func ClipControl(origin uint32, depth uint32) {
	C.glowClipControl(gpClipControl, (C.GLenum)(origin), (C.GLenum)(depth))
}
func ClipPlanefOES(plane uint32, equation *float32) {
	C.glowClipPlanefOES(gpClipPlanefOES, (C.GLenum)(plane), (*C.GLfloat)(unsafe.Pointer(equation)))
}
func ClipPlanexOES(plane uint32, equation *int32) {
	C.glowClipPlanexOES(gpClipPlanexOES, (C.GLenum)(plane), (*C.GLfixed)(unsafe.Pointer(equation)))
}
func Color3xOES(red int32, green int32, blue int32) {
	C.glowColor3xOES(gpColor3xOES, (C.GLfixed)(red), (C.GLfixed)(green), (C.GLfixed)(blue))
}
func Color3xvOES(components *int32) {
	C.glowColor3xvOES(gpColor3xvOES, (*C.GLfixed)(unsafe.Pointer(components)))
}
func Color4xOES(red int32, green int32, blue int32, alpha int32) {
	C.glowColor4xOES(gpColor4xOES, (C.GLfixed)(red), (C.GLfixed)(green), (C.GLfixed)(blue), (C.GLfixed)(alpha))
}
func Color4xvOES(components *int32) {
	C.glowColor4xvOES(gpColor4xvOES, (*C.GLfixed)(unsafe.Pointer(components)))
}
func ColorMask(red bool, green bool, blue bool, alpha bool) {
	C.glowColorMask(gpColorMask, (C.GLboolean)(boolToInt(red)), (C.GLboolean)(boolToInt(green)), (C.GLboolean)(boolToInt(blue)), (C.GLboolean)(boolToInt(alpha)))
}
func ColorMaski(index uint32, r bool, g bool, b bool, a bool) {
	C.glowColorMaski(gpColorMaski, (C.GLuint)(index), (C.GLboolean)(boolToInt(r)), (C.GLboolean)(boolToInt(g)), (C.GLboolean)(boolToInt(b)), (C.GLboolean)(boolToInt(a)))
}

// Compiles a shader object
func CompileShader(shader uint32) {
	C.glowCompileShader(gpCompileShader, (C.GLuint)(shader))
}
func CompileShaderIncludeARB(shader uint32, count int32, path **uint8, length *int32) {
	C.glowCompileShaderIncludeARB(gpCompileShaderIncludeARB, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(path)), (*C.GLint)(unsafe.Pointer(length)))
}

// specify a one-dimensional texture image in a compressed format
func CompressedTexImage1D(target uint32, level int32, internalformat uint32, width int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage1D(gpCompressedTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a two-dimensional texture image in a compressed format
func CompressedTexImage2D(target uint32, level int32, internalformat uint32, width int32, height int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage2D(gpCompressedTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture image in a compressed format
func CompressedTexImage3D(target uint32, level int32, internalformat uint32, width int32, height int32, depth int32, border int32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexImage3D(gpCompressedTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLsizei)(imageSize), data)
}

// specify a one-dimensional texture subimage in a compressed     format
func CompressedTexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage1D(gpCompressedTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a two-dimensional texture subimage in a compressed format
func CompressedTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage2D(gpCompressedTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture subimage in a compressed format
func CompressedTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTexSubImage3D(gpCompressedTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a one-dimensional texture subimage in a compressed     format
func CompressedTextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTextureSubImage1D(gpCompressedTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a two-dimensional texture subimage in a compressed format
func CompressedTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTextureSubImage2D(gpCompressedTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}

// specify a three-dimensional texture subimage in a compressed format
func CompressedTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, imageSize int32, data unsafe.Pointer) {
	C.glowCompressedTextureSubImage3D(gpCompressedTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLsizei)(imageSize), data)
}
func ConvolutionParameterxOES(target uint32, pname uint32, param int32) {
	C.glowConvolutionParameterxOES(gpConvolutionParameterxOES, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfixed)(param))
}
func ConvolutionParameterxvOES(target uint32, pname uint32, params *int32) {
	C.glowConvolutionParameterxvOES(gpConvolutionParameterxvOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// copy all or part of the data store of a buffer object to the data store of another buffer object
func CopyBufferSubData(readTarget uint32, writeTarget uint32, readOffset int, writeOffset int, size int) {
	C.glowCopyBufferSubData(gpCopyBufferSubData, (C.GLenum)(readTarget), (C.GLenum)(writeTarget), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizeiptr)(size))
}

// perform a raw data copy between two images
func CopyImageSubData(srcName uint32, srcTarget uint32, srcLevel int32, srcX int32, srcY int32, srcZ int32, dstName uint32, dstTarget uint32, dstLevel int32, dstX int32, dstY int32, dstZ int32, srcWidth int32, srcHeight int32, srcDepth int32) {
	C.glowCopyImageSubData(gpCopyImageSubData, (C.GLuint)(srcName), (C.GLenum)(srcTarget), (C.GLint)(srcLevel), (C.GLint)(srcX), (C.GLint)(srcY), (C.GLint)(srcZ), (C.GLuint)(dstName), (C.GLenum)(dstTarget), (C.GLint)(dstLevel), (C.GLint)(dstX), (C.GLint)(dstY), (C.GLint)(dstZ), (C.GLsizei)(srcWidth), (C.GLsizei)(srcHeight), (C.GLsizei)(srcDepth))
}

// copy all or part of the data store of a buffer object to the data store of another buffer object
func CopyNamedBufferSubData(readBuffer uint32, writeBuffer uint32, readOffset int, writeOffset int, size int32) {
	C.glowCopyNamedBufferSubData(gpCopyNamedBufferSubData, (C.GLuint)(readBuffer), (C.GLuint)(writeBuffer), (C.GLintptr)(readOffset), (C.GLintptr)(writeOffset), (C.GLsizei)(size))
}

// copy pixels into a 1D texture image
func CopyTexImage1D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, border int32) {
	C.glowCopyTexImage1D(gpCopyTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLint)(border))
}

// copy pixels into a 2D texture image
func CopyTexImage2D(target uint32, level int32, internalformat uint32, x int32, y int32, width int32, height int32, border int32) {
	C.glowCopyTexImage2D(gpCopyTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(internalformat), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border))
}

// copy a one-dimensional texture subimage
func CopyTexSubImage1D(target uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	C.glowCopyTexSubImage1D(gpCopyTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}

// copy a two-dimensional texture subimage
func CopyTexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTexSubImage2D(gpCopyTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// copy a three-dimensional texture subimage
func CopyTexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTexSubImage3D(gpCopyTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// copy a one-dimensional texture subimage
func CopyTextureSubImage1D(texture uint32, level int32, xoffset int32, x int32, y int32, width int32) {
	C.glowCopyTextureSubImage1D(gpCopyTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width))
}

// copy a two-dimensional texture subimage
func CopyTextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTextureSubImage2D(gpCopyTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// copy a three-dimensional texture subimage
func CopyTextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, x int32, y int32, width int32, height int32) {
	C.glowCopyTextureSubImage3D(gpCopyTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// create buffer objects
func CreateBuffers(n int32, buffers *uint32) {
	C.glowCreateBuffers(gpCreateBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}

// create framebuffer objects
func CreateFramebuffers(n int32, framebuffers *uint32) {
	C.glowCreateFramebuffers(gpCreateFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func CreatePerfQueryINTEL(queryId uint32, queryHandle *uint32) {
	C.glowCreatePerfQueryINTEL(gpCreatePerfQueryINTEL, (C.GLuint)(queryId), (*C.GLuint)(unsafe.Pointer(queryHandle)))
}

// Creates a program object
func CreateProgram() uint32 {
	ret := C.glowCreateProgram(gpCreateProgram)
	return (uint32)(ret)
}

// create program pipeline objects
func CreateProgramPipelines(n int32, pipelines *uint32) {
	C.glowCreateProgramPipelines(gpCreateProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// create query objects
func CreateQueries(target uint32, n int32, ids *uint32) {
	C.glowCreateQueries(gpCreateQueries, (C.GLenum)(target), (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// create renderbuffer objects
func CreateRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowCreateRenderbuffers(gpCreateRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// create sampler objects
func CreateSamplers(n int32, samplers *uint32) {
	C.glowCreateSamplers(gpCreateSamplers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// Creates a shader object
func CreateShader(xtype uint32) uint32 {
	ret := C.glowCreateShader(gpCreateShader, (C.GLenum)(xtype))
	return (uint32)(ret)
}
func CreateShaderProgramEXT(xtype uint32, xstring *uint8) uint32 {
	ret := C.glowCreateShaderProgramEXT(gpCreateShaderProgramEXT, (C.GLenum)(xtype), (*C.GLchar)(unsafe.Pointer(xstring)))
	return (uint32)(ret)
}

// create a stand-alone program from an array of null-terminated source code strings
func CreateShaderProgramv(xtype uint32, count int32, strings **uint8) uint32 {
	ret := C.glowCreateShaderProgramv(gpCreateShaderProgramv, (C.GLenum)(xtype), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings)))
	return (uint32)(ret)
}
func CreateShaderProgramvEXT(xtype uint32, count int32, strings **uint8) uint32 {
	ret := C.glowCreateShaderProgramvEXT(gpCreateShaderProgramvEXT, (C.GLenum)(xtype), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(strings)))
	return (uint32)(ret)
}
func CreateSyncFromCLeventARB(context *C.struct__cl_context, event *C.struct__cl_event, flags uint32) unsafe.Pointer {
	ret := C.glowCreateSyncFromCLeventARB(gpCreateSyncFromCLeventARB, (*C.struct__cl_context)(unsafe.Pointer(context)), (*C.struct__cl_event)(unsafe.Pointer(event)), (C.GLbitfield)(flags))
	return (unsafe.Pointer)(ret)
}

// create texture objects
func CreateTextures(target uint32, n int32, textures *uint32) {
	C.glowCreateTextures(gpCreateTextures, (C.GLenum)(target), (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// create transform feedback objects
func CreateTransformFeedbacks(n int32, ids *uint32) {
	C.glowCreateTransformFeedbacks(gpCreateTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// create vertex array objects
func CreateVertexArrays(n int32, arrays *uint32) {
	C.glowCreateVertexArrays(gpCreateVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// specify whether front- or back-facing facets can be culled
func CullFace(mode uint32) {
	C.glowCullFace(gpCullFace, (C.GLenum)(mode))
}

// specify a callback to receive debugging messages from the GL
func DebugMessageCallback(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallback(gpDebugMessageCallback, (C.GLDEBUGPROC)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageCallbackARB(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallbackARB(gpDebugMessageCallbackARB, (C.GLDEBUGPROCARB)(unsafe.Pointer(&callback)), userParam)
}
func DebugMessageCallbackKHR(callback DebugProc, userParam unsafe.Pointer) {
	userDebugCallback = callback
	C.glowDebugMessageCallbackKHR(gpDebugMessageCallbackKHR, (C.GLDEBUGPROCKHR)(unsafe.Pointer(&callback)), userParam)
}

// control the reporting of debug messages in a debug context
func DebugMessageControl(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControl(gpDebugMessageControl, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageControlARB(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControlARB(gpDebugMessageControlARB, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}
func DebugMessageControlKHR(source uint32, xtype uint32, severity uint32, count int32, ids *uint32, enabled bool) {
	C.glowDebugMessageControlKHR(gpDebugMessageControlKHR, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLenum)(severity), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(ids)), (C.GLboolean)(boolToInt(enabled)))
}

// inject an application-supplied message into the debug message queue
func DebugMessageInsert(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsert(gpDebugMessageInsert, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}
func DebugMessageInsertARB(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsertARB(gpDebugMessageInsertARB, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}
func DebugMessageInsertKHR(source uint32, xtype uint32, id uint32, severity uint32, length int32, buf *uint8) {
	C.glowDebugMessageInsertKHR(gpDebugMessageInsertKHR, (C.GLenum)(source), (C.GLenum)(xtype), (C.GLuint)(id), (C.GLenum)(severity), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(buf)))
}

// delete named buffer objects
func DeleteBuffers(n int32, buffers *uint32) {
	C.glowDeleteBuffers(gpDeleteBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func DeleteFencesNV(n int32, fences *uint32) {
	C.glowDeleteFencesNV(gpDeleteFencesNV, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(fences)))
}

// delete framebuffer objects
func DeleteFramebuffers(n int32, framebuffers *uint32) {
	C.glowDeleteFramebuffers(gpDeleteFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func DeleteNamedStringARB(namelen int32, name *uint8) {
	C.glowDeleteNamedStringARB(gpDeleteNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)))
}
func DeletePerfMonitorsAMD(n int32, monitors *uint32) {
	C.glowDeletePerfMonitorsAMD(gpDeletePerfMonitorsAMD, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(monitors)))
}
func DeletePerfQueryINTEL(queryHandle uint32) {
	C.glowDeletePerfQueryINTEL(gpDeletePerfQueryINTEL, (C.GLuint)(queryHandle))
}

// Deletes a program object
func DeleteProgram(program uint32) {
	C.glowDeleteProgram(gpDeleteProgram, (C.GLuint)(program))
}

// delete program pipeline objects
func DeleteProgramPipelines(n int32, pipelines *uint32) {
	C.glowDeleteProgramPipelines(gpDeleteProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}
func DeleteProgramPipelinesEXT(n int32, pipelines *uint32) {
	C.glowDeleteProgramPipelinesEXT(gpDeleteProgramPipelinesEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// delete named query objects
func DeleteQueries(n int32, ids *uint32) {
	C.glowDeleteQueries(gpDeleteQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// delete renderbuffer objects
func DeleteRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowDeleteRenderbuffers(gpDeleteRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// delete named sampler objects
func DeleteSamplers(count int32, samplers *uint32) {
	C.glowDeleteSamplers(gpDeleteSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// Deletes a shader object
func DeleteShader(shader uint32) {
	C.glowDeleteShader(gpDeleteShader, (C.GLuint)(shader))
}

// delete a sync object
func DeleteSync(sync unsafe.Pointer) {
	C.glowDeleteSync(gpDeleteSync, (C.GLsync)(sync))
}

// delete named textures
func DeleteTextures(n int32, textures *uint32) {
	C.glowDeleteTextures(gpDeleteTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// delete transform feedback objects
func DeleteTransformFeedbacks(n int32, ids *uint32) {
	C.glowDeleteTransformFeedbacks(gpDeleteTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// delete vertex array objects
func DeleteVertexArrays(n int32, arrays *uint32) {
	C.glowDeleteVertexArrays(gpDeleteVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// specify the value used for depth buffer comparisons
func DepthFunc(xfunc uint32) {
	C.glowDepthFunc(gpDepthFunc, (C.GLenum)(xfunc))
}

// enable or disable writing into the depth buffer
func DepthMask(flag bool) {
	C.glowDepthMask(gpDepthMask, (C.GLboolean)(boolToInt(flag)))
}

// specify mapping of depth values from normalized device coordinates to window coordinates
func DepthRange(near float64, far float64) {
	C.glowDepthRange(gpDepthRange, (C.GLdouble)(near), (C.GLdouble)(far))
}
func DepthRangeArrayv(first uint32, count int32, v *float64) {
	C.glowDepthRangeArrayv(gpDepthRangeArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(v)))
}

// specify mapping of depth values from normalized device coordinates to window coordinates for a specified viewport
func DepthRangeIndexed(index uint32, n float64, f float64) {
	C.glowDepthRangeIndexed(gpDepthRangeIndexed, (C.GLuint)(index), (C.GLdouble)(n), (C.GLdouble)(f))
}
func DepthRangef(n float32, f float32) {
	C.glowDepthRangef(gpDepthRangef, (C.GLfloat)(n), (C.GLfloat)(f))
}
func DepthRangefOES(n float32, f float32) {
	C.glowDepthRangefOES(gpDepthRangefOES, (C.GLclampf)(n), (C.GLclampf)(f))
}
func DepthRangexOES(n int32, f int32) {
	C.glowDepthRangexOES(gpDepthRangexOES, (C.GLfixed)(n), (C.GLfixed)(f))
}

// Detaches a shader object from a program object to which it is attached
func DetachShader(program uint32, shader uint32) {
	C.glowDetachShader(gpDetachShader, (C.GLuint)(program), (C.GLuint)(shader))
}
func Disable(cap uint32) {
	C.glowDisable(gpDisable, (C.GLenum)(cap))
}

// Enable or disable a generic vertex attribute     array
func DisableVertexArrayAttrib(vaobj uint32, index uint32) {
	C.glowDisableVertexArrayAttrib(gpDisableVertexArrayAttrib, (C.GLuint)(vaobj), (C.GLuint)(index))
}

// Enable or disable a generic vertex attribute     array
func DisableVertexAttribArray(index uint32) {
	C.glowDisableVertexAttribArray(gpDisableVertexAttribArray, (C.GLuint)(index))
}
func Disablei(target uint32, index uint32) {
	C.glowDisablei(gpDisablei, (C.GLenum)(target), (C.GLuint)(index))
}

// launch one or more compute work groups
func DispatchCompute(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32) {
	C.glowDispatchCompute(gpDispatchCompute, (C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z))
}
func DispatchComputeGroupSizeARB(num_groups_x uint32, num_groups_y uint32, num_groups_z uint32, group_size_x uint32, group_size_y uint32, group_size_z uint32) {
	C.glowDispatchComputeGroupSizeARB(gpDispatchComputeGroupSizeARB, (C.GLuint)(num_groups_x), (C.GLuint)(num_groups_y), (C.GLuint)(num_groups_z), (C.GLuint)(group_size_x), (C.GLuint)(group_size_y), (C.GLuint)(group_size_z))
}

// launch one or more compute work groups using parameters stored in a buffer
func DispatchComputeIndirect(indirect int) {
	C.glowDispatchComputeIndirect(gpDispatchComputeIndirect, (C.GLintptr)(indirect))
}

// render primitives from array data
func DrawArrays(mode uint32, first int32, count int32) {
	C.glowDrawArrays(gpDrawArrays, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count))
}

// render primitives from array data, taking parameters from memory
func DrawArraysIndirect(mode uint32, indirect unsafe.Pointer) {
	C.glowDrawArraysIndirect(gpDrawArraysIndirect, (C.GLenum)(mode), indirect)
}

// draw multiple instances of a range of elements
func DrawArraysInstanced(mode uint32, first int32, count int32, instancecount int32) {
	C.glowDrawArraysInstanced(gpDrawArraysInstanced, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount))
}

// draw multiple instances of a range of elements with offset applied to instanced attributes
func DrawArraysInstancedBaseInstance(mode uint32, first int32, count int32, instancecount int32, baseinstance uint32) {
	C.glowDrawArraysInstancedBaseInstance(gpDrawArraysInstancedBaseInstance, (C.GLenum)(mode), (C.GLint)(first), (C.GLsizei)(count), (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}
func DrawArraysInstancedEXT(mode uint32, start int32, count int32, primcount int32) {
	C.glowDrawArraysInstancedEXT(gpDrawArraysInstancedEXT, (C.GLenum)(mode), (C.GLint)(start), (C.GLsizei)(count), (C.GLsizei)(primcount))
}

// specify which color buffers are to be drawn into
func DrawBuffer(buf uint32) {
	C.glowDrawBuffer(gpDrawBuffer, (C.GLenum)(buf))
}

// Specifies a list of color buffers to be drawn     into
func DrawBuffers(n int32, bufs *uint32) {
	C.glowDrawBuffers(gpDrawBuffers, (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}

// render primitives from array data
func DrawElements(mode uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	C.glowDrawElements(gpDrawElements, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices)
}

// render primitives from array data with a per-element offset
func DrawElementsBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawElementsBaseVertex(gpDrawElementsBaseVertex, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}

// render indexed primitives from array data, taking parameters from memory
func DrawElementsIndirect(mode uint32, xtype uint32, indirect unsafe.Pointer) {
	C.glowDrawElementsIndirect(gpDrawElementsIndirect, (C.GLenum)(mode), (C.GLenum)(xtype), indirect)
}

// draw multiple instances of a set of elements
func DrawElementsInstanced(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32) {
	C.glowDrawElementsInstanced(gpDrawElementsInstanced, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount))
}

// draw multiple instances of a set of elements with offset applied to instanced attributes
func DrawElementsInstancedBaseInstance(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, baseinstance uint32) {
	C.glowDrawElementsInstancedBaseInstance(gpDrawElementsInstancedBaseInstance, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLuint)(baseinstance))
}

// render multiple instances of a set of primitives from array data with a per-element offset
func DrawElementsInstancedBaseVertex(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32) {
	C.glowDrawElementsInstancedBaseVertex(gpDrawElementsInstancedBaseVertex, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex))
}

// render multiple instances of a set of primitives from array data with a per-element offset
func DrawElementsInstancedBaseVertexBaseInstance(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, instancecount int32, basevertex int32, baseinstance uint32) {
	C.glowDrawElementsInstancedBaseVertexBaseInstance(gpDrawElementsInstancedBaseVertexBaseInstance, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(instancecount), (C.GLint)(basevertex), (C.GLuint)(baseinstance))
}
func DrawElementsInstancedEXT(mode uint32, count int32, xtype uint32, indices unsafe.Pointer, primcount int32) {
	C.glowDrawElementsInstancedEXT(gpDrawElementsInstancedEXT, (C.GLenum)(mode), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLsizei)(primcount))
}

// render primitives from array data
func DrawRangeElements(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer) {
	C.glowDrawRangeElements(gpDrawRangeElements, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(xtype), indices)
}

// render primitives from array data with a per-element offset
func DrawRangeElementsBaseVertex(mode uint32, start uint32, end uint32, count int32, xtype uint32, indices unsafe.Pointer, basevertex int32) {
	C.glowDrawRangeElementsBaseVertex(gpDrawRangeElementsBaseVertex, (C.GLenum)(mode), (C.GLuint)(start), (C.GLuint)(end), (C.GLsizei)(count), (C.GLenum)(xtype), indices, (C.GLint)(basevertex))
}

// render primitives using a count derived from a transform feedback object
func DrawTransformFeedback(mode uint32, id uint32) {
	C.glowDrawTransformFeedback(gpDrawTransformFeedback, (C.GLenum)(mode), (C.GLuint)(id))
}

// render multiple instances of primitives using a count derived from a transform feedback object
func DrawTransformFeedbackInstanced(mode uint32, id uint32, instancecount int32) {
	C.glowDrawTransformFeedbackInstanced(gpDrawTransformFeedbackInstanced, (C.GLenum)(mode), (C.GLuint)(id), (C.GLsizei)(instancecount))
}

// render primitives using a count derived from a specifed stream of a transform feedback object
func DrawTransformFeedbackStream(mode uint32, id uint32, stream uint32) {
	C.glowDrawTransformFeedbackStream(gpDrawTransformFeedbackStream, (C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream))
}

// render multiple instances of primitives using a count derived from a specifed stream of a transform feedback object
func DrawTransformFeedbackStreamInstanced(mode uint32, id uint32, stream uint32, instancecount int32) {
	C.glowDrawTransformFeedbackStreamInstanced(gpDrawTransformFeedbackStreamInstanced, (C.GLenum)(mode), (C.GLuint)(id), (C.GLuint)(stream), (C.GLsizei)(instancecount))
}

// enable or disable server-side GL capabilities
func Enable(cap uint32) {
	C.glowEnable(gpEnable, (C.GLenum)(cap))
}

// Enable or disable a generic vertex attribute     array
func EnableVertexArrayAttrib(vaobj uint32, index uint32) {
	C.glowEnableVertexArrayAttrib(gpEnableVertexArrayAttrib, (C.GLuint)(vaobj), (C.GLuint)(index))
}

// Enable or disable a generic vertex attribute     array
func EnableVertexAttribArray(index uint32) {
	C.glowEnableVertexAttribArray(gpEnableVertexAttribArray, (C.GLuint)(index))
}
func Enablei(target uint32, index uint32) {
	C.glowEnablei(gpEnablei, (C.GLenum)(target), (C.GLuint)(index))
}
func EndConditionalRender() {
	C.glowEndConditionalRender(gpEndConditionalRender)
}
func EndPerfMonitorAMD(monitor uint32) {
	C.glowEndPerfMonitorAMD(gpEndPerfMonitorAMD, (C.GLuint)(monitor))
}
func EndPerfQueryINTEL(queryHandle uint32) {
	C.glowEndPerfQueryINTEL(gpEndPerfQueryINTEL, (C.GLuint)(queryHandle))
}
func EndQuery(target uint32) {
	C.glowEndQuery(gpEndQuery, (C.GLenum)(target))
}
func EndQueryIndexed(target uint32, index uint32) {
	C.glowEndQueryIndexed(gpEndQueryIndexed, (C.GLenum)(target), (C.GLuint)(index))
}
func EndTransformFeedback() {
	C.glowEndTransformFeedback(gpEndTransformFeedback)
}
func EvalCoord1xOES(u int32) {
	C.glowEvalCoord1xOES(gpEvalCoord1xOES, (C.GLfixed)(u))
}
func EvalCoord1xvOES(coords *int32) {
	C.glowEvalCoord1xvOES(gpEvalCoord1xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func EvalCoord2xOES(u int32, v int32) {
	C.glowEvalCoord2xOES(gpEvalCoord2xOES, (C.GLfixed)(u), (C.GLfixed)(v))
}
func EvalCoord2xvOES(coords *int32) {
	C.glowEvalCoord2xvOES(gpEvalCoord2xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func FeedbackBufferxOES(n int32, xtype uint32, buffer *int32) {
	C.glowFeedbackBufferxOES(gpFeedbackBufferxOES, (C.GLsizei)(n), (C.GLenum)(xtype), (*C.GLfixed)(unsafe.Pointer(buffer)))
}

// create a new sync object and insert it into the GL command stream
func FenceSync(condition uint32, flags uint32) unsafe.Pointer {
	ret := C.glowFenceSync(gpFenceSync, (C.GLenum)(condition), (C.GLbitfield)(flags))
	return (unsafe.Pointer)(ret)
}

// block until all GL execution is complete
func Finish() {
	C.glowFinish(gpFinish)
}
func FinishFenceNV(fence uint32) {
	C.glowFinishFenceNV(gpFinishFenceNV, (C.GLuint)(fence))
}

// force execution of GL commands in finite time
func Flush() {
	C.glowFlush(gpFlush)
}

// indicate modifications to a range of a mapped buffer
func FlushMappedBufferRange(target uint32, offset int, length int) {
	C.glowFlushMappedBufferRange(gpFlushMappedBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

// indicate modifications to a range of a mapped buffer
func FlushMappedNamedBufferRange(buffer uint32, offset int, length int32) {
	C.glowFlushMappedNamedBufferRange(gpFlushMappedNamedBufferRange, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(length))
}
func FogxOES(pname uint32, param int32) {
	C.glowFogxOES(gpFogxOES, (C.GLenum)(pname), (C.GLfixed)(param))
}
func FogxvOES(pname uint32, param *int32) {
	C.glowFogxvOES(gpFogxvOES, (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(param)))
}

// set a named parameter of a framebuffer object
func FramebufferParameteri(target uint32, pname uint32, param int32) {
	C.glowFramebufferParameteri(gpFramebufferParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}

// attach a renderbuffer as a logical buffer of a framebuffer object
func FramebufferRenderbuffer(target uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	C.glowFramebufferRenderbuffer(gpFramebufferRenderbuffer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}

// attach a level of a texture object as a logical buffer of a framebuffer object
func FramebufferTexture(target uint32, attachment uint32, texture uint32, level int32) {
	C.glowFramebufferTexture(gpFramebufferTexture, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture1D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	C.glowFramebufferTexture1D(gpFramebufferTexture1D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	C.glowFramebufferTexture2D(gpFramebufferTexture2D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level))
}
func FramebufferTexture3D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32, zoffset int32) {
	C.glowFramebufferTexture3D(gpFramebufferTexture3D, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(textarget), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(zoffset))
}

// attach a single layer of a texture object as a logical buffer of a framebuffer object
func FramebufferTextureLayer(target uint32, attachment uint32, texture uint32, level int32, layer int32) {
	C.glowFramebufferTextureLayer(gpFramebufferTextureLayer, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}

// define front- and back-facing polygons
func FrontFace(mode uint32) {
	C.glowFrontFace(gpFrontFace, (C.GLenum)(mode))
}
func FrustumfOES(l float32, r float32, b float32, t float32, n float32, f float32) {
	C.glowFrustumfOES(gpFrustumfOES, (C.GLfloat)(l), (C.GLfloat)(r), (C.GLfloat)(b), (C.GLfloat)(t), (C.GLfloat)(n), (C.GLfloat)(f))
}
func FrustumxOES(l int32, r int32, b int32, t int32, n int32, f int32) {
	C.glowFrustumxOES(gpFrustumxOES, (C.GLfixed)(l), (C.GLfixed)(r), (C.GLfixed)(b), (C.GLfixed)(t), (C.GLfixed)(n), (C.GLfixed)(f))
}

// generate buffer object names
func GenBuffers(n int32, buffers *uint32) {
	C.glowGenBuffers(gpGenBuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(buffers)))
}
func GenFencesNV(n int32, fences *uint32) {
	C.glowGenFencesNV(gpGenFencesNV, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(fences)))
}

// generate framebuffer object names
func GenFramebuffers(n int32, framebuffers *uint32) {
	C.glowGenFramebuffers(gpGenFramebuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(framebuffers)))
}
func GenPerfMonitorsAMD(n int32, monitors *uint32) {
	C.glowGenPerfMonitorsAMD(gpGenPerfMonitorsAMD, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(monitors)))
}

// reserve program pipeline object names
func GenProgramPipelines(n int32, pipelines *uint32) {
	C.glowGenProgramPipelines(gpGenProgramPipelines, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}
func GenProgramPipelinesEXT(n int32, pipelines *uint32) {
	C.glowGenProgramPipelinesEXT(gpGenProgramPipelinesEXT, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(pipelines)))
}

// generate query object names
func GenQueries(n int32, ids *uint32) {
	C.glowGenQueries(gpGenQueries, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// generate renderbuffer object names
func GenRenderbuffers(n int32, renderbuffers *uint32) {
	C.glowGenRenderbuffers(gpGenRenderbuffers, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(renderbuffers)))
}

// generate sampler object names
func GenSamplers(count int32, samplers *uint32) {
	C.glowGenSamplers(gpGenSamplers, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(samplers)))
}

// generate texture names
func GenTextures(n int32, textures *uint32) {
	C.glowGenTextures(gpGenTextures, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)))
}

// reserve transform feedback object names
func GenTransformFeedbacks(n int32, ids *uint32) {
	C.glowGenTransformFeedbacks(gpGenTransformFeedbacks, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(ids)))
}

// generate vertex array object names
func GenVertexArrays(n int32, arrays *uint32) {
	C.glowGenVertexArrays(gpGenVertexArrays, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(arrays)))
}

// generate mipmaps for a specified texture object
func GenerateMipmap(target uint32) {
	C.glowGenerateMipmap(gpGenerateMipmap, (C.GLenum)(target))
}

// generate mipmaps for a specified texture object
func GenerateTextureMipmap(texture uint32) {
	C.glowGenerateTextureMipmap(gpGenerateTextureMipmap, (C.GLuint)(texture))
}

// retrieve information about the set of active atomic counter buffers for a program
func GetActiveAtomicCounterBufferiv(program uint32, bufferIndex uint32, pname uint32, params *int32) {
	C.glowGetActiveAtomicCounterBufferiv(gpGetActiveAtomicCounterBufferiv, (C.GLuint)(program), (C.GLuint)(bufferIndex), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns information about an active attribute variable for the specified program object
func GetActiveAttrib(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetActiveAttrib(gpGetActiveAttrib, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}

// query the name of an active shader subroutine
func GetActiveSubroutineName(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8) {
	C.glowGetActiveSubroutineName(gpGetActiveSubroutineName, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(name)))
}

// query the name of an active shader subroutine uniform
func GetActiveSubroutineUniformName(program uint32, shadertype uint32, index uint32, bufsize int32, length *int32, name *uint8) {
	C.glowGetActiveSubroutineUniformName(gpGetActiveSubroutineUniformName, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLsizei)(bufsize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetActiveSubroutineUniformiv(program uint32, shadertype uint32, index uint32, pname uint32, values *int32) {
	C.glowGetActiveSubroutineUniformiv(gpGetActiveSubroutineUniformiv, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(values)))
}

// Returns information about an active uniform variable for the specified program object
func GetActiveUniform(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetActiveUniform(gpGetActiveUniform, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}

// retrieve the name of an active uniform block
func GetActiveUniformBlockName(program uint32, uniformBlockIndex uint32, bufSize int32, length *int32, uniformBlockName *uint8) {
	C.glowGetActiveUniformBlockName(gpGetActiveUniformBlockName, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(uniformBlockName)))
}
func GetActiveUniformBlockiv(program uint32, uniformBlockIndex uint32, pname uint32, params *int32) {
	C.glowGetActiveUniformBlockiv(gpGetActiveUniformBlockiv, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query the name of an active uniform
func GetActiveUniformName(program uint32, uniformIndex uint32, bufSize int32, length *int32, uniformName *uint8) {
	C.glowGetActiveUniformName(gpGetActiveUniformName, (C.GLuint)(program), (C.GLuint)(uniformIndex), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(uniformName)))
}

// Returns information about several active uniform variables for the specified program object
func GetActiveUniformsiv(program uint32, uniformCount int32, uniformIndices *uint32, pname uint32, params *int32) {
	C.glowGetActiveUniformsiv(gpGetActiveUniformsiv, (C.GLuint)(program), (C.GLsizei)(uniformCount), (*C.GLuint)(unsafe.Pointer(uniformIndices)), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns the handles of the shader objects attached to a program object
func GetAttachedShaders(program uint32, maxCount int32, count *int32, shaders *uint32) {
	C.glowGetAttachedShaders(gpGetAttachedShaders, (C.GLuint)(program), (C.GLsizei)(maxCount), (*C.GLsizei)(unsafe.Pointer(count)), (*C.GLuint)(unsafe.Pointer(shaders)))
}

// Returns the location of an attribute variable
func GetAttribLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetAttribLocation(gpGetAttribLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}
func GetBooleani_v(target uint32, index uint32, data *bool) {
	C.glowGetBooleani_v(gpGetBooleani_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLboolean)(unsafe.Pointer(data)))
}
func GetBooleanv(pname uint32, data *bool) {
	C.glowGetBooleanv(gpGetBooleanv, (C.GLenum)(pname), (*C.GLboolean)(unsafe.Pointer(data)))
}

// return parameters of a buffer object
func GetBufferParameteri64v(target uint32, pname uint32, params *int64) {
	C.glowGetBufferParameteri64v(gpGetBufferParameteri64v, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}

// return parameters of a buffer object
func GetBufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetBufferParameteriv(gpGetBufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return the pointer to a mapped buffer object's data store
func GetBufferPointerv(target uint32, pname uint32, params *unsafe.Pointer) {
	C.glowGetBufferPointerv(gpGetBufferPointerv, (C.GLenum)(target), (C.GLenum)(pname), params)
}

// returns a subset of a buffer object's data store
func GetBufferSubData(target uint32, offset int, size int, data unsafe.Pointer) {
	C.glowGetBufferSubData(gpGetBufferSubData, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(size), data)
}
func GetClipPlanefOES(plane uint32, equation *float32) {
	C.glowGetClipPlanefOES(gpGetClipPlanefOES, (C.GLenum)(plane), (*C.GLfloat)(unsafe.Pointer(equation)))
}
func GetClipPlanexOES(plane uint32, equation *int32) {
	C.glowGetClipPlanexOES(gpGetClipPlanexOES, (C.GLenum)(plane), (*C.GLfixed)(unsafe.Pointer(equation)))
}

// return a compressed texture image
func GetCompressedTexImage(target uint32, level int32, img unsafe.Pointer) {
	C.glowGetCompressedTexImage(gpGetCompressedTexImage, (C.GLenum)(target), (C.GLint)(level), img)
}

// return a compressed texture image
func GetCompressedTextureImage(texture uint32, level int32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetCompressedTextureImage(gpGetCompressedTextureImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLsizei)(bufSize), pixels)
}

// retrieve a sub-region of a compressed texture image from a     compressed texture object
func GetCompressedTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetCompressedTextureSubImage(gpGetCompressedTextureSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLsizei)(bufSize), pixels)
}
func GetConvolutionParameterxvOES(target uint32, pname uint32, params *int32) {
	C.glowGetConvolutionParameterxvOES(gpGetConvolutionParameterxvOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// retrieve messages from the debug message log
func GetDebugMessageLog(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLog(gpGetDebugMessageLog, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDebugMessageLogARB(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLogARB(gpGetDebugMessageLogARB, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDebugMessageLogKHR(count uint32, bufSize int32, sources *uint32, types *uint32, ids *uint32, severities *uint32, lengths *int32, messageLog *uint8) uint32 {
	ret := C.glowGetDebugMessageLogKHR(gpGetDebugMessageLogKHR, (C.GLuint)(count), (C.GLsizei)(bufSize), (*C.GLenum)(unsafe.Pointer(sources)), (*C.GLenum)(unsafe.Pointer(types)), (*C.GLuint)(unsafe.Pointer(ids)), (*C.GLenum)(unsafe.Pointer(severities)), (*C.GLsizei)(unsafe.Pointer(lengths)), (*C.GLchar)(unsafe.Pointer(messageLog)))
	return (uint32)(ret)
}
func GetDoublei_v(target uint32, index uint32, data *float64) {
	C.glowGetDoublei_v(gpGetDoublei_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(data)))
}
func GetDoublev(pname uint32, data *float64) {
	C.glowGetDoublev(gpGetDoublev, (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(data)))
}

// return error information
func GetError() uint32 {
	ret := C.glowGetError(gpGetError)
	return (uint32)(ret)
}
func GetFenceivNV(fence uint32, pname uint32, params *int32) {
	C.glowGetFenceivNV(gpGetFenceivNV, (C.GLuint)(fence), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetFirstPerfQueryIdINTEL(queryId *uint32) {
	C.glowGetFirstPerfQueryIdINTEL(gpGetFirstPerfQueryIdINTEL, (*C.GLuint)(unsafe.Pointer(queryId)))
}
func GetFixedvOES(pname uint32, params *int32) {
	C.glowGetFixedvOES(gpGetFixedvOES, (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func GetFloati_v(target uint32, index uint32, data *float32) {
	C.glowGetFloati_v(gpGetFloati_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(data)))
}
func GetFloatv(pname uint32, data *float32) {
	C.glowGetFloatv(gpGetFloatv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(data)))
}

// query the bindings of color indices to user-defined varying out variables
func GetFragDataIndex(program uint32, name *uint8) int32 {
	ret := C.glowGetFragDataIndex(gpGetFragDataIndex, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the bindings of color numbers to user-defined varying out variables
func GetFragDataLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetFragDataLocation(gpGetFragDataLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// retrieve information about attachments of a framebuffer object
func GetFramebufferAttachmentParameteriv(target uint32, attachment uint32, pname uint32, params *int32) {
	C.glowGetFramebufferAttachmentParameteriv(gpGetFramebufferAttachmentParameteriv, (C.GLenum)(target), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query a named parameter of a framebuffer object
func GetFramebufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetFramebufferParameteriv(gpGetFramebufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// check if the rendering context has not been lost due to software or hardware issues
func GetGraphicsResetStatus() uint32 {
	ret := C.glowGetGraphicsResetStatus(gpGetGraphicsResetStatus)
	return (uint32)(ret)
}
func GetGraphicsResetStatusARB() uint32 {
	ret := C.glowGetGraphicsResetStatusARB(gpGetGraphicsResetStatusARB)
	return (uint32)(ret)
}
func GetGraphicsResetStatusKHR() uint32 {
	ret := C.glowGetGraphicsResetStatusKHR(gpGetGraphicsResetStatusKHR)
	return (uint32)(ret)
}
func GetHistogramParameterxvOES(target uint32, pname uint32, params *int32) {
	C.glowGetHistogramParameterxvOES(gpGetHistogramParameterxvOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func GetImageHandleARB(texture uint32, level int32, layered bool, layer int32, format uint32) uint64 {
	ret := C.glowGetImageHandleARB(gpGetImageHandleARB, (C.GLuint)(texture), (C.GLint)(level), (C.GLboolean)(boolToInt(layered)), (C.GLint)(layer), (C.GLenum)(format))
	return (uint64)(ret)
}
func GetInteger64i_v(target uint32, index uint32, data *int64) {
	C.glowGetInteger64i_v(gpGetInteger64i_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetInteger64v(pname uint32, data *int64) {
	C.glowGetInteger64v(gpGetInteger64v, (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(data)))
}
func GetIntegeri_v(target uint32, index uint32, data *int32) {
	C.glowGetIntegeri_v(gpGetIntegeri_v, (C.GLenum)(target), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(data)))
}
func GetIntegerv(pname uint32, data *int32) {
	C.glowGetIntegerv(gpGetIntegerv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(data)))
}
func GetInternalformati64v(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int64) {
	C.glowGetInternalformati64v(gpGetInternalformati64v, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetInternalformativ(target uint32, internalformat uint32, pname uint32, bufSize int32, params *int32) {
	C.glowGetInternalformativ(gpGetInternalformativ, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetLightxOES(light uint32, pname uint32, params *int32) {
	C.glowGetLightxOES(gpGetLightxOES, (C.GLenum)(light), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func GetLightxvOES(light uint32, pname uint32, params *int32) {
	C.glowGetLightxvOES(gpGetLightxvOES, (C.GLenum)(light), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func GetMapxvOES(target uint32, query uint32, v *int32) {
	C.glowGetMapxvOES(gpGetMapxvOES, (C.GLenum)(target), (C.GLenum)(query), (*C.GLfixed)(unsafe.Pointer(v)))
}
func GetMaterialxOES(face uint32, pname uint32, param int32) {
	C.glowGetMaterialxOES(gpGetMaterialxOES, (C.GLenum)(face), (C.GLenum)(pname), (C.GLfixed)(param))
}
func GetMaterialxvOES(face uint32, pname uint32, params *int32) {
	C.glowGetMaterialxvOES(gpGetMaterialxvOES, (C.GLenum)(face), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// retrieve the location of a sample
func GetMultisamplefv(pname uint32, index uint32, val *float32) {
	C.glowGetMultisamplefv(gpGetMultisamplefv, (C.GLenum)(pname), (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(val)))
}

// return parameters of a buffer object
func GetNamedBufferParameteri64v(buffer uint32, pname uint32, params *int64) {
	C.glowGetNamedBufferParameteri64v(gpGetNamedBufferParameteri64v, (C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}

// return parameters of a buffer object
func GetNamedBufferParameteriv(buffer uint32, pname uint32, params *int32) {
	C.glowGetNamedBufferParameteriv(gpGetNamedBufferParameteriv, (C.GLuint)(buffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return the pointer to a mapped buffer object's data store
func GetNamedBufferPointerv(buffer uint32, pname uint32, params *unsafe.Pointer) {
	C.glowGetNamedBufferPointerv(gpGetNamedBufferPointerv, (C.GLuint)(buffer), (C.GLenum)(pname), params)
}

// returns a subset of a buffer object's data store
func GetNamedBufferSubData(buffer uint32, offset int, size int32, data unsafe.Pointer) {
	C.glowGetNamedBufferSubData(gpGetNamedBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), data)
}

// retrieve information about attachments of a framebuffer object
func GetNamedFramebufferAttachmentParameteriv(framebuffer uint32, attachment uint32, pname uint32, params *int32) {
	C.glowGetNamedFramebufferAttachmentParameteriv(gpGetNamedFramebufferAttachmentParameteriv, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query a named parameter of a framebuffer object
func GetNamedFramebufferParameteriv(framebuffer uint32, pname uint32, param *int32) {
	C.glowGetNamedFramebufferParameteriv(gpGetNamedFramebufferParameteriv, (C.GLuint)(framebuffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// query a named parameter of a renderbuffer object
func GetNamedRenderbufferParameteriv(renderbuffer uint32, pname uint32, params *int32) {
	C.glowGetNamedRenderbufferParameteriv(gpGetNamedRenderbufferParameteriv, (C.GLuint)(renderbuffer), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetNamedStringARB(namelen int32, name *uint8, bufSize int32, stringlen *int32, xstring *uint8) {
	C.glowGetNamedStringARB(gpGetNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(stringlen)), (*C.GLchar)(unsafe.Pointer(xstring)))
}
func GetNamedStringivARB(namelen int32, name *uint8, pname uint32, params *int32) {
	C.glowGetNamedStringivARB(gpGetNamedStringivARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetNextPerfQueryIdINTEL(queryId uint32, nextQueryId *uint32) {
	C.glowGetNextPerfQueryIdINTEL(gpGetNextPerfQueryIdINTEL, (C.GLuint)(queryId), (*C.GLuint)(unsafe.Pointer(nextQueryId)))
}

// retrieve the label of a named object identified within a namespace
func GetObjectLabel(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabel(gpGetObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectLabelEXT(xtype uint32, object uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabelEXT(gpGetObjectLabelEXT, (C.GLenum)(xtype), (C.GLuint)(object), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectLabelKHR(identifier uint32, name uint32, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectLabelKHR(gpGetObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}

// retrieve the label of a sync object identified by a pointer
func GetObjectPtrLabel(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectPtrLabel(gpGetObjectPtrLabel, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetObjectPtrLabelKHR(ptr unsafe.Pointer, bufSize int32, length *int32, label *uint8) {
	C.glowGetObjectPtrLabelKHR(gpGetObjectPtrLabelKHR, ptr, (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(label)))
}
func GetPerfCounterInfoINTEL(queryId uint32, counterId uint32, counterNameLength uint32, counterName *uint8, counterDescLength uint32, counterDesc *uint8, counterOffset *uint32, counterDataSize *uint32, counterTypeEnum *uint32, counterDataTypeEnum *uint32, rawCounterMaxValue *uint64) {
	C.glowGetPerfCounterInfoINTEL(gpGetPerfCounterInfoINTEL, (C.GLuint)(queryId), (C.GLuint)(counterId), (C.GLuint)(counterNameLength), (*C.GLchar)(unsafe.Pointer(counterName)), (C.GLuint)(counterDescLength), (*C.GLchar)(unsafe.Pointer(counterDesc)), (*C.GLuint)(unsafe.Pointer(counterOffset)), (*C.GLuint)(unsafe.Pointer(counterDataSize)), (*C.GLuint)(unsafe.Pointer(counterTypeEnum)), (*C.GLuint)(unsafe.Pointer(counterDataTypeEnum)), (*C.GLuint64)(unsafe.Pointer(rawCounterMaxValue)))
}
func GetPerfMonitorCounterDataAMD(monitor uint32, pname uint32, dataSize int32, data *uint32, bytesWritten *int32) {
	C.glowGetPerfMonitorCounterDataAMD(gpGetPerfMonitorCounterDataAMD, (C.GLuint)(monitor), (C.GLenum)(pname), (C.GLsizei)(dataSize), (*C.GLuint)(unsafe.Pointer(data)), (*C.GLint)(unsafe.Pointer(bytesWritten)))
}
func GetPerfMonitorCounterInfoAMD(group uint32, counter uint32, pname uint32, data unsafe.Pointer) {
	C.glowGetPerfMonitorCounterInfoAMD(gpGetPerfMonitorCounterInfoAMD, (C.GLuint)(group), (C.GLuint)(counter), (C.GLenum)(pname), data)
}
func GetPerfMonitorCounterStringAMD(group uint32, counter uint32, bufSize int32, length *int32, counterString *uint8) {
	C.glowGetPerfMonitorCounterStringAMD(gpGetPerfMonitorCounterStringAMD, (C.GLuint)(group), (C.GLuint)(counter), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(counterString)))
}
func GetPerfMonitorCountersAMD(group uint32, numCounters *int32, maxActiveCounters *int32, counterSize int32, counters *uint32) {
	C.glowGetPerfMonitorCountersAMD(gpGetPerfMonitorCountersAMD, (C.GLuint)(group), (*C.GLint)(unsafe.Pointer(numCounters)), (*C.GLint)(unsafe.Pointer(maxActiveCounters)), (C.GLsizei)(counterSize), (*C.GLuint)(unsafe.Pointer(counters)))
}
func GetPerfMonitorGroupStringAMD(group uint32, bufSize int32, length *int32, groupString *uint8) {
	C.glowGetPerfMonitorGroupStringAMD(gpGetPerfMonitorGroupStringAMD, (C.GLuint)(group), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(groupString)))
}
func GetPerfMonitorGroupsAMD(numGroups *int32, groupsSize int32, groups *uint32) {
	C.glowGetPerfMonitorGroupsAMD(gpGetPerfMonitorGroupsAMD, (*C.GLint)(unsafe.Pointer(numGroups)), (C.GLsizei)(groupsSize), (*C.GLuint)(unsafe.Pointer(groups)))
}
func GetPerfQueryDataINTEL(queryHandle uint32, flags uint32, dataSize int32, data unsafe.Pointer, bytesWritten *uint32) {
	C.glowGetPerfQueryDataINTEL(gpGetPerfQueryDataINTEL, (C.GLuint)(queryHandle), (C.GLuint)(flags), (C.GLsizei)(dataSize), data, (*C.GLuint)(unsafe.Pointer(bytesWritten)))
}
func GetPerfQueryIdByNameINTEL(queryName *uint8, queryId *uint32) {
	C.glowGetPerfQueryIdByNameINTEL(gpGetPerfQueryIdByNameINTEL, (*C.GLchar)(unsafe.Pointer(queryName)), (*C.GLuint)(unsafe.Pointer(queryId)))
}
func GetPerfQueryInfoINTEL(queryId uint32, queryNameLength uint32, queryName *uint8, dataSize *uint32, noCounters *uint32, noInstances *uint32, capsMask *uint32) {
	C.glowGetPerfQueryInfoINTEL(gpGetPerfQueryInfoINTEL, (C.GLuint)(queryId), (C.GLuint)(queryNameLength), (*C.GLchar)(unsafe.Pointer(queryName)), (*C.GLuint)(unsafe.Pointer(dataSize)), (*C.GLuint)(unsafe.Pointer(noCounters)), (*C.GLuint)(unsafe.Pointer(noInstances)), (*C.GLuint)(unsafe.Pointer(capsMask)))
}
func GetPixelMapxv(xmap uint32, size int32, values *int32) {
	C.glowGetPixelMapxv(gpGetPixelMapxv, (C.GLenum)(xmap), (C.GLint)(size), (*C.GLfixed)(unsafe.Pointer(values)))
}

// return the address of the specified pointer
func GetPointerv(pname uint32, params *unsafe.Pointer) {
	C.glowGetPointerv(gpGetPointerv, (C.GLenum)(pname), params)
}
func GetPointervKHR(pname uint32, params *unsafe.Pointer) {
	C.glowGetPointervKHR(gpGetPointervKHR, (C.GLenum)(pname), params)
}

// return a binary representation of a program object's compiled and linked executable source
func GetProgramBinary(program uint32, bufSize int32, length *int32, binaryFormat *uint32, binary unsafe.Pointer) {
	C.glowGetProgramBinary(gpGetProgramBinary, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLenum)(unsafe.Pointer(binaryFormat)), binary)
}

// Returns the information log for a program object
func GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramInfoLog(gpGetProgramInfoLog, (C.GLuint)(program), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramInterfaceiv(program uint32, programInterface uint32, pname uint32, params *int32) {
	C.glowGetProgramInterfaceiv(gpGetProgramInterfaceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// retrieve the info log string from a program pipeline object
func GetProgramPipelineInfoLog(pipeline uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramPipelineInfoLog(gpGetProgramPipelineInfoLog, (C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramPipelineInfoLogEXT(pipeline uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetProgramPipelineInfoLogEXT(gpGetProgramPipelineInfoLogEXT, (C.GLuint)(pipeline), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}
func GetProgramPipelineiv(pipeline uint32, pname uint32, params *int32) {
	C.glowGetProgramPipelineiv(gpGetProgramPipelineiv, (C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetProgramPipelineivEXT(pipeline uint32, pname uint32, params *int32) {
	C.glowGetProgramPipelineivEXT(gpGetProgramPipelineivEXT, (C.GLuint)(pipeline), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query the index of a named resource within a program
func GetProgramResourceIndex(program uint32, programInterface uint32, name *uint8) uint32 {
	ret := C.glowGetProgramResourceIndex(gpGetProgramResourceIndex, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (uint32)(ret)
}

// query the location of a named resource within a program
func GetProgramResourceLocation(program uint32, programInterface uint32, name *uint8) int32 {
	ret := C.glowGetProgramResourceLocation(gpGetProgramResourceLocation, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the fragment color index of a named variable within a program
func GetProgramResourceLocationIndex(program uint32, programInterface uint32, name *uint8) int32 {
	ret := C.glowGetProgramResourceLocationIndex(gpGetProgramResourceLocationIndex, (C.GLuint)(program), (C.GLenum)(programInterface), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the name of an indexed resource within a program
func GetProgramResourceName(program uint32, programInterface uint32, index uint32, bufSize int32, length *int32, name *uint8) {
	C.glowGetProgramResourceName(gpGetProgramResourceName, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetProgramResourceiv(program uint32, programInterface uint32, index uint32, propCount int32, props *uint32, bufSize int32, length *int32, params *int32) {
	C.glowGetProgramResourceiv(gpGetProgramResourceiv, (C.GLuint)(program), (C.GLenum)(programInterface), (C.GLuint)(index), (C.GLsizei)(propCount), (*C.GLenum)(unsafe.Pointer(props)), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(params)))
}
func GetProgramStageiv(program uint32, shadertype uint32, pname uint32, values *int32) {
	C.glowGetProgramStageiv(gpGetProgramStageiv, (C.GLuint)(program), (C.GLenum)(shadertype), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(values)))
}

// Returns a parameter from a program object
func GetProgramiv(program uint32, pname uint32, params *int32) {
	C.glowGetProgramiv(gpGetProgramiv, (C.GLuint)(program), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return parameters of an indexed query object target
func GetQueryIndexediv(target uint32, index uint32, pname uint32, params *int32) {
	C.glowGetQueryIndexediv(gpGetQueryIndexediv, (C.GLenum)(target), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjecti64v(id uint32, pname uint32, params *int64) {
	C.glowGetQueryObjecti64v(gpGetQueryObjecti64v, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(params)))
}
func GetQueryObjectiv(id uint32, pname uint32, params *int32) {
	C.glowGetQueryObjectiv(gpGetQueryObjectiv, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetQueryObjectui64v(id uint32, pname uint32, params *uint64) {
	C.glowGetQueryObjectui64v(gpGetQueryObjectui64v, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint64)(unsafe.Pointer(params)))
}
func GetQueryObjectuiv(id uint32, pname uint32, params *uint32) {
	C.glowGetQueryObjectuiv(gpGetQueryObjectuiv, (C.GLuint)(id), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}

// return parameters of a query object target
func GetQueryiv(target uint32, pname uint32, params *int32) {
	C.glowGetQueryiv(gpGetQueryiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// query a named parameter of a renderbuffer object
func GetRenderbufferParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetRenderbufferParameteriv(gpGetRenderbufferParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIiv(sampler uint32, pname uint32, params *int32) {
	C.glowGetSamplerParameterIiv(gpGetSamplerParameterIiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetSamplerParameterIuiv(sampler uint32, pname uint32, params *uint32) {
	C.glowGetSamplerParameterIuiv(gpGetSamplerParameterIuiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetSamplerParameterfv(sampler uint32, pname uint32, params *float32) {
	C.glowGetSamplerParameterfv(gpGetSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetSamplerParameteriv(sampler uint32, pname uint32, params *int32) {
	C.glowGetSamplerParameteriv(gpGetSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Returns the information log for a shader object
func GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8) {
	C.glowGetShaderInfoLog(gpGetShaderInfoLog, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(infoLog)))
}

// retrieve the range and precision for numeric formats supported by the shader compiler
func GetShaderPrecisionFormat(shadertype uint32, precisiontype uint32, xrange *int32, precision *int32) {
	C.glowGetShaderPrecisionFormat(gpGetShaderPrecisionFormat, (C.GLenum)(shadertype), (C.GLenum)(precisiontype), (*C.GLint)(unsafe.Pointer(xrange)), (*C.GLint)(unsafe.Pointer(precision)))
}

// Returns the source code string from a shader object
func GetShaderSource(shader uint32, bufSize int32, length *int32, source *uint8) {
	C.glowGetShaderSource(gpGetShaderSource, (C.GLuint)(shader), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLchar)(unsafe.Pointer(source)))
}

// Returns a parameter from a shader object
func GetShaderiv(shader uint32, pname uint32, params *int32) {
	C.glowGetShaderiv(gpGetShaderiv, (C.GLuint)(shader), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// return a string describing the current GL connection
func GetString(name uint32) *uint8 {
	ret := C.glowGetString(gpGetString, (C.GLenum)(name))
	return (*uint8)(ret)
}
func GetStringi(name uint32, index uint32) *uint8 {
	ret := C.glowGetStringi(gpGetStringi, (C.GLenum)(name), (C.GLuint)(index))
	return (*uint8)(ret)
}

// retrieve the index of a subroutine uniform of a given shader stage within a program
func GetSubroutineIndex(program uint32, shadertype uint32, name *uint8) uint32 {
	ret := C.glowGetSubroutineIndex(gpGetSubroutineIndex, (C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(unsafe.Pointer(name)))
	return (uint32)(ret)
}

// retrieve the location of a subroutine uniform of a given shader stage within a program
func GetSubroutineUniformLocation(program uint32, shadertype uint32, name *uint8) int32 {
	ret := C.glowGetSubroutineUniformLocation(gpGetSubroutineUniformLocation, (C.GLuint)(program), (C.GLenum)(shadertype), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}

// query the properties of a sync object
func GetSynciv(sync unsafe.Pointer, pname uint32, bufSize int32, length *int32, values *int32) {
	C.glowGetSynciv(gpGetSynciv, (C.GLsync)(sync), (C.GLenum)(pname), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLint)(unsafe.Pointer(values)))
}
func GetTexEnvxvOES(target uint32, pname uint32, params *int32) {
	C.glowGetTexEnvxvOES(gpGetTexEnvxvOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func GetTexGenxvOES(coord uint32, pname uint32, params *int32) {
	C.glowGetTexGenxvOES(gpGetTexGenxvOES, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// return a texture image
func GetTexImage(target uint32, level int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowGetTexImage(gpGetTexImage, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}
func GetTexLevelParameterfv(target uint32, level int32, pname uint32, params *float32) {
	C.glowGetTexLevelParameterfv(gpGetTexLevelParameterfv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexLevelParameteriv(target uint32, level int32, pname uint32, params *int32) {
	C.glowGetTexLevelParameteriv(gpGetTexLevelParameteriv, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexLevelParameterxvOES(target uint32, level int32, pname uint32, params *int32) {
	C.glowGetTexLevelParameterxvOES(gpGetTexLevelParameterxvOES, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func GetTexParameterIiv(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameterIiv(gpGetTexParameterIiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterIuiv(target uint32, pname uint32, params *uint32) {
	C.glowGetTexParameterIuiv(gpGetTexParameterIuiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTexParameterfv(target uint32, pname uint32, params *float32) {
	C.glowGetTexParameterfv(gpGetTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTexParameteriv(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameteriv(gpGetTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTexParameterxvOES(target uint32, pname uint32, params *int32) {
	C.glowGetTexParameterxvOES(gpGetTexParameterxvOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func GetTextureHandleARB(texture uint32) uint64 {
	ret := C.glowGetTextureHandleARB(gpGetTextureHandleARB, (C.GLuint)(texture))
	return (uint64)(ret)
}

// return a texture image
func GetTextureImage(texture uint32, level int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetTextureImage(gpGetTextureImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), pixels)
}
func GetTextureLevelParameterfv(texture uint32, level int32, pname uint32, params *float32) {
	C.glowGetTextureLevelParameterfv(gpGetTextureLevelParameterfv, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTextureLevelParameteriv(texture uint32, level int32, pname uint32, params *int32) {
	C.glowGetTextureLevelParameteriv(gpGetTextureLevelParameteriv, (C.GLuint)(texture), (C.GLint)(level), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureParameterIiv(texture uint32, pname uint32, params *int32) {
	C.glowGetTextureParameterIiv(gpGetTextureParameterIiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureParameterIuiv(texture uint32, pname uint32, params *uint32) {
	C.glowGetTextureParameterIuiv(gpGetTextureParameterIuiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetTextureParameterfv(texture uint32, pname uint32, params *float32) {
	C.glowGetTextureParameterfv(gpGetTextureParameterfv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetTextureParameteriv(texture uint32, pname uint32, params *int32) {
	C.glowGetTextureParameteriv(gpGetTextureParameteriv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetTextureSamplerHandleARB(texture uint32, sampler uint32) uint64 {
	ret := C.glowGetTextureSamplerHandleARB(gpGetTextureSamplerHandleARB, (C.GLuint)(texture), (C.GLuint)(sampler))
	return (uint64)(ret)
}

// retrieve a sub-region of a texture image from a texture     object
func GetTextureSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, bufSize int32, pixels unsafe.Pointer) {
	C.glowGetTextureSubImage(gpGetTextureSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), pixels)
}

// retrieve information about varying variables selected for transform feedback
func GetTransformFeedbackVarying(program uint32, index uint32, bufSize int32, length *int32, size *int32, xtype *uint32, name *uint8) {
	C.glowGetTransformFeedbackVarying(gpGetTransformFeedbackVarying, (C.GLuint)(program), (C.GLuint)(index), (C.GLsizei)(bufSize), (*C.GLsizei)(unsafe.Pointer(length)), (*C.GLsizei)(unsafe.Pointer(size)), (*C.GLenum)(unsafe.Pointer(xtype)), (*C.GLchar)(unsafe.Pointer(name)))
}
func GetTransformFeedbacki64_v(xfb uint32, pname uint32, index uint32, param *int64) {
	C.glowGetTransformFeedbacki64_v(gpGetTransformFeedbacki64_v, (C.GLuint)(xfb), (C.GLenum)(pname), (C.GLuint)(index), (*C.GLint64)(unsafe.Pointer(param)))
}
func GetTransformFeedbacki_v(xfb uint32, pname uint32, index uint32, param *int32) {
	C.glowGetTransformFeedbacki_v(gpGetTransformFeedbacki_v, (C.GLuint)(xfb), (C.GLenum)(pname), (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(param)))
}

// query the state of a transform feedback object.
func GetTransformFeedbackiv(xfb uint32, pname uint32, param *int32) {
	C.glowGetTransformFeedbackiv(gpGetTransformFeedbackiv, (C.GLuint)(xfb), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// retrieve the index of a named uniform block
func GetUniformBlockIndex(program uint32, uniformBlockName *uint8) uint32 {
	ret := C.glowGetUniformBlockIndex(gpGetUniformBlockIndex, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(uniformBlockName)))
	return (uint32)(ret)
}

// retrieve the index of a named uniform block
func GetUniformIndices(program uint32, uniformCount int32, uniformNames **uint8, uniformIndices *uint32) {
	C.glowGetUniformIndices(gpGetUniformIndices, (C.GLuint)(program), (C.GLsizei)(uniformCount), (**C.GLchar)(unsafe.Pointer(uniformNames)), (*C.GLuint)(unsafe.Pointer(uniformIndices)))
}

// Returns the location of a uniform variable
func GetUniformLocation(program uint32, name *uint8) int32 {
	ret := C.glowGetUniformLocation(gpGetUniformLocation, (C.GLuint)(program), (*C.GLchar)(unsafe.Pointer(name)))
	return (int32)(ret)
}
func GetUniformSubroutineuiv(shadertype uint32, location int32, params *uint32) {
	C.glowGetUniformSubroutineuiv(gpGetUniformSubroutineuiv, (C.GLenum)(shadertype), (C.GLint)(location), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetUniformdv(program uint32, location int32, params *float64) {
	C.glowGetUniformdv(gpGetUniformdv, (C.GLuint)(program), (C.GLint)(location), (*C.GLdouble)(unsafe.Pointer(params)))
}

// Returns the value of a uniform variable
func GetUniformfv(program uint32, location int32, params *float32) {
	C.glowGetUniformfv(gpGetUniformfv, (C.GLuint)(program), (C.GLint)(location), (*C.GLfloat)(unsafe.Pointer(params)))
}

// Returns the value of a uniform variable
func GetUniformiv(program uint32, location int32, params *int32) {
	C.glowGetUniformiv(gpGetUniformiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLint)(unsafe.Pointer(params)))
}
func GetUniformuiv(program uint32, location int32, params *uint32) {
	C.glowGetUniformuiv(gpGetUniformuiv, (C.GLuint)(program), (C.GLint)(location), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetVertexArrayIndexed64iv(vaobj uint32, index uint32, pname uint32, param *int64) {
	C.glowGetVertexArrayIndexed64iv(gpGetVertexArrayIndexed64iv, (C.GLuint)(vaobj), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint64)(unsafe.Pointer(param)))
}
func GetVertexArrayIndexediv(vaobj uint32, index uint32, pname uint32, param *int32) {
	C.glowGetVertexArrayIndexediv(gpGetVertexArrayIndexediv, (C.GLuint)(vaobj), (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// retrieve parameters of a vertex array object
func GetVertexArrayiv(vaobj uint32, pname uint32, param *int32) {
	C.glowGetVertexArrayiv(gpGetVertexArrayiv, (C.GLuint)(vaobj), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribIiv(index uint32, pname uint32, params *int32) {
	C.glowGetVertexAttribIiv(gpGetVertexAttribIiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribIuiv(index uint32, pname uint32, params *uint32) {
	C.glowGetVertexAttribIuiv(gpGetVertexAttribIuiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribLdv(index uint32, pname uint32, params *float64) {
	C.glowGetVertexAttribLdv(gpGetVertexAttribLdv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetVertexAttribLui64vARB(index uint32, pname uint32, params *uint64) {
	C.glowGetVertexAttribLui64vARB(gpGetVertexAttribLui64vARB, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLuint64EXT)(unsafe.Pointer(params)))
}

// return the address of the specified generic vertex attribute pointer
func GetVertexAttribPointerv(index uint32, pname uint32, pointer *unsafe.Pointer) {
	C.glowGetVertexAttribPointerv(gpGetVertexAttribPointerv, (C.GLuint)(index), (C.GLenum)(pname), pointer)
}

// Return a generic vertex attribute parameter
func GetVertexAttribdv(index uint32, pname uint32, params *float64) {
	C.glowGetVertexAttribdv(gpGetVertexAttribdv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLdouble)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribfv(index uint32, pname uint32, params *float32) {
	C.glowGetVertexAttribfv(gpGetVertexAttribfv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}

// Return a generic vertex attribute parameter
func GetVertexAttribiv(index uint32, pname uint32, params *int32) {
	C.glowGetVertexAttribiv(gpGetVertexAttribiv, (C.GLuint)(index), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnCompressedTexImageARB(target uint32, lod int32, bufSize int32, img unsafe.Pointer) {
	C.glowGetnCompressedTexImageARB(gpGetnCompressedTexImageARB, (C.GLenum)(target), (C.GLint)(lod), (C.GLsizei)(bufSize), img)
}
func GetnTexImageARB(target uint32, level int32, format uint32, xtype uint32, bufSize int32, img unsafe.Pointer) {
	C.glowGetnTexImageARB(gpGetnTexImageARB, (C.GLenum)(target), (C.GLint)(level), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), img)
}
func GetnUniformdvARB(program uint32, location int32, bufSize int32, params *float64) {
	C.glowGetnUniformdvARB(gpGetnUniformdvARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLdouble)(unsafe.Pointer(params)))
}
func GetnUniformfv(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfv(gpGetnUniformfv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvARB(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfvARB(gpGetnUniformfvARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformfvKHR(program uint32, location int32, bufSize int32, params *float32) {
	C.glowGetnUniformfvKHR(gpGetnUniformfvKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLfloat)(unsafe.Pointer(params)))
}
func GetnUniformiv(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformiv(gpGetnUniformiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivARB(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformivARB(gpGetnUniformivARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformivKHR(program uint32, location int32, bufSize int32, params *int32) {
	C.glowGetnUniformivKHR(gpGetnUniformivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLint)(unsafe.Pointer(params)))
}
func GetnUniformuiv(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuiv(gpGetnUniformuiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetnUniformuivARB(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuivARB(gpGetnUniformuivARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}
func GetnUniformuivKHR(program uint32, location int32, bufSize int32, params *uint32) {
	C.glowGetnUniformuivKHR(gpGetnUniformuivKHR, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(bufSize), (*C.GLuint)(unsafe.Pointer(params)))
}

// specify implementation-specific hints
func Hint(target uint32, mode uint32) {
	C.glowHint(gpHint, (C.GLenum)(target), (C.GLenum)(mode))
}
func IndexxOES(component int32) {
	C.glowIndexxOES(gpIndexxOES, (C.GLfixed)(component))
}
func IndexxvOES(component *int32) {
	C.glowIndexxvOES(gpIndexxvOES, (*C.GLfixed)(unsafe.Pointer(component)))
}
func InsertEventMarkerEXT(length int32, marker *uint8) {
	C.glowInsertEventMarkerEXT(gpInsertEventMarkerEXT, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(marker)))
}

// invalidate the content of a buffer object's data store
func InvalidateBufferData(buffer uint32) {
	C.glowInvalidateBufferData(gpInvalidateBufferData, (C.GLuint)(buffer))
}

// invalidate a region of a buffer object's data store
func InvalidateBufferSubData(buffer uint32, offset int, length int) {
	C.glowInvalidateBufferSubData(gpInvalidateBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(length))
}

// invalidate the content of some or all of a framebuffer's attachments
func InvalidateFramebuffer(target uint32, numAttachments int32, attachments *uint32) {
	C.glowInvalidateFramebuffer(gpInvalidateFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}

// invalidate the content of some or all of a framebuffer's attachments
func InvalidateNamedFramebufferData(framebuffer uint32, numAttachments int32, attachments *uint32) {
	C.glowInvalidateNamedFramebufferData(gpInvalidateNamedFramebufferData, (C.GLuint)(framebuffer), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)))
}

// invalidate the content of a region of some or all of a framebuffer's attachments
func InvalidateNamedFramebufferSubData(framebuffer uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32) {
	C.glowInvalidateNamedFramebufferSubData(gpInvalidateNamedFramebufferSubData, (C.GLuint)(framebuffer), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// invalidate the content of a region of some or all of a framebuffer's attachments
func InvalidateSubFramebuffer(target uint32, numAttachments int32, attachments *uint32, x int32, y int32, width int32, height int32) {
	C.glowInvalidateSubFramebuffer(gpInvalidateSubFramebuffer, (C.GLenum)(target), (C.GLsizei)(numAttachments), (*C.GLenum)(unsafe.Pointer(attachments)), (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}

// invalidate the entirety a texture image
func InvalidateTexImage(texture uint32, level int32) {
	C.glowInvalidateTexImage(gpInvalidateTexImage, (C.GLuint)(texture), (C.GLint)(level))
}

// invalidate a region of a texture image
func InvalidateTexSubImage(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32) {
	C.glowInvalidateTexSubImage(gpInvalidateTexSubImage, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

// determine if a name corresponds to a buffer object
func IsBuffer(buffer uint32) bool {
	ret := C.glowIsBuffer(gpIsBuffer, (C.GLuint)(buffer))
	return ret == TRUE
}
func IsEnabled(cap uint32) bool {
	ret := C.glowIsEnabled(gpIsEnabled, (C.GLenum)(cap))
	return ret == TRUE
}
func IsEnabledi(target uint32, index uint32) bool {
	ret := C.glowIsEnabledi(gpIsEnabledi, (C.GLenum)(target), (C.GLuint)(index))
	return ret == TRUE
}
func IsFenceNV(fence uint32) bool {
	ret := C.glowIsFenceNV(gpIsFenceNV, (C.GLuint)(fence))
	return ret == TRUE
}

// determine if a name corresponds to a framebuffer object
func IsFramebuffer(framebuffer uint32) bool {
	ret := C.glowIsFramebuffer(gpIsFramebuffer, (C.GLuint)(framebuffer))
	return ret == TRUE
}
func IsImageHandleResidentARB(handle uint64) bool {
	ret := C.glowIsImageHandleResidentARB(gpIsImageHandleResidentARB, (C.GLuint64)(handle))
	return ret == TRUE
}
func IsNamedStringARB(namelen int32, name *uint8) bool {
	ret := C.glowIsNamedStringARB(gpIsNamedStringARB, (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)))
	return ret == TRUE
}

// Determines if a name corresponds to a program object
func IsProgram(program uint32) bool {
	ret := C.glowIsProgram(gpIsProgram, (C.GLuint)(program))
	return ret == TRUE
}

// determine if a name corresponds to a program pipeline object
func IsProgramPipeline(pipeline uint32) bool {
	ret := C.glowIsProgramPipeline(gpIsProgramPipeline, (C.GLuint)(pipeline))
	return ret == TRUE
}
func IsProgramPipelineEXT(pipeline uint32) bool {
	ret := C.glowIsProgramPipelineEXT(gpIsProgramPipelineEXT, (C.GLuint)(pipeline))
	return ret == TRUE
}

// determine if a name corresponds to a query object
func IsQuery(id uint32) bool {
	ret := C.glowIsQuery(gpIsQuery, (C.GLuint)(id))
	return ret == TRUE
}

// determine if a name corresponds to a renderbuffer object
func IsRenderbuffer(renderbuffer uint32) bool {
	ret := C.glowIsRenderbuffer(gpIsRenderbuffer, (C.GLuint)(renderbuffer))
	return ret == TRUE
}

// determine if a name corresponds to a sampler object
func IsSampler(sampler uint32) bool {
	ret := C.glowIsSampler(gpIsSampler, (C.GLuint)(sampler))
	return ret == TRUE
}

// Determines if a name corresponds to a shader object
func IsShader(shader uint32) bool {
	ret := C.glowIsShader(gpIsShader, (C.GLuint)(shader))
	return ret == TRUE
}

// determine if a name corresponds to a sync object
func IsSync(sync unsafe.Pointer) bool {
	ret := C.glowIsSync(gpIsSync, (C.GLsync)(sync))
	return ret == TRUE
}

// determine if a name corresponds to a texture
func IsTexture(texture uint32) bool {
	ret := C.glowIsTexture(gpIsTexture, (C.GLuint)(texture))
	return ret == TRUE
}
func IsTextureHandleResidentARB(handle uint64) bool {
	ret := C.glowIsTextureHandleResidentARB(gpIsTextureHandleResidentARB, (C.GLuint64)(handle))
	return ret == TRUE
}

// determine if a name corresponds to a transform feedback object
func IsTransformFeedback(id uint32) bool {
	ret := C.glowIsTransformFeedback(gpIsTransformFeedback, (C.GLuint)(id))
	return ret == TRUE
}

// determine if a name corresponds to a vertex array object
func IsVertexArray(array uint32) bool {
	ret := C.glowIsVertexArray(gpIsVertexArray, (C.GLuint)(array))
	return ret == TRUE
}
func LabelObjectEXT(xtype uint32, object uint32, length int32, label *uint8) {
	C.glowLabelObjectEXT(gpLabelObjectEXT, (C.GLenum)(xtype), (C.GLuint)(object), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func LightModelxOES(pname uint32, param int32) {
	C.glowLightModelxOES(gpLightModelxOES, (C.GLenum)(pname), (C.GLfixed)(param))
}
func LightModelxvOES(pname uint32, param *int32) {
	C.glowLightModelxvOES(gpLightModelxvOES, (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(param)))
}
func LightxOES(light uint32, pname uint32, param int32) {
	C.glowLightxOES(gpLightxOES, (C.GLenum)(light), (C.GLenum)(pname), (C.GLfixed)(param))
}
func LightxvOES(light uint32, pname uint32, params *int32) {
	C.glowLightxvOES(gpLightxvOES, (C.GLenum)(light), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// specify the width of rasterized lines
func LineWidth(width float32) {
	C.glowLineWidth(gpLineWidth, (C.GLfloat)(width))
}
func LineWidthxOES(width int32) {
	C.glowLineWidthxOES(gpLineWidthxOES, (C.GLfixed)(width))
}

// Links a program object
func LinkProgram(program uint32) {
	C.glowLinkProgram(gpLinkProgram, (C.GLuint)(program))
}
func LoadMatrixxOES(m *int32) {
	C.glowLoadMatrixxOES(gpLoadMatrixxOES, (*C.GLfixed)(unsafe.Pointer(m)))
}
func LoadTransposeMatrixxOES(m *int32) {
	C.glowLoadTransposeMatrixxOES(gpLoadTransposeMatrixxOES, (*C.GLfixed)(unsafe.Pointer(m)))
}

// specify a logical pixel operation for rendering
func LogicOp(opcode uint32) {
	C.glowLogicOp(gpLogicOp, (C.GLenum)(opcode))
}
func MakeImageHandleNonResidentARB(handle uint64) {
	C.glowMakeImageHandleNonResidentARB(gpMakeImageHandleNonResidentARB, (C.GLuint64)(handle))
}
func MakeImageHandleResidentARB(handle uint64, access uint32) {
	C.glowMakeImageHandleResidentARB(gpMakeImageHandleResidentARB, (C.GLuint64)(handle), (C.GLenum)(access))
}
func MakeTextureHandleNonResidentARB(handle uint64) {
	C.glowMakeTextureHandleNonResidentARB(gpMakeTextureHandleNonResidentARB, (C.GLuint64)(handle))
}
func MakeTextureHandleResidentARB(handle uint64) {
	C.glowMakeTextureHandleResidentARB(gpMakeTextureHandleResidentARB, (C.GLuint64)(handle))
}
func Map1xOES(target uint32, u1 int32, u2 int32, stride int32, order int32, points int32) {
	C.glowMap1xOES(gpMap1xOES, (C.GLenum)(target), (C.GLfixed)(u1), (C.GLfixed)(u2), (C.GLint)(stride), (C.GLint)(order), (C.GLfixed)(points))
}
func Map2xOES(target uint32, u1 int32, u2 int32, ustride int32, uorder int32, v1 int32, v2 int32, vstride int32, vorder int32, points int32) {
	C.glowMap2xOES(gpMap2xOES, (C.GLenum)(target), (C.GLfixed)(u1), (C.GLfixed)(u2), (C.GLint)(ustride), (C.GLint)(uorder), (C.GLfixed)(v1), (C.GLfixed)(v2), (C.GLint)(vstride), (C.GLint)(vorder), (C.GLfixed)(points))
}

// map all of a buffer object's data store into the client's address space
func MapBuffer(target uint32, access uint32) unsafe.Pointer {
	ret := C.glowMapBuffer(gpMapBuffer, (C.GLenum)(target), (C.GLenum)(access))
	return (unsafe.Pointer)(ret)
}

// map all or part of a buffer object's data store into the client's address space
func MapBufferRange(target uint32, offset int, length int, access uint32) unsafe.Pointer {
	ret := C.glowMapBufferRange(gpMapBufferRange, (C.GLenum)(target), (C.GLintptr)(offset), (C.GLsizeiptr)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func MapGrid1xOES(n int32, u1 int32, u2 int32) {
	C.glowMapGrid1xOES(gpMapGrid1xOES, (C.GLint)(n), (C.GLfixed)(u1), (C.GLfixed)(u2))
}
func MapGrid2xOES(n int32, u1 int32, u2 int32, v1 int32, v2 int32) {
	C.glowMapGrid2xOES(gpMapGrid2xOES, (C.GLint)(n), (C.GLfixed)(u1), (C.GLfixed)(u2), (C.GLfixed)(v1), (C.GLfixed)(v2))
}

// map all of a buffer object's data store into the client's address space
func MapNamedBuffer(buffer uint32, access uint32) unsafe.Pointer {
	ret := C.glowMapNamedBuffer(gpMapNamedBuffer, (C.GLuint)(buffer), (C.GLenum)(access))
	return (unsafe.Pointer)(ret)
}

// map all or part of a buffer object's data store into the client's address space
func MapNamedBufferRange(buffer uint32, offset int, length int32, access uint32) unsafe.Pointer {
	ret := C.glowMapNamedBufferRange(gpMapNamedBufferRange, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(length), (C.GLbitfield)(access))
	return (unsafe.Pointer)(ret)
}
func MaterialxOES(face uint32, pname uint32, param int32) {
	C.glowMaterialxOES(gpMaterialxOES, (C.GLenum)(face), (C.GLenum)(pname), (C.GLfixed)(param))
}
func MaterialxvOES(face uint32, pname uint32, param *int32) {
	C.glowMaterialxvOES(gpMaterialxvOES, (C.GLenum)(face), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(param)))
}

// defines a barrier ordering memory transactions
func MemoryBarrier(barriers uint32) {
	C.glowMemoryBarrier(gpMemoryBarrier, (C.GLbitfield)(barriers))
}
func MemoryBarrierByRegion(barriers uint32) {
	C.glowMemoryBarrierByRegion(gpMemoryBarrierByRegion, (C.GLbitfield)(barriers))
}

// specifies minimum rate at which sample shaing takes place
func MinSampleShading(value float32) {
	C.glowMinSampleShading(gpMinSampleShading, (C.GLfloat)(value))
}
func MinSampleShadingARB(value float32) {
	C.glowMinSampleShadingARB(gpMinSampleShadingARB, (C.GLfloat)(value))
}
func MultMatrixxOES(m *int32) {
	C.glowMultMatrixxOES(gpMultMatrixxOES, (*C.GLfixed)(unsafe.Pointer(m)))
}
func MultTransposeMatrixxOES(m *int32) {
	C.glowMultTransposeMatrixxOES(gpMultTransposeMatrixxOES, (*C.GLfixed)(unsafe.Pointer(m)))
}

// render multiple sets of primitives from array data
func MultiDrawArrays(mode uint32, first *int32, count *int32, drawcount int32) {
	C.glowMultiDrawArrays(gpMultiDrawArrays, (C.GLenum)(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLsizei)(drawcount))
}
func MultiDrawArraysEXT(mode uint32, first *int32, count *int32, primcount int32) {
	C.glowMultiDrawArraysEXT(gpMultiDrawArraysEXT, (C.GLenum)(mode), (*C.GLint)(unsafe.Pointer(first)), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLsizei)(primcount))
}

// render multiple sets of primitives from array data, taking parameters from memory
func MultiDrawArraysIndirect(mode uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.glowMultiDrawArraysIndirect(gpMultiDrawArraysIndirect, (C.GLenum)(mode), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func MultiDrawArraysIndirectCountARB(mode uint32, indirect int, drawcount int, maxdrawcount int32, stride int32) {
	C.glowMultiDrawArraysIndirectCountARB(gpMultiDrawArraysIndirectCountARB, (C.GLenum)(mode), (C.GLintptr)(indirect), (C.GLintptr)(drawcount), (C.GLsizei)(maxdrawcount), (C.GLsizei)(stride))
}

// render multiple sets of primitives by specifying indices of array data elements
func MultiDrawElements(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32) {
	C.glowMultiDrawElements(gpMultiDrawElements, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(xtype), indices, (C.GLsizei)(drawcount))
}

// render multiple sets of primitives by specifying indices of array data elements and an index to apply to each index
func MultiDrawElementsBaseVertex(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, drawcount int32, basevertex *int32) {
	C.glowMultiDrawElementsBaseVertex(gpMultiDrawElementsBaseVertex, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(xtype), indices, (C.GLsizei)(drawcount), (*C.GLint)(unsafe.Pointer(basevertex)))
}
func MultiDrawElementsEXT(mode uint32, count *int32, xtype uint32, indices *unsafe.Pointer, primcount int32) {
	C.glowMultiDrawElementsEXT(gpMultiDrawElementsEXT, (C.GLenum)(mode), (*C.GLsizei)(unsafe.Pointer(count)), (C.GLenum)(xtype), indices, (C.GLsizei)(primcount))
}

// render indexed primitives from array data, taking parameters from memory
func MultiDrawElementsIndirect(mode uint32, xtype uint32, indirect unsafe.Pointer, drawcount int32, stride int32) {
	C.glowMultiDrawElementsIndirect(gpMultiDrawElementsIndirect, (C.GLenum)(mode), (C.GLenum)(xtype), indirect, (C.GLsizei)(drawcount), (C.GLsizei)(stride))
}
func MultiDrawElementsIndirectCountARB(mode uint32, xtype uint32, indirect int, drawcount int, maxdrawcount int32, stride int32) {
	C.glowMultiDrawElementsIndirectCountARB(gpMultiDrawElementsIndirectCountARB, (C.GLenum)(mode), (C.GLenum)(xtype), (C.GLintptr)(indirect), (C.GLintptr)(drawcount), (C.GLsizei)(maxdrawcount), (C.GLsizei)(stride))
}
func MultiTexCoord1bOES(texture uint32, s int8) {
	C.glowMultiTexCoord1bOES(gpMultiTexCoord1bOES, (C.GLenum)(texture), (C.GLbyte)(s))
}
func MultiTexCoord1bvOES(texture uint32, coords *int8) {
	C.glowMultiTexCoord1bvOES(gpMultiTexCoord1bvOES, (C.GLenum)(texture), (*C.GLbyte)(unsafe.Pointer(coords)))
}
func MultiTexCoord1xOES(texture uint32, s int32) {
	C.glowMultiTexCoord1xOES(gpMultiTexCoord1xOES, (C.GLenum)(texture), (C.GLfixed)(s))
}
func MultiTexCoord1xvOES(texture uint32, coords *int32) {
	C.glowMultiTexCoord1xvOES(gpMultiTexCoord1xvOES, (C.GLenum)(texture), (*C.GLfixed)(unsafe.Pointer(coords)))
}
func MultiTexCoord2bOES(texture uint32, s int8, t int8) {
	C.glowMultiTexCoord2bOES(gpMultiTexCoord2bOES, (C.GLenum)(texture), (C.GLbyte)(s), (C.GLbyte)(t))
}
func MultiTexCoord2bvOES(texture uint32, coords *int8) {
	C.glowMultiTexCoord2bvOES(gpMultiTexCoord2bvOES, (C.GLenum)(texture), (*C.GLbyte)(unsafe.Pointer(coords)))
}
func MultiTexCoord2xOES(texture uint32, s int32, t int32) {
	C.glowMultiTexCoord2xOES(gpMultiTexCoord2xOES, (C.GLenum)(texture), (C.GLfixed)(s), (C.GLfixed)(t))
}
func MultiTexCoord2xvOES(texture uint32, coords *int32) {
	C.glowMultiTexCoord2xvOES(gpMultiTexCoord2xvOES, (C.GLenum)(texture), (*C.GLfixed)(unsafe.Pointer(coords)))
}
func MultiTexCoord3bOES(texture uint32, s int8, t int8, r int8) {
	C.glowMultiTexCoord3bOES(gpMultiTexCoord3bOES, (C.GLenum)(texture), (C.GLbyte)(s), (C.GLbyte)(t), (C.GLbyte)(r))
}
func MultiTexCoord3bvOES(texture uint32, coords *int8) {
	C.glowMultiTexCoord3bvOES(gpMultiTexCoord3bvOES, (C.GLenum)(texture), (*C.GLbyte)(unsafe.Pointer(coords)))
}
func MultiTexCoord3xOES(texture uint32, s int32, t int32, r int32) {
	C.glowMultiTexCoord3xOES(gpMultiTexCoord3xOES, (C.GLenum)(texture), (C.GLfixed)(s), (C.GLfixed)(t), (C.GLfixed)(r))
}
func MultiTexCoord3xvOES(texture uint32, coords *int32) {
	C.glowMultiTexCoord3xvOES(gpMultiTexCoord3xvOES, (C.GLenum)(texture), (*C.GLfixed)(unsafe.Pointer(coords)))
}
func MultiTexCoord4bOES(texture uint32, s int8, t int8, r int8, q int8) {
	C.glowMultiTexCoord4bOES(gpMultiTexCoord4bOES, (C.GLenum)(texture), (C.GLbyte)(s), (C.GLbyte)(t), (C.GLbyte)(r), (C.GLbyte)(q))
}
func MultiTexCoord4bvOES(texture uint32, coords *int8) {
	C.glowMultiTexCoord4bvOES(gpMultiTexCoord4bvOES, (C.GLenum)(texture), (*C.GLbyte)(unsafe.Pointer(coords)))
}
func MultiTexCoord4xOES(texture uint32, s int32, t int32, r int32, q int32) {
	C.glowMultiTexCoord4xOES(gpMultiTexCoord4xOES, (C.GLenum)(texture), (C.GLfixed)(s), (C.GLfixed)(t), (C.GLfixed)(r), (C.GLfixed)(q))
}
func MultiTexCoord4xvOES(texture uint32, coords *int32) {
	C.glowMultiTexCoord4xvOES(gpMultiTexCoord4xvOES, (C.GLenum)(texture), (*C.GLfixed)(unsafe.Pointer(coords)))
}

// creates and initializes a buffer object's data     store
func NamedBufferData(buffer uint32, size int32, data unsafe.Pointer, usage uint32) {
	C.glowNamedBufferData(gpNamedBufferData, (C.GLuint)(buffer), (C.GLsizei)(size), data, (C.GLenum)(usage))
}
func NamedBufferPageCommitmentARB(buffer uint32, offset int, size int32, commit bool) {
	C.glowNamedBufferPageCommitmentARB(gpNamedBufferPageCommitmentARB, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}
func NamedBufferPageCommitmentEXT(buffer uint32, offset int, size int32, commit bool) {
	C.glowNamedBufferPageCommitmentEXT(gpNamedBufferPageCommitmentEXT, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), (C.GLboolean)(boolToInt(commit)))
}

// creates and initializes a buffer object's immutable data     store
func NamedBufferStorage(buffer uint32, size int32, data unsafe.Pointer, flags uint32) {
	C.glowNamedBufferStorage(gpNamedBufferStorage, (C.GLuint)(buffer), (C.GLsizei)(size), data, (C.GLbitfield)(flags))
}

// updates a subset of a buffer object's data store
func NamedBufferSubData(buffer uint32, offset int, size int32, data unsafe.Pointer) {
	C.glowNamedBufferSubData(gpNamedBufferSubData, (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size), data)
}

// specify which color buffers are to be drawn into
func NamedFramebufferDrawBuffer(framebuffer uint32, buf uint32) {
	C.glowNamedFramebufferDrawBuffer(gpNamedFramebufferDrawBuffer, (C.GLuint)(framebuffer), (C.GLenum)(buf))
}

// Specifies a list of color buffers to be drawn     into
func NamedFramebufferDrawBuffers(framebuffer uint32, n int32, bufs *uint32) {
	C.glowNamedFramebufferDrawBuffers(gpNamedFramebufferDrawBuffers, (C.GLuint)(framebuffer), (C.GLsizei)(n), (*C.GLenum)(unsafe.Pointer(bufs)))
}

// set a named parameter of a framebuffer object
func NamedFramebufferParameteri(framebuffer uint32, pname uint32, param int32) {
	C.glowNamedFramebufferParameteri(gpNamedFramebufferParameteri, (C.GLuint)(framebuffer), (C.GLenum)(pname), (C.GLint)(param))
}

// select a color buffer source for pixels
func NamedFramebufferReadBuffer(framebuffer uint32, src uint32) {
	C.glowNamedFramebufferReadBuffer(gpNamedFramebufferReadBuffer, (C.GLuint)(framebuffer), (C.GLenum)(src))
}

// attach a renderbuffer as a logical buffer of a framebuffer object
func NamedFramebufferRenderbuffer(framebuffer uint32, attachment uint32, renderbuffertarget uint32, renderbuffer uint32) {
	C.glowNamedFramebufferRenderbuffer(gpNamedFramebufferRenderbuffer, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLenum)(renderbuffertarget), (C.GLuint)(renderbuffer))
}
func NamedFramebufferTexture(framebuffer uint32, attachment uint32, texture uint32, level int32) {
	C.glowNamedFramebufferTexture(gpNamedFramebufferTexture, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level))
}

// attach a single layer of a texture object as a logical buffer of a framebuffer object
func NamedFramebufferTextureLayer(framebuffer uint32, attachment uint32, texture uint32, level int32, layer int32) {
	C.glowNamedFramebufferTextureLayer(gpNamedFramebufferTextureLayer, (C.GLuint)(framebuffer), (C.GLenum)(attachment), (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(layer))
}

// establish data storage, format and dimensions of a     renderbuffer object's image
func NamedRenderbufferStorage(renderbuffer uint32, internalformat uint32, width int32, height int32) {
	C.glowNamedRenderbufferStorage(gpNamedRenderbufferStorage, (C.GLuint)(renderbuffer), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// establish data storage, format, dimensions and sample count of     a renderbuffer object's image
func NamedRenderbufferStorageMultisample(renderbuffer uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowNamedRenderbufferStorageMultisample(gpNamedRenderbufferStorageMultisample, (C.GLuint)(renderbuffer), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}
func NamedStringARB(xtype uint32, namelen int32, name *uint8, stringlen int32, xstring *uint8) {
	C.glowNamedStringARB(gpNamedStringARB, (C.GLenum)(xtype), (C.GLint)(namelen), (*C.GLchar)(unsafe.Pointer(name)), (C.GLint)(stringlen), (*C.GLchar)(unsafe.Pointer(xstring)))
}
func Normal3xOES(nx int32, ny int32, nz int32) {
	C.glowNormal3xOES(gpNormal3xOES, (C.GLfixed)(nx), (C.GLfixed)(ny), (C.GLfixed)(nz))
}
func Normal3xvOES(coords *int32) {
	C.glowNormal3xvOES(gpNormal3xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}

// label a named object identified within a namespace
func ObjectLabel(identifier uint32, name uint32, length int32, label *uint8) {
	C.glowObjectLabel(gpObjectLabel, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func ObjectLabelKHR(identifier uint32, name uint32, length int32, label *uint8) {
	C.glowObjectLabelKHR(gpObjectLabelKHR, (C.GLenum)(identifier), (C.GLuint)(name), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}

// label a a sync object identified by a pointer
func ObjectPtrLabel(ptr unsafe.Pointer, length int32, label *uint8) {
	C.glowObjectPtrLabel(gpObjectPtrLabel, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func ObjectPtrLabelKHR(ptr unsafe.Pointer, length int32, label *uint8) {
	C.glowObjectPtrLabelKHR(gpObjectPtrLabelKHR, ptr, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(label)))
}
func OrthofOES(l float32, r float32, b float32, t float32, n float32, f float32) {
	C.glowOrthofOES(gpOrthofOES, (C.GLfloat)(l), (C.GLfloat)(r), (C.GLfloat)(b), (C.GLfloat)(t), (C.GLfloat)(n), (C.GLfloat)(f))
}
func OrthoxOES(l int32, r int32, b int32, t int32, n int32, f int32) {
	C.glowOrthoxOES(gpOrthoxOES, (C.GLfixed)(l), (C.GLfixed)(r), (C.GLfixed)(b), (C.GLfixed)(t), (C.GLfixed)(n), (C.GLfixed)(f))
}
func PassThroughxOES(token int32) {
	C.glowPassThroughxOES(gpPassThroughxOES, (C.GLfixed)(token))
}
func PatchParameterfv(pname uint32, values *float32) {
	C.glowPatchParameterfv(gpPatchParameterfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(values)))
}
func PatchParameteri(pname uint32, value int32) {
	C.glowPatchParameteri(gpPatchParameteri, (C.GLenum)(pname), (C.GLint)(value))
}

// pause transform feedback operations
func PauseTransformFeedback() {
	C.glowPauseTransformFeedback(gpPauseTransformFeedback)
}
func PixelMapx(xmap uint32, size int32, values *int32) {
	C.glowPixelMapx(gpPixelMapx, (C.GLenum)(xmap), (C.GLint)(size), (*C.GLfixed)(unsafe.Pointer(values)))
}
func PixelStoref(pname uint32, param float32) {
	C.glowPixelStoref(gpPixelStoref, (C.GLenum)(pname), (C.GLfloat)(param))
}
func PixelStorei(pname uint32, param int32) {
	C.glowPixelStorei(gpPixelStorei, (C.GLenum)(pname), (C.GLint)(param))
}
func PixelStorex(pname uint32, param int32) {
	C.glowPixelStorex(gpPixelStorex, (C.GLenum)(pname), (C.GLfixed)(param))
}
func PixelTransferxOES(pname uint32, param int32) {
	C.glowPixelTransferxOES(gpPixelTransferxOES, (C.GLenum)(pname), (C.GLfixed)(param))
}
func PixelZoomxOES(xfactor int32, yfactor int32) {
	C.glowPixelZoomxOES(gpPixelZoomxOES, (C.GLfixed)(xfactor), (C.GLfixed)(yfactor))
}
func PointParameterf(pname uint32, param float32) {
	C.glowPointParameterf(gpPointParameterf, (C.GLenum)(pname), (C.GLfloat)(param))
}
func PointParameterfv(pname uint32, params *float32) {
	C.glowPointParameterfv(gpPointParameterfv, (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func PointParameteri(pname uint32, param int32) {
	C.glowPointParameteri(gpPointParameteri, (C.GLenum)(pname), (C.GLint)(param))
}
func PointParameteriv(pname uint32, params *int32) {
	C.glowPointParameteriv(gpPointParameteriv, (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func PointParameterxOES(pname uint32, param int32) {
	C.glowPointParameterxOES(gpPointParameterxOES, (C.GLenum)(pname), (C.GLfixed)(param))
}
func PointParameterxvOES(pname uint32, params *int32) {
	C.glowPointParameterxvOES(gpPointParameterxvOES, (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// specify the diameter of rasterized points
func PointSize(size float32) {
	C.glowPointSize(gpPointSize, (C.GLfloat)(size))
}
func PointSizexOES(size int32) {
	C.glowPointSizexOES(gpPointSizexOES, (C.GLfixed)(size))
}

// select a polygon rasterization mode
func PolygonMode(face uint32, mode uint32) {
	C.glowPolygonMode(gpPolygonMode, (C.GLenum)(face), (C.GLenum)(mode))
}

// set the scale and units used to calculate depth values
func PolygonOffset(factor float32, units float32) {
	C.glowPolygonOffset(gpPolygonOffset, (C.GLfloat)(factor), (C.GLfloat)(units))
}
func PolygonOffsetxOES(factor int32, units int32) {
	C.glowPolygonOffsetxOES(gpPolygonOffsetxOES, (C.GLfixed)(factor), (C.GLfixed)(units))
}

// pop the active debug group
func PopDebugGroup() {
	C.glowPopDebugGroup(gpPopDebugGroup)
}
func PopDebugGroupKHR() {
	C.glowPopDebugGroupKHR(gpPopDebugGroupKHR)
}
func PopGroupMarkerEXT() {
	C.glowPopGroupMarkerEXT(gpPopGroupMarkerEXT)
}

// specify the primitive restart index
func PrimitiveRestartIndex(index uint32) {
	C.glowPrimitiveRestartIndex(gpPrimitiveRestartIndex, (C.GLuint)(index))
}
func PrioritizeTexturesxOES(n int32, textures *uint32, priorities *int32) {
	C.glowPrioritizeTexturesxOES(gpPrioritizeTexturesxOES, (C.GLsizei)(n), (*C.GLuint)(unsafe.Pointer(textures)), (*C.GLfixed)(unsafe.Pointer(priorities)))
}

// load a program object with a program binary
func ProgramBinary(program uint32, binaryFormat uint32, binary unsafe.Pointer, length int32) {
	C.glowProgramBinary(gpProgramBinary, (C.GLuint)(program), (C.GLenum)(binaryFormat), binary, (C.GLsizei)(length))
}
func ProgramParameteri(program uint32, pname uint32, value int32) {
	C.glowProgramParameteri(gpProgramParameteri, (C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}
func ProgramParameteriEXT(program uint32, pname uint32, value int32) {
	C.glowProgramParameteriEXT(gpProgramParameteriEXT, (C.GLuint)(program), (C.GLenum)(pname), (C.GLint)(value))
}
func ProgramUniform1d(program uint32, location int32, v0 float64) {
	C.glowProgramUniform1d(gpProgramUniform1d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0))
}
func ProgramUniform1dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform1dv(gpProgramUniform1dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1f(program uint32, location int32, v0 float32) {
	C.glowProgramUniform1f(gpProgramUniform1f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}
func ProgramUniform1fEXT(program uint32, location int32, v0 float32) {
	C.glowProgramUniform1fEXT(gpProgramUniform1fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform1fv(gpProgramUniform1fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform1fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform1fvEXT(gpProgramUniform1fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1i(program uint32, location int32, v0 int32) {
	C.glowProgramUniform1i(gpProgramUniform1i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}
func ProgramUniform1iEXT(program uint32, location int32, v0 int32) {
	C.glowProgramUniform1iEXT(gpProgramUniform1iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform1iv(gpProgramUniform1iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform1ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform1ivEXT(gpProgramUniform1ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1ui(program uint32, location int32, v0 uint32) {
	C.glowProgramUniform1ui(gpProgramUniform1ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}
func ProgramUniform1uiEXT(program uint32, location int32, v0 uint32) {
	C.glowProgramUniform1uiEXT(gpProgramUniform1uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform1uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform1uiv(gpProgramUniform1uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform1uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform1uivEXT(gpProgramUniform1uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform2d(program uint32, location int32, v0 float64, v1 float64) {
	C.glowProgramUniform2d(gpProgramUniform2d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1))
}
func ProgramUniform2dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform2dv(gpProgramUniform2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2f(program uint32, location int32, v0 float32, v1 float32) {
	C.glowProgramUniform2f(gpProgramUniform2f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}
func ProgramUniform2fEXT(program uint32, location int32, v0 float32, v1 float32) {
	C.glowProgramUniform2fEXT(gpProgramUniform2fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform2fv(gpProgramUniform2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform2fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform2fvEXT(gpProgramUniform2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2i(program uint32, location int32, v0 int32, v1 int32) {
	C.glowProgramUniform2i(gpProgramUniform2i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}
func ProgramUniform2iEXT(program uint32, location int32, v0 int32, v1 int32) {
	C.glowProgramUniform2iEXT(gpProgramUniform2iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform2iv(gpProgramUniform2iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform2ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform2ivEXT(gpProgramUniform2ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2ui(program uint32, location int32, v0 uint32, v1 uint32) {
	C.glowProgramUniform2ui(gpProgramUniform2ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}
func ProgramUniform2uiEXT(program uint32, location int32, v0 uint32, v1 uint32) {
	C.glowProgramUniform2uiEXT(gpProgramUniform2uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform2uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform2uiv(gpProgramUniform2uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform2uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform2uivEXT(gpProgramUniform2uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform3d(program uint32, location int32, v0 float64, v1 float64, v2 float64) {
	C.glowProgramUniform3d(gpProgramUniform3d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2))
}
func ProgramUniform3dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform3dv(gpProgramUniform3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3f(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	C.glowProgramUniform3f(gpProgramUniform3f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}
func ProgramUniform3fEXT(program uint32, location int32, v0 float32, v1 float32, v2 float32) {
	C.glowProgramUniform3fEXT(gpProgramUniform3fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform3fv(gpProgramUniform3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform3fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform3fvEXT(gpProgramUniform3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3i(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	C.glowProgramUniform3i(gpProgramUniform3i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}
func ProgramUniform3iEXT(program uint32, location int32, v0 int32, v1 int32, v2 int32) {
	C.glowProgramUniform3iEXT(gpProgramUniform3iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform3iv(gpProgramUniform3iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform3ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform3ivEXT(gpProgramUniform3ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowProgramUniform3ui(gpProgramUniform3ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}
func ProgramUniform3uiEXT(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowProgramUniform3uiEXT(gpProgramUniform3uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform3uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform3uiv(gpProgramUniform3uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform3uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform3uivEXT(gpProgramUniform3uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform4d(program uint32, location int32, v0 float64, v1 float64, v2 float64, v3 float64) {
	C.glowProgramUniform4d(gpProgramUniform4d, (C.GLuint)(program), (C.GLint)(location), (C.GLdouble)(v0), (C.GLdouble)(v1), (C.GLdouble)(v2), (C.GLdouble)(v3))
}
func ProgramUniform4dv(program uint32, location int32, count int32, value *float64) {
	C.glowProgramUniform4dv(gpProgramUniform4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4f(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowProgramUniform4f(gpProgramUniform4f, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}
func ProgramUniform4fEXT(program uint32, location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowProgramUniform4fEXT(gpProgramUniform4fEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4fv(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform4fv(gpProgramUniform4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniform4fvEXT(program uint32, location int32, count int32, value *float32) {
	C.glowProgramUniform4fvEXT(gpProgramUniform4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4i(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowProgramUniform4i(gpProgramUniform4i, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}
func ProgramUniform4iEXT(program uint32, location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowProgramUniform4iEXT(gpProgramUniform4iEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4iv(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform4iv(gpProgramUniform4iv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}
func ProgramUniform4ivEXT(program uint32, location int32, count int32, value *int32) {
	C.glowProgramUniform4ivEXT(gpProgramUniform4ivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4ui(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowProgramUniform4ui(gpProgramUniform4ui, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}
func ProgramUniform4uiEXT(program uint32, location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowProgramUniform4uiEXT(gpProgramUniform4uiEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniform4uiv(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform4uiv(gpProgramUniform4uiv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniform4uivEXT(program uint32, location int32, count int32, value *uint32) {
	C.glowProgramUniform4uivEXT(gpProgramUniform4uivEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func ProgramUniformHandleui64ARB(program uint32, location int32, value uint64) {
	C.glowProgramUniformHandleui64ARB(gpProgramUniformHandleui64ARB, (C.GLuint)(program), (C.GLint)(location), (C.GLuint64)(value))
}
func ProgramUniformHandleui64vARB(program uint32, location int32, count int32, values *uint64) {
	C.glowProgramUniformHandleui64vARB(gpProgramUniformHandleui64vARB, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(values)))
}
func ProgramUniformMatrix2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix2dv(gpProgramUniformMatrix2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2fv(gpProgramUniformMatrix2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2fvEXT(gpProgramUniformMatrix2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix2x3dv(gpProgramUniformMatrix2x3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x3fv(gpProgramUniformMatrix2x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x3fvEXT(gpProgramUniformMatrix2x3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix2x4dv(gpProgramUniformMatrix2x4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix2x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x4fv(gpProgramUniformMatrix2x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix2x4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix2x4fvEXT(gpProgramUniformMatrix2x4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix3dv(gpProgramUniformMatrix3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3fv(gpProgramUniformMatrix3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3fvEXT(gpProgramUniformMatrix3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix3x2dv(gpProgramUniformMatrix3x2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x2fv(gpProgramUniformMatrix3x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x2fvEXT(gpProgramUniformMatrix3x2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix3x4dv(gpProgramUniformMatrix3x4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix3x4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x4fv(gpProgramUniformMatrix3x4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix3x4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix3x4fvEXT(gpProgramUniformMatrix3x4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix4dv(gpProgramUniformMatrix4dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4fv(gpProgramUniformMatrix4fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4fvEXT(gpProgramUniformMatrix4fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x2dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix4x2dv(gpProgramUniformMatrix4x2dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x2fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x2fv(gpProgramUniformMatrix4x2fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x2fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x2fvEXT(gpProgramUniformMatrix4x2fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x3dv(program uint32, location int32, count int32, transpose bool, value *float64) {
	C.glowProgramUniformMatrix4x3dv(gpProgramUniformMatrix4x3dv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for a specified program object
func ProgramUniformMatrix4x3fv(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x3fv(gpProgramUniformMatrix4x3fv, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func ProgramUniformMatrix4x3fvEXT(program uint32, location int32, count int32, transpose bool, value *float32) {
	C.glowProgramUniformMatrix4x3fvEXT(gpProgramUniformMatrix4x3fvEXT, (C.GLuint)(program), (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}

// specifiy the vertex to be used as the source of data for flat shaded varyings
func ProvokingVertex(mode uint32) {
	C.glowProvokingVertex(gpProvokingVertex, (C.GLenum)(mode))
}

// push a named debug group into the command stream
func PushDebugGroup(source uint32, id uint32, length int32, message *uint8) {
	C.glowPushDebugGroup(gpPushDebugGroup, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(message)))
}
func PushDebugGroupKHR(source uint32, id uint32, length int32, message *uint8) {
	C.glowPushDebugGroupKHR(gpPushDebugGroupKHR, (C.GLenum)(source), (C.GLuint)(id), (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(message)))
}
func PushGroupMarkerEXT(length int32, marker *uint8) {
	C.glowPushGroupMarkerEXT(gpPushGroupMarkerEXT, (C.GLsizei)(length), (*C.GLchar)(unsafe.Pointer(marker)))
}

// record the GL time into a query object after all previous commands have reached the GL server but have not yet necessarily executed.
func QueryCounter(id uint32, target uint32) {
	C.glowQueryCounter(gpQueryCounter, (C.GLuint)(id), (C.GLenum)(target))
}
func QueryMatrixxOES(mantissa *int32, exponent *int32) uint32 {
	ret := C.glowQueryMatrixxOES(gpQueryMatrixxOES, (*C.GLfixed)(unsafe.Pointer(mantissa)), (*C.GLint)(unsafe.Pointer(exponent)))
	return (uint32)(ret)
}
func RasterPos2xOES(x int32, y int32) {
	C.glowRasterPos2xOES(gpRasterPos2xOES, (C.GLfixed)(x), (C.GLfixed)(y))
}
func RasterPos2xvOES(coords *int32) {
	C.glowRasterPos2xvOES(gpRasterPos2xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func RasterPos3xOES(x int32, y int32, z int32) {
	C.glowRasterPos3xOES(gpRasterPos3xOES, (C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
}
func RasterPos3xvOES(coords *int32) {
	C.glowRasterPos3xvOES(gpRasterPos3xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func RasterPos4xOES(x int32, y int32, z int32, w int32) {
	C.glowRasterPos4xOES(gpRasterPos4xOES, (C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z), (C.GLfixed)(w))
}
func RasterPos4xvOES(coords *int32) {
	C.glowRasterPos4xvOES(gpRasterPos4xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}

// select a color buffer source for pixels
func ReadBuffer(src uint32) {
	C.glowReadBuffer(gpReadBuffer, (C.GLenum)(src))
}

// read a block of pixels from the frame buffer
func ReadPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowReadPixels(gpReadPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// read a block of pixels from the frame buffer
func ReadnPixels(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixels(gpReadnPixels, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func ReadnPixelsARB(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixelsARB(gpReadnPixelsARB, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func ReadnPixelsKHR(x int32, y int32, width int32, height int32, format uint32, xtype uint32, bufSize int32, data unsafe.Pointer) {
	C.glowReadnPixelsKHR(gpReadnPixelsKHR, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), (C.GLsizei)(bufSize), data)
}
func RectxOES(x1 int32, y1 int32, x2 int32, y2 int32) {
	C.glowRectxOES(gpRectxOES, (C.GLfixed)(x1), (C.GLfixed)(y1), (C.GLfixed)(x2), (C.GLfixed)(y2))
}
func RectxvOES(v1 *int32, v2 *int32) {
	C.glowRectxvOES(gpRectxvOES, (*C.GLfixed)(unsafe.Pointer(v1)), (*C.GLfixed)(unsafe.Pointer(v2)))
}

// release resources consumed by the implementation's shader compiler
func ReleaseShaderCompiler() {
	C.glowReleaseShaderCompiler(gpReleaseShaderCompiler)
}

// establish data storage, format and dimensions of a     renderbuffer object's image
func RenderbufferStorage(target uint32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorage(gpRenderbufferStorage, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// establish data storage, format, dimensions and sample count of     a renderbuffer object's image
func RenderbufferStorageMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32) {
	C.glowRenderbufferStorageMultisample(gpRenderbufferStorageMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// resume transform feedback operations
func ResumeTransformFeedback() {
	C.glowResumeTransformFeedback(gpResumeTransformFeedback)
}
func RotatexOES(angle int32, x int32, y int32, z int32) {
	C.glowRotatexOES(gpRotatexOES, (C.GLfixed)(angle), (C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
}

// specify multisample coverage parameters
func SampleCoverage(value float32, invert bool) {
	C.glowSampleCoverage(gpSampleCoverage, (C.GLfloat)(value), (C.GLboolean)(boolToInt(invert)))
}
func SampleCoverageOES(value int32, invert bool) {
	C.glowSampleCoverageOES(gpSampleCoverageOES, (C.GLfixed)(value), (C.GLboolean)(boolToInt(invert)))
}
func SampleCoveragexOES(value int32, invert bool) {
	C.glowSampleCoveragexOES(gpSampleCoveragexOES, (C.GLclampx)(value), (C.GLboolean)(boolToInt(invert)))
}

// set the value of a sub-word of the sample mask
func SampleMaski(maskNumber uint32, mask uint32) {
	C.glowSampleMaski(gpSampleMaski, (C.GLuint)(maskNumber), (C.GLbitfield)(mask))
}
func SamplerParameterIiv(sampler uint32, pname uint32, param *int32) {
	C.glowSamplerParameterIiv(gpSamplerParameterIiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func SamplerParameterIuiv(sampler uint32, pname uint32, param *uint32) {
	C.glowSamplerParameterIuiv(gpSamplerParameterIuiv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(param)))
}
func SamplerParameterf(sampler uint32, pname uint32, param float32) {
	C.glowSamplerParameterf(gpSamplerParameterf, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLfloat)(param))
}
func SamplerParameterfv(sampler uint32, pname uint32, param *float32) {
	C.glowSamplerParameterfv(gpSamplerParameterfv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(param)))
}
func SamplerParameteri(sampler uint32, pname uint32, param int32) {
	C.glowSamplerParameteri(gpSamplerParameteri, (C.GLuint)(sampler), (C.GLenum)(pname), (C.GLint)(param))
}
func SamplerParameteriv(sampler uint32, pname uint32, param *int32) {
	C.glowSamplerParameteriv(gpSamplerParameteriv, (C.GLuint)(sampler), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}
func ScalexOES(x int32, y int32, z int32) {
	C.glowScalexOES(gpScalexOES, (C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
}

// define the scissor box
func Scissor(x int32, y int32, width int32, height int32) {
	C.glowScissor(gpScissor, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorArrayv(first uint32, count int32, v *int32) {
	C.glowScissorArrayv(gpScissorArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(v)))
}

// define the scissor box for a specific viewport
func ScissorIndexed(index uint32, left int32, bottom int32, width int32, height int32) {
	C.glowScissorIndexed(gpScissorIndexed, (C.GLuint)(index), (C.GLint)(left), (C.GLint)(bottom), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ScissorIndexedv(index uint32, v *int32) {
	C.glowScissorIndexedv(gpScissorIndexedv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func SelectPerfMonitorCountersAMD(monitor uint32, enable bool, group uint32, numCounters int32, counterList *uint32) {
	C.glowSelectPerfMonitorCountersAMD(gpSelectPerfMonitorCountersAMD, (C.GLuint)(monitor), (C.GLboolean)(boolToInt(enable)), (C.GLuint)(group), (C.GLint)(numCounters), (*C.GLuint)(unsafe.Pointer(counterList)))
}
func SetFenceNV(fence uint32, condition uint32) {
	C.glowSetFenceNV(gpSetFenceNV, (C.GLuint)(fence), (C.GLenum)(condition))
}

// load pre-compiled shader binaries
func ShaderBinary(count int32, shaders *uint32, binaryformat uint32, binary unsafe.Pointer, length int32) {
	C.glowShaderBinary(gpShaderBinary, (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(shaders)), (C.GLenum)(binaryformat), binary, (C.GLsizei)(length))
}

// Replaces the source code in a shader object
func ShaderSource(shader uint32, count int32, xstring **uint8, length *int32) {
	C.glowShaderSource(gpShaderSource, (C.GLuint)(shader), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(xstring)), (*C.GLint)(unsafe.Pointer(length)))
}

// change an active shader storage block binding
func ShaderStorageBlockBinding(program uint32, storageBlockIndex uint32, storageBlockBinding uint32) {
	C.glowShaderStorageBlockBinding(gpShaderStorageBlockBinding, (C.GLuint)(program), (C.GLuint)(storageBlockIndex), (C.GLuint)(storageBlockBinding))
}

// set front and back function and reference value for stencil testing
func StencilFunc(xfunc uint32, ref int32, mask uint32) {
	C.glowStencilFunc(gpStencilFunc, (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}

// set front and/or back function and reference value for stencil testing
func StencilFuncSeparate(face uint32, xfunc uint32, ref int32, mask uint32) {
	C.glowStencilFuncSeparate(gpStencilFuncSeparate, (C.GLenum)(face), (C.GLenum)(xfunc), (C.GLint)(ref), (C.GLuint)(mask))
}

// control the front and back writing of individual bits in the stencil planes
func StencilMask(mask uint32) {
	C.glowStencilMask(gpStencilMask, (C.GLuint)(mask))
}

// control the front and/or back writing of individual bits in the stencil planes
func StencilMaskSeparate(face uint32, mask uint32) {
	C.glowStencilMaskSeparate(gpStencilMaskSeparate, (C.GLenum)(face), (C.GLuint)(mask))
}

// set front and back stencil test actions
func StencilOp(fail uint32, zfail uint32, zpass uint32) {
	C.glowStencilOp(gpStencilOp, (C.GLenum)(fail), (C.GLenum)(zfail), (C.GLenum)(zpass))
}

// set front and/or back stencil test actions
func StencilOpSeparate(face uint32, sfail uint32, dpfail uint32, dppass uint32) {
	C.glowStencilOpSeparate(gpStencilOpSeparate, (C.GLenum)(face), (C.GLenum)(sfail), (C.GLenum)(dpfail), (C.GLenum)(dppass))
}
func TestFenceNV(fence uint32) bool {
	ret := C.glowTestFenceNV(gpTestFenceNV, (C.GLuint)(fence))
	return ret == TRUE
}

// attach a buffer object's data store to a buffer texture object
func TexBuffer(target uint32, internalformat uint32, buffer uint32) {
	C.glowTexBuffer(gpTexBuffer, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}

// attach a range of a buffer object's data store to a buffer texture object
func TexBufferRange(target uint32, internalformat uint32, buffer uint32, offset int, size int) {
	C.glowTexBufferRange(gpTexBufferRange, (C.GLenum)(target), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizeiptr)(size))
}
func TexCoord1bOES(s int8) {
	C.glowTexCoord1bOES(gpTexCoord1bOES, (C.GLbyte)(s))
}
func TexCoord1bvOES(coords *int8) {
	C.glowTexCoord1bvOES(gpTexCoord1bvOES, (*C.GLbyte)(unsafe.Pointer(coords)))
}
func TexCoord1xOES(s int32) {
	C.glowTexCoord1xOES(gpTexCoord1xOES, (C.GLfixed)(s))
}
func TexCoord1xvOES(coords *int32) {
	C.glowTexCoord1xvOES(gpTexCoord1xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func TexCoord2bOES(s int8, t int8) {
	C.glowTexCoord2bOES(gpTexCoord2bOES, (C.GLbyte)(s), (C.GLbyte)(t))
}
func TexCoord2bvOES(coords *int8) {
	C.glowTexCoord2bvOES(gpTexCoord2bvOES, (*C.GLbyte)(unsafe.Pointer(coords)))
}
func TexCoord2xOES(s int32, t int32) {
	C.glowTexCoord2xOES(gpTexCoord2xOES, (C.GLfixed)(s), (C.GLfixed)(t))
}
func TexCoord2xvOES(coords *int32) {
	C.glowTexCoord2xvOES(gpTexCoord2xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func TexCoord3bOES(s int8, t int8, r int8) {
	C.glowTexCoord3bOES(gpTexCoord3bOES, (C.GLbyte)(s), (C.GLbyte)(t), (C.GLbyte)(r))
}
func TexCoord3bvOES(coords *int8) {
	C.glowTexCoord3bvOES(gpTexCoord3bvOES, (*C.GLbyte)(unsafe.Pointer(coords)))
}
func TexCoord3xOES(s int32, t int32, r int32) {
	C.glowTexCoord3xOES(gpTexCoord3xOES, (C.GLfixed)(s), (C.GLfixed)(t), (C.GLfixed)(r))
}
func TexCoord3xvOES(coords *int32) {
	C.glowTexCoord3xvOES(gpTexCoord3xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func TexCoord4bOES(s int8, t int8, r int8, q int8) {
	C.glowTexCoord4bOES(gpTexCoord4bOES, (C.GLbyte)(s), (C.GLbyte)(t), (C.GLbyte)(r), (C.GLbyte)(q))
}
func TexCoord4bvOES(coords *int8) {
	C.glowTexCoord4bvOES(gpTexCoord4bvOES, (*C.GLbyte)(unsafe.Pointer(coords)))
}
func TexCoord4xOES(s int32, t int32, r int32, q int32) {
	C.glowTexCoord4xOES(gpTexCoord4xOES, (C.GLfixed)(s), (C.GLfixed)(t), (C.GLfixed)(r), (C.GLfixed)(q))
}
func TexCoord4xvOES(coords *int32) {
	C.glowTexCoord4xvOES(gpTexCoord4xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func TexEnvxOES(target uint32, pname uint32, param int32) {
	C.glowTexEnvxOES(gpTexEnvxOES, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfixed)(param))
}
func TexEnvxvOES(target uint32, pname uint32, params *int32) {
	C.glowTexEnvxvOES(gpTexEnvxvOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}
func TexGenxOES(coord uint32, pname uint32, param int32) {
	C.glowTexGenxOES(gpTexGenxOES, (C.GLenum)(coord), (C.GLenum)(pname), (C.GLfixed)(param))
}
func TexGenxvOES(coord uint32, pname uint32, params *int32) {
	C.glowTexGenxvOES(gpTexGenxvOES, (C.GLenum)(coord), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// specify a one-dimensional texture image
func TexImage1D(target uint32, level int32, internalformat int32, width int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage1D(gpTexImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a two-dimensional texture image
func TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage2D(gpTexImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// establish the data storage, format, dimensions, and number of samples of a multisample texture's image
func TexImage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	C.glowTexImage2DMultisample(gpTexImage2DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// specify a three-dimensional texture image
func TexImage3D(target uint32, level int32, internalformat int32, width int32, height int32, depth int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexImage3D(gpTexImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLint)(border), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// establish the data storage, format, dimensions, and number of samples of a multisample texture's image
func TexImage3DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.glowTexImage3DMultisample(gpTexImage3DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}
func TexPageCommitmentARB(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, resident bool) {
	C.glowTexPageCommitmentARB(gpTexPageCommitmentARB, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(resident)))
}
func TexParameterIiv(target uint32, pname uint32, params *int32) {
	C.glowTexParameterIiv(gpTexParameterIiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexParameterIuiv(target uint32, pname uint32, params *uint32) {
	C.glowTexParameterIuiv(gpTexParameterIuiv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TexParameterf(target uint32, pname uint32, param float32) {
	C.glowTexParameterf(gpTexParameterf, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TexParameterfv(target uint32, pname uint32, params *float32) {
	C.glowTexParameterfv(gpTexParameterfv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(params)))
}
func TexParameteri(target uint32, pname uint32, param int32) {
	C.glowTexParameteri(gpTexParameteri, (C.GLenum)(target), (C.GLenum)(pname), (C.GLint)(param))
}
func TexParameteriv(target uint32, pname uint32, params *int32) {
	C.glowTexParameteriv(gpTexParameteriv, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TexParameterxOES(target uint32, pname uint32, param int32) {
	C.glowTexParameterxOES(gpTexParameterxOES, (C.GLenum)(target), (C.GLenum)(pname), (C.GLfixed)(param))
}
func TexParameterxvOES(target uint32, pname uint32, params *int32) {
	C.glowTexParameterxvOES(gpTexParameterxvOES, (C.GLenum)(target), (C.GLenum)(pname), (*C.GLfixed)(unsafe.Pointer(params)))
}

// simultaneously specify storage for all levels of a one-dimensional texture
func TexStorage1D(target uint32, levels int32, internalformat uint32, width int32) {
	C.glowTexStorage1D(gpTexStorage1D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}

// simultaneously specify storage for all levels of a two-dimensional or one-dimensional array texture
func TexStorage2D(target uint32, levels int32, internalformat uint32, width int32, height int32) {
	C.glowTexStorage2D(gpTexStorage2D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// specify storage for a two-dimensional multisample texture
func TexStorage2DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	C.glowTexStorage2DMultisample(gpTexStorage2DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// simultaneously specify storage for all levels of a three-dimensional, two-dimensional array or cube-map array texture
func TexStorage3D(target uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	C.glowTexStorage3D(gpTexStorage3D, (C.GLenum)(target), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

// specify storage for a two-dimensional multisample array texture
func TexStorage3DMultisample(target uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.glowTexStorage3DMultisample(gpTexStorage3DMultisample, (C.GLenum)(target), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// specify a one-dimensional texture subimage
func TexSubImage1D(target uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage1D(gpTexSubImage1D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a two-dimensional texture subimage
func TexSubImage2D(target uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage2D(gpTexSubImage2D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a three-dimensional texture subimage
func TexSubImage3D(target uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTexSubImage3D(gpTexSubImage3D, (C.GLenum)(target), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// controls the ordering of reads and writes to rendered fragments across drawing commands
func TextureBarrier() {
	C.glowTextureBarrier(gpTextureBarrier)
}

// attach a buffer object's data store to a buffer texture object
func TextureBuffer(texture uint32, internalformat uint32, buffer uint32) {
	C.glowTextureBuffer(gpTextureBuffer, (C.GLuint)(texture), (C.GLenum)(internalformat), (C.GLuint)(buffer))
}

// attach a range of a buffer object's data store to a buffer texture object
func TextureBufferRange(texture uint32, internalformat uint32, buffer uint32, offset int, size int32) {
	C.glowTextureBufferRange(gpTextureBufferRange, (C.GLuint)(texture), (C.GLenum)(internalformat), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size))
}
func TextureParameterIiv(texture uint32, pname uint32, params *int32) {
	C.glowTextureParameterIiv(gpTextureParameterIiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(params)))
}
func TextureParameterIuiv(texture uint32, pname uint32, params *uint32) {
	C.glowTextureParameterIuiv(gpTextureParameterIuiv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLuint)(unsafe.Pointer(params)))
}
func TextureParameterf(texture uint32, pname uint32, param float32) {
	C.glowTextureParameterf(gpTextureParameterf, (C.GLuint)(texture), (C.GLenum)(pname), (C.GLfloat)(param))
}
func TextureParameterfv(texture uint32, pname uint32, param *float32) {
	C.glowTextureParameterfv(gpTextureParameterfv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLfloat)(unsafe.Pointer(param)))
}
func TextureParameteri(texture uint32, pname uint32, param int32) {
	C.glowTextureParameteri(gpTextureParameteri, (C.GLuint)(texture), (C.GLenum)(pname), (C.GLint)(param))
}
func TextureParameteriv(texture uint32, pname uint32, param *int32) {
	C.glowTextureParameteriv(gpTextureParameteriv, (C.GLuint)(texture), (C.GLenum)(pname), (*C.GLint)(unsafe.Pointer(param)))
}

// simultaneously specify storage for all levels of a one-dimensional texture
func TextureStorage1D(texture uint32, levels int32, internalformat uint32, width int32) {
	C.glowTextureStorage1D(gpTextureStorage1D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width))
}

// simultaneously specify storage for all levels of a two-dimensional or one-dimensional array texture
func TextureStorage2D(texture uint32, levels int32, internalformat uint32, width int32, height int32) {
	C.glowTextureStorage2D(gpTextureStorage2D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height))
}

// specify storage for a two-dimensional multisample texture
func TextureStorage2DMultisample(texture uint32, samples int32, internalformat uint32, width int32, height int32, fixedsamplelocations bool) {
	C.glowTextureStorage2DMultisample(gpTextureStorage2DMultisample, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// simultaneously specify storage for all levels of a three-dimensional, two-dimensional array or cube-map array texture
func TextureStorage3D(texture uint32, levels int32, internalformat uint32, width int32, height int32, depth int32) {
	C.glowTextureStorage3D(gpTextureStorage3D, (C.GLuint)(texture), (C.GLsizei)(levels), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth))
}

// specify storage for a two-dimensional multisample array texture
func TextureStorage3DMultisample(texture uint32, samples int32, internalformat uint32, width int32, height int32, depth int32, fixedsamplelocations bool) {
	C.glowTextureStorage3DMultisample(gpTextureStorage3DMultisample, (C.GLuint)(texture), (C.GLsizei)(samples), (C.GLenum)(internalformat), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLboolean)(boolToInt(fixedsamplelocations)))
}

// specify a one-dimensional texture subimage
func TextureSubImage1D(texture uint32, level int32, xoffset int32, width int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTextureSubImage1D(gpTextureSubImage1D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLsizei)(width), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a two-dimensional texture subimage
func TextureSubImage2D(texture uint32, level int32, xoffset int32, yoffset int32, width int32, height int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTextureSubImage2D(gpTextureSubImage2D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// specify a three-dimensional texture subimage
func TextureSubImage3D(texture uint32, level int32, xoffset int32, yoffset int32, zoffset int32, width int32, height int32, depth int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	C.glowTextureSubImage3D(gpTextureSubImage3D, (C.GLuint)(texture), (C.GLint)(level), (C.GLint)(xoffset), (C.GLint)(yoffset), (C.GLint)(zoffset), (C.GLsizei)(width), (C.GLsizei)(height), (C.GLsizei)(depth), (C.GLenum)(format), (C.GLenum)(xtype), pixels)
}

// initialize a texture as a data alias of another texture's data store
func TextureView(texture uint32, target uint32, origtexture uint32, internalformat uint32, minlevel uint32, numlevels uint32, minlayer uint32, numlayers uint32) {
	C.glowTextureView(gpTextureView, (C.GLuint)(texture), (C.GLenum)(target), (C.GLuint)(origtexture), (C.GLenum)(internalformat), (C.GLuint)(minlevel), (C.GLuint)(numlevels), (C.GLuint)(minlayer), (C.GLuint)(numlayers))
}

// bind a buffer object to a transform feedback buffer object
func TransformFeedbackBufferBase(xfb uint32, index uint32, buffer uint32) {
	C.glowTransformFeedbackBufferBase(gpTransformFeedbackBufferBase, (C.GLuint)(xfb), (C.GLuint)(index), (C.GLuint)(buffer))
}

// bind a range within a buffer object to a transform feedback buffer object
func TransformFeedbackBufferRange(xfb uint32, index uint32, buffer uint32, offset int, size int32) {
	C.glowTransformFeedbackBufferRange(gpTransformFeedbackBufferRange, (C.GLuint)(xfb), (C.GLuint)(index), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(size))
}

// specify values to record in transform feedback buffers
func TransformFeedbackVaryings(program uint32, count int32, varyings **uint8, bufferMode uint32) {
	C.glowTransformFeedbackVaryings(gpTransformFeedbackVaryings, (C.GLuint)(program), (C.GLsizei)(count), (**C.GLchar)(unsafe.Pointer(varyings)), (C.GLenum)(bufferMode))
}
func TranslatexOES(x int32, y int32, z int32) {
	C.glowTranslatexOES(gpTranslatexOES, (C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
}
func Uniform1d(location int32, x float64) {
	C.glowUniform1d(gpUniform1d, (C.GLint)(location), (C.GLdouble)(x))
}
func Uniform1dv(location int32, count int32, value *float64) {
	C.glowUniform1dv(gpUniform1dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1f(location int32, v0 float32) {
	C.glowUniform1f(gpUniform1f, (C.GLint)(location), (C.GLfloat)(v0))
}

// Specify the value of a uniform variable for the current program object
func Uniform1fv(location int32, count int32, value *float32) {
	C.glowUniform1fv(gpUniform1fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1i(location int32, v0 int32) {
	C.glowUniform1i(gpUniform1i, (C.GLint)(location), (C.GLint)(v0))
}

// Specify the value of a uniform variable for the current program object
func Uniform1iv(location int32, count int32, value *int32) {
	C.glowUniform1iv(gpUniform1iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform1ui(location int32, v0 uint32) {
	C.glowUniform1ui(gpUniform1ui, (C.GLint)(location), (C.GLuint)(v0))
}

// Specify the value of a uniform variable for the current program object
func Uniform1uiv(location int32, count int32, value *uint32) {
	C.glowUniform1uiv(gpUniform1uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform2d(location int32, x float64, y float64) {
	C.glowUniform2d(gpUniform2d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y))
}
func Uniform2dv(location int32, count int32, value *float64) {
	C.glowUniform2dv(gpUniform2dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2f(location int32, v0 float32, v1 float32) {
	C.glowUniform2f(gpUniform2f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2fv(location int32, count int32, value *float32) {
	C.glowUniform2fv(gpUniform2fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2i(location int32, v0 int32, v1 int32) {
	C.glowUniform2i(gpUniform2i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2iv(location int32, count int32, value *int32) {
	C.glowUniform2iv(gpUniform2iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform2ui(location int32, v0 uint32, v1 uint32) {
	C.glowUniform2ui(gpUniform2ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1))
}

// Specify the value of a uniform variable for the current program object
func Uniform2uiv(location int32, count int32, value *uint32) {
	C.glowUniform2uiv(gpUniform2uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform3d(location int32, x float64, y float64, z float64) {
	C.glowUniform3d(gpUniform3d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func Uniform3dv(location int32, count int32, value *float64) {
	C.glowUniform3dv(gpUniform3dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3f(location int32, v0 float32, v1 float32, v2 float32) {
	C.glowUniform3f(gpUniform3f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2))
}

// Specify the value of a uniform variable for the current program object
func Uniform3fv(location int32, count int32, value *float32) {
	C.glowUniform3fv(gpUniform3fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3i(location int32, v0 int32, v1 int32, v2 int32) {
	C.glowUniform3i(gpUniform3i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2))
}

// Specify the value of a uniform variable for the current program object
func Uniform3iv(location int32, count int32, value *int32) {
	C.glowUniform3iv(gpUniform3iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform3ui(location int32, v0 uint32, v1 uint32, v2 uint32) {
	C.glowUniform3ui(gpUniform3ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2))
}

// Specify the value of a uniform variable for the current program object
func Uniform3uiv(location int32, count int32, value *uint32) {
	C.glowUniform3uiv(gpUniform3uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}
func Uniform4d(location int32, x float64, y float64, z float64, w float64) {
	C.glowUniform4d(gpUniform4d, (C.GLint)(location), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func Uniform4dv(location int32, count int32, value *float64) {
	C.glowUniform4dv(gpUniform4dv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4f(location int32, v0 float32, v1 float32, v2 float32, v3 float32) {
	C.glowUniform4f(gpUniform4f, (C.GLint)(location), (C.GLfloat)(v0), (C.GLfloat)(v1), (C.GLfloat)(v2), (C.GLfloat)(v3))
}

// Specify the value of a uniform variable for the current program object
func Uniform4fv(location int32, count int32, value *float32) {
	C.glowUniform4fv(gpUniform4fv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4i(location int32, v0 int32, v1 int32, v2 int32, v3 int32) {
	C.glowUniform4i(gpUniform4i, (C.GLint)(location), (C.GLint)(v0), (C.GLint)(v1), (C.GLint)(v2), (C.GLint)(v3))
}

// Specify the value of a uniform variable for the current program object
func Uniform4iv(location int32, count int32, value *int32) {
	C.glowUniform4iv(gpUniform4iv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLint)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func Uniform4ui(location int32, v0 uint32, v1 uint32, v2 uint32, v3 uint32) {
	C.glowUniform4ui(gpUniform4ui, (C.GLint)(location), (C.GLuint)(v0), (C.GLuint)(v1), (C.GLuint)(v2), (C.GLuint)(v3))
}

// Specify the value of a uniform variable for the current program object
func Uniform4uiv(location int32, count int32, value *uint32) {
	C.glowUniform4uiv(gpUniform4uiv, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(value)))
}

// assign a binding point to an active uniform block
func UniformBlockBinding(program uint32, uniformBlockIndex uint32, uniformBlockBinding uint32) {
	C.glowUniformBlockBinding(gpUniformBlockBinding, (C.GLuint)(program), (C.GLuint)(uniformBlockIndex), (C.GLuint)(uniformBlockBinding))
}
func UniformHandleui64ARB(location int32, value uint64) {
	C.glowUniformHandleui64ARB(gpUniformHandleui64ARB, (C.GLint)(location), (C.GLuint64)(value))
}
func UniformHandleui64vARB(location int32, count int32, value *uint64) {
	C.glowUniformHandleui64vARB(gpUniformHandleui64vARB, (C.GLint)(location), (C.GLsizei)(count), (*C.GLuint64)(unsafe.Pointer(value)))
}
func UniformMatrix2dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix2dv(gpUniformMatrix2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2fv(gpUniformMatrix2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x3dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix2x3dv(gpUniformMatrix2x3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x3fv(gpUniformMatrix2x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix2x4dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix2x4dv(gpUniformMatrix2x4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix2x4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix2x4fv(gpUniformMatrix2x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix3dv(gpUniformMatrix3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3fv(gpUniformMatrix3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x2dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix3x2dv(gpUniformMatrix3x2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x2fv(gpUniformMatrix3x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix3x4dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix3x4dv(gpUniformMatrix3x4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix3x4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix3x4fv(gpUniformMatrix3x4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix4dv(gpUniformMatrix4dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4fv(gpUniformMatrix4fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x2dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix4x2dv(gpUniformMatrix4x2dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x2fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x2fv(gpUniformMatrix4x2fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformMatrix4x3dv(location int32, count int32, transpose bool, value *float64) {
	C.glowUniformMatrix4x3dv(gpUniformMatrix4x3dv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLdouble)(unsafe.Pointer(value)))
}

// Specify the value of a uniform variable for the current program object
func UniformMatrix4x3fv(location int32, count int32, transpose bool, value *float32) {
	C.glowUniformMatrix4x3fv(gpUniformMatrix4x3fv, (C.GLint)(location), (C.GLsizei)(count), (C.GLboolean)(boolToInt(transpose)), (*C.GLfloat)(unsafe.Pointer(value)))
}
func UniformSubroutinesuiv(shadertype uint32, count int32, indices *uint32) {
	C.glowUniformSubroutinesuiv(gpUniformSubroutinesuiv, (C.GLenum)(shadertype), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(indices)))
}

// release the mapping of a buffer object's data store into the client's address space
func UnmapBuffer(target uint32) bool {
	ret := C.glowUnmapBuffer(gpUnmapBuffer, (C.GLenum)(target))
	return ret == TRUE
}

// release the mapping of a buffer object's data store into the client's address space
func UnmapNamedBuffer(buffer uint32) bool {
	ret := C.glowUnmapNamedBuffer(gpUnmapNamedBuffer, (C.GLuint)(buffer))
	return ret == TRUE
}

// Installs a program object as part of current rendering state
func UseProgram(program uint32) {
	C.glowUseProgram(gpUseProgram, (C.GLuint)(program))
}

// bind stages of a program object to a program pipeline
func UseProgramStages(pipeline uint32, stages uint32, program uint32) {
	C.glowUseProgramStages(gpUseProgramStages, (C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}
func UseProgramStagesEXT(pipeline uint32, stages uint32, program uint32) {
	C.glowUseProgramStagesEXT(gpUseProgramStagesEXT, (C.GLuint)(pipeline), (C.GLbitfield)(stages), (C.GLuint)(program))
}
func UseShaderProgramEXT(xtype uint32, program uint32) {
	C.glowUseShaderProgramEXT(gpUseShaderProgramEXT, (C.GLenum)(xtype), (C.GLuint)(program))
}

// Validates a program object
func ValidateProgram(program uint32) {
	C.glowValidateProgram(gpValidateProgram, (C.GLuint)(program))
}

// validate a program pipeline object against current GL state
func ValidateProgramPipeline(pipeline uint32) {
	C.glowValidateProgramPipeline(gpValidateProgramPipeline, (C.GLuint)(pipeline))
}
func ValidateProgramPipelineEXT(pipeline uint32) {
	C.glowValidateProgramPipelineEXT(gpValidateProgramPipelineEXT, (C.GLuint)(pipeline))
}
func Vertex2bOES(x int8, y int8) {
	C.glowVertex2bOES(gpVertex2bOES, (C.GLbyte)(x), (C.GLbyte)(y))
}
func Vertex2bvOES(coords *int8) {
	C.glowVertex2bvOES(gpVertex2bvOES, (*C.GLbyte)(unsafe.Pointer(coords)))
}
func Vertex2xOES(x int32) {
	C.glowVertex2xOES(gpVertex2xOES, (C.GLfixed)(x))
}
func Vertex2xvOES(coords *int32) {
	C.glowVertex2xvOES(gpVertex2xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func Vertex3bOES(x int8, y int8, z int8) {
	C.glowVertex3bOES(gpVertex3bOES, (C.GLbyte)(x), (C.GLbyte)(y), (C.GLbyte)(z))
}
func Vertex3bvOES(coords *int8) {
	C.glowVertex3bvOES(gpVertex3bvOES, (*C.GLbyte)(unsafe.Pointer(coords)))
}
func Vertex3xOES(x int32, y int32) {
	C.glowVertex3xOES(gpVertex3xOES, (C.GLfixed)(x), (C.GLfixed)(y))
}
func Vertex3xvOES(coords *int32) {
	C.glowVertex3xvOES(gpVertex3xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func Vertex4bOES(x int8, y int8, z int8, w int8) {
	C.glowVertex4bOES(gpVertex4bOES, (C.GLbyte)(x), (C.GLbyte)(y), (C.GLbyte)(z), (C.GLbyte)(w))
}
func Vertex4bvOES(coords *int8) {
	C.glowVertex4bvOES(gpVertex4bvOES, (*C.GLbyte)(unsafe.Pointer(coords)))
}
func Vertex4xOES(x int32, y int32, z int32) {
	C.glowVertex4xOES(gpVertex4xOES, (C.GLfixed)(x), (C.GLfixed)(y), (C.GLfixed)(z))
}
func Vertex4xvOES(coords *int32) {
	C.glowVertex4xvOES(gpVertex4xvOES, (*C.GLfixed)(unsafe.Pointer(coords)))
}
func VertexArrayAttribBinding(vaobj uint32, attribindex uint32, bindingindex uint32) {
	C.glowVertexArrayAttribBinding(gpVertexArrayAttribBinding, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}

// specify the organization of vertex arrays
func VertexArrayAttribFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	C.glowVertexArrayAttribFormat(gpVertexArrayAttribFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(relativeoffset))
}
func VertexArrayAttribIFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexArrayAttribIFormat(gpVertexArrayAttribIFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}
func VertexArrayAttribLFormat(vaobj uint32, attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexArrayAttribLFormat(gpVertexArrayAttribLFormat, (C.GLuint)(vaobj), (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}

// modify the rate at which generic vertex attributes     advance
func VertexArrayBindingDivisor(vaobj uint32, bindingindex uint32, divisor uint32) {
	C.glowVertexArrayBindingDivisor(gpVertexArrayBindingDivisor, (C.GLuint)(vaobj), (C.GLuint)(bindingindex), (C.GLuint)(divisor))
}

// configures element array buffer binding of a vertex array object
func VertexArrayElementBuffer(vaobj uint32, buffer uint32) {
	C.glowVertexArrayElementBuffer(gpVertexArrayElementBuffer, (C.GLuint)(vaobj), (C.GLuint)(buffer))
}

// bind a buffer to a vertex buffer bind point
func VertexArrayVertexBuffer(vaobj uint32, bindingindex uint32, buffer uint32, offset int, stride int32) {
	C.glowVertexArrayVertexBuffer(gpVertexArrayVertexBuffer, (C.GLuint)(vaobj), (C.GLuint)(bindingindex), (C.GLuint)(buffer), (C.GLintptr)(offset), (C.GLsizei)(stride))
}

// attach multiple buffer objects to a vertex array object
func VertexArrayVertexBuffers(vaobj uint32, first uint32, count int32, buffers *uint32, offsets *int, strides *int32) {
	C.glowVertexArrayVertexBuffers(gpVertexArrayVertexBuffers, (C.GLuint)(vaobj), (C.GLuint)(first), (C.GLsizei)(count), (*C.GLuint)(unsafe.Pointer(buffers)), (*C.GLintptr)(unsafe.Pointer(offsets)), (*C.GLsizei)(unsafe.Pointer(strides)))
}
func VertexAttrib1d(index uint32, x float64) {
	C.glowVertexAttrib1d(gpVertexAttrib1d, (C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttrib1dv(index uint32, v *float64) {
	C.glowVertexAttrib1dv(gpVertexAttrib1dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib1f(index uint32, x float32) {
	C.glowVertexAttrib1f(gpVertexAttrib1f, (C.GLuint)(index), (C.GLfloat)(x))
}
func VertexAttrib1fv(index uint32, v *float32) {
	C.glowVertexAttrib1fv(gpVertexAttrib1fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib1s(index uint32, x int16) {
	C.glowVertexAttrib1s(gpVertexAttrib1s, (C.GLuint)(index), (C.GLshort)(x))
}
func VertexAttrib1sv(index uint32, v *int16) {
	C.glowVertexAttrib1sv(gpVertexAttrib1sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib2d(index uint32, x float64, y float64) {
	C.glowVertexAttrib2d(gpVertexAttrib2d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttrib2dv(index uint32, v *float64) {
	C.glowVertexAttrib2dv(gpVertexAttrib2dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib2f(index uint32, x float32, y float32) {
	C.glowVertexAttrib2f(gpVertexAttrib2f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y))
}
func VertexAttrib2fv(index uint32, v *float32) {
	C.glowVertexAttrib2fv(gpVertexAttrib2fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib2s(index uint32, x int16, y int16) {
	C.glowVertexAttrib2s(gpVertexAttrib2s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y))
}
func VertexAttrib2sv(index uint32, v *int16) {
	C.glowVertexAttrib2sv(gpVertexAttrib2sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib3d(index uint32, x float64, y float64, z float64) {
	C.glowVertexAttrib3d(gpVertexAttrib3d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttrib3dv(index uint32, v *float64) {
	C.glowVertexAttrib3dv(gpVertexAttrib3dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib3f(index uint32, x float32, y float32, z float32) {
	C.glowVertexAttrib3f(gpVertexAttrib3f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z))
}
func VertexAttrib3fv(index uint32, v *float32) {
	C.glowVertexAttrib3fv(gpVertexAttrib3fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib3s(index uint32, x int16, y int16, z int16) {
	C.glowVertexAttrib3s(gpVertexAttrib3s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z))
}
func VertexAttrib3sv(index uint32, v *int16) {
	C.glowVertexAttrib3sv(gpVertexAttrib3sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4Nbv(index uint32, v *int8) {
	C.glowVertexAttrib4Nbv(gpVertexAttrib4Nbv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttrib4Niv(index uint32, v *int32) {
	C.glowVertexAttrib4Niv(gpVertexAttrib4Niv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttrib4Nsv(index uint32, v *int16) {
	C.glowVertexAttrib4Nsv(gpVertexAttrib4Nsv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4Nub(index uint32, x uint8, y uint8, z uint8, w uint8) {
	C.glowVertexAttrib4Nub(gpVertexAttrib4Nub, (C.GLuint)(index), (C.GLubyte)(x), (C.GLubyte)(y), (C.GLubyte)(z), (C.GLubyte)(w))
}
func VertexAttrib4Nubv(index uint32, v *uint8) {
	C.glowVertexAttrib4Nubv(gpVertexAttrib4Nubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttrib4Nuiv(index uint32, v *uint32) {
	C.glowVertexAttrib4Nuiv(gpVertexAttrib4Nuiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttrib4Nusv(index uint32, v *uint16) {
	C.glowVertexAttrib4Nusv(gpVertexAttrib4Nusv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}
func VertexAttrib4bv(index uint32, v *int8) {
	C.glowVertexAttrib4bv(gpVertexAttrib4bv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttrib4d(index uint32, x float64, y float64, z float64, w float64) {
	C.glowVertexAttrib4d(gpVertexAttrib4d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttrib4dv(index uint32, v *float64) {
	C.glowVertexAttrib4dv(gpVertexAttrib4dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttrib4f(index uint32, x float32, y float32, z float32, w float32) {
	C.glowVertexAttrib4f(gpVertexAttrib4f, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(z), (C.GLfloat)(w))
}
func VertexAttrib4fv(index uint32, v *float32) {
	C.glowVertexAttrib4fv(gpVertexAttrib4fv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}
func VertexAttrib4iv(index uint32, v *int32) {
	C.glowVertexAttrib4iv(gpVertexAttrib4iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttrib4s(index uint32, x int16, y int16, z int16, w int16) {
	C.glowVertexAttrib4s(gpVertexAttrib4s, (C.GLuint)(index), (C.GLshort)(x), (C.GLshort)(y), (C.GLshort)(z), (C.GLshort)(w))
}
func VertexAttrib4sv(index uint32, v *int16) {
	C.glowVertexAttrib4sv(gpVertexAttrib4sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttrib4ubv(index uint32, v *uint8) {
	C.glowVertexAttrib4ubv(gpVertexAttrib4ubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttrib4uiv(index uint32, v *uint32) {
	C.glowVertexAttrib4uiv(gpVertexAttrib4uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttrib4usv(index uint32, v *uint16) {
	C.glowVertexAttrib4usv(gpVertexAttrib4usv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}

// associate a vertex attribute and a vertex buffer binding for a vertex array object
func VertexAttribBinding(attribindex uint32, bindingindex uint32) {
	C.glowVertexAttribBinding(gpVertexAttribBinding, (C.GLuint)(attribindex), (C.GLuint)(bindingindex))
}

// modify the rate at which generic vertex attributes advance during instanced rendering
func VertexAttribDivisor(index uint32, divisor uint32) {
	C.glowVertexAttribDivisor(gpVertexAttribDivisor, (C.GLuint)(index), (C.GLuint)(divisor))
}

// specify the organization of vertex arrays
func VertexAttribFormat(attribindex uint32, size int32, xtype uint32, normalized bool, relativeoffset uint32) {
	C.glowVertexAttribFormat(gpVertexAttribFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(relativeoffset))
}
func VertexAttribI1i(index uint32, x int32) {
	C.glowVertexAttribI1i(gpVertexAttribI1i, (C.GLuint)(index), (C.GLint)(x))
}
func VertexAttribI1iv(index uint32, v *int32) {
	C.glowVertexAttribI1iv(gpVertexAttribI1iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI1ui(index uint32, x uint32) {
	C.glowVertexAttribI1ui(gpVertexAttribI1ui, (C.GLuint)(index), (C.GLuint)(x))
}
func VertexAttribI1uiv(index uint32, v *uint32) {
	C.glowVertexAttribI1uiv(gpVertexAttribI1uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI2i(index uint32, x int32, y int32) {
	C.glowVertexAttribI2i(gpVertexAttribI2i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y))
}
func VertexAttribI2iv(index uint32, v *int32) {
	C.glowVertexAttribI2iv(gpVertexAttribI2iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI2ui(index uint32, x uint32, y uint32) {
	C.glowVertexAttribI2ui(gpVertexAttribI2ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y))
}
func VertexAttribI2uiv(index uint32, v *uint32) {
	C.glowVertexAttribI2uiv(gpVertexAttribI2uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI3i(index uint32, x int32, y int32, z int32) {
	C.glowVertexAttribI3i(gpVertexAttribI3i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z))
}
func VertexAttribI3iv(index uint32, v *int32) {
	C.glowVertexAttribI3iv(gpVertexAttribI3iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI3ui(index uint32, x uint32, y uint32, z uint32) {
	C.glowVertexAttribI3ui(gpVertexAttribI3ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z))
}
func VertexAttribI3uiv(index uint32, v *uint32) {
	C.glowVertexAttribI3uiv(gpVertexAttribI3uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI4bv(index uint32, v *int8) {
	C.glowVertexAttribI4bv(gpVertexAttribI4bv, (C.GLuint)(index), (*C.GLbyte)(unsafe.Pointer(v)))
}
func VertexAttribI4i(index uint32, x int32, y int32, z int32, w int32) {
	C.glowVertexAttribI4i(gpVertexAttribI4i, (C.GLuint)(index), (C.GLint)(x), (C.GLint)(y), (C.GLint)(z), (C.GLint)(w))
}
func VertexAttribI4iv(index uint32, v *int32) {
	C.glowVertexAttribI4iv(gpVertexAttribI4iv, (C.GLuint)(index), (*C.GLint)(unsafe.Pointer(v)))
}
func VertexAttribI4sv(index uint32, v *int16) {
	C.glowVertexAttribI4sv(gpVertexAttribI4sv, (C.GLuint)(index), (*C.GLshort)(unsafe.Pointer(v)))
}
func VertexAttribI4ubv(index uint32, v *uint8) {
	C.glowVertexAttribI4ubv(gpVertexAttribI4ubv, (C.GLuint)(index), (*C.GLubyte)(unsafe.Pointer(v)))
}
func VertexAttribI4ui(index uint32, x uint32, y uint32, z uint32, w uint32) {
	C.glowVertexAttribI4ui(gpVertexAttribI4ui, (C.GLuint)(index), (C.GLuint)(x), (C.GLuint)(y), (C.GLuint)(z), (C.GLuint)(w))
}
func VertexAttribI4uiv(index uint32, v *uint32) {
	C.glowVertexAttribI4uiv(gpVertexAttribI4uiv, (C.GLuint)(index), (*C.GLuint)(unsafe.Pointer(v)))
}
func VertexAttribI4usv(index uint32, v *uint16) {
	C.glowVertexAttribI4usv(gpVertexAttribI4usv, (C.GLuint)(index), (*C.GLushort)(unsafe.Pointer(v)))
}
func VertexAttribIFormat(attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexAttribIFormat(gpVertexAttribIFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}
func VertexAttribIPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribIPointer(gpVertexAttribIPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}
func VertexAttribL1d(index uint32, x float64) {
	C.glowVertexAttribL1d(gpVertexAttribL1d, (C.GLuint)(index), (C.GLdouble)(x))
}
func VertexAttribL1dv(index uint32, v *float64) {
	C.glowVertexAttribL1dv(gpVertexAttribL1dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL1ui64ARB(index uint32, x uint64) {
	C.glowVertexAttribL1ui64ARB(gpVertexAttribL1ui64ARB, (C.GLuint)(index), (C.GLuint64EXT)(x))
}
func VertexAttribL1ui64vARB(index uint32, v *uint64) {
	C.glowVertexAttribL1ui64vARB(gpVertexAttribL1ui64vARB, (C.GLuint)(index), (*C.GLuint64EXT)(unsafe.Pointer(v)))
}
func VertexAttribL2d(index uint32, x float64, y float64) {
	C.glowVertexAttribL2d(gpVertexAttribL2d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y))
}
func VertexAttribL2dv(index uint32, v *float64) {
	C.glowVertexAttribL2dv(gpVertexAttribL2dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL3d(index uint32, x float64, y float64, z float64) {
	C.glowVertexAttribL3d(gpVertexAttribL3d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z))
}
func VertexAttribL3dv(index uint32, v *float64) {
	C.glowVertexAttribL3dv(gpVertexAttribL3dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribL4d(index uint32, x float64, y float64, z float64, w float64) {
	C.glowVertexAttribL4d(gpVertexAttribL4d, (C.GLuint)(index), (C.GLdouble)(x), (C.GLdouble)(y), (C.GLdouble)(z), (C.GLdouble)(w))
}
func VertexAttribL4dv(index uint32, v *float64) {
	C.glowVertexAttribL4dv(gpVertexAttribL4dv, (C.GLuint)(index), (*C.GLdouble)(unsafe.Pointer(v)))
}
func VertexAttribLFormat(attribindex uint32, size int32, xtype uint32, relativeoffset uint32) {
	C.glowVertexAttribLFormat(gpVertexAttribLFormat, (C.GLuint)(attribindex), (C.GLint)(size), (C.GLenum)(xtype), (C.GLuint)(relativeoffset))
}
func VertexAttribLPointer(index uint32, size int32, xtype uint32, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribLPointer(gpVertexAttribLPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLsizei)(stride), pointer)
}
func VertexAttribP1ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP1ui(gpVertexAttribP1ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP1uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP1uiv(gpVertexAttribP1uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP2ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP2ui(gpVertexAttribP2ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP2uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP2uiv(gpVertexAttribP2uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP3ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP3ui(gpVertexAttribP3ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP3uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP3uiv(gpVertexAttribP3uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}
func VertexAttribP4ui(index uint32, xtype uint32, normalized bool, value uint32) {
	C.glowVertexAttribP4ui(gpVertexAttribP4ui, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLuint)(value))
}
func VertexAttribP4uiv(index uint32, xtype uint32, normalized bool, value *uint32) {
	C.glowVertexAttribP4uiv(gpVertexAttribP4uiv, (C.GLuint)(index), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (*C.GLuint)(unsafe.Pointer(value)))
}

// define an array of generic vertex attribute data
func VertexAttribPointer(index uint32, size int32, xtype uint32, normalized bool, stride int32, pointer unsafe.Pointer) {
	C.glowVertexAttribPointer(gpVertexAttribPointer, (C.GLuint)(index), (C.GLint)(size), (C.GLenum)(xtype), (C.GLboolean)(boolToInt(normalized)), (C.GLsizei)(stride), pointer)
}

// modify the rate at which generic vertex attributes     advance
func VertexBindingDivisor(bindingindex uint32, divisor uint32) {
	C.glowVertexBindingDivisor(gpVertexBindingDivisor, (C.GLuint)(bindingindex), (C.GLuint)(divisor))
}

// set the viewport
func Viewport(x int32, y int32, width int32, height int32) {
	C.glowViewport(gpViewport, (C.GLint)(x), (C.GLint)(y), (C.GLsizei)(width), (C.GLsizei)(height))
}
func ViewportArrayv(first uint32, count int32, v *float32) {
	C.glowViewportArrayv(gpViewportArrayv, (C.GLuint)(first), (C.GLsizei)(count), (*C.GLfloat)(unsafe.Pointer(v)))
}
func ViewportIndexedf(index uint32, x float32, y float32, w float32, h float32) {
	C.glowViewportIndexedf(gpViewportIndexedf, (C.GLuint)(index), (C.GLfloat)(x), (C.GLfloat)(y), (C.GLfloat)(w), (C.GLfloat)(h))
}
func ViewportIndexedfv(index uint32, v *float32) {
	C.glowViewportIndexedfv(gpViewportIndexedfv, (C.GLuint)(index), (*C.GLfloat)(unsafe.Pointer(v)))
}

// instruct the GL server to block until the specified sync object becomes signaled
func WaitSync(sync unsafe.Pointer, flags uint32, timeout uint64) {
	C.glowWaitSync(gpWaitSync, (C.GLsync)(sync), (C.GLbitfield)(flags), (C.GLuint64)(timeout))
}
func Init() error {
	return InitWithProcAddrFunc(auto.GetProcAddress)
}
func InitWithProcAddrFunc(getProcAddr procaddr.GetProcAddressFunc) error {
	gpAccumxOES = (C.GPACCUMXOES)(getProcAddr("glAccumxOES"))
	gpActiveProgramEXT = (C.GPACTIVEPROGRAMEXT)(getProcAddr("glActiveProgramEXT"))
	gpActiveShaderProgram = (C.GPACTIVESHADERPROGRAM)(getProcAddr("glActiveShaderProgram"))
	if gpActiveShaderProgram == nil {
		return errors.New("glActiveShaderProgram")
	}
	gpActiveShaderProgramEXT = (C.GPACTIVESHADERPROGRAMEXT)(getProcAddr("glActiveShaderProgramEXT"))
	gpActiveTexture = (C.GPACTIVETEXTURE)(getProcAddr("glActiveTexture"))
	if gpActiveTexture == nil {
		return errors.New("glActiveTexture")
	}
	gpAlphaFuncxOES = (C.GPALPHAFUNCXOES)(getProcAddr("glAlphaFuncxOES"))
	gpAttachShader = (C.GPATTACHSHADER)(getProcAddr("glAttachShader"))
	if gpAttachShader == nil {
		return errors.New("glAttachShader")
	}
	gpBeginConditionalRender = (C.GPBEGINCONDITIONALRENDER)(getProcAddr("glBeginConditionalRender"))
	if gpBeginConditionalRender == nil {
		return errors.New("glBeginConditionalRender")
	}
	gpBeginPerfMonitorAMD = (C.GPBEGINPERFMONITORAMD)(getProcAddr("glBeginPerfMonitorAMD"))
	gpBeginPerfQueryINTEL = (C.GPBEGINPERFQUERYINTEL)(getProcAddr("glBeginPerfQueryINTEL"))
	gpBeginQuery = (C.GPBEGINQUERY)(getProcAddr("glBeginQuery"))
	if gpBeginQuery == nil {
		return errors.New("glBeginQuery")
	}
	gpBeginQueryIndexed = (C.GPBEGINQUERYINDEXED)(getProcAddr("glBeginQueryIndexed"))
	if gpBeginQueryIndexed == nil {
		return errors.New("glBeginQueryIndexed")
	}
	gpBeginTransformFeedback = (C.GPBEGINTRANSFORMFEEDBACK)(getProcAddr("glBeginTransformFeedback"))
	if gpBeginTransformFeedback == nil {
		return errors.New("glBeginTransformFeedback")
	}
	gpBindAttribLocation = (C.GPBINDATTRIBLOCATION)(getProcAddr("glBindAttribLocation"))
	if gpBindAttribLocation == nil {
		return errors.New("glBindAttribLocation")
	}
	gpBindBuffer = (C.GPBINDBUFFER)(getProcAddr("glBindBuffer"))
	if gpBindBuffer == nil {
		return errors.New("glBindBuffer")
	}
	gpBindBufferBase = (C.GPBINDBUFFERBASE)(getProcAddr("glBindBufferBase"))
	if gpBindBufferBase == nil {
		return errors.New("glBindBufferBase")
	}
	gpBindBufferRange = (C.GPBINDBUFFERRANGE)(getProcAddr("glBindBufferRange"))
	if gpBindBufferRange == nil {
		return errors.New("glBindBufferRange")
	}
	gpBindBuffersBase = (C.GPBINDBUFFERSBASE)(getProcAddr("glBindBuffersBase"))
	gpBindBuffersRange = (C.GPBINDBUFFERSRANGE)(getProcAddr("glBindBuffersRange"))
	gpBindFragDataLocation = (C.GPBINDFRAGDATALOCATION)(getProcAddr("glBindFragDataLocation"))
	if gpBindFragDataLocation == nil {
		return errors.New("glBindFragDataLocation")
	}
	gpBindFragDataLocationIndexed = (C.GPBINDFRAGDATALOCATIONINDEXED)(getProcAddr("glBindFragDataLocationIndexed"))
	if gpBindFragDataLocationIndexed == nil {
		return errors.New("glBindFragDataLocationIndexed")
	}
	gpBindFramebuffer = (C.GPBINDFRAMEBUFFER)(getProcAddr("glBindFramebuffer"))
	if gpBindFramebuffer == nil {
		return errors.New("glBindFramebuffer")
	}
	gpBindImageTexture = (C.GPBINDIMAGETEXTURE)(getProcAddr("glBindImageTexture"))
	gpBindImageTextures = (C.GPBINDIMAGETEXTURES)(getProcAddr("glBindImageTextures"))
	gpBindProgramPipeline = (C.GPBINDPROGRAMPIPELINE)(getProcAddr("glBindProgramPipeline"))
	if gpBindProgramPipeline == nil {
		return errors.New("glBindProgramPipeline")
	}
	gpBindProgramPipelineEXT = (C.GPBINDPROGRAMPIPELINEEXT)(getProcAddr("glBindProgramPipelineEXT"))
	gpBindRenderbuffer = (C.GPBINDRENDERBUFFER)(getProcAddr("glBindRenderbuffer"))
	if gpBindRenderbuffer == nil {
		return errors.New("glBindRenderbuffer")
	}
	gpBindSampler = (C.GPBINDSAMPLER)(getProcAddr("glBindSampler"))
	if gpBindSampler == nil {
		return errors.New("glBindSampler")
	}
	gpBindSamplers = (C.GPBINDSAMPLERS)(getProcAddr("glBindSamplers"))
	gpBindTexture = (C.GPBINDTEXTURE)(getProcAddr("glBindTexture"))
	if gpBindTexture == nil {
		return errors.New("glBindTexture")
	}
	gpBindTextureUnit = (C.GPBINDTEXTUREUNIT)(getProcAddr("glBindTextureUnit"))
	gpBindTextures = (C.GPBINDTEXTURES)(getProcAddr("glBindTextures"))
	gpBindTransformFeedback = (C.GPBINDTRANSFORMFEEDBACK)(getProcAddr("glBindTransformFeedback"))
	if gpBindTransformFeedback == nil {
		return errors.New("glBindTransformFeedback")
	}
	gpBindVertexArray = (C.GPBINDVERTEXARRAY)(getProcAddr("glBindVertexArray"))
	if gpBindVertexArray == nil {
		return errors.New("glBindVertexArray")
	}
	gpBindVertexBuffer = (C.GPBINDVERTEXBUFFER)(getProcAddr("glBindVertexBuffer"))
	gpBindVertexBuffers = (C.GPBINDVERTEXBUFFERS)(getProcAddr("glBindVertexBuffers"))
	gpBitmapxOES = (C.GPBITMAPXOES)(getProcAddr("glBitmapxOES"))
	gpBlendBarrierKHR = (C.GPBLENDBARRIERKHR)(getProcAddr("glBlendBarrierKHR"))
	gpBlendBarrierNV = (C.GPBLENDBARRIERNV)(getProcAddr("glBlendBarrierNV"))
	gpBlendColor = (C.GPBLENDCOLOR)(getProcAddr("glBlendColor"))
	if gpBlendColor == nil {
		return errors.New("glBlendColor")
	}
	gpBlendColorxOES = (C.GPBLENDCOLORXOES)(getProcAddr("glBlendColorxOES"))
	gpBlendEquation = (C.GPBLENDEQUATION)(getProcAddr("glBlendEquation"))
	if gpBlendEquation == nil {
		return errors.New("glBlendEquation")
	}
	gpBlendEquationEXT = (C.GPBLENDEQUATIONEXT)(getProcAddr("glBlendEquationEXT"))
	gpBlendEquationSeparate = (C.GPBLENDEQUATIONSEPARATE)(getProcAddr("glBlendEquationSeparate"))
	if gpBlendEquationSeparate == nil {
		return errors.New("glBlendEquationSeparate")
	}
	gpBlendEquationSeparatei = (C.GPBLENDEQUATIONSEPARATEI)(getProcAddr("glBlendEquationSeparatei"))
	if gpBlendEquationSeparatei == nil {
		return errors.New("glBlendEquationSeparatei")
	}
	gpBlendEquationSeparateiARB = (C.GPBLENDEQUATIONSEPARATEIARB)(getProcAddr("glBlendEquationSeparateiARB"))
	gpBlendEquationi = (C.GPBLENDEQUATIONI)(getProcAddr("glBlendEquationi"))
	if gpBlendEquationi == nil {
		return errors.New("glBlendEquationi")
	}
	gpBlendEquationiARB = (C.GPBLENDEQUATIONIARB)(getProcAddr("glBlendEquationiARB"))
	gpBlendFunc = (C.GPBLENDFUNC)(getProcAddr("glBlendFunc"))
	if gpBlendFunc == nil {
		return errors.New("glBlendFunc")
	}
	gpBlendFuncSeparate = (C.GPBLENDFUNCSEPARATE)(getProcAddr("glBlendFuncSeparate"))
	if gpBlendFuncSeparate == nil {
		return errors.New("glBlendFuncSeparate")
	}
	gpBlendFuncSeparatei = (C.GPBLENDFUNCSEPARATEI)(getProcAddr("glBlendFuncSeparatei"))
	if gpBlendFuncSeparatei == nil {
		return errors.New("glBlendFuncSeparatei")
	}
	gpBlendFuncSeparateiARB = (C.GPBLENDFUNCSEPARATEIARB)(getProcAddr("glBlendFuncSeparateiARB"))
	gpBlendFunci = (C.GPBLENDFUNCI)(getProcAddr("glBlendFunci"))
	if gpBlendFunci == nil {
		return errors.New("glBlendFunci")
	}
	gpBlendFunciARB = (C.GPBLENDFUNCIARB)(getProcAddr("glBlendFunciARB"))
	gpBlendParameteriNV = (C.GPBLENDPARAMETERINV)(getProcAddr("glBlendParameteriNV"))
	gpBlitFramebuffer = (C.GPBLITFRAMEBUFFER)(getProcAddr("glBlitFramebuffer"))
	if gpBlitFramebuffer == nil {
		return errors.New("glBlitFramebuffer")
	}
	gpBlitNamedFramebuffer = (C.GPBLITNAMEDFRAMEBUFFER)(getProcAddr("glBlitNamedFramebuffer"))
	gpBufferData = (C.GPBUFFERDATA)(getProcAddr("glBufferData"))
	if gpBufferData == nil {
		return errors.New("glBufferData")
	}
	gpBufferPageCommitmentARB = (C.GPBUFFERPAGECOMMITMENTARB)(getProcAddr("glBufferPageCommitmentARB"))
	gpBufferStorage = (C.GPBUFFERSTORAGE)(getProcAddr("glBufferStorage"))
	gpBufferSubData = (C.GPBUFFERSUBDATA)(getProcAddr("glBufferSubData"))
	if gpBufferSubData == nil {
		return errors.New("glBufferSubData")
	}
	gpCheckFramebufferStatus = (C.GPCHECKFRAMEBUFFERSTATUS)(getProcAddr("glCheckFramebufferStatus"))
	if gpCheckFramebufferStatus == nil {
		return errors.New("glCheckFramebufferStatus")
	}
	gpCheckNamedFramebufferStatus = (C.GPCHECKNAMEDFRAMEBUFFERSTATUS)(getProcAddr("glCheckNamedFramebufferStatus"))
	gpClampColor = (C.GPCLAMPCOLOR)(getProcAddr("glClampColor"))
	if gpClampColor == nil {
		return errors.New("glClampColor")
	}
	gpClear = (C.GPCLEAR)(getProcAddr("glClear"))
	if gpClear == nil {
		return errors.New("glClear")
	}
	gpClearAccumxOES = (C.GPCLEARACCUMXOES)(getProcAddr("glClearAccumxOES"))
	gpClearBufferData = (C.GPCLEARBUFFERDATA)(getProcAddr("glClearBufferData"))
	gpClearBufferSubData = (C.GPCLEARBUFFERSUBDATA)(getProcAddr("glClearBufferSubData"))
	gpClearBufferfi = (C.GPCLEARBUFFERFI)(getProcAddr("glClearBufferfi"))
	if gpClearBufferfi == nil {
		return errors.New("glClearBufferfi")
	}
	gpClearBufferfv = (C.GPCLEARBUFFERFV)(getProcAddr("glClearBufferfv"))
	if gpClearBufferfv == nil {
		return errors.New("glClearBufferfv")
	}
	gpClearBufferiv = (C.GPCLEARBUFFERIV)(getProcAddr("glClearBufferiv"))
	if gpClearBufferiv == nil {
		return errors.New("glClearBufferiv")
	}
	gpClearBufferuiv = (C.GPCLEARBUFFERUIV)(getProcAddr("glClearBufferuiv"))
	if gpClearBufferuiv == nil {
		return errors.New("glClearBufferuiv")
	}
	gpClearColor = (C.GPCLEARCOLOR)(getProcAddr("glClearColor"))
	if gpClearColor == nil {
		return errors.New("glClearColor")
	}
	gpClearColorxOES = (C.GPCLEARCOLORXOES)(getProcAddr("glClearColorxOES"))
	gpClearDepth = (C.GPCLEARDEPTH)(getProcAddr("glClearDepth"))
	if gpClearDepth == nil {
		return errors.New("glClearDepth")
	}
	gpClearDepthf = (C.GPCLEARDEPTHF)(getProcAddr("glClearDepthf"))
	if gpClearDepthf == nil {
		return errors.New("glClearDepthf")
	}
	gpClearDepthfOES = (C.GPCLEARDEPTHFOES)(getProcAddr("glClearDepthfOES"))
	gpClearDepthxOES = (C.GPCLEARDEPTHXOES)(getProcAddr("glClearDepthxOES"))
	gpClearNamedBufferData = (C.GPCLEARNAMEDBUFFERDATA)(getProcAddr("glClearNamedBufferData"))
	gpClearNamedBufferSubData = (C.GPCLEARNAMEDBUFFERSUBDATA)(getProcAddr("glClearNamedBufferSubData"))
	gpClearNamedFramebufferfi = (C.GPCLEARNAMEDFRAMEBUFFERFI)(getProcAddr("glClearNamedFramebufferfi"))
	gpClearNamedFramebufferfv = (C.GPCLEARNAMEDFRAMEBUFFERFV)(getProcAddr("glClearNamedFramebufferfv"))
	gpClearNamedFramebufferiv = (C.GPCLEARNAMEDFRAMEBUFFERIV)(getProcAddr("glClearNamedFramebufferiv"))
	gpClearNamedFramebufferuiv = (C.GPCLEARNAMEDFRAMEBUFFERUIV)(getProcAddr("glClearNamedFramebufferuiv"))
	gpClearStencil = (C.GPCLEARSTENCIL)(getProcAddr("glClearStencil"))
	if gpClearStencil == nil {
		return errors.New("glClearStencil")
	}
	gpClearTexImage = (C.GPCLEARTEXIMAGE)(getProcAddr("glClearTexImage"))
	gpClearTexSubImage = (C.GPCLEARTEXSUBIMAGE)(getProcAddr("glClearTexSubImage"))
	gpClientWaitSync = (C.GPCLIENTWAITSYNC)(getProcAddr("glClientWaitSync"))
	if gpClientWaitSync == nil {
		return errors.New("glClientWaitSync")
	}
	gpClipControl = (C.GPCLIPCONTROL)(getProcAddr("glClipControl"))
	gpClipPlanefOES = (C.GPCLIPPLANEFOES)(getProcAddr("glClipPlanefOES"))
	gpClipPlanexOES = (C.GPCLIPPLANEXOES)(getProcAddr("glClipPlanexOES"))
	gpColor3xOES = (C.GPCOLOR3XOES)(getProcAddr("glColor3xOES"))
	gpColor3xvOES = (C.GPCOLOR3XVOES)(getProcAddr("glColor3xvOES"))
	gpColor4xOES = (C.GPCOLOR4XOES)(getProcAddr("glColor4xOES"))
	gpColor4xvOES = (C.GPCOLOR4XVOES)(getProcAddr("glColor4xvOES"))
	gpColorMask = (C.GPCOLORMASK)(getProcAddr("glColorMask"))
	if gpColorMask == nil {
		return errors.New("glColorMask")
	}
	gpColorMaski = (C.GPCOLORMASKI)(getProcAddr("glColorMaski"))
	if gpColorMaski == nil {
		return errors.New("glColorMaski")
	}
	gpCompileShader = (C.GPCOMPILESHADER)(getProcAddr("glCompileShader"))
	if gpCompileShader == nil {
		return errors.New("glCompileShader")
	}
	gpCompileShaderIncludeARB = (C.GPCOMPILESHADERINCLUDEARB)(getProcAddr("glCompileShaderIncludeARB"))
	gpCompressedTexImage1D = (C.GPCOMPRESSEDTEXIMAGE1D)(getProcAddr("glCompressedTexImage1D"))
	if gpCompressedTexImage1D == nil {
		return errors.New("glCompressedTexImage1D")
	}
	gpCompressedTexImage2D = (C.GPCOMPRESSEDTEXIMAGE2D)(getProcAddr("glCompressedTexImage2D"))
	if gpCompressedTexImage2D == nil {
		return errors.New("glCompressedTexImage2D")
	}
	gpCompressedTexImage3D = (C.GPCOMPRESSEDTEXIMAGE3D)(getProcAddr("glCompressedTexImage3D"))
	if gpCompressedTexImage3D == nil {
		return errors.New("glCompressedTexImage3D")
	}
	gpCompressedTexSubImage1D = (C.GPCOMPRESSEDTEXSUBIMAGE1D)(getProcAddr("glCompressedTexSubImage1D"))
	if gpCompressedTexSubImage1D == nil {
		return errors.New("glCompressedTexSubImage1D")
	}
	gpCompressedTexSubImage2D = (C.GPCOMPRESSEDTEXSUBIMAGE2D)(getProcAddr("glCompressedTexSubImage2D"))
	if gpCompressedTexSubImage2D == nil {
		return errors.New("glCompressedTexSubImage2D")
	}
	gpCompressedTexSubImage3D = (C.GPCOMPRESSEDTEXSUBIMAGE3D)(getProcAddr("glCompressedTexSubImage3D"))
	if gpCompressedTexSubImage3D == nil {
		return errors.New("glCompressedTexSubImage3D")
	}
	gpCompressedTextureSubImage1D = (C.GPCOMPRESSEDTEXTURESUBIMAGE1D)(getProcAddr("glCompressedTextureSubImage1D"))
	gpCompressedTextureSubImage2D = (C.GPCOMPRESSEDTEXTURESUBIMAGE2D)(getProcAddr("glCompressedTextureSubImage2D"))
	gpCompressedTextureSubImage3D = (C.GPCOMPRESSEDTEXTURESUBIMAGE3D)(getProcAddr("glCompressedTextureSubImage3D"))
	gpConvolutionParameterxOES = (C.GPCONVOLUTIONPARAMETERXOES)(getProcAddr("glConvolutionParameterxOES"))
	gpConvolutionParameterxvOES = (C.GPCONVOLUTIONPARAMETERXVOES)(getProcAddr("glConvolutionParameterxvOES"))
	gpCopyBufferSubData = (C.GPCOPYBUFFERSUBDATA)(getProcAddr("glCopyBufferSubData"))
	if gpCopyBufferSubData == nil {
		return errors.New("glCopyBufferSubData")
	}
	gpCopyImageSubData = (C.GPCOPYIMAGESUBDATA)(getProcAddr("glCopyImageSubData"))
	gpCopyNamedBufferSubData = (C.GPCOPYNAMEDBUFFERSUBDATA)(getProcAddr("glCopyNamedBufferSubData"))
	gpCopyTexImage1D = (C.GPCOPYTEXIMAGE1D)(getProcAddr("glCopyTexImage1D"))
	if gpCopyTexImage1D == nil {
		return errors.New("glCopyTexImage1D")
	}
	gpCopyTexImage2D = (C.GPCOPYTEXIMAGE2D)(getProcAddr("glCopyTexImage2D"))
	if gpCopyTexImage2D == nil {
		return errors.New("glCopyTexImage2D")
	}
	gpCopyTexSubImage1D = (C.GPCOPYTEXSUBIMAGE1D)(getProcAddr("glCopyTexSubImage1D"))
	if gpCopyTexSubImage1D == nil {
		return errors.New("glCopyTexSubImage1D")
	}
	gpCopyTexSubImage2D = (C.GPCOPYTEXSUBIMAGE2D)(getProcAddr("glCopyTexSubImage2D"))
	if gpCopyTexSubImage2D == nil {
		return errors.New("glCopyTexSubImage2D")
	}
	gpCopyTexSubImage3D = (C.GPCOPYTEXSUBIMAGE3D)(getProcAddr("glCopyTexSubImage3D"))
	if gpCopyTexSubImage3D == nil {
		return errors.New("glCopyTexSubImage3D")
	}
	gpCopyTextureSubImage1D = (C.GPCOPYTEXTURESUBIMAGE1D)(getProcAddr("glCopyTextureSubImage1D"))
	gpCopyTextureSubImage2D = (C.GPCOPYTEXTURESUBIMAGE2D)(getProcAddr("glCopyTextureSubImage2D"))
	gpCopyTextureSubImage3D = (C.GPCOPYTEXTURESUBIMAGE3D)(getProcAddr("glCopyTextureSubImage3D"))
	gpCreateBuffers = (C.GPCREATEBUFFERS)(getProcAddr("glCreateBuffers"))
	gpCreateFramebuffers = (C.GPCREATEFRAMEBUFFERS)(getProcAddr("glCreateFramebuffers"))
	gpCreatePerfQueryINTEL = (C.GPCREATEPERFQUERYINTEL)(getProcAddr("glCreatePerfQueryINTEL"))
	gpCreateProgram = (C.GPCREATEPROGRAM)(getProcAddr("glCreateProgram"))
	if gpCreateProgram == nil {
		return errors.New("glCreateProgram")
	}
	gpCreateProgramPipelines = (C.GPCREATEPROGRAMPIPELINES)(getProcAddr("glCreateProgramPipelines"))
	gpCreateQueries = (C.GPCREATEQUERIES)(getProcAddr("glCreateQueries"))
	gpCreateRenderbuffers = (C.GPCREATERENDERBUFFERS)(getProcAddr("glCreateRenderbuffers"))
	gpCreateSamplers = (C.GPCREATESAMPLERS)(getProcAddr("glCreateSamplers"))
	gpCreateShader = (C.GPCREATESHADER)(getProcAddr("glCreateShader"))
	if gpCreateShader == nil {
		return errors.New("glCreateShader")
	}
	gpCreateShaderProgramEXT = (C.GPCREATESHADERPROGRAMEXT)(getProcAddr("glCreateShaderProgramEXT"))
	gpCreateShaderProgramv = (C.GPCREATESHADERPROGRAMV)(getProcAddr("glCreateShaderProgramv"))
	if gpCreateShaderProgramv == nil {
		return errors.New("glCreateShaderProgramv")
	}
	gpCreateShaderProgramvEXT = (C.GPCREATESHADERPROGRAMVEXT)(getProcAddr("glCreateShaderProgramvEXT"))
	gpCreateSyncFromCLeventARB = (C.GPCREATESYNCFROMCLEVENTARB)(getProcAddr("glCreateSyncFromCLeventARB"))
	gpCreateTextures = (C.GPCREATETEXTURES)(getProcAddr("glCreateTextures"))
	gpCreateTransformFeedbacks = (C.GPCREATETRANSFORMFEEDBACKS)(getProcAddr("glCreateTransformFeedbacks"))
	gpCreateVertexArrays = (C.GPCREATEVERTEXARRAYS)(getProcAddr("glCreateVertexArrays"))
	gpCullFace = (C.GPCULLFACE)(getProcAddr("glCullFace"))
	if gpCullFace == nil {
		return errors.New("glCullFace")
	}
	gpDebugMessageCallback = (C.GPDEBUGMESSAGECALLBACK)(getProcAddr("glDebugMessageCallback"))
	gpDebugMessageCallbackARB = (C.GPDEBUGMESSAGECALLBACKARB)(getProcAddr("glDebugMessageCallbackARB"))
	gpDebugMessageCallbackKHR = (C.GPDEBUGMESSAGECALLBACKKHR)(getProcAddr("glDebugMessageCallbackKHR"))
	gpDebugMessageControl = (C.GPDEBUGMESSAGECONTROL)(getProcAddr("glDebugMessageControl"))
	gpDebugMessageControlARB = (C.GPDEBUGMESSAGECONTROLARB)(getProcAddr("glDebugMessageControlARB"))
	gpDebugMessageControlKHR = (C.GPDEBUGMESSAGECONTROLKHR)(getProcAddr("glDebugMessageControlKHR"))
	gpDebugMessageInsert = (C.GPDEBUGMESSAGEINSERT)(getProcAddr("glDebugMessageInsert"))
	gpDebugMessageInsertARB = (C.GPDEBUGMESSAGEINSERTARB)(getProcAddr("glDebugMessageInsertARB"))
	gpDebugMessageInsertKHR = (C.GPDEBUGMESSAGEINSERTKHR)(getProcAddr("glDebugMessageInsertKHR"))
	gpDeleteBuffers = (C.GPDELETEBUFFERS)(getProcAddr("glDeleteBuffers"))
	if gpDeleteBuffers == nil {
		return errors.New("glDeleteBuffers")
	}
	gpDeleteFencesNV = (C.GPDELETEFENCESNV)(getProcAddr("glDeleteFencesNV"))
	gpDeleteFramebuffers = (C.GPDELETEFRAMEBUFFERS)(getProcAddr("glDeleteFramebuffers"))
	if gpDeleteFramebuffers == nil {
		return errors.New("glDeleteFramebuffers")
	}
	gpDeleteNamedStringARB = (C.GPDELETENAMEDSTRINGARB)(getProcAddr("glDeleteNamedStringARB"))
	gpDeletePerfMonitorsAMD = (C.GPDELETEPERFMONITORSAMD)(getProcAddr("glDeletePerfMonitorsAMD"))
	gpDeletePerfQueryINTEL = (C.GPDELETEPERFQUERYINTEL)(getProcAddr("glDeletePerfQueryINTEL"))
	gpDeleteProgram = (C.GPDELETEPROGRAM)(getProcAddr("glDeleteProgram"))
	if gpDeleteProgram == nil {
		return errors.New("glDeleteProgram")
	}
	gpDeleteProgramPipelines = (C.GPDELETEPROGRAMPIPELINES)(getProcAddr("glDeleteProgramPipelines"))
	if gpDeleteProgramPipelines == nil {
		return errors.New("glDeleteProgramPipelines")
	}
	gpDeleteProgramPipelinesEXT = (C.GPDELETEPROGRAMPIPELINESEXT)(getProcAddr("glDeleteProgramPipelinesEXT"))
	gpDeleteQueries = (C.GPDELETEQUERIES)(getProcAddr("glDeleteQueries"))
	if gpDeleteQueries == nil {
		return errors.New("glDeleteQueries")
	}
	gpDeleteRenderbuffers = (C.GPDELETERENDERBUFFERS)(getProcAddr("glDeleteRenderbuffers"))
	if gpDeleteRenderbuffers == nil {
		return errors.New("glDeleteRenderbuffers")
	}
	gpDeleteSamplers = (C.GPDELETESAMPLERS)(getProcAddr("glDeleteSamplers"))
	if gpDeleteSamplers == nil {
		return errors.New("glDeleteSamplers")
	}
	gpDeleteShader = (C.GPDELETESHADER)(getProcAddr("glDeleteShader"))
	if gpDeleteShader == nil {
		return errors.New("glDeleteShader")
	}
	gpDeleteSync = (C.GPDELETESYNC)(getProcAddr("glDeleteSync"))
	if gpDeleteSync == nil {
		return errors.New("glDeleteSync")
	}
	gpDeleteTextures = (C.GPDELETETEXTURES)(getProcAddr("glDeleteTextures"))
	if gpDeleteTextures == nil {
		return errors.New("glDeleteTextures")
	}
	gpDeleteTransformFeedbacks = (C.GPDELETETRANSFORMFEEDBACKS)(getProcAddr("glDeleteTransformFeedbacks"))
	if gpDeleteTransformFeedbacks == nil {
		return errors.New("glDeleteTransformFeedbacks")
	}
	gpDeleteVertexArrays = (C.GPDELETEVERTEXARRAYS)(getProcAddr("glDeleteVertexArrays"))
	if gpDeleteVertexArrays == nil {
		return errors.New("glDeleteVertexArrays")
	}
	gpDepthFunc = (C.GPDEPTHFUNC)(getProcAddr("glDepthFunc"))
	if gpDepthFunc == nil {
		return errors.New("glDepthFunc")
	}
	gpDepthMask = (C.GPDEPTHMASK)(getProcAddr("glDepthMask"))
	if gpDepthMask == nil {
		return errors.New("glDepthMask")
	}
	gpDepthRange = (C.GPDEPTHRANGE)(getProcAddr("glDepthRange"))
	if gpDepthRange == nil {
		return errors.New("glDepthRange")
	}
	gpDepthRangeArrayv = (C.GPDEPTHRANGEARRAYV)(getProcAddr("glDepthRangeArrayv"))
	if gpDepthRangeArrayv == nil {
		return errors.New("glDepthRangeArrayv")
	}
	gpDepthRangeIndexed = (C.GPDEPTHRANGEINDEXED)(getProcAddr("glDepthRangeIndexed"))
	if gpDepthRangeIndexed == nil {
		return errors.New("glDepthRangeIndexed")
	}
	gpDepthRangef = (C.GPDEPTHRANGEF)(getProcAddr("glDepthRangef"))
	if gpDepthRangef == nil {
		return errors.New("glDepthRangef")
	}
	gpDepthRangefOES = (C.GPDEPTHRANGEFOES)(getProcAddr("glDepthRangefOES"))
	gpDepthRangexOES = (C.GPDEPTHRANGEXOES)(getProcAddr("glDepthRangexOES"))
	gpDetachShader = (C.GPDETACHSHADER)(getProcAddr("glDetachShader"))
	if gpDetachShader == nil {
		return errors.New("glDetachShader")
	}
	gpDisable = (C.GPDISABLE)(getProcAddr("glDisable"))
	if gpDisable == nil {
		return errors.New("glDisable")
	}
	gpDisableVertexArrayAttrib = (C.GPDISABLEVERTEXARRAYATTRIB)(getProcAddr("glDisableVertexArrayAttrib"))
	gpDisableVertexAttribArray = (C.GPDISABLEVERTEXATTRIBARRAY)(getProcAddr("glDisableVertexAttribArray"))
	if gpDisableVertexAttribArray == nil {
		return errors.New("glDisableVertexAttribArray")
	}
	gpDisablei = (C.GPDISABLEI)(getProcAddr("glDisablei"))
	if gpDisablei == nil {
		return errors.New("glDisablei")
	}
	gpDispatchCompute = (C.GPDISPATCHCOMPUTE)(getProcAddr("glDispatchCompute"))
	gpDispatchComputeGroupSizeARB = (C.GPDISPATCHCOMPUTEGROUPSIZEARB)(getProcAddr("glDispatchComputeGroupSizeARB"))
	gpDispatchComputeIndirect = (C.GPDISPATCHCOMPUTEINDIRECT)(getProcAddr("glDispatchComputeIndirect"))
	gpDrawArrays = (C.GPDRAWARRAYS)(getProcAddr("glDrawArrays"))
	if gpDrawArrays == nil {
		return errors.New("glDrawArrays")
	}
	gpDrawArraysIndirect = (C.GPDRAWARRAYSINDIRECT)(getProcAddr("glDrawArraysIndirect"))
	if gpDrawArraysIndirect == nil {
		return errors.New("glDrawArraysIndirect")
	}
	gpDrawArraysInstanced = (C.GPDRAWARRAYSINSTANCED)(getProcAddr("glDrawArraysInstanced"))
	if gpDrawArraysInstanced == nil {
		return errors.New("glDrawArraysInstanced")
	}
	gpDrawArraysInstancedBaseInstance = (C.GPDRAWARRAYSINSTANCEDBASEINSTANCE)(getProcAddr("glDrawArraysInstancedBaseInstance"))
	gpDrawArraysInstancedEXT = (C.GPDRAWARRAYSINSTANCEDEXT)(getProcAddr("glDrawArraysInstancedEXT"))
	gpDrawBuffer = (C.GPDRAWBUFFER)(getProcAddr("glDrawBuffer"))
	if gpDrawBuffer == nil {
		return errors.New("glDrawBuffer")
	}
	gpDrawBuffers = (C.GPDRAWBUFFERS)(getProcAddr("glDrawBuffers"))
	if gpDrawBuffers == nil {
		return errors.New("glDrawBuffers")
	}
	gpDrawElements = (C.GPDRAWELEMENTS)(getProcAddr("glDrawElements"))
	if gpDrawElements == nil {
		return errors.New("glDrawElements")
	}
	gpDrawElementsBaseVertex = (C.GPDRAWELEMENTSBASEVERTEX)(getProcAddr("glDrawElementsBaseVertex"))
	if gpDrawElementsBaseVertex == nil {
		return errors.New("glDrawElementsBaseVertex")
	}
	gpDrawElementsIndirect = (C.GPDRAWELEMENTSINDIRECT)(getProcAddr("glDrawElementsIndirect"))
	if gpDrawElementsIndirect == nil {
		return errors.New("glDrawElementsIndirect")
	}
	gpDrawElementsInstanced = (C.GPDRAWELEMENTSINSTANCED)(getProcAddr("glDrawElementsInstanced"))
	if gpDrawElementsInstanced == nil {
		return errors.New("glDrawElementsInstanced")
	}
	gpDrawElementsInstancedBaseInstance = (C.GPDRAWELEMENTSINSTANCEDBASEINSTANCE)(getProcAddr("glDrawElementsInstancedBaseInstance"))
	gpDrawElementsInstancedBaseVertex = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEX)(getProcAddr("glDrawElementsInstancedBaseVertex"))
	if gpDrawElementsInstancedBaseVertex == nil {
		return errors.New("glDrawElementsInstancedBaseVertex")
	}
	gpDrawElementsInstancedBaseVertexBaseInstance = (C.GPDRAWELEMENTSINSTANCEDBASEVERTEXBASEINSTANCE)(getProcAddr("glDrawElementsInstancedBaseVertexBaseInstance"))
	gpDrawElementsInstancedEXT = (C.GPDRAWELEMENTSINSTANCEDEXT)(getProcAddr("glDrawElementsInstancedEXT"))
	gpDrawRangeElements = (C.GPDRAWRANGEELEMENTS)(getProcAddr("glDrawRangeElements"))
	if gpDrawRangeElements == nil {
		return errors.New("glDrawRangeElements")
	}
	gpDrawRangeElementsBaseVertex = (C.GPDRAWRANGEELEMENTSBASEVERTEX)(getProcAddr("glDrawRangeElementsBaseVertex"))
	if gpDrawRangeElementsBaseVertex == nil {
		return errors.New("glDrawRangeElementsBaseVertex")
	}
	gpDrawTransformFeedback = (C.GPDRAWTRANSFORMFEEDBACK)(getProcAddr("glDrawTransformFeedback"))
	if gpDrawTransformFeedback == nil {
		return errors.New("glDrawTransformFeedback")
	}
	gpDrawTransformFeedbackInstanced = (C.GPDRAWTRANSFORMFEEDBACKINSTANCED)(getProcAddr("glDrawTransformFeedbackInstanced"))
	gpDrawTransformFeedbackStream = (C.GPDRAWTRANSFORMFEEDBACKSTREAM)(getProcAddr("glDrawTransformFeedbackStream"))
	if gpDrawTransformFeedbackStream == nil {
		return errors.New("glDrawTransformFeedbackStream")
	}
	gpDrawTransformFeedbackStreamInstanced = (C.GPDRAWTRANSFORMFEEDBACKSTREAMINSTANCED)(getProcAddr("glDrawTransformFeedbackStreamInstanced"))
	gpEnable = (C.GPENABLE)(getProcAddr("glEnable"))
	if gpEnable == nil {
		return errors.New("glEnable")
	}
	gpEnableVertexArrayAttrib = (C.GPENABLEVERTEXARRAYATTRIB)(getProcAddr("glEnableVertexArrayAttrib"))
	gpEnableVertexAttribArray = (C.GPENABLEVERTEXATTRIBARRAY)(getProcAddr("glEnableVertexAttribArray"))
	if gpEnableVertexAttribArray == nil {
		return errors.New("glEnableVertexAttribArray")
	}
	gpEnablei = (C.GPENABLEI)(getProcAddr("glEnablei"))
	if gpEnablei == nil {
		return errors.New("glEnablei")
	}
	gpEndConditionalRender = (C.GPENDCONDITIONALRENDER)(getProcAddr("glEndConditionalRender"))
	if gpEndConditionalRender == nil {
		return errors.New("glEndConditionalRender")
	}
	gpEndPerfMonitorAMD = (C.GPENDPERFMONITORAMD)(getProcAddr("glEndPerfMonitorAMD"))
	gpEndPerfQueryINTEL = (C.GPENDPERFQUERYINTEL)(getProcAddr("glEndPerfQueryINTEL"))
	gpEndQuery = (C.GPENDQUERY)(getProcAddr("glEndQuery"))
	if gpEndQuery == nil {
		return errors.New("glEndQuery")
	}
	gpEndQueryIndexed = (C.GPENDQUERYINDEXED)(getProcAddr("glEndQueryIndexed"))
	if gpEndQueryIndexed == nil {
		return errors.New("glEndQueryIndexed")
	}
	gpEndTransformFeedback = (C.GPENDTRANSFORMFEEDBACK)(getProcAddr("glEndTransformFeedback"))
	if gpEndTransformFeedback == nil {
		return errors.New("glEndTransformFeedback")
	}
	gpEvalCoord1xOES = (C.GPEVALCOORD1XOES)(getProcAddr("glEvalCoord1xOES"))
	gpEvalCoord1xvOES = (C.GPEVALCOORD1XVOES)(getProcAddr("glEvalCoord1xvOES"))
	gpEvalCoord2xOES = (C.GPEVALCOORD2XOES)(getProcAddr("glEvalCoord2xOES"))
	gpEvalCoord2xvOES = (C.GPEVALCOORD2XVOES)(getProcAddr("glEvalCoord2xvOES"))
	gpFeedbackBufferxOES = (C.GPFEEDBACKBUFFERXOES)(getProcAddr("glFeedbackBufferxOES"))
	gpFenceSync = (C.GPFENCESYNC)(getProcAddr("glFenceSync"))
	if gpFenceSync == nil {
		return errors.New("glFenceSync")
	}
	gpFinish = (C.GPFINISH)(getProcAddr("glFinish"))
	if gpFinish == nil {
		return errors.New("glFinish")
	}
	gpFinishFenceNV = (C.GPFINISHFENCENV)(getProcAddr("glFinishFenceNV"))
	gpFlush = (C.GPFLUSH)(getProcAddr("glFlush"))
	if gpFlush == nil {
		return errors.New("glFlush")
	}
	gpFlushMappedBufferRange = (C.GPFLUSHMAPPEDBUFFERRANGE)(getProcAddr("glFlushMappedBufferRange"))
	if gpFlushMappedBufferRange == nil {
		return errors.New("glFlushMappedBufferRange")
	}
	gpFlushMappedNamedBufferRange = (C.GPFLUSHMAPPEDNAMEDBUFFERRANGE)(getProcAddr("glFlushMappedNamedBufferRange"))
	gpFogxOES = (C.GPFOGXOES)(getProcAddr("glFogxOES"))
	gpFogxvOES = (C.GPFOGXVOES)(getProcAddr("glFogxvOES"))
	gpFramebufferParameteri = (C.GPFRAMEBUFFERPARAMETERI)(getProcAddr("glFramebufferParameteri"))
	gpFramebufferRenderbuffer = (C.GPFRAMEBUFFERRENDERBUFFER)(getProcAddr("glFramebufferRenderbuffer"))
	if gpFramebufferRenderbuffer == nil {
		return errors.New("glFramebufferRenderbuffer")
	}
	gpFramebufferTexture = (C.GPFRAMEBUFFERTEXTURE)(getProcAddr("glFramebufferTexture"))
	if gpFramebufferTexture == nil {
		return errors.New("glFramebufferTexture")
	}
	gpFramebufferTexture1D = (C.GPFRAMEBUFFERTEXTURE1D)(getProcAddr("glFramebufferTexture1D"))
	if gpFramebufferTexture1D == nil {
		return errors.New("glFramebufferTexture1D")
	}
	gpFramebufferTexture2D = (C.GPFRAMEBUFFERTEXTURE2D)(getProcAddr("glFramebufferTexture2D"))
	if gpFramebufferTexture2D == nil {
		return errors.New("glFramebufferTexture2D")
	}
	gpFramebufferTexture3D = (C.GPFRAMEBUFFERTEXTURE3D)(getProcAddr("glFramebufferTexture3D"))
	if gpFramebufferTexture3D == nil {
		return errors.New("glFramebufferTexture3D")
	}
	gpFramebufferTextureLayer = (C.GPFRAMEBUFFERTEXTURELAYER)(getProcAddr("glFramebufferTextureLayer"))
	if gpFramebufferTextureLayer == nil {
		return errors.New("glFramebufferTextureLayer")
	}
	gpFrontFace = (C.GPFRONTFACE)(getProcAddr("glFrontFace"))
	if gpFrontFace == nil {
		return errors.New("glFrontFace")
	}
	gpFrustumfOES = (C.GPFRUSTUMFOES)(getProcAddr("glFrustumfOES"))
	gpFrustumxOES = (C.GPFRUSTUMXOES)(getProcAddr("glFrustumxOES"))
	gpGenBuffers = (C.GPGENBUFFERS)(getProcAddr("glGenBuffers"))
	if gpGenBuffers == nil {
		return errors.New("glGenBuffers")
	}
	gpGenFencesNV = (C.GPGENFENCESNV)(getProcAddr("glGenFencesNV"))
	gpGenFramebuffers = (C.GPGENFRAMEBUFFERS)(getProcAddr("glGenFramebuffers"))
	if gpGenFramebuffers == nil {
		return errors.New("glGenFramebuffers")
	}
	gpGenPerfMonitorsAMD = (C.GPGENPERFMONITORSAMD)(getProcAddr("glGenPerfMonitorsAMD"))
	gpGenProgramPipelines = (C.GPGENPROGRAMPIPELINES)(getProcAddr("glGenProgramPipelines"))
	if gpGenProgramPipelines == nil {
		return errors.New("glGenProgramPipelines")
	}
	gpGenProgramPipelinesEXT = (C.GPGENPROGRAMPIPELINESEXT)(getProcAddr("glGenProgramPipelinesEXT"))
	gpGenQueries = (C.GPGENQUERIES)(getProcAddr("glGenQueries"))
	if gpGenQueries == nil {
		return errors.New("glGenQueries")
	}
	gpGenRenderbuffers = (C.GPGENRENDERBUFFERS)(getProcAddr("glGenRenderbuffers"))
	if gpGenRenderbuffers == nil {
		return errors.New("glGenRenderbuffers")
	}
	gpGenSamplers = (C.GPGENSAMPLERS)(getProcAddr("glGenSamplers"))
	if gpGenSamplers == nil {
		return errors.New("glGenSamplers")
	}
	gpGenTextures = (C.GPGENTEXTURES)(getProcAddr("glGenTextures"))
	if gpGenTextures == nil {
		return errors.New("glGenTextures")
	}
	gpGenTransformFeedbacks = (C.GPGENTRANSFORMFEEDBACKS)(getProcAddr("glGenTransformFeedbacks"))
	if gpGenTransformFeedbacks == nil {
		return errors.New("glGenTransformFeedbacks")
	}
	gpGenVertexArrays = (C.GPGENVERTEXARRAYS)(getProcAddr("glGenVertexArrays"))
	if gpGenVertexArrays == nil {
		return errors.New("glGenVertexArrays")
	}
	gpGenerateMipmap = (C.GPGENERATEMIPMAP)(getProcAddr("glGenerateMipmap"))
	if gpGenerateMipmap == nil {
		return errors.New("glGenerateMipmap")
	}
	gpGenerateTextureMipmap = (C.GPGENERATETEXTUREMIPMAP)(getProcAddr("glGenerateTextureMipmap"))
	gpGetActiveAtomicCounterBufferiv = (C.GPGETACTIVEATOMICCOUNTERBUFFERIV)(getProcAddr("glGetActiveAtomicCounterBufferiv"))
	gpGetActiveAttrib = (C.GPGETACTIVEATTRIB)(getProcAddr("glGetActiveAttrib"))
	if gpGetActiveAttrib == nil {
		return errors.New("glGetActiveAttrib")
	}
	gpGetActiveSubroutineName = (C.GPGETACTIVESUBROUTINENAME)(getProcAddr("glGetActiveSubroutineName"))
	if gpGetActiveSubroutineName == nil {
		return errors.New("glGetActiveSubroutineName")
	}
	gpGetActiveSubroutineUniformName = (C.GPGETACTIVESUBROUTINEUNIFORMNAME)(getProcAddr("glGetActiveSubroutineUniformName"))
	if gpGetActiveSubroutineUniformName == nil {
		return errors.New("glGetActiveSubroutineUniformName")
	}
	gpGetActiveSubroutineUniformiv = (C.GPGETACTIVESUBROUTINEUNIFORMIV)(getProcAddr("glGetActiveSubroutineUniformiv"))
	if gpGetActiveSubroutineUniformiv == nil {
		return errors.New("glGetActiveSubroutineUniformiv")
	}
	gpGetActiveUniform = (C.GPGETACTIVEUNIFORM)(getProcAddr("glGetActiveUniform"))
	if gpGetActiveUniform == nil {
		return errors.New("glGetActiveUniform")
	}
	gpGetActiveUniformBlockName = (C.GPGETACTIVEUNIFORMBLOCKNAME)(getProcAddr("glGetActiveUniformBlockName"))
	if gpGetActiveUniformBlockName == nil {
		return errors.New("glGetActiveUniformBlockName")
	}
	gpGetActiveUniformBlockiv = (C.GPGETACTIVEUNIFORMBLOCKIV)(getProcAddr("glGetActiveUniformBlockiv"))
	if gpGetActiveUniformBlockiv == nil {
		return errors.New("glGetActiveUniformBlockiv")
	}
	gpGetActiveUniformName = (C.GPGETACTIVEUNIFORMNAME)(getProcAddr("glGetActiveUniformName"))
	if gpGetActiveUniformName == nil {
		return errors.New("glGetActiveUniformName")
	}
	gpGetActiveUniformsiv = (C.GPGETACTIVEUNIFORMSIV)(getProcAddr("glGetActiveUniformsiv"))
	if gpGetActiveUniformsiv == nil {
		return errors.New("glGetActiveUniformsiv")
	}
	gpGetAttachedShaders = (C.GPGETATTACHEDSHADERS)(getProcAddr("glGetAttachedShaders"))
	if gpGetAttachedShaders == nil {
		return errors.New("glGetAttachedShaders")
	}
	gpGetAttribLocation = (C.GPGETATTRIBLOCATION)(getProcAddr("glGetAttribLocation"))
	if gpGetAttribLocation == nil {
		return errors.New("glGetAttribLocation")
	}
	gpGetBooleani_v = (C.GPGETBOOLEANI_V)(getProcAddr("glGetBooleani_v"))
	if gpGetBooleani_v == nil {
		return errors.New("glGetBooleani_v")
	}
	gpGetBooleanv = (C.GPGETBOOLEANV)(getProcAddr("glGetBooleanv"))
	if gpGetBooleanv == nil {
		return errors.New("glGetBooleanv")
	}
	gpGetBufferParameteri64v = (C.GPGETBUFFERPARAMETERI64V)(getProcAddr("glGetBufferParameteri64v"))
	if gpGetBufferParameteri64v == nil {
		return errors.New("glGetBufferParameteri64v")
	}
	gpGetBufferParameteriv = (C.GPGETBUFFERPARAMETERIV)(getProcAddr("glGetBufferParameteriv"))
	if gpGetBufferParameteriv == nil {
		return errors.New("glGetBufferParameteriv")
	}
	gpGetBufferPointerv = (C.GPGETBUFFERPOINTERV)(getProcAddr("glGetBufferPointerv"))
	if gpGetBufferPointerv == nil {
		return errors.New("glGetBufferPointerv")
	}
	gpGetBufferSubData = (C.GPGETBUFFERSUBDATA)(getProcAddr("glGetBufferSubData"))
	if gpGetBufferSubData == nil {
		return errors.New("glGetBufferSubData")
	}
	gpGetClipPlanefOES = (C.GPGETCLIPPLANEFOES)(getProcAddr("glGetClipPlanefOES"))
	gpGetClipPlanexOES = (C.GPGETCLIPPLANEXOES)(getProcAddr("glGetClipPlanexOES"))
	gpGetCompressedTexImage = (C.GPGETCOMPRESSEDTEXIMAGE)(getProcAddr("glGetCompressedTexImage"))
	if gpGetCompressedTexImage == nil {
		return errors.New("glGetCompressedTexImage")
	}
	gpGetCompressedTextureImage = (C.GPGETCOMPRESSEDTEXTUREIMAGE)(getProcAddr("glGetCompressedTextureImage"))
	gpGetCompressedTextureSubImage = (C.GPGETCOMPRESSEDTEXTURESUBIMAGE)(getProcAddr("glGetCompressedTextureSubImage"))
	gpGetConvolutionParameterxvOES = (C.GPGETCONVOLUTIONPARAMETERXVOES)(getProcAddr("glGetConvolutionParameterxvOES"))
	gpGetDebugMessageLog = (C.GPGETDEBUGMESSAGELOG)(getProcAddr("glGetDebugMessageLog"))
	gpGetDebugMessageLogARB = (C.GPGETDEBUGMESSAGELOGARB)(getProcAddr("glGetDebugMessageLogARB"))
	gpGetDebugMessageLogKHR = (C.GPGETDEBUGMESSAGELOGKHR)(getProcAddr("glGetDebugMessageLogKHR"))
	gpGetDoublei_v = (C.GPGETDOUBLEI_V)(getProcAddr("glGetDoublei_v"))
	if gpGetDoublei_v == nil {
		return errors.New("glGetDoublei_v")
	}
	gpGetDoublev = (C.GPGETDOUBLEV)(getProcAddr("glGetDoublev"))
	if gpGetDoublev == nil {
		return errors.New("glGetDoublev")
	}
	gpGetError = (C.GPGETERROR)(getProcAddr("glGetError"))
	if gpGetError == nil {
		return errors.New("glGetError")
	}
	gpGetFenceivNV = (C.GPGETFENCEIVNV)(getProcAddr("glGetFenceivNV"))
	gpGetFirstPerfQueryIdINTEL = (C.GPGETFIRSTPERFQUERYIDINTEL)(getProcAddr("glGetFirstPerfQueryIdINTEL"))
	gpGetFixedvOES = (C.GPGETFIXEDVOES)(getProcAddr("glGetFixedvOES"))
	gpGetFloati_v = (C.GPGETFLOATI_V)(getProcAddr("glGetFloati_v"))
	if gpGetFloati_v == nil {
		return errors.New("glGetFloati_v")
	}
	gpGetFloatv = (C.GPGETFLOATV)(getProcAddr("glGetFloatv"))
	if gpGetFloatv == nil {
		return errors.New("glGetFloatv")
	}
	gpGetFragDataIndex = (C.GPGETFRAGDATAINDEX)(getProcAddr("glGetFragDataIndex"))
	if gpGetFragDataIndex == nil {
		return errors.New("glGetFragDataIndex")
	}
	gpGetFragDataLocation = (C.GPGETFRAGDATALOCATION)(getProcAddr("glGetFragDataLocation"))
	if gpGetFragDataLocation == nil {
		return errors.New("glGetFragDataLocation")
	}
	gpGetFramebufferAttachmentParameteriv = (C.GPGETFRAMEBUFFERATTACHMENTPARAMETERIV)(getProcAddr("glGetFramebufferAttachmentParameteriv"))
	if gpGetFramebufferAttachmentParameteriv == nil {
		return errors.New("glGetFramebufferAttachmentParameteriv")
	}
	gpGetFramebufferParameteriv = (C.GPGETFRAMEBUFFERPARAMETERIV)(getProcAddr("glGetFramebufferParameteriv"))
	gpGetGraphicsResetStatus = (C.GPGETGRAPHICSRESETSTATUS)(getProcAddr("glGetGraphicsResetStatus"))
	gpGetGraphicsResetStatusARB = (C.GPGETGRAPHICSRESETSTATUSARB)(getProcAddr("glGetGraphicsResetStatusARB"))
	gpGetGraphicsResetStatusKHR = (C.GPGETGRAPHICSRESETSTATUSKHR)(getProcAddr("glGetGraphicsResetStatusKHR"))
	gpGetHistogramParameterxvOES = (C.GPGETHISTOGRAMPARAMETERXVOES)(getProcAddr("glGetHistogramParameterxvOES"))
	gpGetImageHandleARB = (C.GPGETIMAGEHANDLEARB)(getProcAddr("glGetImageHandleARB"))
	gpGetInteger64i_v = (C.GPGETINTEGER64I_V)(getProcAddr("glGetInteger64i_v"))
	if gpGetInteger64i_v == nil {
		return errors.New("glGetInteger64i_v")
	}
	gpGetInteger64v = (C.GPGETINTEGER64V)(getProcAddr("glGetInteger64v"))
	if gpGetInteger64v == nil {
		return errors.New("glGetInteger64v")
	}
	gpGetIntegeri_v = (C.GPGETINTEGERI_V)(getProcAddr("glGetIntegeri_v"))
	if gpGetIntegeri_v == nil {
		return errors.New("glGetIntegeri_v")
	}
	gpGetIntegerv = (C.GPGETINTEGERV)(getProcAddr("glGetIntegerv"))
	if gpGetIntegerv == nil {
		return errors.New("glGetIntegerv")
	}
	gpGetInternalformati64v = (C.GPGETINTERNALFORMATI64V)(getProcAddr("glGetInternalformati64v"))
	gpGetInternalformativ = (C.GPGETINTERNALFORMATIV)(getProcAddr("glGetInternalformativ"))
	gpGetLightxOES = (C.GPGETLIGHTXOES)(getProcAddr("glGetLightxOES"))
	gpGetLightxvOES = (C.GPGETLIGHTXVOES)(getProcAddr("glGetLightxvOES"))
	gpGetMapxvOES = (C.GPGETMAPXVOES)(getProcAddr("glGetMapxvOES"))
	gpGetMaterialxOES = (C.GPGETMATERIALXOES)(getProcAddr("glGetMaterialxOES"))
	gpGetMaterialxvOES = (C.GPGETMATERIALXVOES)(getProcAddr("glGetMaterialxvOES"))
	gpGetMultisamplefv = (C.GPGETMULTISAMPLEFV)(getProcAddr("glGetMultisamplefv"))
	if gpGetMultisamplefv == nil {
		return errors.New("glGetMultisamplefv")
	}
	gpGetNamedBufferParameteri64v = (C.GPGETNAMEDBUFFERPARAMETERI64V)(getProcAddr("glGetNamedBufferParameteri64v"))
	gpGetNamedBufferParameteriv = (C.GPGETNAMEDBUFFERPARAMETERIV)(getProcAddr("glGetNamedBufferParameteriv"))
	gpGetNamedBufferPointerv = (C.GPGETNAMEDBUFFERPOINTERV)(getProcAddr("glGetNamedBufferPointerv"))
	gpGetNamedBufferSubData = (C.GPGETNAMEDBUFFERSUBDATA)(getProcAddr("glGetNamedBufferSubData"))
	gpGetNamedFramebufferAttachmentParameteriv = (C.GPGETNAMEDFRAMEBUFFERATTACHMENTPARAMETERIV)(getProcAddr("glGetNamedFramebufferAttachmentParameteriv"))
	gpGetNamedFramebufferParameteriv = (C.GPGETNAMEDFRAMEBUFFERPARAMETERIV)(getProcAddr("glGetNamedFramebufferParameteriv"))
	gpGetNamedRenderbufferParameteriv = (C.GPGETNAMEDRENDERBUFFERPARAMETERIV)(getProcAddr("glGetNamedRenderbufferParameteriv"))
	gpGetNamedStringARB = (C.GPGETNAMEDSTRINGARB)(getProcAddr("glGetNamedStringARB"))
	gpGetNamedStringivARB = (C.GPGETNAMEDSTRINGIVARB)(getProcAddr("glGetNamedStringivARB"))
	gpGetNextPerfQueryIdINTEL = (C.GPGETNEXTPERFQUERYIDINTEL)(getProcAddr("glGetNextPerfQueryIdINTEL"))
	gpGetObjectLabel = (C.GPGETOBJECTLABEL)(getProcAddr("glGetObjectLabel"))
	gpGetObjectLabelEXT = (C.GPGETOBJECTLABELEXT)(getProcAddr("glGetObjectLabelEXT"))
	gpGetObjectLabelKHR = (C.GPGETOBJECTLABELKHR)(getProcAddr("glGetObjectLabelKHR"))
	gpGetObjectPtrLabel = (C.GPGETOBJECTPTRLABEL)(getProcAddr("glGetObjectPtrLabel"))
	gpGetObjectPtrLabelKHR = (C.GPGETOBJECTPTRLABELKHR)(getProcAddr("glGetObjectPtrLabelKHR"))
	gpGetPerfCounterInfoINTEL = (C.GPGETPERFCOUNTERINFOINTEL)(getProcAddr("glGetPerfCounterInfoINTEL"))
	gpGetPerfMonitorCounterDataAMD = (C.GPGETPERFMONITORCOUNTERDATAAMD)(getProcAddr("glGetPerfMonitorCounterDataAMD"))
	gpGetPerfMonitorCounterInfoAMD = (C.GPGETPERFMONITORCOUNTERINFOAMD)(getProcAddr("glGetPerfMonitorCounterInfoAMD"))
	gpGetPerfMonitorCounterStringAMD = (C.GPGETPERFMONITORCOUNTERSTRINGAMD)(getProcAddr("glGetPerfMonitorCounterStringAMD"))
	gpGetPerfMonitorCountersAMD = (C.GPGETPERFMONITORCOUNTERSAMD)(getProcAddr("glGetPerfMonitorCountersAMD"))
	gpGetPerfMonitorGroupStringAMD = (C.GPGETPERFMONITORGROUPSTRINGAMD)(getProcAddr("glGetPerfMonitorGroupStringAMD"))
	gpGetPerfMonitorGroupsAMD = (C.GPGETPERFMONITORGROUPSAMD)(getProcAddr("glGetPerfMonitorGroupsAMD"))
	gpGetPerfQueryDataINTEL = (C.GPGETPERFQUERYDATAINTEL)(getProcAddr("glGetPerfQueryDataINTEL"))
	gpGetPerfQueryIdByNameINTEL = (C.GPGETPERFQUERYIDBYNAMEINTEL)(getProcAddr("glGetPerfQueryIdByNameINTEL"))
	gpGetPerfQueryInfoINTEL = (C.GPGETPERFQUERYINFOINTEL)(getProcAddr("glGetPerfQueryInfoINTEL"))
	gpGetPixelMapxv = (C.GPGETPIXELMAPXV)(getProcAddr("glGetPixelMapxv"))
	gpGetPointerv = (C.GPGETPOINTERV)(getProcAddr("glGetPointerv"))
	gpGetPointervKHR = (C.GPGETPOINTERVKHR)(getProcAddr("glGetPointervKHR"))
	gpGetProgramBinary = (C.GPGETPROGRAMBINARY)(getProcAddr("glGetProgramBinary"))
	if gpGetProgramBinary == nil {
		return errors.New("glGetProgramBinary")
	}
	gpGetProgramInfoLog = (C.GPGETPROGRAMINFOLOG)(getProcAddr("glGetProgramInfoLog"))
	if gpGetProgramInfoLog == nil {
		return errors.New("glGetProgramInfoLog")
	}
	gpGetProgramInterfaceiv = (C.GPGETPROGRAMINTERFACEIV)(getProcAddr("glGetProgramInterfaceiv"))
	gpGetProgramPipelineInfoLog = (C.GPGETPROGRAMPIPELINEINFOLOG)(getProcAddr("glGetProgramPipelineInfoLog"))
	if gpGetProgramPipelineInfoLog == nil {
		return errors.New("glGetProgramPipelineInfoLog")
	}
	gpGetProgramPipelineInfoLogEXT = (C.GPGETPROGRAMPIPELINEINFOLOGEXT)(getProcAddr("glGetProgramPipelineInfoLogEXT"))
	gpGetProgramPipelineiv = (C.GPGETPROGRAMPIPELINEIV)(getProcAddr("glGetProgramPipelineiv"))
	if gpGetProgramPipelineiv == nil {
		return errors.New("glGetProgramPipelineiv")
	}
	gpGetProgramPipelineivEXT = (C.GPGETPROGRAMPIPELINEIVEXT)(getProcAddr("glGetProgramPipelineivEXT"))
	gpGetProgramResourceIndex = (C.GPGETPROGRAMRESOURCEINDEX)(getProcAddr("glGetProgramResourceIndex"))
	gpGetProgramResourceLocation = (C.GPGETPROGRAMRESOURCELOCATION)(getProcAddr("glGetProgramResourceLocation"))
	gpGetProgramResourceLocationIndex = (C.GPGETPROGRAMRESOURCELOCATIONINDEX)(getProcAddr("glGetProgramResourceLocationIndex"))
	gpGetProgramResourceName = (C.GPGETPROGRAMRESOURCENAME)(getProcAddr("glGetProgramResourceName"))
	gpGetProgramResourceiv = (C.GPGETPROGRAMRESOURCEIV)(getProcAddr("glGetProgramResourceiv"))
	gpGetProgramStageiv = (C.GPGETPROGRAMSTAGEIV)(getProcAddr("glGetProgramStageiv"))
	if gpGetProgramStageiv == nil {
		return errors.New("glGetProgramStageiv")
	}
	gpGetProgramiv = (C.GPGETPROGRAMIV)(getProcAddr("glGetProgramiv"))
	if gpGetProgramiv == nil {
		return errors.New("glGetProgramiv")
	}
	gpGetQueryIndexediv = (C.GPGETQUERYINDEXEDIV)(getProcAddr("glGetQueryIndexediv"))
	if gpGetQueryIndexediv == nil {
		return errors.New("glGetQueryIndexediv")
	}
	gpGetQueryObjecti64v = (C.GPGETQUERYOBJECTI64V)(getProcAddr("glGetQueryObjecti64v"))
	if gpGetQueryObjecti64v == nil {
		return errors.New("glGetQueryObjecti64v")
	}
	gpGetQueryObjectiv = (C.GPGETQUERYOBJECTIV)(getProcAddr("glGetQueryObjectiv"))
	if gpGetQueryObjectiv == nil {
		return errors.New("glGetQueryObjectiv")
	}
	gpGetQueryObjectui64v = (C.GPGETQUERYOBJECTUI64V)(getProcAddr("glGetQueryObjectui64v"))
	if gpGetQueryObjectui64v == nil {
		return errors.New("glGetQueryObjectui64v")
	}
	gpGetQueryObjectuiv = (C.GPGETQUERYOBJECTUIV)(getProcAddr("glGetQueryObjectuiv"))
	if gpGetQueryObjectuiv == nil {
		return errors.New("glGetQueryObjectuiv")
	}
	gpGetQueryiv = (C.GPGETQUERYIV)(getProcAddr("glGetQueryiv"))
	if gpGetQueryiv == nil {
		return errors.New("glGetQueryiv")
	}
	gpGetRenderbufferParameteriv = (C.GPGETRENDERBUFFERPARAMETERIV)(getProcAddr("glGetRenderbufferParameteriv"))
	if gpGetRenderbufferParameteriv == nil {
		return errors.New("glGetRenderbufferParameteriv")
	}
	gpGetSamplerParameterIiv = (C.GPGETSAMPLERPARAMETERIIV)(getProcAddr("glGetSamplerParameterIiv"))
	if gpGetSamplerParameterIiv == nil {
		return errors.New("glGetSamplerParameterIiv")
	}
	gpGetSamplerParameterIuiv = (C.GPGETSAMPLERPARAMETERIUIV)(getProcAddr("glGetSamplerParameterIuiv"))
	if gpGetSamplerParameterIuiv == nil {
		return errors.New("glGetSamplerParameterIuiv")
	}
	gpGetSamplerParameterfv = (C.GPGETSAMPLERPARAMETERFV)(getProcAddr("glGetSamplerParameterfv"))
	if gpGetSamplerParameterfv == nil {
		return errors.New("glGetSamplerParameterfv")
	}
	gpGetSamplerParameteriv = (C.GPGETSAMPLERPARAMETERIV)(getProcAddr("glGetSamplerParameteriv"))
	if gpGetSamplerParameteriv == nil {
		return errors.New("glGetSamplerParameteriv")
	}
	gpGetShaderInfoLog = (C.GPGETSHADERINFOLOG)(getProcAddr("glGetShaderInfoLog"))
	if gpGetShaderInfoLog == nil {
		return errors.New("glGetShaderInfoLog")
	}
	gpGetShaderPrecisionFormat = (C.GPGETSHADERPRECISIONFORMAT)(getProcAddr("glGetShaderPrecisionFormat"))
	if gpGetShaderPrecisionFormat == nil {
		return errors.New("glGetShaderPrecisionFormat")
	}
	gpGetShaderSource = (C.GPGETSHADERSOURCE)(getProcAddr("glGetShaderSource"))
	if gpGetShaderSource == nil {
		return errors.New("glGetShaderSource")
	}
	gpGetShaderiv = (C.GPGETSHADERIV)(getProcAddr("glGetShaderiv"))
	if gpGetShaderiv == nil {
		return errors.New("glGetShaderiv")
	}
	gpGetString = (C.GPGETSTRING)(getProcAddr("glGetString"))
	if gpGetString == nil {
		return errors.New("glGetString")
	}
	gpGetStringi = (C.GPGETSTRINGI)(getProcAddr("glGetStringi"))
	if gpGetStringi == nil {
		return errors.New("glGetStringi")
	}
	gpGetSubroutineIndex = (C.GPGETSUBROUTINEINDEX)(getProcAddr("glGetSubroutineIndex"))
	if gpGetSubroutineIndex == nil {
		return errors.New("glGetSubroutineIndex")
	}
	gpGetSubroutineUniformLocation = (C.GPGETSUBROUTINEUNIFORMLOCATION)(getProcAddr("glGetSubroutineUniformLocation"))
	if gpGetSubroutineUniformLocation == nil {
		return errors.New("glGetSubroutineUniformLocation")
	}
	gpGetSynciv = (C.GPGETSYNCIV)(getProcAddr("glGetSynciv"))
	if gpGetSynciv == nil {
		return errors.New("glGetSynciv")
	}
	gpGetTexEnvxvOES = (C.GPGETTEXENVXVOES)(getProcAddr("glGetTexEnvxvOES"))
	gpGetTexGenxvOES = (C.GPGETTEXGENXVOES)(getProcAddr("glGetTexGenxvOES"))
	gpGetTexImage = (C.GPGETTEXIMAGE)(getProcAddr("glGetTexImage"))
	if gpGetTexImage == nil {
		return errors.New("glGetTexImage")
	}
	gpGetTexLevelParameterfv = (C.GPGETTEXLEVELPARAMETERFV)(getProcAddr("glGetTexLevelParameterfv"))
	if gpGetTexLevelParameterfv == nil {
		return errors.New("glGetTexLevelParameterfv")
	}
	gpGetTexLevelParameteriv = (C.GPGETTEXLEVELPARAMETERIV)(getProcAddr("glGetTexLevelParameteriv"))
	if gpGetTexLevelParameteriv == nil {
		return errors.New("glGetTexLevelParameteriv")
	}
	gpGetTexLevelParameterxvOES = (C.GPGETTEXLEVELPARAMETERXVOES)(getProcAddr("glGetTexLevelParameterxvOES"))
	gpGetTexParameterIiv = (C.GPGETTEXPARAMETERIIV)(getProcAddr("glGetTexParameterIiv"))
	if gpGetTexParameterIiv == nil {
		return errors.New("glGetTexParameterIiv")
	}
	gpGetTexParameterIuiv = (C.GPGETTEXPARAMETERIUIV)(getProcAddr("glGetTexParameterIuiv"))
	if gpGetTexParameterIuiv == nil {
		return errors.New("glGetTexParameterIuiv")
	}
	gpGetTexParameterfv = (C.GPGETTEXPARAMETERFV)(getProcAddr("glGetTexParameterfv"))
	if gpGetTexParameterfv == nil {
		return errors.New("glGetTexParameterfv")
	}
	gpGetTexParameteriv = (C.GPGETTEXPARAMETERIV)(getProcAddr("glGetTexParameteriv"))
	if gpGetTexParameteriv == nil {
		return errors.New("glGetTexParameteriv")
	}
	gpGetTexParameterxvOES = (C.GPGETTEXPARAMETERXVOES)(getProcAddr("glGetTexParameterxvOES"))
	gpGetTextureHandleARB = (C.GPGETTEXTUREHANDLEARB)(getProcAddr("glGetTextureHandleARB"))
	gpGetTextureImage = (C.GPGETTEXTUREIMAGE)(getProcAddr("glGetTextureImage"))
	gpGetTextureLevelParameterfv = (C.GPGETTEXTURELEVELPARAMETERFV)(getProcAddr("glGetTextureLevelParameterfv"))
	gpGetTextureLevelParameteriv = (C.GPGETTEXTURELEVELPARAMETERIV)(getProcAddr("glGetTextureLevelParameteriv"))
	gpGetTextureParameterIiv = (C.GPGETTEXTUREPARAMETERIIV)(getProcAddr("glGetTextureParameterIiv"))
	gpGetTextureParameterIuiv = (C.GPGETTEXTUREPARAMETERIUIV)(getProcAddr("glGetTextureParameterIuiv"))
	gpGetTextureParameterfv = (C.GPGETTEXTUREPARAMETERFV)(getProcAddr("glGetTextureParameterfv"))
	gpGetTextureParameteriv = (C.GPGETTEXTUREPARAMETERIV)(getProcAddr("glGetTextureParameteriv"))
	gpGetTextureSamplerHandleARB = (C.GPGETTEXTURESAMPLERHANDLEARB)(getProcAddr("glGetTextureSamplerHandleARB"))
	gpGetTextureSubImage = (C.GPGETTEXTURESUBIMAGE)(getProcAddr("glGetTextureSubImage"))
	gpGetTransformFeedbackVarying = (C.GPGETTRANSFORMFEEDBACKVARYING)(getProcAddr("glGetTransformFeedbackVarying"))
	if gpGetTransformFeedbackVarying == nil {
		return errors.New("glGetTransformFeedbackVarying")
	}
	gpGetTransformFeedbacki64_v = (C.GPGETTRANSFORMFEEDBACKI64_V)(getProcAddr("glGetTransformFeedbacki64_v"))
	gpGetTransformFeedbacki_v = (C.GPGETTRANSFORMFEEDBACKI_V)(getProcAddr("glGetTransformFeedbacki_v"))
	gpGetTransformFeedbackiv = (C.GPGETTRANSFORMFEEDBACKIV)(getProcAddr("glGetTransformFeedbackiv"))
	gpGetUniformBlockIndex = (C.GPGETUNIFORMBLOCKINDEX)(getProcAddr("glGetUniformBlockIndex"))
	if gpGetUniformBlockIndex == nil {
		return errors.New("glGetUniformBlockIndex")
	}
	gpGetUniformIndices = (C.GPGETUNIFORMINDICES)(getProcAddr("glGetUniformIndices"))
	if gpGetUniformIndices == nil {
		return errors.New("glGetUniformIndices")
	}
	gpGetUniformLocation = (C.GPGETUNIFORMLOCATION)(getProcAddr("glGetUniformLocation"))
	if gpGetUniformLocation == nil {
		return errors.New("glGetUniformLocation")
	}
	gpGetUniformSubroutineuiv = (C.GPGETUNIFORMSUBROUTINEUIV)(getProcAddr("glGetUniformSubroutineuiv"))
	if gpGetUniformSubroutineuiv == nil {
		return errors.New("glGetUniformSubroutineuiv")
	}
	gpGetUniformdv = (C.GPGETUNIFORMDV)(getProcAddr("glGetUniformdv"))
	if gpGetUniformdv == nil {
		return errors.New("glGetUniformdv")
	}
	gpGetUniformfv = (C.GPGETUNIFORMFV)(getProcAddr("glGetUniformfv"))
	if gpGetUniformfv == nil {
		return errors.New("glGetUniformfv")
	}
	gpGetUniformiv = (C.GPGETUNIFORMIV)(getProcAddr("glGetUniformiv"))
	if gpGetUniformiv == nil {
		return errors.New("glGetUniformiv")
	}
	gpGetUniformuiv = (C.GPGETUNIFORMUIV)(getProcAddr("glGetUniformuiv"))
	if gpGetUniformuiv == nil {
		return errors.New("glGetUniformuiv")
	}
	gpGetVertexArrayIndexed64iv = (C.GPGETVERTEXARRAYINDEXED64IV)(getProcAddr("glGetVertexArrayIndexed64iv"))
	gpGetVertexArrayIndexediv = (C.GPGETVERTEXARRAYINDEXEDIV)(getProcAddr("glGetVertexArrayIndexediv"))
	gpGetVertexArrayiv = (C.GPGETVERTEXARRAYIV)(getProcAddr("glGetVertexArrayiv"))
	gpGetVertexAttribIiv = (C.GPGETVERTEXATTRIBIIV)(getProcAddr("glGetVertexAttribIiv"))
	if gpGetVertexAttribIiv == nil {
		return errors.New("glGetVertexAttribIiv")
	}
	gpGetVertexAttribIuiv = (C.GPGETVERTEXATTRIBIUIV)(getProcAddr("glGetVertexAttribIuiv"))
	if gpGetVertexAttribIuiv == nil {
		return errors.New("glGetVertexAttribIuiv")
	}
	gpGetVertexAttribLdv = (C.GPGETVERTEXATTRIBLDV)(getProcAddr("glGetVertexAttribLdv"))
	if gpGetVertexAttribLdv == nil {
		return errors.New("glGetVertexAttribLdv")
	}
	gpGetVertexAttribLui64vARB = (C.GPGETVERTEXATTRIBLUI64VARB)(getProcAddr("glGetVertexAttribLui64vARB"))
	gpGetVertexAttribPointerv = (C.GPGETVERTEXATTRIBPOINTERV)(getProcAddr("glGetVertexAttribPointerv"))
	if gpGetVertexAttribPointerv == nil {
		return errors.New("glGetVertexAttribPointerv")
	}
	gpGetVertexAttribdv = (C.GPGETVERTEXATTRIBDV)(getProcAddr("glGetVertexAttribdv"))
	if gpGetVertexAttribdv == nil {
		return errors.New("glGetVertexAttribdv")
	}
	gpGetVertexAttribfv = (C.GPGETVERTEXATTRIBFV)(getProcAddr("glGetVertexAttribfv"))
	if gpGetVertexAttribfv == nil {
		return errors.New("glGetVertexAttribfv")
	}
	gpGetVertexAttribiv = (C.GPGETVERTEXATTRIBIV)(getProcAddr("glGetVertexAttribiv"))
	if gpGetVertexAttribiv == nil {
		return errors.New("glGetVertexAttribiv")
	}
	gpGetnCompressedTexImageARB = (C.GPGETNCOMPRESSEDTEXIMAGEARB)(getProcAddr("glGetnCompressedTexImageARB"))
	gpGetnTexImageARB = (C.GPGETNTEXIMAGEARB)(getProcAddr("glGetnTexImageARB"))
	gpGetnUniformdvARB = (C.GPGETNUNIFORMDVARB)(getProcAddr("glGetnUniformdvARB"))
	gpGetnUniformfv = (C.GPGETNUNIFORMFV)(getProcAddr("glGetnUniformfv"))
	gpGetnUniformfvARB = (C.GPGETNUNIFORMFVARB)(getProcAddr("glGetnUniformfvARB"))
	gpGetnUniformfvKHR = (C.GPGETNUNIFORMFVKHR)(getProcAddr("glGetnUniformfvKHR"))
	gpGetnUniformiv = (C.GPGETNUNIFORMIV)(getProcAddr("glGetnUniformiv"))
	gpGetnUniformivARB = (C.GPGETNUNIFORMIVARB)(getProcAddr("glGetnUniformivARB"))
	gpGetnUniformivKHR = (C.GPGETNUNIFORMIVKHR)(getProcAddr("glGetnUniformivKHR"))
	gpGetnUniformuiv = (C.GPGETNUNIFORMUIV)(getProcAddr("glGetnUniformuiv"))
	gpGetnUniformuivARB = (C.GPGETNUNIFORMUIVARB)(getProcAddr("glGetnUniformuivARB"))
	gpGetnUniformuivKHR = (C.GPGETNUNIFORMUIVKHR)(getProcAddr("glGetnUniformuivKHR"))
	gpHint = (C.GPHINT)(getProcAddr("glHint"))
	if gpHint == nil {
		return errors.New("glHint")
	}
	gpIndexxOES = (C.GPINDEXXOES)(getProcAddr("glIndexxOES"))
	gpIndexxvOES = (C.GPINDEXXVOES)(getProcAddr("glIndexxvOES"))
	gpInsertEventMarkerEXT = (C.GPINSERTEVENTMARKEREXT)(getProcAddr("glInsertEventMarkerEXT"))
	gpInvalidateBufferData = (C.GPINVALIDATEBUFFERDATA)(getProcAddr("glInvalidateBufferData"))
	gpInvalidateBufferSubData = (C.GPINVALIDATEBUFFERSUBDATA)(getProcAddr("glInvalidateBufferSubData"))
	gpInvalidateFramebuffer = (C.GPINVALIDATEFRAMEBUFFER)(getProcAddr("glInvalidateFramebuffer"))
	gpInvalidateNamedFramebufferData = (C.GPINVALIDATENAMEDFRAMEBUFFERDATA)(getProcAddr("glInvalidateNamedFramebufferData"))
	gpInvalidateNamedFramebufferSubData = (C.GPINVALIDATENAMEDFRAMEBUFFERSUBDATA)(getProcAddr("glInvalidateNamedFramebufferSubData"))
	gpInvalidateSubFramebuffer = (C.GPINVALIDATESUBFRAMEBUFFER)(getProcAddr("glInvalidateSubFramebuffer"))
	gpInvalidateTexImage = (C.GPINVALIDATETEXIMAGE)(getProcAddr("glInvalidateTexImage"))
	gpInvalidateTexSubImage = (C.GPINVALIDATETEXSUBIMAGE)(getProcAddr("glInvalidateTexSubImage"))
	gpIsBuffer = (C.GPISBUFFER)(getProcAddr("glIsBuffer"))
	if gpIsBuffer == nil {
		return errors.New("glIsBuffer")
	}
	gpIsEnabled = (C.GPISENABLED)(getProcAddr("glIsEnabled"))
	if gpIsEnabled == nil {
		return errors.New("glIsEnabled")
	}
	gpIsEnabledi = (C.GPISENABLEDI)(getProcAddr("glIsEnabledi"))
	if gpIsEnabledi == nil {
		return errors.New("glIsEnabledi")
	}
	gpIsFenceNV = (C.GPISFENCENV)(getProcAddr("glIsFenceNV"))
	gpIsFramebuffer = (C.GPISFRAMEBUFFER)(getProcAddr("glIsFramebuffer"))
	if gpIsFramebuffer == nil {
		return errors.New("glIsFramebuffer")
	}
	gpIsImageHandleResidentARB = (C.GPISIMAGEHANDLERESIDENTARB)(getProcAddr("glIsImageHandleResidentARB"))
	gpIsNamedStringARB = (C.GPISNAMEDSTRINGARB)(getProcAddr("glIsNamedStringARB"))
	gpIsProgram = (C.GPISPROGRAM)(getProcAddr("glIsProgram"))
	if gpIsProgram == nil {
		return errors.New("glIsProgram")
	}
	gpIsProgramPipeline = (C.GPISPROGRAMPIPELINE)(getProcAddr("glIsProgramPipeline"))
	if gpIsProgramPipeline == nil {
		return errors.New("glIsProgramPipeline")
	}
	gpIsProgramPipelineEXT = (C.GPISPROGRAMPIPELINEEXT)(getProcAddr("glIsProgramPipelineEXT"))
	gpIsQuery = (C.GPISQUERY)(getProcAddr("glIsQuery"))
	if gpIsQuery == nil {
		return errors.New("glIsQuery")
	}
	gpIsRenderbuffer = (C.GPISRENDERBUFFER)(getProcAddr("glIsRenderbuffer"))
	if gpIsRenderbuffer == nil {
		return errors.New("glIsRenderbuffer")
	}
	gpIsSampler = (C.GPISSAMPLER)(getProcAddr("glIsSampler"))
	if gpIsSampler == nil {
		return errors.New("glIsSampler")
	}
	gpIsShader = (C.GPISSHADER)(getProcAddr("glIsShader"))
	if gpIsShader == nil {
		return errors.New("glIsShader")
	}
	gpIsSync = (C.GPISSYNC)(getProcAddr("glIsSync"))
	if gpIsSync == nil {
		return errors.New("glIsSync")
	}
	gpIsTexture = (C.GPISTEXTURE)(getProcAddr("glIsTexture"))
	if gpIsTexture == nil {
		return errors.New("glIsTexture")
	}
	gpIsTextureHandleResidentARB = (C.GPISTEXTUREHANDLERESIDENTARB)(getProcAddr("glIsTextureHandleResidentARB"))
	gpIsTransformFeedback = (C.GPISTRANSFORMFEEDBACK)(getProcAddr("glIsTransformFeedback"))
	if gpIsTransformFeedback == nil {
		return errors.New("glIsTransformFeedback")
	}
	gpIsVertexArray = (C.GPISVERTEXARRAY)(getProcAddr("glIsVertexArray"))
	if gpIsVertexArray == nil {
		return errors.New("glIsVertexArray")
	}
	gpLabelObjectEXT = (C.GPLABELOBJECTEXT)(getProcAddr("glLabelObjectEXT"))
	gpLightModelxOES = (C.GPLIGHTMODELXOES)(getProcAddr("glLightModelxOES"))
	gpLightModelxvOES = (C.GPLIGHTMODELXVOES)(getProcAddr("glLightModelxvOES"))
	gpLightxOES = (C.GPLIGHTXOES)(getProcAddr("glLightxOES"))
	gpLightxvOES = (C.GPLIGHTXVOES)(getProcAddr("glLightxvOES"))
	gpLineWidth = (C.GPLINEWIDTH)(getProcAddr("glLineWidth"))
	if gpLineWidth == nil {
		return errors.New("glLineWidth")
	}
	gpLineWidthxOES = (C.GPLINEWIDTHXOES)(getProcAddr("glLineWidthxOES"))
	gpLinkProgram = (C.GPLINKPROGRAM)(getProcAddr("glLinkProgram"))
	if gpLinkProgram == nil {
		return errors.New("glLinkProgram")
	}
	gpLoadMatrixxOES = (C.GPLOADMATRIXXOES)(getProcAddr("glLoadMatrixxOES"))
	gpLoadTransposeMatrixxOES = (C.GPLOADTRANSPOSEMATRIXXOES)(getProcAddr("glLoadTransposeMatrixxOES"))
	gpLogicOp = (C.GPLOGICOP)(getProcAddr("glLogicOp"))
	if gpLogicOp == nil {
		return errors.New("glLogicOp")
	}
	gpMakeImageHandleNonResidentARB = (C.GPMAKEIMAGEHANDLENONRESIDENTARB)(getProcAddr("glMakeImageHandleNonResidentARB"))
	gpMakeImageHandleResidentARB = (C.GPMAKEIMAGEHANDLERESIDENTARB)(getProcAddr("glMakeImageHandleResidentARB"))
	gpMakeTextureHandleNonResidentARB = (C.GPMAKETEXTUREHANDLENONRESIDENTARB)(getProcAddr("glMakeTextureHandleNonResidentARB"))
	gpMakeTextureHandleResidentARB = (C.GPMAKETEXTUREHANDLERESIDENTARB)(getProcAddr("glMakeTextureHandleResidentARB"))
	gpMap1xOES = (C.GPMAP1XOES)(getProcAddr("glMap1xOES"))
	gpMap2xOES = (C.GPMAP2XOES)(getProcAddr("glMap2xOES"))
	gpMapBuffer = (C.GPMAPBUFFER)(getProcAddr("glMapBuffer"))
	if gpMapBuffer == nil {
		return errors.New("glMapBuffer")
	}
	gpMapBufferRange = (C.GPMAPBUFFERRANGE)(getProcAddr("glMapBufferRange"))
	if gpMapBufferRange == nil {
		return errors.New("glMapBufferRange")
	}
	gpMapGrid1xOES = (C.GPMAPGRID1XOES)(getProcAddr("glMapGrid1xOES"))
	gpMapGrid2xOES = (C.GPMAPGRID2XOES)(getProcAddr("glMapGrid2xOES"))
	gpMapNamedBuffer = (C.GPMAPNAMEDBUFFER)(getProcAddr("glMapNamedBuffer"))
	gpMapNamedBufferRange = (C.GPMAPNAMEDBUFFERRANGE)(getProcAddr("glMapNamedBufferRange"))
	gpMaterialxOES = (C.GPMATERIALXOES)(getProcAddr("glMaterialxOES"))
	gpMaterialxvOES = (C.GPMATERIALXVOES)(getProcAddr("glMaterialxvOES"))
	gpMemoryBarrier = (C.GPMEMORYBARRIER)(getProcAddr("glMemoryBarrier"))
	gpMemoryBarrierByRegion = (C.GPMEMORYBARRIERBYREGION)(getProcAddr("glMemoryBarrierByRegion"))
	gpMinSampleShading = (C.GPMINSAMPLESHADING)(getProcAddr("glMinSampleShading"))
	if gpMinSampleShading == nil {
		return errors.New("glMinSampleShading")
	}
	gpMinSampleShadingARB = (C.GPMINSAMPLESHADINGARB)(getProcAddr("glMinSampleShadingARB"))
	gpMultMatrixxOES = (C.GPMULTMATRIXXOES)(getProcAddr("glMultMatrixxOES"))
	gpMultTransposeMatrixxOES = (C.GPMULTTRANSPOSEMATRIXXOES)(getProcAddr("glMultTransposeMatrixxOES"))
	gpMultiDrawArrays = (C.GPMULTIDRAWARRAYS)(getProcAddr("glMultiDrawArrays"))
	if gpMultiDrawArrays == nil {
		return errors.New("glMultiDrawArrays")
	}
	gpMultiDrawArraysEXT = (C.GPMULTIDRAWARRAYSEXT)(getProcAddr("glMultiDrawArraysEXT"))
	gpMultiDrawArraysIndirect = (C.GPMULTIDRAWARRAYSINDIRECT)(getProcAddr("glMultiDrawArraysIndirect"))
	gpMultiDrawArraysIndirectCountARB = (C.GPMULTIDRAWARRAYSINDIRECTCOUNTARB)(getProcAddr("glMultiDrawArraysIndirectCountARB"))
	gpMultiDrawElements = (C.GPMULTIDRAWELEMENTS)(getProcAddr("glMultiDrawElements"))
	if gpMultiDrawElements == nil {
		return errors.New("glMultiDrawElements")
	}
	gpMultiDrawElementsBaseVertex = (C.GPMULTIDRAWELEMENTSBASEVERTEX)(getProcAddr("glMultiDrawElementsBaseVertex"))
	if gpMultiDrawElementsBaseVertex == nil {
		return errors.New("glMultiDrawElementsBaseVertex")
	}
	gpMultiDrawElementsEXT = (C.GPMULTIDRAWELEMENTSEXT)(getProcAddr("glMultiDrawElementsEXT"))
	gpMultiDrawElementsIndirect = (C.GPMULTIDRAWELEMENTSINDIRECT)(getProcAddr("glMultiDrawElementsIndirect"))
	gpMultiDrawElementsIndirectCountARB = (C.GPMULTIDRAWELEMENTSINDIRECTCOUNTARB)(getProcAddr("glMultiDrawElementsIndirectCountARB"))
	gpMultiTexCoord1bOES = (C.GPMULTITEXCOORD1BOES)(getProcAddr("glMultiTexCoord1bOES"))
	gpMultiTexCoord1bvOES = (C.GPMULTITEXCOORD1BVOES)(getProcAddr("glMultiTexCoord1bvOES"))
	gpMultiTexCoord1xOES = (C.GPMULTITEXCOORD1XOES)(getProcAddr("glMultiTexCoord1xOES"))
	gpMultiTexCoord1xvOES = (C.GPMULTITEXCOORD1XVOES)(getProcAddr("glMultiTexCoord1xvOES"))
	gpMultiTexCoord2bOES = (C.GPMULTITEXCOORD2BOES)(getProcAddr("glMultiTexCoord2bOES"))
	gpMultiTexCoord2bvOES = (C.GPMULTITEXCOORD2BVOES)(getProcAddr("glMultiTexCoord2bvOES"))
	gpMultiTexCoord2xOES = (C.GPMULTITEXCOORD2XOES)(getProcAddr("glMultiTexCoord2xOES"))
	gpMultiTexCoord2xvOES = (C.GPMULTITEXCOORD2XVOES)(getProcAddr("glMultiTexCoord2xvOES"))
	gpMultiTexCoord3bOES = (C.GPMULTITEXCOORD3BOES)(getProcAddr("glMultiTexCoord3bOES"))
	gpMultiTexCoord3bvOES = (C.GPMULTITEXCOORD3BVOES)(getProcAddr("glMultiTexCoord3bvOES"))
	gpMultiTexCoord3xOES = (C.GPMULTITEXCOORD3XOES)(getProcAddr("glMultiTexCoord3xOES"))
	gpMultiTexCoord3xvOES = (C.GPMULTITEXCOORD3XVOES)(getProcAddr("glMultiTexCoord3xvOES"))
	gpMultiTexCoord4bOES = (C.GPMULTITEXCOORD4BOES)(getProcAddr("glMultiTexCoord4bOES"))
	gpMultiTexCoord4bvOES = (C.GPMULTITEXCOORD4BVOES)(getProcAddr("glMultiTexCoord4bvOES"))
	gpMultiTexCoord4xOES = (C.GPMULTITEXCOORD4XOES)(getProcAddr("glMultiTexCoord4xOES"))
	gpMultiTexCoord4xvOES = (C.GPMULTITEXCOORD4XVOES)(getProcAddr("glMultiTexCoord4xvOES"))
	gpNamedBufferData = (C.GPNAMEDBUFFERDATA)(getProcAddr("glNamedBufferData"))
	gpNamedBufferPageCommitmentARB = (C.GPNAMEDBUFFERPAGECOMMITMENTARB)(getProcAddr("glNamedBufferPageCommitmentARB"))
	gpNamedBufferPageCommitmentEXT = (C.GPNAMEDBUFFERPAGECOMMITMENTEXT)(getProcAddr("glNamedBufferPageCommitmentEXT"))
	gpNamedBufferStorage = (C.GPNAMEDBUFFERSTORAGE)(getProcAddr("glNamedBufferStorage"))
	gpNamedBufferSubData = (C.GPNAMEDBUFFERSUBDATA)(getProcAddr("glNamedBufferSubData"))
	gpNamedFramebufferDrawBuffer = (C.GPNAMEDFRAMEBUFFERDRAWBUFFER)(getProcAddr("glNamedFramebufferDrawBuffer"))
	gpNamedFramebufferDrawBuffers = (C.GPNAMEDFRAMEBUFFERDRAWBUFFERS)(getProcAddr("glNamedFramebufferDrawBuffers"))
	gpNamedFramebufferParameteri = (C.GPNAMEDFRAMEBUFFERPARAMETERI)(getProcAddr("glNamedFramebufferParameteri"))
	gpNamedFramebufferReadBuffer = (C.GPNAMEDFRAMEBUFFERREADBUFFER)(getProcAddr("glNamedFramebufferReadBuffer"))
	gpNamedFramebufferRenderbuffer = (C.GPNAMEDFRAMEBUFFERRENDERBUFFER)(getProcAddr("glNamedFramebufferRenderbuffer"))
	gpNamedFramebufferTexture = (C.GPNAMEDFRAMEBUFFERTEXTURE)(getProcAddr("glNamedFramebufferTexture"))
	gpNamedFramebufferTextureLayer = (C.GPNAMEDFRAMEBUFFERTEXTURELAYER)(getProcAddr("glNamedFramebufferTextureLayer"))
	gpNamedRenderbufferStorage = (C.GPNAMEDRENDERBUFFERSTORAGE)(getProcAddr("glNamedRenderbufferStorage"))
	gpNamedRenderbufferStorageMultisample = (C.GPNAMEDRENDERBUFFERSTORAGEMULTISAMPLE)(getProcAddr("glNamedRenderbufferStorageMultisample"))
	gpNamedStringARB = (C.GPNAMEDSTRINGARB)(getProcAddr("glNamedStringARB"))
	gpNormal3xOES = (C.GPNORMAL3XOES)(getProcAddr("glNormal3xOES"))
	gpNormal3xvOES = (C.GPNORMAL3XVOES)(getProcAddr("glNormal3xvOES"))
	gpObjectLabel = (C.GPOBJECTLABEL)(getProcAddr("glObjectLabel"))
	gpObjectLabelKHR = (C.GPOBJECTLABELKHR)(getProcAddr("glObjectLabelKHR"))
	gpObjectPtrLabel = (C.GPOBJECTPTRLABEL)(getProcAddr("glObjectPtrLabel"))
	gpObjectPtrLabelKHR = (C.GPOBJECTPTRLABELKHR)(getProcAddr("glObjectPtrLabelKHR"))
	gpOrthofOES = (C.GPORTHOFOES)(getProcAddr("glOrthofOES"))
	gpOrthoxOES = (C.GPORTHOXOES)(getProcAddr("glOrthoxOES"))
	gpPassThroughxOES = (C.GPPASSTHROUGHXOES)(getProcAddr("glPassThroughxOES"))
	gpPatchParameterfv = (C.GPPATCHPARAMETERFV)(getProcAddr("glPatchParameterfv"))
	if gpPatchParameterfv == nil {
		return errors.New("glPatchParameterfv")
	}
	gpPatchParameteri = (C.GPPATCHPARAMETERI)(getProcAddr("glPatchParameteri"))
	if gpPatchParameteri == nil {
		return errors.New("glPatchParameteri")
	}
	gpPauseTransformFeedback = (C.GPPAUSETRANSFORMFEEDBACK)(getProcAddr("glPauseTransformFeedback"))
	if gpPauseTransformFeedback == nil {
		return errors.New("glPauseTransformFeedback")
	}
	gpPixelMapx = (C.GPPIXELMAPX)(getProcAddr("glPixelMapx"))
	gpPixelStoref = (C.GPPIXELSTOREF)(getProcAddr("glPixelStoref"))
	if gpPixelStoref == nil {
		return errors.New("glPixelStoref")
	}
	gpPixelStorei = (C.GPPIXELSTOREI)(getProcAddr("glPixelStorei"))
	if gpPixelStorei == nil {
		return errors.New("glPixelStorei")
	}
	gpPixelStorex = (C.GPPIXELSTOREX)(getProcAddr("glPixelStorex"))
	gpPixelTransferxOES = (C.GPPIXELTRANSFERXOES)(getProcAddr("glPixelTransferxOES"))
	gpPixelZoomxOES = (C.GPPIXELZOOMXOES)(getProcAddr("glPixelZoomxOES"))
	gpPointParameterf = (C.GPPOINTPARAMETERF)(getProcAddr("glPointParameterf"))
	if gpPointParameterf == nil {
		return errors.New("glPointParameterf")
	}
	gpPointParameterfv = (C.GPPOINTPARAMETERFV)(getProcAddr("glPointParameterfv"))
	if gpPointParameterfv == nil {
		return errors.New("glPointParameterfv")
	}
	gpPointParameteri = (C.GPPOINTPARAMETERI)(getProcAddr("glPointParameteri"))
	if gpPointParameteri == nil {
		return errors.New("glPointParameteri")
	}
	gpPointParameteriv = (C.GPPOINTPARAMETERIV)(getProcAddr("glPointParameteriv"))
	if gpPointParameteriv == nil {
		return errors.New("glPointParameteriv")
	}
	gpPointParameterxOES = (C.GPPOINTPARAMETERXOES)(getProcAddr("glPointParameterxOES"))
	gpPointParameterxvOES = (C.GPPOINTPARAMETERXVOES)(getProcAddr("glPointParameterxvOES"))
	gpPointSize = (C.GPPOINTSIZE)(getProcAddr("glPointSize"))
	if gpPointSize == nil {
		return errors.New("glPointSize")
	}
	gpPointSizexOES = (C.GPPOINTSIZEXOES)(getProcAddr("glPointSizexOES"))
	gpPolygonMode = (C.GPPOLYGONMODE)(getProcAddr("glPolygonMode"))
	if gpPolygonMode == nil {
		return errors.New("glPolygonMode")
	}
	gpPolygonOffset = (C.GPPOLYGONOFFSET)(getProcAddr("glPolygonOffset"))
	if gpPolygonOffset == nil {
		return errors.New("glPolygonOffset")
	}
	gpPolygonOffsetxOES = (C.GPPOLYGONOFFSETXOES)(getProcAddr("glPolygonOffsetxOES"))
	gpPopDebugGroup = (C.GPPOPDEBUGGROUP)(getProcAddr("glPopDebugGroup"))
	gpPopDebugGroupKHR = (C.GPPOPDEBUGGROUPKHR)(getProcAddr("glPopDebugGroupKHR"))
	gpPopGroupMarkerEXT = (C.GPPOPGROUPMARKEREXT)(getProcAddr("glPopGroupMarkerEXT"))
	gpPrimitiveRestartIndex = (C.GPPRIMITIVERESTARTINDEX)(getProcAddr("glPrimitiveRestartIndex"))
	if gpPrimitiveRestartIndex == nil {
		return errors.New("glPrimitiveRestartIndex")
	}
	gpPrioritizeTexturesxOES = (C.GPPRIORITIZETEXTURESXOES)(getProcAddr("glPrioritizeTexturesxOES"))
	gpProgramBinary = (C.GPPROGRAMBINARY)(getProcAddr("glProgramBinary"))
	if gpProgramBinary == nil {
		return errors.New("glProgramBinary")
	}
	gpProgramParameteri = (C.GPPROGRAMPARAMETERI)(getProcAddr("glProgramParameteri"))
	if gpProgramParameteri == nil {
		return errors.New("glProgramParameteri")
	}
	gpProgramParameteriEXT = (C.GPPROGRAMPARAMETERIEXT)(getProcAddr("glProgramParameteriEXT"))
	gpProgramUniform1d = (C.GPPROGRAMUNIFORM1D)(getProcAddr("glProgramUniform1d"))
	if gpProgramUniform1d == nil {
		return errors.New("glProgramUniform1d")
	}
	gpProgramUniform1dv = (C.GPPROGRAMUNIFORM1DV)(getProcAddr("glProgramUniform1dv"))
	if gpProgramUniform1dv == nil {
		return errors.New("glProgramUniform1dv")
	}
	gpProgramUniform1f = (C.GPPROGRAMUNIFORM1F)(getProcAddr("glProgramUniform1f"))
	if gpProgramUniform1f == nil {
		return errors.New("glProgramUniform1f")
	}
	gpProgramUniform1fEXT = (C.GPPROGRAMUNIFORM1FEXT)(getProcAddr("glProgramUniform1fEXT"))
	gpProgramUniform1fv = (C.GPPROGRAMUNIFORM1FV)(getProcAddr("glProgramUniform1fv"))
	if gpProgramUniform1fv == nil {
		return errors.New("glProgramUniform1fv")
	}
	gpProgramUniform1fvEXT = (C.GPPROGRAMUNIFORM1FVEXT)(getProcAddr("glProgramUniform1fvEXT"))
	gpProgramUniform1i = (C.GPPROGRAMUNIFORM1I)(getProcAddr("glProgramUniform1i"))
	if gpProgramUniform1i == nil {
		return errors.New("glProgramUniform1i")
	}
	gpProgramUniform1iEXT = (C.GPPROGRAMUNIFORM1IEXT)(getProcAddr("glProgramUniform1iEXT"))
	gpProgramUniform1iv = (C.GPPROGRAMUNIFORM1IV)(getProcAddr("glProgramUniform1iv"))
	if gpProgramUniform1iv == nil {
		return errors.New("glProgramUniform1iv")
	}
	gpProgramUniform1ivEXT = (C.GPPROGRAMUNIFORM1IVEXT)(getProcAddr("glProgramUniform1ivEXT"))
	gpProgramUniform1ui = (C.GPPROGRAMUNIFORM1UI)(getProcAddr("glProgramUniform1ui"))
	if gpProgramUniform1ui == nil {
		return errors.New("glProgramUniform1ui")
	}
	gpProgramUniform1uiEXT = (C.GPPROGRAMUNIFORM1UIEXT)(getProcAddr("glProgramUniform1uiEXT"))
	gpProgramUniform1uiv = (C.GPPROGRAMUNIFORM1UIV)(getProcAddr("glProgramUniform1uiv"))
	if gpProgramUniform1uiv == nil {
		return errors.New("glProgramUniform1uiv")
	}
	gpProgramUniform1uivEXT = (C.GPPROGRAMUNIFORM1UIVEXT)(getProcAddr("glProgramUniform1uivEXT"))
	gpProgramUniform2d = (C.GPPROGRAMUNIFORM2D)(getProcAddr("glProgramUniform2d"))
	if gpProgramUniform2d == nil {
		return errors.New("glProgramUniform2d")
	}
	gpProgramUniform2dv = (C.GPPROGRAMUNIFORM2DV)(getProcAddr("glProgramUniform2dv"))
	if gpProgramUniform2dv == nil {
		return errors.New("glProgramUniform2dv")
	}
	gpProgramUniform2f = (C.GPPROGRAMUNIFORM2F)(getProcAddr("glProgramUniform2f"))
	if gpProgramUniform2f == nil {
		return errors.New("glProgramUniform2f")
	}
	gpProgramUniform2fEXT = (C.GPPROGRAMUNIFORM2FEXT)(getProcAddr("glProgramUniform2fEXT"))
	gpProgramUniform2fv = (C.GPPROGRAMUNIFORM2FV)(getProcAddr("glProgramUniform2fv"))
	if gpProgramUniform2fv == nil {
		return errors.New("glProgramUniform2fv")
	}
	gpProgramUniform2fvEXT = (C.GPPROGRAMUNIFORM2FVEXT)(getProcAddr("glProgramUniform2fvEXT"))
	gpProgramUniform2i = (C.GPPROGRAMUNIFORM2I)(getProcAddr("glProgramUniform2i"))
	if gpProgramUniform2i == nil {
		return errors.New("glProgramUniform2i")
	}
	gpProgramUniform2iEXT = (C.GPPROGRAMUNIFORM2IEXT)(getProcAddr("glProgramUniform2iEXT"))
	gpProgramUniform2iv = (C.GPPROGRAMUNIFORM2IV)(getProcAddr("glProgramUniform2iv"))
	if gpProgramUniform2iv == nil {
		return errors.New("glProgramUniform2iv")
	}
	gpProgramUniform2ivEXT = (C.GPPROGRAMUNIFORM2IVEXT)(getProcAddr("glProgramUniform2ivEXT"))
	gpProgramUniform2ui = (C.GPPROGRAMUNIFORM2UI)(getProcAddr("glProgramUniform2ui"))
	if gpProgramUniform2ui == nil {
		return errors.New("glProgramUniform2ui")
	}
	gpProgramUniform2uiEXT = (C.GPPROGRAMUNIFORM2UIEXT)(getProcAddr("glProgramUniform2uiEXT"))
	gpProgramUniform2uiv = (C.GPPROGRAMUNIFORM2UIV)(getProcAddr("glProgramUniform2uiv"))
	if gpProgramUniform2uiv == nil {
		return errors.New("glProgramUniform2uiv")
	}
	gpProgramUniform2uivEXT = (C.GPPROGRAMUNIFORM2UIVEXT)(getProcAddr("glProgramUniform2uivEXT"))
	gpProgramUniform3d = (C.GPPROGRAMUNIFORM3D)(getProcAddr("glProgramUniform3d"))
	if gpProgramUniform3d == nil {
		return errors.New("glProgramUniform3d")
	}
	gpProgramUniform3dv = (C.GPPROGRAMUNIFORM3DV)(getProcAddr("glProgramUniform3dv"))
	if gpProgramUniform3dv == nil {
		return errors.New("glProgramUniform3dv")
	}
	gpProgramUniform3f = (C.GPPROGRAMUNIFORM3F)(getProcAddr("glProgramUniform3f"))
	if gpProgramUniform3f == nil {
		return errors.New("glProgramUniform3f")
	}
	gpProgramUniform3fEXT = (C.GPPROGRAMUNIFORM3FEXT)(getProcAddr("glProgramUniform3fEXT"))
	gpProgramUniform3fv = (C.GPPROGRAMUNIFORM3FV)(getProcAddr("glProgramUniform3fv"))
	if gpProgramUniform3fv == nil {
		return errors.New("glProgramUniform3fv")
	}
	gpProgramUniform3fvEXT = (C.GPPROGRAMUNIFORM3FVEXT)(getProcAddr("glProgramUniform3fvEXT"))
	gpProgramUniform3i = (C.GPPROGRAMUNIFORM3I)(getProcAddr("glProgramUniform3i"))
	if gpProgramUniform3i == nil {
		return errors.New("glProgramUniform3i")
	}
	gpProgramUniform3iEXT = (C.GPPROGRAMUNIFORM3IEXT)(getProcAddr("glProgramUniform3iEXT"))
	gpProgramUniform3iv = (C.GPPROGRAMUNIFORM3IV)(getProcAddr("glProgramUniform3iv"))
	if gpProgramUniform3iv == nil {
		return errors.New("glProgramUniform3iv")
	}
	gpProgramUniform3ivEXT = (C.GPPROGRAMUNIFORM3IVEXT)(getProcAddr("glProgramUniform3ivEXT"))
	gpProgramUniform3ui = (C.GPPROGRAMUNIFORM3UI)(getProcAddr("glProgramUniform3ui"))
	if gpProgramUniform3ui == nil {
		return errors.New("glProgramUniform3ui")
	}
	gpProgramUniform3uiEXT = (C.GPPROGRAMUNIFORM3UIEXT)(getProcAddr("glProgramUniform3uiEXT"))
	gpProgramUniform3uiv = (C.GPPROGRAMUNIFORM3UIV)(getProcAddr("glProgramUniform3uiv"))
	if gpProgramUniform3uiv == nil {
		return errors.New("glProgramUniform3uiv")
	}
	gpProgramUniform3uivEXT = (C.GPPROGRAMUNIFORM3UIVEXT)(getProcAddr("glProgramUniform3uivEXT"))
	gpProgramUniform4d = (C.GPPROGRAMUNIFORM4D)(getProcAddr("glProgramUniform4d"))
	if gpProgramUniform4d == nil {
		return errors.New("glProgramUniform4d")
	}
	gpProgramUniform4dv = (C.GPPROGRAMUNIFORM4DV)(getProcAddr("glProgramUniform4dv"))
	if gpProgramUniform4dv == nil {
		return errors.New("glProgramUniform4dv")
	}
	gpProgramUniform4f = (C.GPPROGRAMUNIFORM4F)(getProcAddr("glProgramUniform4f"))
	if gpProgramUniform4f == nil {
		return errors.New("glProgramUniform4f")
	}
	gpProgramUniform4fEXT = (C.GPPROGRAMUNIFORM4FEXT)(getProcAddr("glProgramUniform4fEXT"))
	gpProgramUniform4fv = (C.GPPROGRAMUNIFORM4FV)(getProcAddr("glProgramUniform4fv"))
	if gpProgramUniform4fv == nil {
		return errors.New("glProgramUniform4fv")
	}
	gpProgramUniform4fvEXT = (C.GPPROGRAMUNIFORM4FVEXT)(getProcAddr("glProgramUniform4fvEXT"))
	gpProgramUniform4i = (C.GPPROGRAMUNIFORM4I)(getProcAddr("glProgramUniform4i"))
	if gpProgramUniform4i == nil {
		return errors.New("glProgramUniform4i")
	}
	gpProgramUniform4iEXT = (C.GPPROGRAMUNIFORM4IEXT)(getProcAddr("glProgramUniform4iEXT"))
	gpProgramUniform4iv = (C.GPPROGRAMUNIFORM4IV)(getProcAddr("glProgramUniform4iv"))
	if gpProgramUniform4iv == nil {
		return errors.New("glProgramUniform4iv")
	}
	gpProgramUniform4ivEXT = (C.GPPROGRAMUNIFORM4IVEXT)(getProcAddr("glProgramUniform4ivEXT"))
	gpProgramUniform4ui = (C.GPPROGRAMUNIFORM4UI)(getProcAddr("glProgramUniform4ui"))
	if gpProgramUniform4ui == nil {
		return errors.New("glProgramUniform4ui")
	}
	gpProgramUniform4uiEXT = (C.GPPROGRAMUNIFORM4UIEXT)(getProcAddr("glProgramUniform4uiEXT"))
	gpProgramUniform4uiv = (C.GPPROGRAMUNIFORM4UIV)(getProcAddr("glProgramUniform4uiv"))
	if gpProgramUniform4uiv == nil {
		return errors.New("glProgramUniform4uiv")
	}
	gpProgramUniform4uivEXT = (C.GPPROGRAMUNIFORM4UIVEXT)(getProcAddr("glProgramUniform4uivEXT"))
	gpProgramUniformHandleui64ARB = (C.GPPROGRAMUNIFORMHANDLEUI64ARB)(getProcAddr("glProgramUniformHandleui64ARB"))
	gpProgramUniformHandleui64vARB = (C.GPPROGRAMUNIFORMHANDLEUI64VARB)(getProcAddr("glProgramUniformHandleui64vARB"))
	gpProgramUniformMatrix2dv = (C.GPPROGRAMUNIFORMMATRIX2DV)(getProcAddr("glProgramUniformMatrix2dv"))
	if gpProgramUniformMatrix2dv == nil {
		return errors.New("glProgramUniformMatrix2dv")
	}
	gpProgramUniformMatrix2fv = (C.GPPROGRAMUNIFORMMATRIX2FV)(getProcAddr("glProgramUniformMatrix2fv"))
	if gpProgramUniformMatrix2fv == nil {
		return errors.New("glProgramUniformMatrix2fv")
	}
	gpProgramUniformMatrix2fvEXT = (C.GPPROGRAMUNIFORMMATRIX2FVEXT)(getProcAddr("glProgramUniformMatrix2fvEXT"))
	gpProgramUniformMatrix2x3dv = (C.GPPROGRAMUNIFORMMATRIX2X3DV)(getProcAddr("glProgramUniformMatrix2x3dv"))
	if gpProgramUniformMatrix2x3dv == nil {
		return errors.New("glProgramUniformMatrix2x3dv")
	}
	gpProgramUniformMatrix2x3fv = (C.GPPROGRAMUNIFORMMATRIX2X3FV)(getProcAddr("glProgramUniformMatrix2x3fv"))
	if gpProgramUniformMatrix2x3fv == nil {
		return errors.New("glProgramUniformMatrix2x3fv")
	}
	gpProgramUniformMatrix2x3fvEXT = (C.GPPROGRAMUNIFORMMATRIX2X3FVEXT)(getProcAddr("glProgramUniformMatrix2x3fvEXT"))
	gpProgramUniformMatrix2x4dv = (C.GPPROGRAMUNIFORMMATRIX2X4DV)(getProcAddr("glProgramUniformMatrix2x4dv"))
	if gpProgramUniformMatrix2x4dv == nil {
		return errors.New("glProgramUniformMatrix2x4dv")
	}
	gpProgramUniformMatrix2x4fv = (C.GPPROGRAMUNIFORMMATRIX2X4FV)(getProcAddr("glProgramUniformMatrix2x4fv"))
	if gpProgramUniformMatrix2x4fv == nil {
		return errors.New("glProgramUniformMatrix2x4fv")
	}
	gpProgramUniformMatrix2x4fvEXT = (C.GPPROGRAMUNIFORMMATRIX2X4FVEXT)(getProcAddr("glProgramUniformMatrix2x4fvEXT"))
	gpProgramUniformMatrix3dv = (C.GPPROGRAMUNIFORMMATRIX3DV)(getProcAddr("glProgramUniformMatrix3dv"))
	if gpProgramUniformMatrix3dv == nil {
		return errors.New("glProgramUniformMatrix3dv")
	}
	gpProgramUniformMatrix3fv = (C.GPPROGRAMUNIFORMMATRIX3FV)(getProcAddr("glProgramUniformMatrix3fv"))
	if gpProgramUniformMatrix3fv == nil {
		return errors.New("glProgramUniformMatrix3fv")
	}
	gpProgramUniformMatrix3fvEXT = (C.GPPROGRAMUNIFORMMATRIX3FVEXT)(getProcAddr("glProgramUniformMatrix3fvEXT"))
	gpProgramUniformMatrix3x2dv = (C.GPPROGRAMUNIFORMMATRIX3X2DV)(getProcAddr("glProgramUniformMatrix3x2dv"))
	if gpProgramUniformMatrix3x2dv == nil {
		return errors.New("glProgramUniformMatrix3x2dv")
	}
	gpProgramUniformMatrix3x2fv = (C.GPPROGRAMUNIFORMMATRIX3X2FV)(getProcAddr("glProgramUniformMatrix3x2fv"))
	if gpProgramUniformMatrix3x2fv == nil {
		return errors.New("glProgramUniformMatrix3x2fv")
	}
	gpProgramUniformMatrix3x2fvEXT = (C.GPPROGRAMUNIFORMMATRIX3X2FVEXT)(getProcAddr("glProgramUniformMatrix3x2fvEXT"))
	gpProgramUniformMatrix3x4dv = (C.GPPROGRAMUNIFORMMATRIX3X4DV)(getProcAddr("glProgramUniformMatrix3x4dv"))
	if gpProgramUniformMatrix3x4dv == nil {
		return errors.New("glProgramUniformMatrix3x4dv")
	}
	gpProgramUniformMatrix3x4fv = (C.GPPROGRAMUNIFORMMATRIX3X4FV)(getProcAddr("glProgramUniformMatrix3x4fv"))
	if gpProgramUniformMatrix3x4fv == nil {
		return errors.New("glProgramUniformMatrix3x4fv")
	}
	gpProgramUniformMatrix3x4fvEXT = (C.GPPROGRAMUNIFORMMATRIX3X4FVEXT)(getProcAddr("glProgramUniformMatrix3x4fvEXT"))
	gpProgramUniformMatrix4dv = (C.GPPROGRAMUNIFORMMATRIX4DV)(getProcAddr("glProgramUniformMatrix4dv"))
	if gpProgramUniformMatrix4dv == nil {
		return errors.New("glProgramUniformMatrix4dv")
	}
	gpProgramUniformMatrix4fv = (C.GPPROGRAMUNIFORMMATRIX4FV)(getProcAddr("glProgramUniformMatrix4fv"))
	if gpProgramUniformMatrix4fv == nil {
		return errors.New("glProgramUniformMatrix4fv")
	}
	gpProgramUniformMatrix4fvEXT = (C.GPPROGRAMUNIFORMMATRIX4FVEXT)(getProcAddr("glProgramUniformMatrix4fvEXT"))
	gpProgramUniformMatrix4x2dv = (C.GPPROGRAMUNIFORMMATRIX4X2DV)(getProcAddr("glProgramUniformMatrix4x2dv"))
	if gpProgramUniformMatrix4x2dv == nil {
		return errors.New("glProgramUniformMatrix4x2dv")
	}
	gpProgramUniformMatrix4x2fv = (C.GPPROGRAMUNIFORMMATRIX4X2FV)(getProcAddr("glProgramUniformMatrix4x2fv"))
	if gpProgramUniformMatrix4x2fv == nil {
		return errors.New("glProgramUniformMatrix4x2fv")
	}
	gpProgramUniformMatrix4x2fvEXT = (C.GPPROGRAMUNIFORMMATRIX4X2FVEXT)(getProcAddr("glProgramUniformMatrix4x2fvEXT"))
	gpProgramUniformMatrix4x3dv = (C.GPPROGRAMUNIFORMMATRIX4X3DV)(getProcAddr("glProgramUniformMatrix4x3dv"))
	if gpProgramUniformMatrix4x3dv == nil {
		return errors.New("glProgramUniformMatrix4x3dv")
	}
	gpProgramUniformMatrix4x3fv = (C.GPPROGRAMUNIFORMMATRIX4X3FV)(getProcAddr("glProgramUniformMatrix4x3fv"))
	if gpProgramUniformMatrix4x3fv == nil {
		return errors.New("glProgramUniformMatrix4x3fv")
	}
	gpProgramUniformMatrix4x3fvEXT = (C.GPPROGRAMUNIFORMMATRIX4X3FVEXT)(getProcAddr("glProgramUniformMatrix4x3fvEXT"))
	gpProvokingVertex = (C.GPPROVOKINGVERTEX)(getProcAddr("glProvokingVertex"))
	if gpProvokingVertex == nil {
		return errors.New("glProvokingVertex")
	}
	gpPushDebugGroup = (C.GPPUSHDEBUGGROUP)(getProcAddr("glPushDebugGroup"))
	gpPushDebugGroupKHR = (C.GPPUSHDEBUGGROUPKHR)(getProcAddr("glPushDebugGroupKHR"))
	gpPushGroupMarkerEXT = (C.GPPUSHGROUPMARKEREXT)(getProcAddr("glPushGroupMarkerEXT"))
	gpQueryCounter = (C.GPQUERYCOUNTER)(getProcAddr("glQueryCounter"))
	if gpQueryCounter == nil {
		return errors.New("glQueryCounter")
	}
	gpQueryMatrixxOES = (C.GPQUERYMATRIXXOES)(getProcAddr("glQueryMatrixxOES"))
	gpRasterPos2xOES = (C.GPRASTERPOS2XOES)(getProcAddr("glRasterPos2xOES"))
	gpRasterPos2xvOES = (C.GPRASTERPOS2XVOES)(getProcAddr("glRasterPos2xvOES"))
	gpRasterPos3xOES = (C.GPRASTERPOS3XOES)(getProcAddr("glRasterPos3xOES"))
	gpRasterPos3xvOES = (C.GPRASTERPOS3XVOES)(getProcAddr("glRasterPos3xvOES"))
	gpRasterPos4xOES = (C.GPRASTERPOS4XOES)(getProcAddr("glRasterPos4xOES"))
	gpRasterPos4xvOES = (C.GPRASTERPOS4XVOES)(getProcAddr("glRasterPos4xvOES"))
	gpReadBuffer = (C.GPREADBUFFER)(getProcAddr("glReadBuffer"))
	if gpReadBuffer == nil {
		return errors.New("glReadBuffer")
	}
	gpReadPixels = (C.GPREADPIXELS)(getProcAddr("glReadPixels"))
	if gpReadPixels == nil {
		return errors.New("glReadPixels")
	}
	gpReadnPixels = (C.GPREADNPIXELS)(getProcAddr("glReadnPixels"))
	gpReadnPixelsARB = (C.GPREADNPIXELSARB)(getProcAddr("glReadnPixelsARB"))
	gpReadnPixelsKHR = (C.GPREADNPIXELSKHR)(getProcAddr("glReadnPixelsKHR"))
	gpRectxOES = (C.GPRECTXOES)(getProcAddr("glRectxOES"))
	gpRectxvOES = (C.GPRECTXVOES)(getProcAddr("glRectxvOES"))
	gpReleaseShaderCompiler = (C.GPRELEASESHADERCOMPILER)(getProcAddr("glReleaseShaderCompiler"))
	if gpReleaseShaderCompiler == nil {
		return errors.New("glReleaseShaderCompiler")
	}
	gpRenderbufferStorage = (C.GPRENDERBUFFERSTORAGE)(getProcAddr("glRenderbufferStorage"))
	if gpRenderbufferStorage == nil {
		return errors.New("glRenderbufferStorage")
	}
	gpRenderbufferStorageMultisample = (C.GPRENDERBUFFERSTORAGEMULTISAMPLE)(getProcAddr("glRenderbufferStorageMultisample"))
	if gpRenderbufferStorageMultisample == nil {
		return errors.New("glRenderbufferStorageMultisample")
	}
	gpResumeTransformFeedback = (C.GPRESUMETRANSFORMFEEDBACK)(getProcAddr("glResumeTransformFeedback"))
	if gpResumeTransformFeedback == nil {
		return errors.New("glResumeTransformFeedback")
	}
	gpRotatexOES = (C.GPROTATEXOES)(getProcAddr("glRotatexOES"))
	gpSampleCoverage = (C.GPSAMPLECOVERAGE)(getProcAddr("glSampleCoverage"))
	if gpSampleCoverage == nil {
		return errors.New("glSampleCoverage")
	}
	gpSampleCoverageOES = (C.GPSAMPLECOVERAGEOES)(getProcAddr("glSampleCoverageOES"))
	gpSampleCoveragexOES = (C.GPSAMPLECOVERAGEXOES)(getProcAddr("glSampleCoveragexOES"))
	gpSampleMaski = (C.GPSAMPLEMASKI)(getProcAddr("glSampleMaski"))
	if gpSampleMaski == nil {
		return errors.New("glSampleMaski")
	}
	gpSamplerParameterIiv = (C.GPSAMPLERPARAMETERIIV)(getProcAddr("glSamplerParameterIiv"))
	if gpSamplerParameterIiv == nil {
		return errors.New("glSamplerParameterIiv")
	}
	gpSamplerParameterIuiv = (C.GPSAMPLERPARAMETERIUIV)(getProcAddr("glSamplerParameterIuiv"))
	if gpSamplerParameterIuiv == nil {
		return errors.New("glSamplerParameterIuiv")
	}
	gpSamplerParameterf = (C.GPSAMPLERPARAMETERF)(getProcAddr("glSamplerParameterf"))
	if gpSamplerParameterf == nil {
		return errors.New("glSamplerParameterf")
	}
	gpSamplerParameterfv = (C.GPSAMPLERPARAMETERFV)(getProcAddr("glSamplerParameterfv"))
	if gpSamplerParameterfv == nil {
		return errors.New("glSamplerParameterfv")
	}
	gpSamplerParameteri = (C.GPSAMPLERPARAMETERI)(getProcAddr("glSamplerParameteri"))
	if gpSamplerParameteri == nil {
		return errors.New("glSamplerParameteri")
	}
	gpSamplerParameteriv = (C.GPSAMPLERPARAMETERIV)(getProcAddr("glSamplerParameteriv"))
	if gpSamplerParameteriv == nil {
		return errors.New("glSamplerParameteriv")
	}
	gpScalexOES = (C.GPSCALEXOES)(getProcAddr("glScalexOES"))
	gpScissor = (C.GPSCISSOR)(getProcAddr("glScissor"))
	if gpScissor == nil {
		return errors.New("glScissor")
	}
	gpScissorArrayv = (C.GPSCISSORARRAYV)(getProcAddr("glScissorArrayv"))
	if gpScissorArrayv == nil {
		return errors.New("glScissorArrayv")
	}
	gpScissorIndexed = (C.GPSCISSORINDEXED)(getProcAddr("glScissorIndexed"))
	if gpScissorIndexed == nil {
		return errors.New("glScissorIndexed")
	}
	gpScissorIndexedv = (C.GPSCISSORINDEXEDV)(getProcAddr("glScissorIndexedv"))
	if gpScissorIndexedv == nil {
		return errors.New("glScissorIndexedv")
	}
	gpSelectPerfMonitorCountersAMD = (C.GPSELECTPERFMONITORCOUNTERSAMD)(getProcAddr("glSelectPerfMonitorCountersAMD"))
	gpSetFenceNV = (C.GPSETFENCENV)(getProcAddr("glSetFenceNV"))
	gpShaderBinary = (C.GPSHADERBINARY)(getProcAddr("glShaderBinary"))
	if gpShaderBinary == nil {
		return errors.New("glShaderBinary")
	}
	gpShaderSource = (C.GPSHADERSOURCE)(getProcAddr("glShaderSource"))
	if gpShaderSource == nil {
		return errors.New("glShaderSource")
	}
	gpShaderStorageBlockBinding = (C.GPSHADERSTORAGEBLOCKBINDING)(getProcAddr("glShaderStorageBlockBinding"))
	gpStencilFunc = (C.GPSTENCILFUNC)(getProcAddr("glStencilFunc"))
	if gpStencilFunc == nil {
		return errors.New("glStencilFunc")
	}
	gpStencilFuncSeparate = (C.GPSTENCILFUNCSEPARATE)(getProcAddr("glStencilFuncSeparate"))
	if gpStencilFuncSeparate == nil {
		return errors.New("glStencilFuncSeparate")
	}
	gpStencilMask = (C.GPSTENCILMASK)(getProcAddr("glStencilMask"))
	if gpStencilMask == nil {
		return errors.New("glStencilMask")
	}
	gpStencilMaskSeparate = (C.GPSTENCILMASKSEPARATE)(getProcAddr("glStencilMaskSeparate"))
	if gpStencilMaskSeparate == nil {
		return errors.New("glStencilMaskSeparate")
	}
	gpStencilOp = (C.GPSTENCILOP)(getProcAddr("glStencilOp"))
	if gpStencilOp == nil {
		return errors.New("glStencilOp")
	}
	gpStencilOpSeparate = (C.GPSTENCILOPSEPARATE)(getProcAddr("glStencilOpSeparate"))
	if gpStencilOpSeparate == nil {
		return errors.New("glStencilOpSeparate")
	}
	gpTestFenceNV = (C.GPTESTFENCENV)(getProcAddr("glTestFenceNV"))
	gpTexBuffer = (C.GPTEXBUFFER)(getProcAddr("glTexBuffer"))
	if gpTexBuffer == nil {
		return errors.New("glTexBuffer")
	}
	gpTexBufferRange = (C.GPTEXBUFFERRANGE)(getProcAddr("glTexBufferRange"))
	gpTexCoord1bOES = (C.GPTEXCOORD1BOES)(getProcAddr("glTexCoord1bOES"))
	gpTexCoord1bvOES = (C.GPTEXCOORD1BVOES)(getProcAddr("glTexCoord1bvOES"))
	gpTexCoord1xOES = (C.GPTEXCOORD1XOES)(getProcAddr("glTexCoord1xOES"))
	gpTexCoord1xvOES = (C.GPTEXCOORD1XVOES)(getProcAddr("glTexCoord1xvOES"))
	gpTexCoord2bOES = (C.GPTEXCOORD2BOES)(getProcAddr("glTexCoord2bOES"))
	gpTexCoord2bvOES = (C.GPTEXCOORD2BVOES)(getProcAddr("glTexCoord2bvOES"))
	gpTexCoord2xOES = (C.GPTEXCOORD2XOES)(getProcAddr("glTexCoord2xOES"))
	gpTexCoord2xvOES = (C.GPTEXCOORD2XVOES)(getProcAddr("glTexCoord2xvOES"))
	gpTexCoord3bOES = (C.GPTEXCOORD3BOES)(getProcAddr("glTexCoord3bOES"))
	gpTexCoord3bvOES = (C.GPTEXCOORD3BVOES)(getProcAddr("glTexCoord3bvOES"))
	gpTexCoord3xOES = (C.GPTEXCOORD3XOES)(getProcAddr("glTexCoord3xOES"))
	gpTexCoord3xvOES = (C.GPTEXCOORD3XVOES)(getProcAddr("glTexCoord3xvOES"))
	gpTexCoord4bOES = (C.GPTEXCOORD4BOES)(getProcAddr("glTexCoord4bOES"))
	gpTexCoord4bvOES = (C.GPTEXCOORD4BVOES)(getProcAddr("glTexCoord4bvOES"))
	gpTexCoord4xOES = (C.GPTEXCOORD4XOES)(getProcAddr("glTexCoord4xOES"))
	gpTexCoord4xvOES = (C.GPTEXCOORD4XVOES)(getProcAddr("glTexCoord4xvOES"))
	gpTexEnvxOES = (C.GPTEXENVXOES)(getProcAddr("glTexEnvxOES"))
	gpTexEnvxvOES = (C.GPTEXENVXVOES)(getProcAddr("glTexEnvxvOES"))
	gpTexGenxOES = (C.GPTEXGENXOES)(getProcAddr("glTexGenxOES"))
	gpTexGenxvOES = (C.GPTEXGENXVOES)(getProcAddr("glTexGenxvOES"))
	gpTexImage1D = (C.GPTEXIMAGE1D)(getProcAddr("glTexImage1D"))
	if gpTexImage1D == nil {
		return errors.New("glTexImage1D")
	}
	gpTexImage2D = (C.GPTEXIMAGE2D)(getProcAddr("glTexImage2D"))
	if gpTexImage2D == nil {
		return errors.New("glTexImage2D")
	}
	gpTexImage2DMultisample = (C.GPTEXIMAGE2DMULTISAMPLE)(getProcAddr("glTexImage2DMultisample"))
	if gpTexImage2DMultisample == nil {
		return errors.New("glTexImage2DMultisample")
	}
	gpTexImage3D = (C.GPTEXIMAGE3D)(getProcAddr("glTexImage3D"))
	if gpTexImage3D == nil {
		return errors.New("glTexImage3D")
	}
	gpTexImage3DMultisample = (C.GPTEXIMAGE3DMULTISAMPLE)(getProcAddr("glTexImage3DMultisample"))
	if gpTexImage3DMultisample == nil {
		return errors.New("glTexImage3DMultisample")
	}
	gpTexPageCommitmentARB = (C.GPTEXPAGECOMMITMENTARB)(getProcAddr("glTexPageCommitmentARB"))
	gpTexParameterIiv = (C.GPTEXPARAMETERIIV)(getProcAddr("glTexParameterIiv"))
	if gpTexParameterIiv == nil {
		return errors.New("glTexParameterIiv")
	}
	gpTexParameterIuiv = (C.GPTEXPARAMETERIUIV)(getProcAddr("glTexParameterIuiv"))
	if gpTexParameterIuiv == nil {
		return errors.New("glTexParameterIuiv")
	}
	gpTexParameterf = (C.GPTEXPARAMETERF)(getProcAddr("glTexParameterf"))
	if gpTexParameterf == nil {
		return errors.New("glTexParameterf")
	}
	gpTexParameterfv = (C.GPTEXPARAMETERFV)(getProcAddr("glTexParameterfv"))
	if gpTexParameterfv == nil {
		return errors.New("glTexParameterfv")
	}
	gpTexParameteri = (C.GPTEXPARAMETERI)(getProcAddr("glTexParameteri"))
	if gpTexParameteri == nil {
		return errors.New("glTexParameteri")
	}
	gpTexParameteriv = (C.GPTEXPARAMETERIV)(getProcAddr("glTexParameteriv"))
	if gpTexParameteriv == nil {
		return errors.New("glTexParameteriv")
	}
	gpTexParameterxOES = (C.GPTEXPARAMETERXOES)(getProcAddr("glTexParameterxOES"))
	gpTexParameterxvOES = (C.GPTEXPARAMETERXVOES)(getProcAddr("glTexParameterxvOES"))
	gpTexStorage1D = (C.GPTEXSTORAGE1D)(getProcAddr("glTexStorage1D"))
	gpTexStorage2D = (C.GPTEXSTORAGE2D)(getProcAddr("glTexStorage2D"))
	gpTexStorage2DMultisample = (C.GPTEXSTORAGE2DMULTISAMPLE)(getProcAddr("glTexStorage2DMultisample"))
	gpTexStorage3D = (C.GPTEXSTORAGE3D)(getProcAddr("glTexStorage3D"))
	gpTexStorage3DMultisample = (C.GPTEXSTORAGE3DMULTISAMPLE)(getProcAddr("glTexStorage3DMultisample"))
	gpTexSubImage1D = (C.GPTEXSUBIMAGE1D)(getProcAddr("glTexSubImage1D"))
	if gpTexSubImage1D == nil {
		return errors.New("glTexSubImage1D")
	}
	gpTexSubImage2D = (C.GPTEXSUBIMAGE2D)(getProcAddr("glTexSubImage2D"))
	if gpTexSubImage2D == nil {
		return errors.New("glTexSubImage2D")
	}
	gpTexSubImage3D = (C.GPTEXSUBIMAGE3D)(getProcAddr("glTexSubImage3D"))
	if gpTexSubImage3D == nil {
		return errors.New("glTexSubImage3D")
	}
	gpTextureBarrier = (C.GPTEXTUREBARRIER)(getProcAddr("glTextureBarrier"))
	gpTextureBuffer = (C.GPTEXTUREBUFFER)(getProcAddr("glTextureBuffer"))
	gpTextureBufferRange = (C.GPTEXTUREBUFFERRANGE)(getProcAddr("glTextureBufferRange"))
	gpTextureParameterIiv = (C.GPTEXTUREPARAMETERIIV)(getProcAddr("glTextureParameterIiv"))
	gpTextureParameterIuiv = (C.GPTEXTUREPARAMETERIUIV)(getProcAddr("glTextureParameterIuiv"))
	gpTextureParameterf = (C.GPTEXTUREPARAMETERF)(getProcAddr("glTextureParameterf"))
	gpTextureParameterfv = (C.GPTEXTUREPARAMETERFV)(getProcAddr("glTextureParameterfv"))
	gpTextureParameteri = (C.GPTEXTUREPARAMETERI)(getProcAddr("glTextureParameteri"))
	gpTextureParameteriv = (C.GPTEXTUREPARAMETERIV)(getProcAddr("glTextureParameteriv"))
	gpTextureStorage1D = (C.GPTEXTURESTORAGE1D)(getProcAddr("glTextureStorage1D"))
	gpTextureStorage2D = (C.GPTEXTURESTORAGE2D)(getProcAddr("glTextureStorage2D"))
	gpTextureStorage2DMultisample = (C.GPTEXTURESTORAGE2DMULTISAMPLE)(getProcAddr("glTextureStorage2DMultisample"))
	gpTextureStorage3D = (C.GPTEXTURESTORAGE3D)(getProcAddr("glTextureStorage3D"))
	gpTextureStorage3DMultisample = (C.GPTEXTURESTORAGE3DMULTISAMPLE)(getProcAddr("glTextureStorage3DMultisample"))
	gpTextureSubImage1D = (C.GPTEXTURESUBIMAGE1D)(getProcAddr("glTextureSubImage1D"))
	gpTextureSubImage2D = (C.GPTEXTURESUBIMAGE2D)(getProcAddr("glTextureSubImage2D"))
	gpTextureSubImage3D = (C.GPTEXTURESUBIMAGE3D)(getProcAddr("glTextureSubImage3D"))
	gpTextureView = (C.GPTEXTUREVIEW)(getProcAddr("glTextureView"))
	gpTransformFeedbackBufferBase = (C.GPTRANSFORMFEEDBACKBUFFERBASE)(getProcAddr("glTransformFeedbackBufferBase"))
	gpTransformFeedbackBufferRange = (C.GPTRANSFORMFEEDBACKBUFFERRANGE)(getProcAddr("glTransformFeedbackBufferRange"))
	gpTransformFeedbackVaryings = (C.GPTRANSFORMFEEDBACKVARYINGS)(getProcAddr("glTransformFeedbackVaryings"))
	if gpTransformFeedbackVaryings == nil {
		return errors.New("glTransformFeedbackVaryings")
	}
	gpTranslatexOES = (C.GPTRANSLATEXOES)(getProcAddr("glTranslatexOES"))
	gpUniform1d = (C.GPUNIFORM1D)(getProcAddr("glUniform1d"))
	if gpUniform1d == nil {
		return errors.New("glUniform1d")
	}
	gpUniform1dv = (C.GPUNIFORM1DV)(getProcAddr("glUniform1dv"))
	if gpUniform1dv == nil {
		return errors.New("glUniform1dv")
	}
	gpUniform1f = (C.GPUNIFORM1F)(getProcAddr("glUniform1f"))
	if gpUniform1f == nil {
		return errors.New("glUniform1f")
	}
	gpUniform1fv = (C.GPUNIFORM1FV)(getProcAddr("glUniform1fv"))
	if gpUniform1fv == nil {
		return errors.New("glUniform1fv")
	}
	gpUniform1i = (C.GPUNIFORM1I)(getProcAddr("glUniform1i"))
	if gpUniform1i == nil {
		return errors.New("glUniform1i")
	}
	gpUniform1iv = (C.GPUNIFORM1IV)(getProcAddr("glUniform1iv"))
	if gpUniform1iv == nil {
		return errors.New("glUniform1iv")
	}
	gpUniform1ui = (C.GPUNIFORM1UI)(getProcAddr("glUniform1ui"))
	if gpUniform1ui == nil {
		return errors.New("glUniform1ui")
	}
	gpUniform1uiv = (C.GPUNIFORM1UIV)(getProcAddr("glUniform1uiv"))
	if gpUniform1uiv == nil {
		return errors.New("glUniform1uiv")
	}
	gpUniform2d = (C.GPUNIFORM2D)(getProcAddr("glUniform2d"))
	if gpUniform2d == nil {
		return errors.New("glUniform2d")
	}
	gpUniform2dv = (C.GPUNIFORM2DV)(getProcAddr("glUniform2dv"))
	if gpUniform2dv == nil {
		return errors.New("glUniform2dv")
	}
	gpUniform2f = (C.GPUNIFORM2F)(getProcAddr("glUniform2f"))
	if gpUniform2f == nil {
		return errors.New("glUniform2f")
	}
	gpUniform2fv = (C.GPUNIFORM2FV)(getProcAddr("glUniform2fv"))
	if gpUniform2fv == nil {
		return errors.New("glUniform2fv")
	}
	gpUniform2i = (C.GPUNIFORM2I)(getProcAddr("glUniform2i"))
	if gpUniform2i == nil {
		return errors.New("glUniform2i")
	}
	gpUniform2iv = (C.GPUNIFORM2IV)(getProcAddr("glUniform2iv"))
	if gpUniform2iv == nil {
		return errors.New("glUniform2iv")
	}
	gpUniform2ui = (C.GPUNIFORM2UI)(getProcAddr("glUniform2ui"))
	if gpUniform2ui == nil {
		return errors.New("glUniform2ui")
	}
	gpUniform2uiv = (C.GPUNIFORM2UIV)(getProcAddr("glUniform2uiv"))
	if gpUniform2uiv == nil {
		return errors.New("glUniform2uiv")
	}
	gpUniform3d = (C.GPUNIFORM3D)(getProcAddr("glUniform3d"))
	if gpUniform3d == nil {
		return errors.New("glUniform3d")
	}
	gpUniform3dv = (C.GPUNIFORM3DV)(getProcAddr("glUniform3dv"))
	if gpUniform3dv == nil {
		return errors.New("glUniform3dv")
	}
	gpUniform3f = (C.GPUNIFORM3F)(getProcAddr("glUniform3f"))
	if gpUniform3f == nil {
		return errors.New("glUniform3f")
	}
	gpUniform3fv = (C.GPUNIFORM3FV)(getProcAddr("glUniform3fv"))
	if gpUniform3fv == nil {
		return errors.New("glUniform3fv")
	}
	gpUniform3i = (C.GPUNIFORM3I)(getProcAddr("glUniform3i"))
	if gpUniform3i == nil {
		return errors.New("glUniform3i")
	}
	gpUniform3iv = (C.GPUNIFORM3IV)(getProcAddr("glUniform3iv"))
	if gpUniform3iv == nil {
		return errors.New("glUniform3iv")
	}
	gpUniform3ui = (C.GPUNIFORM3UI)(getProcAddr("glUniform3ui"))
	if gpUniform3ui == nil {
		return errors.New("glUniform3ui")
	}
	gpUniform3uiv = (C.GPUNIFORM3UIV)(getProcAddr("glUniform3uiv"))
	if gpUniform3uiv == nil {
		return errors.New("glUniform3uiv")
	}
	gpUniform4d = (C.GPUNIFORM4D)(getProcAddr("glUniform4d"))
	if gpUniform4d == nil {
		return errors.New("glUniform4d")
	}
	gpUniform4dv = (C.GPUNIFORM4DV)(getProcAddr("glUniform4dv"))
	if gpUniform4dv == nil {
		return errors.New("glUniform4dv")
	}
	gpUniform4f = (C.GPUNIFORM4F)(getProcAddr("glUniform4f"))
	if gpUniform4f == nil {
		return errors.New("glUniform4f")
	}
	gpUniform4fv = (C.GPUNIFORM4FV)(getProcAddr("glUniform4fv"))
	if gpUniform4fv == nil {
		return errors.New("glUniform4fv")
	}
	gpUniform4i = (C.GPUNIFORM4I)(getProcAddr("glUniform4i"))
	if gpUniform4i == nil {
		return errors.New("glUniform4i")
	}
	gpUniform4iv = (C.GPUNIFORM4IV)(getProcAddr("glUniform4iv"))
	if gpUniform4iv == nil {
		return errors.New("glUniform4iv")
	}
	gpUniform4ui = (C.GPUNIFORM4UI)(getProcAddr("glUniform4ui"))
	if gpUniform4ui == nil {
		return errors.New("glUniform4ui")
	}
	gpUniform4uiv = (C.GPUNIFORM4UIV)(getProcAddr("glUniform4uiv"))
	if gpUniform4uiv == nil {
		return errors.New("glUniform4uiv")
	}
	gpUniformBlockBinding = (C.GPUNIFORMBLOCKBINDING)(getProcAddr("glUniformBlockBinding"))
	if gpUniformBlockBinding == nil {
		return errors.New("glUniformBlockBinding")
	}
	gpUniformHandleui64ARB = (C.GPUNIFORMHANDLEUI64ARB)(getProcAddr("glUniformHandleui64ARB"))
	gpUniformHandleui64vARB = (C.GPUNIFORMHANDLEUI64VARB)(getProcAddr("glUniformHandleui64vARB"))
	gpUniformMatrix2dv = (C.GPUNIFORMMATRIX2DV)(getProcAddr("glUniformMatrix2dv"))
	if gpUniformMatrix2dv == nil {
		return errors.New("glUniformMatrix2dv")
	}
	gpUniformMatrix2fv = (C.GPUNIFORMMATRIX2FV)(getProcAddr("glUniformMatrix2fv"))
	if gpUniformMatrix2fv == nil {
		return errors.New("glUniformMatrix2fv")
	}
	gpUniformMatrix2x3dv = (C.GPUNIFORMMATRIX2X3DV)(getProcAddr("glUniformMatrix2x3dv"))
	if gpUniformMatrix2x3dv == nil {
		return errors.New("glUniformMatrix2x3dv")
	}
	gpUniformMatrix2x3fv = (C.GPUNIFORMMATRIX2X3FV)(getProcAddr("glUniformMatrix2x3fv"))
	if gpUniformMatrix2x3fv == nil {
		return errors.New("glUniformMatrix2x3fv")
	}
	gpUniformMatrix2x4dv = (C.GPUNIFORMMATRIX2X4DV)(getProcAddr("glUniformMatrix2x4dv"))
	if gpUniformMatrix2x4dv == nil {
		return errors.New("glUniformMatrix2x4dv")
	}
	gpUniformMatrix2x4fv = (C.GPUNIFORMMATRIX2X4FV)(getProcAddr("glUniformMatrix2x4fv"))
	if gpUniformMatrix2x4fv == nil {
		return errors.New("glUniformMatrix2x4fv")
	}
	gpUniformMatrix3dv = (C.GPUNIFORMMATRIX3DV)(getProcAddr("glUniformMatrix3dv"))
	if gpUniformMatrix3dv == nil {
		return errors.New("glUniformMatrix3dv")
	}
	gpUniformMatrix3fv = (C.GPUNIFORMMATRIX3FV)(getProcAddr("glUniformMatrix3fv"))
	if gpUniformMatrix3fv == nil {
		return errors.New("glUniformMatrix3fv")
	}
	gpUniformMatrix3x2dv = (C.GPUNIFORMMATRIX3X2DV)(getProcAddr("glUniformMatrix3x2dv"))
	if gpUniformMatrix3x2dv == nil {
		return errors.New("glUniformMatrix3x2dv")
	}
	gpUniformMatrix3x2fv = (C.GPUNIFORMMATRIX3X2FV)(getProcAddr("glUniformMatrix3x2fv"))
	if gpUniformMatrix3x2fv == nil {
		return errors.New("glUniformMatrix3x2fv")
	}
	gpUniformMatrix3x4dv = (C.GPUNIFORMMATRIX3X4DV)(getProcAddr("glUniformMatrix3x4dv"))
	if gpUniformMatrix3x4dv == nil {
		return errors.New("glUniformMatrix3x4dv")
	}
	gpUniformMatrix3x4fv = (C.GPUNIFORMMATRIX3X4FV)(getProcAddr("glUniformMatrix3x4fv"))
	if gpUniformMatrix3x4fv == nil {
		return errors.New("glUniformMatrix3x4fv")
	}
	gpUniformMatrix4dv = (C.GPUNIFORMMATRIX4DV)(getProcAddr("glUniformMatrix4dv"))
	if gpUniformMatrix4dv == nil {
		return errors.New("glUniformMatrix4dv")
	}
	gpUniformMatrix4fv = (C.GPUNIFORMMATRIX4FV)(getProcAddr("glUniformMatrix4fv"))
	if gpUniformMatrix4fv == nil {
		return errors.New("glUniformMatrix4fv")
	}
	gpUniformMatrix4x2dv = (C.GPUNIFORMMATRIX4X2DV)(getProcAddr("glUniformMatrix4x2dv"))
	if gpUniformMatrix4x2dv == nil {
		return errors.New("glUniformMatrix4x2dv")
	}
	gpUniformMatrix4x2fv = (C.GPUNIFORMMATRIX4X2FV)(getProcAddr("glUniformMatrix4x2fv"))
	if gpUniformMatrix4x2fv == nil {
		return errors.New("glUniformMatrix4x2fv")
	}
	gpUniformMatrix4x3dv = (C.GPUNIFORMMATRIX4X3DV)(getProcAddr("glUniformMatrix4x3dv"))
	if gpUniformMatrix4x3dv == nil {
		return errors.New("glUniformMatrix4x3dv")
	}
	gpUniformMatrix4x3fv = (C.GPUNIFORMMATRIX4X3FV)(getProcAddr("glUniformMatrix4x3fv"))
	if gpUniformMatrix4x3fv == nil {
		return errors.New("glUniformMatrix4x3fv")
	}
	gpUniformSubroutinesuiv = (C.GPUNIFORMSUBROUTINESUIV)(getProcAddr("glUniformSubroutinesuiv"))
	if gpUniformSubroutinesuiv == nil {
		return errors.New("glUniformSubroutinesuiv")
	}
	gpUnmapBuffer = (C.GPUNMAPBUFFER)(getProcAddr("glUnmapBuffer"))
	if gpUnmapBuffer == nil {
		return errors.New("glUnmapBuffer")
	}
	gpUnmapNamedBuffer = (C.GPUNMAPNAMEDBUFFER)(getProcAddr("glUnmapNamedBuffer"))
	gpUseProgram = (C.GPUSEPROGRAM)(getProcAddr("glUseProgram"))
	if gpUseProgram == nil {
		return errors.New("glUseProgram")
	}
	gpUseProgramStages = (C.GPUSEPROGRAMSTAGES)(getProcAddr("glUseProgramStages"))
	if gpUseProgramStages == nil {
		return errors.New("glUseProgramStages")
	}
	gpUseProgramStagesEXT = (C.GPUSEPROGRAMSTAGESEXT)(getProcAddr("glUseProgramStagesEXT"))
	gpUseShaderProgramEXT = (C.GPUSESHADERPROGRAMEXT)(getProcAddr("glUseShaderProgramEXT"))
	gpValidateProgram = (C.GPVALIDATEPROGRAM)(getProcAddr("glValidateProgram"))
	if gpValidateProgram == nil {
		return errors.New("glValidateProgram")
	}
	gpValidateProgramPipeline = (C.GPVALIDATEPROGRAMPIPELINE)(getProcAddr("glValidateProgramPipeline"))
	if gpValidateProgramPipeline == nil {
		return errors.New("glValidateProgramPipeline")
	}
	gpValidateProgramPipelineEXT = (C.GPVALIDATEPROGRAMPIPELINEEXT)(getProcAddr("glValidateProgramPipelineEXT"))
	gpVertex2bOES = (C.GPVERTEX2BOES)(getProcAddr("glVertex2bOES"))
	gpVertex2bvOES = (C.GPVERTEX2BVOES)(getProcAddr("glVertex2bvOES"))
	gpVertex2xOES = (C.GPVERTEX2XOES)(getProcAddr("glVertex2xOES"))
	gpVertex2xvOES = (C.GPVERTEX2XVOES)(getProcAddr("glVertex2xvOES"))
	gpVertex3bOES = (C.GPVERTEX3BOES)(getProcAddr("glVertex3bOES"))
	gpVertex3bvOES = (C.GPVERTEX3BVOES)(getProcAddr("glVertex3bvOES"))
	gpVertex3xOES = (C.GPVERTEX3XOES)(getProcAddr("glVertex3xOES"))
	gpVertex3xvOES = (C.GPVERTEX3XVOES)(getProcAddr("glVertex3xvOES"))
	gpVertex4bOES = (C.GPVERTEX4BOES)(getProcAddr("glVertex4bOES"))
	gpVertex4bvOES = (C.GPVERTEX4BVOES)(getProcAddr("glVertex4bvOES"))
	gpVertex4xOES = (C.GPVERTEX4XOES)(getProcAddr("glVertex4xOES"))
	gpVertex4xvOES = (C.GPVERTEX4XVOES)(getProcAddr("glVertex4xvOES"))
	gpVertexArrayAttribBinding = (C.GPVERTEXARRAYATTRIBBINDING)(getProcAddr("glVertexArrayAttribBinding"))
	gpVertexArrayAttribFormat = (C.GPVERTEXARRAYATTRIBFORMAT)(getProcAddr("glVertexArrayAttribFormat"))
	gpVertexArrayAttribIFormat = (C.GPVERTEXARRAYATTRIBIFORMAT)(getProcAddr("glVertexArrayAttribIFormat"))
	gpVertexArrayAttribLFormat = (C.GPVERTEXARRAYATTRIBLFORMAT)(getProcAddr("glVertexArrayAttribLFormat"))
	gpVertexArrayBindingDivisor = (C.GPVERTEXARRAYBINDINGDIVISOR)(getProcAddr("glVertexArrayBindingDivisor"))
	gpVertexArrayElementBuffer = (C.GPVERTEXARRAYELEMENTBUFFER)(getProcAddr("glVertexArrayElementBuffer"))
	gpVertexArrayVertexBuffer = (C.GPVERTEXARRAYVERTEXBUFFER)(getProcAddr("glVertexArrayVertexBuffer"))
	gpVertexArrayVertexBuffers = (C.GPVERTEXARRAYVERTEXBUFFERS)(getProcAddr("glVertexArrayVertexBuffers"))
	gpVertexAttrib1d = (C.GPVERTEXATTRIB1D)(getProcAddr("glVertexAttrib1d"))
	if gpVertexAttrib1d == nil {
		return errors.New("glVertexAttrib1d")
	}
	gpVertexAttrib1dv = (C.GPVERTEXATTRIB1DV)(getProcAddr("glVertexAttrib1dv"))
	if gpVertexAttrib1dv == nil {
		return errors.New("glVertexAttrib1dv")
	}
	gpVertexAttrib1f = (C.GPVERTEXATTRIB1F)(getProcAddr("glVertexAttrib1f"))
	if gpVertexAttrib1f == nil {
		return errors.New("glVertexAttrib1f")
	}
	gpVertexAttrib1fv = (C.GPVERTEXATTRIB1FV)(getProcAddr("glVertexAttrib1fv"))
	if gpVertexAttrib1fv == nil {
		return errors.New("glVertexAttrib1fv")
	}
	gpVertexAttrib1s = (C.GPVERTEXATTRIB1S)(getProcAddr("glVertexAttrib1s"))
	if gpVertexAttrib1s == nil {
		return errors.New("glVertexAttrib1s")
	}
	gpVertexAttrib1sv = (C.GPVERTEXATTRIB1SV)(getProcAddr("glVertexAttrib1sv"))
	if gpVertexAttrib1sv == nil {
		return errors.New("glVertexAttrib1sv")
	}
	gpVertexAttrib2d = (C.GPVERTEXATTRIB2D)(getProcAddr("glVertexAttrib2d"))
	if gpVertexAttrib2d == nil {
		return errors.New("glVertexAttrib2d")
	}
	gpVertexAttrib2dv = (C.GPVERTEXATTRIB2DV)(getProcAddr("glVertexAttrib2dv"))
	if gpVertexAttrib2dv == nil {
		return errors.New("glVertexAttrib2dv")
	}
	gpVertexAttrib2f = (C.GPVERTEXATTRIB2F)(getProcAddr("glVertexAttrib2f"))
	if gpVertexAttrib2f == nil {
		return errors.New("glVertexAttrib2f")
	}
	gpVertexAttrib2fv = (C.GPVERTEXATTRIB2FV)(getProcAddr("glVertexAttrib2fv"))
	if gpVertexAttrib2fv == nil {
		return errors.New("glVertexAttrib2fv")
	}
	gpVertexAttrib2s = (C.GPVERTEXATTRIB2S)(getProcAddr("glVertexAttrib2s"))
	if gpVertexAttrib2s == nil {
		return errors.New("glVertexAttrib2s")
	}
	gpVertexAttrib2sv = (C.GPVERTEXATTRIB2SV)(getProcAddr("glVertexAttrib2sv"))
	if gpVertexAttrib2sv == nil {
		return errors.New("glVertexAttrib2sv")
	}
	gpVertexAttrib3d = (C.GPVERTEXATTRIB3D)(getProcAddr("glVertexAttrib3d"))
	if gpVertexAttrib3d == nil {
		return errors.New("glVertexAttrib3d")
	}
	gpVertexAttrib3dv = (C.GPVERTEXATTRIB3DV)(getProcAddr("glVertexAttrib3dv"))
	if gpVertexAttrib3dv == nil {
		return errors.New("glVertexAttrib3dv")
	}
	gpVertexAttrib3f = (C.GPVERTEXATTRIB3F)(getProcAddr("glVertexAttrib3f"))
	if gpVertexAttrib3f == nil {
		return errors.New("glVertexAttrib3f")
	}
	gpVertexAttrib3fv = (C.GPVERTEXATTRIB3FV)(getProcAddr("glVertexAttrib3fv"))
	if gpVertexAttrib3fv == nil {
		return errors.New("glVertexAttrib3fv")
	}
	gpVertexAttrib3s = (C.GPVERTEXATTRIB3S)(getProcAddr("glVertexAttrib3s"))
	if gpVertexAttrib3s == nil {
		return errors.New("glVertexAttrib3s")
	}
	gpVertexAttrib3sv = (C.GPVERTEXATTRIB3SV)(getProcAddr("glVertexAttrib3sv"))
	if gpVertexAttrib3sv == nil {
		return errors.New("glVertexAttrib3sv")
	}
	gpVertexAttrib4Nbv = (C.GPVERTEXATTRIB4NBV)(getProcAddr("glVertexAttrib4Nbv"))
	if gpVertexAttrib4Nbv == nil {
		return errors.New("glVertexAttrib4Nbv")
	}
	gpVertexAttrib4Niv = (C.GPVERTEXATTRIB4NIV)(getProcAddr("glVertexAttrib4Niv"))
	if gpVertexAttrib4Niv == nil {
		return errors.New("glVertexAttrib4Niv")
	}
	gpVertexAttrib4Nsv = (C.GPVERTEXATTRIB4NSV)(getProcAddr("glVertexAttrib4Nsv"))
	if gpVertexAttrib4Nsv == nil {
		return errors.New("glVertexAttrib4Nsv")
	}
	gpVertexAttrib4Nub = (C.GPVERTEXATTRIB4NUB)(getProcAddr("glVertexAttrib4Nub"))
	if gpVertexAttrib4Nub == nil {
		return errors.New("glVertexAttrib4Nub")
	}
	gpVertexAttrib4Nubv = (C.GPVERTEXATTRIB4NUBV)(getProcAddr("glVertexAttrib4Nubv"))
	if gpVertexAttrib4Nubv == nil {
		return errors.New("glVertexAttrib4Nubv")
	}
	gpVertexAttrib4Nuiv = (C.GPVERTEXATTRIB4NUIV)(getProcAddr("glVertexAttrib4Nuiv"))
	if gpVertexAttrib4Nuiv == nil {
		return errors.New("glVertexAttrib4Nuiv")
	}
	gpVertexAttrib4Nusv = (C.GPVERTEXATTRIB4NUSV)(getProcAddr("glVertexAttrib4Nusv"))
	if gpVertexAttrib4Nusv == nil {
		return errors.New("glVertexAttrib4Nusv")
	}
	gpVertexAttrib4bv = (C.GPVERTEXATTRIB4BV)(getProcAddr("glVertexAttrib4bv"))
	if gpVertexAttrib4bv == nil {
		return errors.New("glVertexAttrib4bv")
	}
	gpVertexAttrib4d = (C.GPVERTEXATTRIB4D)(getProcAddr("glVertexAttrib4d"))
	if gpVertexAttrib4d == nil {
		return errors.New("glVertexAttrib4d")
	}
	gpVertexAttrib4dv = (C.GPVERTEXATTRIB4DV)(getProcAddr("glVertexAttrib4dv"))
	if gpVertexAttrib4dv == nil {
		return errors.New("glVertexAttrib4dv")
	}
	gpVertexAttrib4f = (C.GPVERTEXATTRIB4F)(getProcAddr("glVertexAttrib4f"))
	if gpVertexAttrib4f == nil {
		return errors.New("glVertexAttrib4f")
	}
	gpVertexAttrib4fv = (C.GPVERTEXATTRIB4FV)(getProcAddr("glVertexAttrib4fv"))
	if gpVertexAttrib4fv == nil {
		return errors.New("glVertexAttrib4fv")
	}
	gpVertexAttrib4iv = (C.GPVERTEXATTRIB4IV)(getProcAddr("glVertexAttrib4iv"))
	if gpVertexAttrib4iv == nil {
		return errors.New("glVertexAttrib4iv")
	}
	gpVertexAttrib4s = (C.GPVERTEXATTRIB4S)(getProcAddr("glVertexAttrib4s"))
	if gpVertexAttrib4s == nil {
		return errors.New("glVertexAttrib4s")
	}
	gpVertexAttrib4sv = (C.GPVERTEXATTRIB4SV)(getProcAddr("glVertexAttrib4sv"))
	if gpVertexAttrib4sv == nil {
		return errors.New("glVertexAttrib4sv")
	}
	gpVertexAttrib4ubv = (C.GPVERTEXATTRIB4UBV)(getProcAddr("glVertexAttrib4ubv"))
	if gpVertexAttrib4ubv == nil {
		return errors.New("glVertexAttrib4ubv")
	}
	gpVertexAttrib4uiv = (C.GPVERTEXATTRIB4UIV)(getProcAddr("glVertexAttrib4uiv"))
	if gpVertexAttrib4uiv == nil {
		return errors.New("glVertexAttrib4uiv")
	}
	gpVertexAttrib4usv = (C.GPVERTEXATTRIB4USV)(getProcAddr("glVertexAttrib4usv"))
	if gpVertexAttrib4usv == nil {
		return errors.New("glVertexAttrib4usv")
	}
	gpVertexAttribBinding = (C.GPVERTEXATTRIBBINDING)(getProcAddr("glVertexAttribBinding"))
	gpVertexAttribDivisor = (C.GPVERTEXATTRIBDIVISOR)(getProcAddr("glVertexAttribDivisor"))
	if gpVertexAttribDivisor == nil {
		return errors.New("glVertexAttribDivisor")
	}
	gpVertexAttribFormat = (C.GPVERTEXATTRIBFORMAT)(getProcAddr("glVertexAttribFormat"))
	gpVertexAttribI1i = (C.GPVERTEXATTRIBI1I)(getProcAddr("glVertexAttribI1i"))
	if gpVertexAttribI1i == nil {
		return errors.New("glVertexAttribI1i")
	}
	gpVertexAttribI1iv = (C.GPVERTEXATTRIBI1IV)(getProcAddr("glVertexAttribI1iv"))
	if gpVertexAttribI1iv == nil {
		return errors.New("glVertexAttribI1iv")
	}
	gpVertexAttribI1ui = (C.GPVERTEXATTRIBI1UI)(getProcAddr("glVertexAttribI1ui"))
	if gpVertexAttribI1ui == nil {
		return errors.New("glVertexAttribI1ui")
	}
	gpVertexAttribI1uiv = (C.GPVERTEXATTRIBI1UIV)(getProcAddr("glVertexAttribI1uiv"))
	if gpVertexAttribI1uiv == nil {
		return errors.New("glVertexAttribI1uiv")
	}
	gpVertexAttribI2i = (C.GPVERTEXATTRIBI2I)(getProcAddr("glVertexAttribI2i"))
	if gpVertexAttribI2i == nil {
		return errors.New("glVertexAttribI2i")
	}
	gpVertexAttribI2iv = (C.GPVERTEXATTRIBI2IV)(getProcAddr("glVertexAttribI2iv"))
	if gpVertexAttribI2iv == nil {
		return errors.New("glVertexAttribI2iv")
	}
	gpVertexAttribI2ui = (C.GPVERTEXATTRIBI2UI)(getProcAddr("glVertexAttribI2ui"))
	if gpVertexAttribI2ui == nil {
		return errors.New("glVertexAttribI2ui")
	}
	gpVertexAttribI2uiv = (C.GPVERTEXATTRIBI2UIV)(getProcAddr("glVertexAttribI2uiv"))
	if gpVertexAttribI2uiv == nil {
		return errors.New("glVertexAttribI2uiv")
	}
	gpVertexAttribI3i = (C.GPVERTEXATTRIBI3I)(getProcAddr("glVertexAttribI3i"))
	if gpVertexAttribI3i == nil {
		return errors.New("glVertexAttribI3i")
	}
	gpVertexAttribI3iv = (C.GPVERTEXATTRIBI3IV)(getProcAddr("glVertexAttribI3iv"))
	if gpVertexAttribI3iv == nil {
		return errors.New("glVertexAttribI3iv")
	}
	gpVertexAttribI3ui = (C.GPVERTEXATTRIBI3UI)(getProcAddr("glVertexAttribI3ui"))
	if gpVertexAttribI3ui == nil {
		return errors.New("glVertexAttribI3ui")
	}
	gpVertexAttribI3uiv = (C.GPVERTEXATTRIBI3UIV)(getProcAddr("glVertexAttribI3uiv"))
	if gpVertexAttribI3uiv == nil {
		return errors.New("glVertexAttribI3uiv")
	}
	gpVertexAttribI4bv = (C.GPVERTEXATTRIBI4BV)(getProcAddr("glVertexAttribI4bv"))
	if gpVertexAttribI4bv == nil {
		return errors.New("glVertexAttribI4bv")
	}
	gpVertexAttribI4i = (C.GPVERTEXATTRIBI4I)(getProcAddr("glVertexAttribI4i"))
	if gpVertexAttribI4i == nil {
		return errors.New("glVertexAttribI4i")
	}
	gpVertexAttribI4iv = (C.GPVERTEXATTRIBI4IV)(getProcAddr("glVertexAttribI4iv"))
	if gpVertexAttribI4iv == nil {
		return errors.New("glVertexAttribI4iv")
	}
	gpVertexAttribI4sv = (C.GPVERTEXATTRIBI4SV)(getProcAddr("glVertexAttribI4sv"))
	if gpVertexAttribI4sv == nil {
		return errors.New("glVertexAttribI4sv")
	}
	gpVertexAttribI4ubv = (C.GPVERTEXATTRIBI4UBV)(getProcAddr("glVertexAttribI4ubv"))
	if gpVertexAttribI4ubv == nil {
		return errors.New("glVertexAttribI4ubv")
	}
	gpVertexAttribI4ui = (C.GPVERTEXATTRIBI4UI)(getProcAddr("glVertexAttribI4ui"))
	if gpVertexAttribI4ui == nil {
		return errors.New("glVertexAttribI4ui")
	}
	gpVertexAttribI4uiv = (C.GPVERTEXATTRIBI4UIV)(getProcAddr("glVertexAttribI4uiv"))
	if gpVertexAttribI4uiv == nil {
		return errors.New("glVertexAttribI4uiv")
	}
	gpVertexAttribI4usv = (C.GPVERTEXATTRIBI4USV)(getProcAddr("glVertexAttribI4usv"))
	if gpVertexAttribI4usv == nil {
		return errors.New("glVertexAttribI4usv")
	}
	gpVertexAttribIFormat = (C.GPVERTEXATTRIBIFORMAT)(getProcAddr("glVertexAttribIFormat"))
	gpVertexAttribIPointer = (C.GPVERTEXATTRIBIPOINTER)(getProcAddr("glVertexAttribIPointer"))
	if gpVertexAttribIPointer == nil {
		return errors.New("glVertexAttribIPointer")
	}
	gpVertexAttribL1d = (C.GPVERTEXATTRIBL1D)(getProcAddr("glVertexAttribL1d"))
	if gpVertexAttribL1d == nil {
		return errors.New("glVertexAttribL1d")
	}
	gpVertexAttribL1dv = (C.GPVERTEXATTRIBL1DV)(getProcAddr("glVertexAttribL1dv"))
	if gpVertexAttribL1dv == nil {
		return errors.New("glVertexAttribL1dv")
	}
	gpVertexAttribL1ui64ARB = (C.GPVERTEXATTRIBL1UI64ARB)(getProcAddr("glVertexAttribL1ui64ARB"))
	gpVertexAttribL1ui64vARB = (C.GPVERTEXATTRIBL1UI64VARB)(getProcAddr("glVertexAttribL1ui64vARB"))
	gpVertexAttribL2d = (C.GPVERTEXATTRIBL2D)(getProcAddr("glVertexAttribL2d"))
	if gpVertexAttribL2d == nil {
		return errors.New("glVertexAttribL2d")
	}
	gpVertexAttribL2dv = (C.GPVERTEXATTRIBL2DV)(getProcAddr("glVertexAttribL2dv"))
	if gpVertexAttribL2dv == nil {
		return errors.New("glVertexAttribL2dv")
	}
	gpVertexAttribL3d = (C.GPVERTEXATTRIBL3D)(getProcAddr("glVertexAttribL3d"))
	if gpVertexAttribL3d == nil {
		return errors.New("glVertexAttribL3d")
	}
	gpVertexAttribL3dv = (C.GPVERTEXATTRIBL3DV)(getProcAddr("glVertexAttribL3dv"))
	if gpVertexAttribL3dv == nil {
		return errors.New("glVertexAttribL3dv")
	}
	gpVertexAttribL4d = (C.GPVERTEXATTRIBL4D)(getProcAddr("glVertexAttribL4d"))
	if gpVertexAttribL4d == nil {
		return errors.New("glVertexAttribL4d")
	}
	gpVertexAttribL4dv = (C.GPVERTEXATTRIBL4DV)(getProcAddr("glVertexAttribL4dv"))
	if gpVertexAttribL4dv == nil {
		return errors.New("glVertexAttribL4dv")
	}
	gpVertexAttribLFormat = (C.GPVERTEXATTRIBLFORMAT)(getProcAddr("glVertexAttribLFormat"))
	gpVertexAttribLPointer = (C.GPVERTEXATTRIBLPOINTER)(getProcAddr("glVertexAttribLPointer"))
	if gpVertexAttribLPointer == nil {
		return errors.New("glVertexAttribLPointer")
	}
	gpVertexAttribP1ui = (C.GPVERTEXATTRIBP1UI)(getProcAddr("glVertexAttribP1ui"))
	if gpVertexAttribP1ui == nil {
		return errors.New("glVertexAttribP1ui")
	}
	gpVertexAttribP1uiv = (C.GPVERTEXATTRIBP1UIV)(getProcAddr("glVertexAttribP1uiv"))
	if gpVertexAttribP1uiv == nil {
		return errors.New("glVertexAttribP1uiv")
	}
	gpVertexAttribP2ui = (C.GPVERTEXATTRIBP2UI)(getProcAddr("glVertexAttribP2ui"))
	if gpVertexAttribP2ui == nil {
		return errors.New("glVertexAttribP2ui")
	}
	gpVertexAttribP2uiv = (C.GPVERTEXATTRIBP2UIV)(getProcAddr("glVertexAttribP2uiv"))
	if gpVertexAttribP2uiv == nil {
		return errors.New("glVertexAttribP2uiv")
	}
	gpVertexAttribP3ui = (C.GPVERTEXATTRIBP3UI)(getProcAddr("glVertexAttribP3ui"))
	if gpVertexAttribP3ui == nil {
		return errors.New("glVertexAttribP3ui")
	}
	gpVertexAttribP3uiv = (C.GPVERTEXATTRIBP3UIV)(getProcAddr("glVertexAttribP3uiv"))
	if gpVertexAttribP3uiv == nil {
		return errors.New("glVertexAttribP3uiv")
	}
	gpVertexAttribP4ui = (C.GPVERTEXATTRIBP4UI)(getProcAddr("glVertexAttribP4ui"))
	if gpVertexAttribP4ui == nil {
		return errors.New("glVertexAttribP4ui")
	}
	gpVertexAttribP4uiv = (C.GPVERTEXATTRIBP4UIV)(getProcAddr("glVertexAttribP4uiv"))
	if gpVertexAttribP4uiv == nil {
		return errors.New("glVertexAttribP4uiv")
	}
	gpVertexAttribPointer = (C.GPVERTEXATTRIBPOINTER)(getProcAddr("glVertexAttribPointer"))
	if gpVertexAttribPointer == nil {
		return errors.New("glVertexAttribPointer")
	}
	gpVertexBindingDivisor = (C.GPVERTEXBINDINGDIVISOR)(getProcAddr("glVertexBindingDivisor"))
	gpViewport = (C.GPVIEWPORT)(getProcAddr("glViewport"))
	if gpViewport == nil {
		return errors.New("glViewport")
	}
	gpViewportArrayv = (C.GPVIEWPORTARRAYV)(getProcAddr("glViewportArrayv"))
	if gpViewportArrayv == nil {
		return errors.New("glViewportArrayv")
	}
	gpViewportIndexedf = (C.GPVIEWPORTINDEXEDF)(getProcAddr("glViewportIndexedf"))
	if gpViewportIndexedf == nil {
		return errors.New("glViewportIndexedf")
	}
	gpViewportIndexedfv = (C.GPVIEWPORTINDEXEDFV)(getProcAddr("glViewportIndexedfv"))
	if gpViewportIndexedfv == nil {
		return errors.New("glViewportIndexedfv")
	}
	gpWaitSync = (C.GPWAITSYNC)(getProcAddr("glWaitSync"))
	if gpWaitSync == nil {
		return errors.New("glWaitSync")
	}
	return nil
}
