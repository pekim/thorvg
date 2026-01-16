package thorvg

type Scene struct {
	paintCommon
}

/*
SceneNew Creates a new Scene object.

This function allocates and returns a new Scene instance.
To properly destroy the Scene object, use @ref tvg_paint_rel().

	@return A pointer to the newly created Scene object.

	@see tvg_paint_rel()
*/
func SceneNew() Scene {
	return Scene{
		paintCommon: paintCommon{
			paint_: tvg_scene_new(),
		},
	}
}

/*
Add adds a paint object to the scene.

This function appends a paint object to the scene.

	@param[in] scene A Tvg_Paint pointer to the scene object.
	@param[in] paint A pointer to the Paint object to be added into the scene.

	@note The ownership of the @p paint object is transferred to the scene upon addition.

	@see tvg_scene_remove()
	@see tvg_scene_insert()
*/
func (scene Scene) Add(paint Paint) error {
	return tvg_scene_add(scene.paint_, paint.paint()).error()
}

/*
Insert adds a paint object to the scene.

This function appends a paint object to the scene. The new paint object @p target will
be inserted immediately before the specified paint object @p at in the scene.

	@param[in] scene A Tvg_Paint pointer to the scene object.
	@param[in] target A pointer to the Paint object to be added into the scene.
	@param[in] at A pointer to an existing Paint object in the scene before which
							the new paint object will be added. This parameter must not be @c nullptr.

	@note The ownership of the @p paint object is transferred to the scene upon addition.

	@see tvg_scene_remove()
	@see tvg_scene_add()
	@since 1.0
*/
func (scene Scene) Insert(target Paint, at Paint) error {
	return tvg_scene_insert(scene.paint_, target.paint(), at.paint()).error()
}

/*
Remove removes a paint object from the scene.

This function removes a specified paint object from the scene. If no paint
object is specified (i.e., the default @c nullptr is used), the function
performs to clear all paints from the scene.

	@param[in] scene A Tvg_Paint pointer to the scene object.
	@param[in] paint A pointer to the Paint object to be removed from the scene.
								If @c nullptr, remove all the paints from the scene.

	@see tvg_scene_add()
	@since 1.0
*/
func (scene Scene) Remove(paint Paint) error {
	return tvg_scene_remove(scene.paint_, paint.paint()).error()
}

/*
ClearEffects resets all previously applied scene effects.

This function clears all effects that have been applied to the scene,
restoring it to its original state without any post-processing.

	@param[in] scene A pointer to the Tvg_Paint scene object.

	@since 1.0
*/
func (scene Scene) ClearEffects() error {
	return tvg_scene_clear_effects(scene.paint_).error()
}

/*
AddEffectGaussianBlur applies a Gaussian blur effect to the scene.

This function applies a Gaussian blur filter to the scene as a post-processing effect.
The blur can be applied in different directions with configurable border handling and quality settings.

	@param[in] scene A pointer to the Tvg_Paint scene object.
	@param[in] sigma The blur radius (sigma) value. Must be greater than 0.
	@param[in] direction Blur direction: 0 = both directions, 1 = horizontal only, 2 = vertical only.
	@param[in] border Border handling method: 0 = duplicate, 1 = wrap.
	@param[in] quality Blur quality level [0 - 100].

	@since 1.0
*/
func (scene Scene) AddEffectGaussianBlur(sigma float64, direction int, border int, quality int) error {
	return tvg_scene_add_effect_gaussian_blur(
		scene.paint_, sigma,
		int32(direction), int32(border), int32(quality),
	).error()
}

/*
AddEffectDropShadow applies a drop shadow effect to the scene.

This function applies a drop shadow with a Gaussian blur to the scene. The shadow
can be customized using color, opacity, angle, distance, blur radius (sigma),
and quality parameters.

	@param[in] scene A pointer to the Tvg_Paint scene object.
	@param[in] r Red channel value of the shadow color [0 - 255].
	@param[in] g Green channel value of the shadow color [0 - 255].
	@param[in] b Blue channel value of the shadow color [0 - 255].
	@param[in] a Alpha (opacity) channel value of the shadow [0 - 255].
	@param[in] angle Shadow direction in degrees [0 - 360].
	@param[in] distance Distance of the shadow from the original object.
	@param[in] sigma Gaussian blur sigma value for the shadow. Must be > 0.
	@param[in] quality Blur quality level [0 - 100].

	@since 1.0
*/
func (scene Scene) AddEffectDropShadow(
	r int, g int, b int, a int,
	angle float64, distance float64, sigma float64, quality int,
) error {
	return tvg_scene_add_effect_drop_shadow(scene.paint_,
		int32(r), int32(g), int32(b), int32(a),
		angle, distance, sigma, int32(quality),
	).error()
}

/*
AddEffectFill applies a fill color effect to the scene.

This function overrides the scene's content colors with the specified fill color.

	@param[in] scene A pointer to the Tvg_Paint scene object.
	@param[in] r Red color channel value [0 - 255].
	@param[in] g Green color channel value [0 - 255].
	@param[in] b Blue color channel value [0 - 255].
	@param[in] a Alpha (opacity) channel value [0 - 255].

	@since 1.0
*/
func (scene Scene) AddEffectFill(r int, g int, b int, a int) error {
	return tvg_scene_add_effect_fill(scene.paint_, int32(r), int32(g), int32(b), int32(a)).error()
}

/*
AddEffectTint applies a tint effect to the scene.

This function tints the current scene using specified black and white color values,
modulated by a given intensity.

	@param[in] scene A pointer to the Tvg_Paint scene object.
	@param[in] black_r Red component of the black color [0 - 255].
	@param[in] black_g Green component of the black color [0 - 255].
	@param[in] black_b Blue component of the black color [0 - 255].
	@param[in] white_r Red component of the white color [0 - 255].
	@param[in] white_g Green component of the white color [0 - 255].
	@param[in] white_b Blue component of the white color [0 - 255].
	@param[in] intensity Tint intensity value [0 - 100].

	@since 1.0
*/
func (scene Scene) AddEffectTint(
	black_r int, black_g int, black_b int,
	white_r int, white_g int, white_b int,
	intensity float64,
) error {
	return tvg_scene_add_effect_tint(scene.paint_,
		int32(black_r), int32(black_g), int32(black_b),
		int32(white_r), int32(white_g), int32(white_b),
		intensity,
	).error()
}

/*
AddEffectTritone applies a tritone color effect to the scene.

This function applies a tritone color effect to the given scene using three sets of RGB values
representing shadow, midtone, and highlight colors.

	@param[in] scene A pointer to the Tvg_Paint scene object.
	@param[in] shadow_r Red component of the shadow color [0 - 255].
	@param[in] shadow_g Green component of the shadow color [0 - 255].
	@param[in] shadow_b Blue component of the shadow color [0 - 255].
	@param[in] midtone_r Red component of the midtone color [0 - 255].
	@param[in] midtone_g Green component of the midtone color [0 - 255].
	@param[in] midtone_b Blue component of the midtone color [0 - 255].
	@param[in] highlight_r Red component of the highlight color [0 - 255].
	@param[in] highlight_g Green component of the highlight color [0 - 255].
	@param[in] highlight_b Blue component of the highlight color [0 - 255].
	@param[in] blend A blending factor that determines the mix between the original color and the tritone colors [0 - 255].

	@since 1.0
*/
func (scene Scene) AddEffectTritone(
	shadow_r int, shadow_g int, shadow_b int,
	midtone_r int, midtone_g int, midtone_b int,
	highlight_r int, highlight_g int, highlight_b int,
	blend int,
) error {
	return tvg_scene_add_effect_tritone(scene.paint_,
		int32(shadow_r), int32(shadow_g), int32(shadow_b),
		int32(midtone_r), int32(midtone_g), int32(midtone_b),
		int32(highlight_r), int32(highlight_g), int32(highlight_b),
		int32(blend),
	).error()
}

/*
Duplicate duplicates a Scene.

Creates a new object and sets its all properties as in the original object.

	@param[in] paint The Tvg_Paint object to be copied.

	@return A copied Tvg_Paint object if succeed, @c nullptr otherwise.
*/
func (scene Scene) Duplicate() Scene {
	return Scene{
		paintCommon: scene.duplicate(),
	}
}
